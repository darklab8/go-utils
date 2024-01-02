package utils

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/darklab8/darklab_goutils/goutils/logus/utils_logus"
)

func TmpRender(templateRef *template.Template, data interface{}) string {
	header := bytes.Buffer{}
	err := templateRef.Execute(&header, data)
	utils_logus.Log.CheckFatal(err, "failed to render template")
	return header.String()
}

func TmpInit(content string) *template.Template {
	funcs := map[string]any{
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix}

	var err error
	templateRef, err := template.New("test").Funcs(funcs).Parse(content)
	utils_logus.Log.CheckFatal(err, "failed to init template")
	return templateRef
}
