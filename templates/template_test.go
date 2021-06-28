package templates

import "testing"

func TestGetWinTemplate(t *testing.T) {
	err := GetWinTemplate()
	if err != nil {
		t.Error(err)
	}
}
