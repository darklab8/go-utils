package enverant

/*
Manager for getting values from Environment variables
*/

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/darklab8/go-utils/utils/ptr"
)

type Enverant struct {
	file_envs        map[string]interface{}
	validate_missing bool

	Description        string
	prefix             string
	all_prefixed_keys  map[string]*EnvParamsData
	used_prefixed_keys map[string]bool // useful for validation of finding not used keys
}

type EnvParamsData struct {
	*ValueParams
	PrefixedKey string
	Value       string
}

func (e *Enverant) GetParams() []*EnvParamsData {
	var infos []*EnvParamsData
	for _, value := range e.all_prefixed_keys {
		infos = append(infos, value)
	}
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].PrefixedKey < infos[j].PrefixedKey
	})
	return infos
}

func (e *Enverant) ValidetNoUnused() {

	if e.prefix == "" {
		fmt.Println("e.prefix is empty. Don't validate not used env vars if not using prefix")
		os.Exit(1)
	}

	os.Environ()

	env := os.Environ()

	for _, v := range env {
		// split the key and value
		key := strings.Split(v, "=")[0]

		if !strings.HasPrefix(key, e.prefix) {
			continue
		}

		// exception as thing used in typelog
		if key == e.prefix+"LOG_LEVEL" {
			continue
		}

		if _, ok := e.used_prefixed_keys[key]; !ok {
			fmt.Println("found darkstat not used env var. Change key to smth else, key=", key)
			os.Exit(1)
		}
	}
}

type EnverantOption func(m *Enverant)

