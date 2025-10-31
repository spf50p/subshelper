package tpl

import (
	"bytes"
	"html/template"
)

func Execute(v any, tplName string, tplContent string) (string, error) {
	tpl := template.New(tplName)
	_, err := tpl.Parse(tplContent)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, v)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
