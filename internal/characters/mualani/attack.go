package mualani

import (
	"fmt"
	"slices"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

const normalHitNum = 3
const stackDelayAfterBite = 15

var (
	attackFrames   [][]int
	attackHitmarks = []int{19, 16, 32}
	attackHitboxes = []float64{1.0, 1.0, 1.5}

	sharkBiteFrames      [][]int
	sharkBiteHitmarks    = []int{2, 2, 2, 38}
	sharkBiteHitboxes    = 1.0
	sharkMissileHitboxes = 5.0
)

func init() {
	attackFrames = make([][]int, normalHitNum)

	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0], 36)
	attackFrames[0][action.ActionAttack] = 29
	attackFrames[0][action.ActionCharge] = 20

	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1], 33)
	attackFrames[1][action.ActionAttack] = 31
	attackFrames[1][action.ActionCharge] = 22

	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2], 62)
	attackFrames[2][action.ActionWalk] = 61
	attackFrames[2][action.ActionCharge] = 51

	sharkBiteFrames = make([][]int, 4)

	sharkBiteFrames[0] = frames.InitAbilSlice(40)

	sharkBiteFrames[1] = frames.InitAbilSlice(40)

	sharkBiteFrames[2] = frames.InitAbilSlice(40)

	sharkBiteFrames[3] = frames.InitAbilSlice(81)
}

func (c *char) Attack(p map[string]int) (action.Info, error) {
	if c.nightsoulPoints > 0 {
		return c.sharkBite(p), nil
	}

	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       fmt.Sprintf("Normal %v", c.NormalCounter),
		AttackTag:  attacks.AttackTagNormal,
		ICDTag:     attacks.ICDTagNormalAttack,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeDefault,
		Element:    attributes.Hydro,
		Durability: 25,
		Mult:       attack[c.NormalCounter][c.TalentLvlAttack()],
	}

	c.Core.QueueAttack(
		ai,
		combat.NewCircleHitOnTarget(
			c.Core.Combat.PrimaryTarget(),
			nil,
			attackHitboxes[c.NormalCounter],
		),
		attackHitmarks[c.NormalCounter],
		attackHitmarks[c.NormalCounter],
	)

	defer c.AdvanceNormalIndex()

	return action.Info{
		Frames:          frames.NewAttackFunc(c.Character, attackFrames),
		AnimationLength: attackFrames[c.NormalCounter][action.InvalidAction],
		CanQueueAfter:   attackFrames[c.NormalCounter][action.ActionSwap],
		State:           action.NormalAttackState,
	}, nil
}

func (c *char) sharkBite(p map[string]int) action.Info {
	c.NormalCounter = 0
	c.momentumSrc = c.Core.F
	momentumStacks := c.momentumStacks
	c.momentumStacks = 0

	nextMomentumFrame := max(c.lastStackFrame+0.7*60-c.Core.F, sharkBiteFrames[momentumStacks][action.ActionSwap]+stackDelayAfterBite)
	c.QueueCharTask(c.momentumStackGain(c.momentumSrc), nextMomentumFrame)
	c.QueueCharTask(func() {
		mult := bite[c.TalentLvlSkill()] + momentumBonus[c.TalentLvlSkill()]*float64(momentumStacks) + c.c1()
		ai := combat.AttackInfo{
			ActorIndex: c.Index,
			Abil:       fmt.Sprintf("Sharky's Bite (%v momentum)", momentumStacks),
			AttackTag:  attacks.AttackTagNormal,
			ICDTag:     attacks.ICDTagNone,
			ICDGroup:   attacks.ICDGroupDefault,
			StrikeType: attacks.StrikeTypeDefault,
			Element:    attributes.Hydro,
			Durability: 25,
		}

		if momentumStacks >= 3 {
			ai.Abil = "Sharky's Surging Bite"
			mult += surgingBite[c.TalentLvlSkill()]
		}

		ap := combat.NewCircleHitOnTarget(c.Core.Combat.PrimaryTarget(), nil, sharkBiteHitboxes)

		enemiesBite := c.Core.Combat.EnemiesWithinArea(
			ap,
			nil,
		)

		markOfPrey := false
		for _, e := range enemiesBite {
			if e.StatusIsActive(markedAsPreyKey) {
				markOfPrey = true
			}
		}
		totalEnemies := len(enemiesBite)

		var enemiesMissile []combat.Enemy
		if markOfPrey {
			ap := combat.NewCircleHitOnTarget(c.Core.Combat.PrimaryTarget(), nil, sharkMissileHitboxes)
			enemiesMissile = c.Core.Combat.EnemiesWithinArea(
				ap,
				func(e combat.Enemy) bool { return !slices.Contains(enemiesBite, e) },
			)
			totalEnemies += len(enemiesMissile)
		}

		mult *= max(1.14-0.14*float64(totalEnemies), 0.72)
		ai.FlatDmg = mult * c.MaxHP()
		c.Core.QueueAttack(
			ai,
			combat.NewCircleHitOnTarget(
				c.Core.Combat.PrimaryTarget(),
				nil,
				sharkBiteHitboxes,
			),
			0,
			0,
			c.particleCB,
			c.a1cb(),
		)

		for _, e := range enemiesMissile {
			c.Core.QueueAttack(
				ai,
				combat.NewSingleTargetHit(e.Key()),
				0,
				20,
			)
		}

		c.SetCDWithDelay(action.ActionAttack, 1.8*60, 0)
	}, sharkBiteHitmarks[momentumStacks])

	return action.Info{
		Frames:          frames.NewAbilFunc(sharkBiteFrames[momentumStacks]),
		AnimationLength: sharkBiteFrames[momentumStacks][action.InvalidAction],
		CanQueueAfter:   sharkBiteFrames[momentumStacks][action.ActionSwap],
		State:           action.NormalAttackState,
	}
}
