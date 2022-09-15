package dao

import (
	"context"
	"github.com/Bpazy/behappy/dto"
	"github.com/Bpazy/behappy/ent"
	"github.com/Bpazy/behappy/ent/hero"
	"github.com/Bpazy/behappy/util/berrors"
	"strconv"
)

func GetHeroName(heroID int) string {
	h := client.Hero.Query().
		Where(hero.HeroID(heroID)).
		FirstX(context.Background())
	if h != nil {
		return h.LocalizedName
	}
	return strconv.Itoa(heroID)
}

func HasHeroData() bool {
	h := client.Hero.Query().FirstX(context.Background())
	return h != nil
}

func AddHeros(heros []dto.HeroDto) {
	ctx := context.Background()
	tx := berrors.Unwrap(client.Tx(ctx))
	tx.Hero.Delete().ExecX(ctx)
	bulk := make([]*ent.HeroCreate, len(heros))
	for i, h := range heros {
		bulk[i] = tx.Hero.Create().SetHeroID(h.ID).SetName(h.Name).SetLocalizedName(h.LocalizedName)
	}
	tx.Hero.CreateBulk(bulk...).SaveX(ctx)
	berrors.Must(tx.Commit())
}
