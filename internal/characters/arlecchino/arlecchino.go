package arlecchino

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/model"
)

func init() {
	core.RegisterCharFunc(keys.Arlecchino, NewChar)
}

type char struct {
	*tmpl.Character
	skillDebt             float64
	skillDebtMax          float64
	initialDirectiveLevel int
	bondOnBurst           float64
}

func NewChar(s *core.Core, w *character.CharWrapper, _ info.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = base.SkillDetails.BurstEnergyCost
	c.NormalHitNum = normalHitNum
	c.NormalCon = 3
	c.BurstCon = 5

	w.Character = &c

	return nil
}

func (c *char) Init() error {
	c.naBuff()
	c.passive()
	c.a1OnKill()
	c.a4()

	c.c2()
	c.c4()
	c.c6()
	return nil
}

func (c *char) NextQueueItemIsValid(a action.Action, p map[string]int) error {
	// can use charge without attack beforehand unlike most of the other polearm users
	if a == action.ActionCharge {
		return nil
	}
	return c.Character.NextQueueItemIsValid(a, p)
}

func (c *char) AnimationStartDelay(k model.AnimationDelayKey) int {
	if k == model.AnimationXingqiuN0StartDelay {
		return 7
	}
	return c.Character.AnimationStartDelay(k)
}

func (c *char) getTotalAtk() float64 {
	stats, _ := c.Stats()
	return c.Base.Atk*(1+stats[attributes.ATKP]) + stats[attributes.ATK]
}
