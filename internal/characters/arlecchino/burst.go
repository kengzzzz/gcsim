package arlecchino

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

const burstHitmarks = 130

var (
	burstFrames []int
)

func init() {
	burstFrames = frames.InitAbilSlice(146)
}

func (c *char) Burst(p map[string]int) (action.Info, error) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Balemoon Rising",
		AttackTag:  attacks.AttackTagElementalBurst,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeDefault,
		Element:    attributes.Pyro,
		Durability: 25,
		Mult:       burst[c.TalentLvlSkill()],
	}
	skillArea := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 10)
	c.Core.QueueAttack(ai, skillArea, burstHitmarks, burstHitmarks)

	c.QueueCharTask(func() {
		c.ModifyHPDebtByRatio(0.15)
	}, burstHitmarks)
	// add cooldown to sim
	c.SetCDWithDelay(action.ActionBurst, 15*60, 0)
	// use up energy
	c.ConsumeEnergy(8)

	return action.Info{
		Frames:          frames.NewAbilFunc(burstFrames),
		AnimationLength: burstFrames[action.InvalidAction],
		CanQueueAfter:   burstFrames[action.ActionSwap], // earliest cancel
		State:           action.BurstState,
	}, nil
}