func NewEnverant(opts ...EnverantOption) *Enverant {
	e := &Enverant{
		file_envs:          map[string]interface{}{},
		all_prefixed_keys:  map[string]*EnvParamsData{},
		used_prefixed_keys: map[string]bool{},
	}
	if path, ok := e.LookupEnv("ENVERANT_ENV_FILE", &ValueParams{Description: "where to seek dev env file for inputing env vars through env file"}); ok {
		e.file_envs = ReadJson(path)
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func WithValidate(validate bool) EnverantOption {
	return func(m *Enverant) {
		m.validate_missing = validate
	}
}

func WithDescription(Description string) EnverantOption {
	return func(m *Enverant) {
		m.Description = Description
	}
}

func WithPrefix(prefix string) EnverantOption {
	return func(m *Enverant) {
		m.prefix = prefix
	}
}

func (m *Enverant) GetValidating() *Enverant {
	var clone *Enverant = &Enverant{}
	*clone = *m
	clone.validate_missing = true
	return clone
}

func EnrichStr(value string) string {
	// unhardcode later
	if strings.Contains(value, "${env:HOME}") {
		value = strings.ReplaceAll(value, "${env:HOME}", os.Getenv("HOME"))
	}
	return value
}

type ValueParams struct {
	VarType     VarType
	Default     any
	Description string
}
type ValueOption func(m *ValueParams)

/*
WithDesc - is used to build help information
*/
func WithDesc(description string) ValueOption {
	return func(m *ValueParams) {
		m.Description = description
	}
}

func OrStr(default_ string) ValueOption {
	return func(m *ValueParams) {
		m.Default = default_
	}
}

func OrInt(default_ int) ValueOption {
	return func(m *ValueParams) {
		m.Default = default_
	}
}

func OrBool(default_ bool) ValueOption {
	return func(m *ValueParams) {
		m.Default = default_
	}
}

func (e *Enverant) GetStrOr(key string, default_ string, opts ...ValueOption) string {
	value, _ := e.GetString(key, append([]ValueOption{OrStr(default_)}, opts...)...)
	return value
}

func (e *Enverant) GetStr(key string, opts ...ValueOption) string {
	if value, ok := e.GetString(key, opts...); ok {
		return value
	}
	return ""
}

func (e *Enverant) GetPtrStr(key string, opts ...ValueOption) *string {
	if value, ok := e.GetString(key, opts...); ok {
		return ptr.Ptr(value)
	}
	return nil
}

func (e *Enverant) GetString(key string, opts ...ValueOption) (string, bool) {
	params := &ValueParams{
		VarType: VarStr,
	}
	for _, opt := range opts {
		opt(params)
	}

	if value, ok := e.LookupEnv(key, params); ok {
		found_str := EnrichStr(value)

		if strings.Contains(found_str, "pass[") {
			_, err_to_find_pass := exec.LookPath("pass")
			if err_to_find_pass != nil {
				fmt.Println("pass not found in PATH")
				return "", false
			}

			keys := strings.Split(found_str, " ")
			json_key := strings.ReplaceAll(keys[0], "pass[", "")
			json_key = strings.ReplaceAll(json_key, "]", "")

			pass_key := keys[1]

			cmd := exec.Command("pass", pass_key)

			output, err := cmd.CombinedOutput()
			if err != nil {
				panic(fmt.Sprintln("failed to get pass output, err=", err.Error()))
			}

			var result map[string]string
			err = json.Unmarshal(output, &result)
			if err != nil {
				panic(err)
			}

			found_str = result[json_key]
		}

		return found_str, true
	}

	if params.Default != nil {
		return params.Default.(string), true
	}

	if e.validate_missing {
		panic(fmt.Sprintln("enverant value is not defined, key=", key))
	}

	return "", false
}

func (e *Enverant) GetBoolOr(key string, default_ bool, opts ...ValueOption) bool {
	value, _ := e.GetBoolean(key, append([]ValueOption{OrBool(default_)}, opts...)...)
	return value
}

func (e *Enverant) GetBool(key string, opts ...ValueOption) bool {
	if value, ok := e.GetBoolean(key, opts...); ok {
		return value
	}
	return false
}

func (e *Enverant) GetPtrBool(key string, opts ...ValueOption) *bool {
	if value, ok := e.GetBoolean(key, opts...); ok {
		return ptr.Ptr(value)
	}
	return nil
}

func (e *Enverant) LookupEnv(key string, params *ValueParams) (string, bool) {
	info := &EnvParamsData{
		ValueParams: params,
		PrefixedKey: e.prefix + key,
	}
	e.all_prefixed_keys[key] = info

	if e.prefix != "" {
		if value, ok := os.LookupEnv(e.prefix + key); ok {
			e.used_prefixed_keys[e.prefix+key] = true
			info.Value = value
			return value, ok
		}
	}

	if value, ok := os.LookupEnv(key); ok {
		info.Value = value
		return value, ok
	}

	return "", false
}

type VarType int64

const (
	VarBool VarType = iota
	VarInt
	VarStr
)

func (v VarType) ToStr() string {
	switch v {
	case VarBool:
		return "bool"
	case VarInt:
		return "int"
	case VarStr:
		return "str"

	}
	panic("undefined")
}

func (e *Enverant) GetBoolean(key string, opts ...ValueOption) (bool, bool) {
	params := &ValueParams{
		VarType: VarBool,
	}
	for _, opt := range opts {
		opt(params)
	}

	if value, ok := e.LookupEnv(key, params); ok {
		return value == "true", true
	}

	if params.Default != nil {
		return params.Default.(bool), true
	}

	if e.validate_missing {
		panic(fmt.Sprintln("enverant value is not defined, key=", key))
	}

	return false, false
}

func (e *Enverant) GetIntOr(key string, default_ int, opts ...ValueOption) int {
	value, _ := e.GetInteger(key, append([]ValueOption{OrInt(default_)}, opts...)...)
	return value
}

func (e *Enverant) GetInt(key string, opts ...ValueOption) int {
	if value, ok := e.GetInteger(key, opts...); ok {
		return value
	}
	return 0
}

func (e *Enverant) GetPtrInt(key string, opts ...ValueOption) *int {
	if value, ok := e.GetInteger(key, opts...); ok {
		return ptr.Ptr(value)
	}
	return nil
}

func (e *Enverant) GetInteger(key string, opts ...ValueOption) (int, bool) {
	params := &ValueParams{
		VarType: VarInt,
	}
	for _, opt := range opts {
		opt(params)
	}

	if value, ok := e.LookupEnv(key, params); ok {
		int_value, err := strconv.Atoi(value)
		if err != nil {
			panic(fmt.Sprintln(err, "expected to be int, key=", key))
		}
		return int_value, true
	}

	if params.Default != nil {
		return params.Default.(int), true
	}

	if e.validate_missing {
		panic(fmt.Sprintln("enverant value is not defined, key=", key))
	}

	return 0, false
}
