package templates

import (
	"embed"
	_ "embed"
	"html/template"
	"os"
)

//go:embed *.tmpl
var winTemplate embed.FS

func GetWinTemplate() error {
	tpl, err := template.ParseFS(winTemplate, "*.tmpl")
	if err != nil {
		return err
	}

	err = tpl.ExecuteTemplate(os.Stdout, "singleWin.tmpl", map[string]interface{}{
		"Win":        true,
		"Name":       "测试",
		"HeroName":   "斧王",
		"MatchID":    "1111111",
		"MatchLevel": "Very High",
		"Kills":      "1",
		"Deaths":     "1",
		"Assists":    "1",
	})
	if err != nil {
		return err
	}
	return nil
}
