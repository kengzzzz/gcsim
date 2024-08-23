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

var (
	attackFrames   [][]int
	attackHitmarks = []int{19, 16, 32}
	attackHitboxes = []float64{1.0, 1.0, 1.5}

	sharkBiteFrames      [][]int
	sharkBiteHitmarks    = 19
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

	sharkBiteFrames = make([][]int, 1)

	sharkBiteFrames[0] = frames.InitNormalCancelSlice(sharkBiteHitmarks, 36)
	sharkBiteFrames[0][action.ActionAttack] = 61
	sharkBiteFrames[0][action.ActionCharge] = 51
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
	c.SetCDWithDelay(action.ActionAttack, 1.8*60, 0)
	c.QueueCharTask(func() {
		mult := bite[c.TalentLvlSkill()] + momentumBonus[c.TalentLvlSkill()]*float64(c.momentumStacks) + c.c1()
		ai := combat.AttackInfo{
			ActorIndex: c.Index,
			Abil:       fmt.Sprintf("Sharky's Bite (%v momentum)", c.momentumStacks),
			AttackTag:  attacks.AttackTagNormal,
			ICDTag:     attacks.ICDTagNone,
			ICDGroup:   attacks.ICDGroupDefault,
			StrikeType: attacks.StrikeTypeDefault,
			Element:    attributes.Hydro,
			Durability: 25,
		}

		if c.momentumStacks >= 3 {
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
	}, sharkBiteHitmarks)

	return action.Info{
		Frames:          frames.NewAttackFunc(c.Character, sharkBiteFrames),
		AnimationLength: sharkBiteFrames[0][action.InvalidAction],
		CanQueueAfter:   sharkBiteFrames[0][action.ActionSwap],
		State:           action.NormalAttackState,
	}
}
