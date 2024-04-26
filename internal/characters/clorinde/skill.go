package clorinde

import (
	"fmt"
	"math"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/info"
)

var (
	skillFrames            []int
	skillDashNoBOLFrames   []int
	skillDashLowBOLFrames  []int
	skillDashFullBOLFrames []int
)

const (
	skillStateKey = "clorinde-night-watch"
	tolerance     = 0.0000001

	// TODO: all hit marks
	skillDashNoBOLHitmark   = 24
	skillDashLowBOLHitmark  = 24
	skillDashFullBOLHitmark = 24
)

func init() {
	// TODO: all these frames are gusses
	skillFrames = frames.InitAbilSlice(32)
	skillFrames[action.ActionSkill] = 27

	skillDashNoBOLFrames = frames.InitAbilSlice(28)
	skillDashLowBOLFrames = frames.InitAbilSlice(28)
	skillDashFullBOLFrames = frames.InitAbilSlice(28)
}

func (c *char) Skill(p map[string]int) (action.Info, error) {
	// first press activates skill state
	// sequential presses pew pew stuff
	if c.StatusIsActive(skillStateKey) {
		return c.skillDash(p)
	}

	c.AddStatus(skillStateKey, 60*int(skillStateDuration[0]), true)

	return action.Info{
		Frames:          frames.NewAbilFunc(skillFrames),
		AnimationLength: skillFrames[action.InvalidAction],
		CanQueueAfter:   skillFrames[action.ActionSkill],
		State:           action.SkillState,
	}, nil
}

func (c *char) skillDash(p map[string]int) (action.Info, error) {
	// depending on BOL lvl it does either 1 hit or 3 hit
	ratio := c.currentHPDebtRatio()
	switch {
	case math.Abs(ratio-1) < tolerance:
		return c.skillDashFullBOL(p)
	case math.Abs(ratio) < tolerance:
		return c.skillDashNoBOL(p)
	default:
		return c.skillDashRegular(p)
	}
}

func (c *char) currentHPDebtRatio() float64 {
	return c.CurrentHPDebt() / c.MaxHP()
}

func (c *char) gainBOLOnAttack() {
	c.ModifyHPDebtByRatio(skillBOLGain[c.TalentLvlSkill()])
}

func (c *char) skillDashNoBOL(_ map[string]int) (action.Info, error) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Skill Dash (No BOL)",
		AttackTag:  attacks.AttackTagNormal,
		ICDTag:     attacks.ICDTagNormalAttack,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeSlash,
		Element:    attributes.Electro,
		Durability: 25,
		Mult:       skillLungeNoBOL[c.TalentLvlSkill()],
	}
	// TODO: what's the size of this??
	ap := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 0.6)
	// TODO: assume no snapshotting on this
	c.Core.QueueAttack(ai, ap, skillDashNoBOLHitmark, skillDashNoBOLHitmark)
	// TODO: no idea if this counts as a normal attack state or not. pretend it does for now
	return action.Info{
		Frames:          frames.NewAbilFunc(skillDashNoBOLFrames),
		AnimationLength: skillDashNoBOLFrames[action.InvalidAction],
		CanQueueAfter:   skillDashNoBOLFrames[action.InvalidAction], //TODO: fastest cancel?
		State:           action.NormalAttackState,
	}, nil
}

func (c *char) skillDashFullBOL(_ map[string]int) (action.Info, error) {
	for i := 0; i < 3; i++ {
		ai := combat.AttackInfo{
			ActorIndex: c.Index,
			Abil:       fmt.Sprintf("Skill Dash (Full BOL): %v", i+1),
			AttackTag:  attacks.AttackTagNormal,
			ICDTag:     attacks.ICDTagNormalAttack,
			ICDGroup:   attacks.ICDGroupDefault,
			StrikeType: attacks.StrikeTypeSlash,
			Element:    attributes.Electro,
			Durability: 25,
			Mult:       skillLungeFullBOL[c.TalentLvlSkill()],
		}
		// TODO: what's the size of this??
		ap := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 0.8)
		// TODO: assume no snapshotting on this
		c.Core.QueueAttack(ai, ap, skillDashFullBOLHitmark, skillDashFullBOLHitmark)
	}

	// TODO: timing on this heal?
	c.skillHeal(skillLungeFullBOLHeal[0])

	// TODO: no idea if this counts as a normal attack state or not. pretend it does for now
	return action.Info{
		Frames:          frames.NewAbilFunc(skillDashFullBOLFrames),
		AnimationLength: skillDashFullBOLFrames[action.InvalidAction],
		CanQueueAfter:   skillDashFullBOLFrames[action.InvalidAction], //TODO: fastest cancel?
		State:           action.NormalAttackState,
	}, nil
}

func (c *char) skillDashRegular(_ map[string]int) (action.Info, error) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Skill Dash (< 100% BOL)",
		AttackTag:  attacks.AttackTagNormal,
		ICDTag:     attacks.ICDTagNormalAttack,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeSlash,
		Element:    attributes.Electro,
		Durability: 25,
		Mult:       skillLungeLowBOL[c.TalentLvlSkill()],
	}
	// TODO: what's the size of this??
	ap := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 0.8)
	// TODO: assume no snapshotting on this
	c.Core.QueueAttack(ai, ap, skillDashLowBOLHitmark, skillDashLowBOLHitmark)

	// TODO: timing on this heal?
	c.skillHeal(skillLungeLowBOLHeal[0])

	// TODO: no idea if this counts as a normal attack state or not. pretend it does for now
	return action.Info{
		Frames:          frames.NewAbilFunc(skillDashLowBOLFrames),
		AnimationLength: skillDashLowBOLFrames[action.InvalidAction],
		CanQueueAfter:   skillDashLowBOLFrames[action.InvalidAction], //TODO: fastest cancel?
		State:           action.NormalAttackState,
	}, nil
}

func (c *char) skillHeal(bolMult float64) {
	amt := c.CurrentHPDebt() * bolMult
	c.Character.Heal(&info.HealInfo{
		Caller:  c.Index,
		Target:  c.Index,
		Message: "Clorinde Skill", // TODO: fix naming...
		Src:     amt,
		Bonus:   c.Stat(attributes.Heal), // TODO: confirms that it scales with healing %
	})
}
