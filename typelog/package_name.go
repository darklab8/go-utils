package typelog

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

/*
All of this to write file + line of code.
*/

const TP_SKIP = "typelog_skip"

var (
	CtxSkip context.Context = context.Background()
)

func (l *Logger) GetFile() string {
	ctx_skip := 3
	if ctx_skip := CtxSkip.Value(TP_SKIP); ctx_skip != nil {
		ctx_skip = ctx_skip.(int)
	}
	return getFileCaller(ctx_skip)
}

func getFileCaller(skip int) string {
	pc, filename, line, ok := runtime.Caller(skip)

	if !ok {
		return ""
	}
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return ""
	}

	location := filepath.Join(parsePackageName(fn.Name()), filepath.Base(filename)) + fmt.Sprintf(":%d", line)
	root_package := rootPackage(fn.Name())
	cleaned_path := strings.ReplaceAll(location, root_package, "")
	final := filepath.Base(root_package) + cleaned_path
	return final
}

func rootPackage(funcName string) string {
	lastSlash := strings.LastIndex(funcName, "/")
	if lastSlash < 0 {
		lastSlash = 0
	} else {
		lastSlash++
	}
	rest := funcName[lastSlash:]
	if dot := strings.Index(rest, "."); dot >= 0 {
		funcName = funcName[:lastSlash+dot]
	}

	parts := strings.Split(funcName, "/")

	if strings.Contains(parts[0], ".") && len(parts) >= 3 {
		return strings.Join(parts[:3], "/")
	}

	return parts[0]
}

func parsePackageName(funcName string) string {
	lastSlash := strings.LastIndex(funcName, "/")
	if lastSlash < 0 {
		lastSlash = 0
	} else {
		lastSlash++
	}

	rest := funcName[lastSlash:]
	dot := strings.Index(rest, ".")
	if dot < 0 {
		return rest
	}
	return funcName[:lastSlash+dot]
}

func SetCallerSkip(skip int) {
	if extra_skip := CtxSkip.Value(TP_SKIP); extra_skip != nil {
		return
	}
	CtxSkip = context.WithValue(context.Background(), TP_SKIP, skip)
}

func UnsetCallerSkip() {
	CtxSkip = context.Background()
}
