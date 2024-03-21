package arlecchino

import (
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/targets"
)

const c2IcdKey = "arlecchino-c2-icd"
const c4IcdKey = "arlecchino-c4-icd"

func (c *char) c2() {
	c.initialDirectiveLevel = 1
	if c.Base.Cons >= 2 && c.Base.Ascension >= 1 {
		c.initialDirectiveLevel = 2
	}
}

func (c *char) c2OnAbsorbLevel3() {
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

func (c *char) c4() {
	c.bondOnBurst = 0.15
	if c.Base.Cons >= 4 && c.Base.Ascension >= 1 {
		c.bondOnBurst = 0.25
	}
}

func (c *char) c4cb(a combat.AttackCB) {
	if c.Base.Cons < 4 || c.Base.Ascension < 1 {
		return
	}

	if a.Target.Type() != targets.TargettableEnemy {
		return
	}
	level := a.Target.GetTag(directiveKey)

	if level == 0 {
		return
	}

	if level >= 3 {
		return
	}
	a.Target.SetTag(directiveKey, level+1)
	c.Core.Log.NewEvent("Directive upgraded (C4)", glog.LogCharacterEvent, c.Index).
		Write("new_level", level+1).
		Write("src", "c4")
}

func (c *char) c4OnAbsorb() {
	if c.Base.Cons < 4 {
		return
	}

	if c.StatusIsActive(c4IcdKey) {
		return
	}

	c.AddStatus(c4IcdKey, 10*60, true)
	c.ReduceActionCooldown(action.ActionBurst, 2*60)
	c.AddEnergy("arlecchino-c4", 15)
}
