package dao

import (
	"context"
	"github.com/Bpazy/behappy/ent"
	"github.com/Bpazy/behappy/ent/hero"
	"github.com/Bpazy/behappy/models"
	"strconv"
)

func GetHeroName(heroID int) string {
	h := client.Hero.Query().
		Where(hero.ID(heroID)).
		FirstX(context.TODO())
	if h != nil {
		return h.LocalizedName
	}
	return strconv.Itoa(heroID)
}

func HasHeroData() bool {
	h := client.Hero.Query().FirstX(context.TODO())
	return h != nil
}

func AddHeros(heros []models.HeroDto) {
	bulk := make([]*ent.HeroCreate, len(heros))
	for i, h := range heros {
		bulk[i] = client.Hero.Create().SetHeroID(h.ID).SetName(h.Name).SetLocalizedName(h.LocalizedName)
	}
	client.Hero.CreateBulk(bulk...).SaveX(context.TODO())
}
