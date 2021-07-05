package templates

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"
)

//go:embed *.tmpl
var messageTemplates embed.FS

//GetSingleMessage 获取单排的模板消息
func GetSingleMessage(data map[string]interface{}) (string, error) {
	tpl, err := template.ParseFS(messageTemplates, "*.tmpl")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tpl.ExecuteTemplate(buf, "single.tmpl", data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

//GetMultiMessage 获取组排的模板消息
func GetMultiMessage(data map[string]interface{}) (string, error) {
	tpl, err := template.ParseFS(messageTemplates, "*.tmpl")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tpl.ExecuteTemplate(buf, "multi.tmpl", data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
