package dao

import (
	"github.com/Bpazy/really/models"
	"strconv"
)

func GetHeroName(heroID int) string {
	var hero models.Hero
	db.Where(&models.Hero{
		ID: heroID,
	}).Find(&hero)

	if hero.ID != 0 {
		return hero.LocalizedName
	}
	return strconv.Itoa(heroID)
}

func HasHeroData() bool {
	err := db.First(&models.Hero{}).Error
	return err == nil
}

func AddHeros(heros []models.Hero) {
	for _, hero := range heros {
		db.Create(&hero)
	}
}
