package dao

import (
	"github.com/Bpazy/behappy/models"
	"math/rand"
)

func RandJoke() *models.Joke {
	var count int64
	db.Model(&models.Joke{}).Count(&count)
	if count == 0 {
		return nil
	}

	i := rand.Intn(int(count)) + 1
	var joke models.Joke
	db.Model(&models.Joke{}).Offset(i).First(&joke)

	return &joke
}
