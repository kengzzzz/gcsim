package mualani

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/targets"
)

var skillFrames []int

const (
	particleICD    = 9999 * 60
	particleICDKey = "mualani-particle-icd"

	nightblessing = "nightblessing"

	// momentumStackICDKey = "momentum-icd"
	markedAsPreyKey = "markedAsPrey"
	markedAsPreyDur = 10 * 60
	skillDelay      = 21
)

func init() {
	skillFrames = frames.InitAbilSlice(42)
	skillFrames[action.ActionCharge] = 21
	skillFrames[action.ActionBurst] = 30
	skillFrames[action.ActionDash] = 29
	skillFrames[action.ActionJump] = 32
	skillFrames[action.ActionWalk] = 41
	skillFrames[action.ActionSwap] = 29
	// skill -> skill is unknown
}

func (c *char) reduceNightsoulPoints(val int) {
	c.nightsoulPoints = max(c.nightsoulPoints-val, 0)
	if c.nightsoulPoints <= 0 {
		c.cancelNightsoul()
	}
}

func (c *char) cancelNightsoul() {
	c.DeleteStatus(nightblessing)
	c.SetCDWithDelay(action.ActionSkill, 6*60, 0)
	c.ResetActionCooldown(action.ActionAttack)
	c.momentumStacks = 0
	c.momentumSrc = -1
	c.nightsoulPoints = 0
	c.nightsoulSrc = -1
}

func (c *char) nightsoulPointReduceFunc(src int) func() {
	return func() {
		if c.nightsoulSrc != src {
			return
		}

		if c.nightsoulPoints <= 0 {
			return
		}

		c.reduceNightsoulPoints(1)

		// reduce 1 point per 6f
		c.QueueCharTask(c.nightsoulPointReduceFunc(src), 6)
	}
}

func (c *char) momentumStackGain(src int) func() {
	return func() {
		if c.momentumSrc != src {
			return
		}

		if c.nightsoulPoints <= 0 {
			return
		}

		if c.Core.Player.Active() != c.Index {
			c.QueueCharTask(c.momentumStackGain(src), 0.7*60)
			return
		}

		ap := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 1.0)
		enemies := c.Core.Combat.Enemies()
		enemiesCollided := 0
		for _, e := range enemies {
			enemy, ok := e.(combat.Enemy)
			if !ok {
				continue
			}

			willLand, _ := e.AttackWillLand(ap)
			if willLand {
				enemy.AddStatus(markedAsPreyKey, markedAsPreyDur, true)
				enemiesCollided++
			}
		}
		c.lastStackFrame = c.Core.F
		c.momentumStacks = min(c.momentumStacks+enemiesCollided, 3.0)
		c.QueueCharTask(c.momentumStackGain(src), 0.7*60)
	}
}

func (c *char) Skill(p map[string]int) (action.Info, error) {
	if c.nightsoulPoints > 0 {
		c.cancelNightsoul()
		return action.Info{
			Frames:          func(_ action.Action) int { return 1 },
			AnimationLength: 1,
			CanQueueAfter:   1, // earliest cancel
			State:           action.SkillState,
		}, nil
	}

	c.QueueCharTask(func() {
		// set to high value so that `if .mualani.status.nightblessing` will work
		c.AddStatus(nightblessing, 99999, true)
		c.DeleteStatus(particleICDKey)
		c.a1Count = 0
		c.c1Done = false
		c.c2()
		c.nightsoulPoints = 60
		c.nightsoulSrc = c.Core.F
		c.QueueCharTask(c.nightsoulPointReduceFunc(c.nightsoulSrc), 6)
		c.momentumSrc = c.Core.F
		c.QueueCharTask(c.momentumStackGain(c.momentumSrc), 0)
	}, skillDelay)

	return action.Info{
		Frames:          frames.NewAbilFunc(skillFrames),
		AnimationLength: skillFrames[action.InvalidAction],
		CanQueueAfter:   skillFrames[action.ActionCharge], // earliest cancel
		State:           action.SkillState,
	}, nil
}

func (c *char) particleCB(a combat.AttackCB) {
	if a.Target.Type() != targets.TargettableEnemy {
		return
	}
	if c.StatusIsActive(particleICDKey) {
		return
	}
	c.AddStatus(particleICDKey, particleICD, true)

	count := 4.0
	if c.Core.Rand.Float64() < .5 {
		count = 5
	}
	c.Core.QueueParticle(c.Base.Key.String(), count, attributes.Hydro, c.ParticleDelay)
}
