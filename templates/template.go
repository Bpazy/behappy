package templates

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"
)

//go:embed *.tmpl
var winTemplate embed.FS

//GetSingleMessage 获取单排的模板消息
func GetSingleMessage(data map[string]interface{}) (string, error) {
	tpl, err := template.ParseFS(winTemplate, "*.tmpl")
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = tpl.ExecuteTemplate(buf, "message.tmpl", data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
