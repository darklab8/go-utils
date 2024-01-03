package utils_types

type FilePath string

func (f FilePath) ToString() string { return string(f) }

type RegExp string

type TemplateExpression string
