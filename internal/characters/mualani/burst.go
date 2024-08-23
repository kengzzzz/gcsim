package mualani

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
)

const burstHitmarks = 110

var (
	burstFrames []int
)

func init() {
	burstFrames = frames.InitAbilSlice(146)
	burstFrames[action.ActionAttack] = 113
	burstFrames[action.ActionCharge] = 124
	burstFrames[action.ActionDash] = 111
	burstFrames[action.ActionJump] = 113
	burstFrames[action.ActionSwap] = 145
}

func (c *char) Burst(p map[string]int) (action.Info, error) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Stormburst Shot",
		AttackTag:  attacks.AttackTagElementalBurst,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeDefault,
		Element:    attributes.Hydro,
		Durability: 25,
		Mult:       0,
		FlatDmg:    burst[c.TalentLvlBurst()] * c.MaxHP(),
	}
	skillArea := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), geometry.Point{Y: 6}, 6.5)

	c.QueueCharTask(func() { c.ResetActionCooldown(action.ActionSkill) }, 107)

	c.QueueCharTask(func() {
		// the A4 stacks can change during the burst
		ai.FlatDmg += c.a4amount()
		c.Core.QueueAttack(ai, skillArea, 0, 0)
	}, burstHitmarks)

	c.SetCDWithDelay(action.ActionBurst, 15*60, 0)
	c.ConsumeEnergy(12)

	return action.Info{
		Frames:          frames.NewAbilFunc(burstFrames),
		AnimationLength: burstFrames[action.InvalidAction],
		CanQueueAfter:   burstFrames[action.ActionSwap], // earliest cancel
		State:           action.BurstState,
	}, nil
}
