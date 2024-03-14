package arlecchino

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/enemy"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

const healMod = 10000

var a1Directive = []float64{0.0, 0.2, 0.25, 0.7}

func (c *char) passive() {
	m := make([]float64, attributes.EndStatType)
	// zeroes out healing from all other sources besides arlecchino's heal
	m[attributes.Heal] = -healMod
	c.AddStatMod(character.StatMod{
		Base:         modifier.NewBase("arlecchino-passive", -1),
		AffectedStat: attributes.Heal,
		Amount: func() ([]float64, bool) {
			return m, true
		},
	})
}

func (c *char) a1OnKill() {
	c.Core.Events.Subscribe(event.OnTargetDied, func(args ...interface{}) bool {
		trg, ok := args[0].(*enemy.Enemy)
		// ignore if not an enemy
		if !ok {
			return false
		}
		if trg.StatusIsActive(directiveKey) {
			c.ModifyHPDebtByRatio(0.7)
		}
		return false
	}, "arlechinno-a1-onkill")
}

func (c *char) a1Upgrade(e combat.Enemy, src int) {
	if c.Base.Ascension < 1 {
		return
	}
	e.QueueEnemyTask(func() {
		level := e.GetTag(directiveKey)
		if level == 0 {
			return
		}
		if e.GetTag(directiveSrcKey) != src {
			return
		}
		e.SetTag(directiveKey, min(level+1, 3))
	}, 3*60)
}

func (c *char) a4() {
	if c.Base.Ascension < 4 {
		return
	}
	// Resistances are not implemented
}
