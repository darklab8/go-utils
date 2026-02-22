package utils

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/darklab8/go-utils/utils/utils_logus"
	"github.com/darklab8/go-utils/utils/utils_types"
)

func TmpRender(templateRef *template.Template, data interface{}) string {
	header := bytes.Buffer{}
	err := templateRef.Execute(&header, data)
	utils_logus.Log.CheckPanic(err, "failed to render template")
	return header.String()
}

type Tmp struct {
	ExtraFuncs map[string]any
}
type TmpOpt func(t *Tmp)

func TmpInit(content utils_types.TemplateExpression, opts ...TmpOpt) *template.Template {
	funcs := map[string]any{
		"capitalize": strings.Title,
		"contains":   strings.Contains,
		"hasPrefix":  strings.HasPrefix,
		"hasSuffix":  strings.HasSuffix,
		"derefI": func(i *int) int {
			if i == nil {
				return -1
			}
			return *i
		},
		"derefI64": func(i *int64) int64 {
			if i == nil {
				return -1
			}
			return *i
		},
		"derefF64": func(i *float64) float64 {
			if i == nil {
				return -1
			}
			return *i
		},
		"derefS": func(i *string) string {
			if i == nil {
				return "nil"
			}
			return *i
		},
	}

	t := &Tmp{}
	for _, opt := range opts {
		opt(t)
	}

	if t.ExtraFuncs != nil {
		for key, value := range t.ExtraFuncs {
			funcs[key] = value
		}
	}

	var err error
	templateRef, err := template.New("test").Funcs(funcs).Parse(string(content))
	utils_logus.Log.CheckPanic(err, "failed to init template")
	return templateRef
}
