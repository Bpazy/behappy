package kda

import (
	"fmt"
	"strings"
)

func GetKda(kills, deaths, assist int) string {
	if deaths == 0 {
		deaths = 1
	}
	kda := (float64(kills) + float64(assist)) / float64(deaths)
	s := fmt.Sprintf("%.1f", kda)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}
