package arlecchino

import (
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

const c2IcdKey = "arlecchino-c2-icd"

func (c *char) c2() {
	c.initialDirectiveLevel = 1
	if c.Base.Cons >= 2 && c.Base.Ascension >= 1 {
		c.initialDirectiveLevel = 2
	}
}

func (c *char) c2OnAbsorb() {
	// Check is redundant? Can't reach level 3 directives without A1
	if c.Base.Cons < 2 || c.Base.Ascension < 1 {
		return
	}

	if c.StatusIsActive(c2IcdKey) {
		return
	}

	c.AddStatus(c2IcdKey, 10*60, true)
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Balemoon Bloodfire (C2)",
		AttackTag:  attacks.AttackTagNone,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeDefault,
		Element:    attributes.Pyro,
		Durability: 25,
		Mult:       9.00,
	}
	c.Core.QueueAttack(
		ai,
		combat.NewCircleHit(
			c.Core.Combat.Player(),
			c.Core.Combat.PrimaryTarget(),
			nil,
			1.2,
		),
		4,
		4,
	)
}
