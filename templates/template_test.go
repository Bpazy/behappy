package templates

import (
	"log"
	"testing"
)

func TestGetWinTemplate(t *testing.T) {
	data := map[string]interface{}{
		"Win":        true,
		"Name":       "测试",
		"HeroName":   "斧王",
		"MatchID":    "1111111",
		"MatchLevel": "Very High",
		"Kills":      1,
		"Deaths":     1,
		"Assists":    1,
		"winTimes":   1,
		"loseTimes":  0,
	}
	output, err := GetSingleMessage(data)
	if err != nil {
		t.Error(err)
	}
	log.Println(output)

	data["Win"] = false
	output, err = GetSingleMessage(data)
	if err != nil {
		t.Error(err)
	}
	log.Println(output)
}
