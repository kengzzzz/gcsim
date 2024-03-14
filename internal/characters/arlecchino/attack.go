package arlecchino

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/geometry"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

var (
	attackFrames          [][]int
	attackHitmarks        = [][]int{{17}, {15}, {15}, {14, 31}, {16}, {39}}
	attackHitlagHaltFrame = [][]float64{{0.01}, {0.01}, {0.01}, {0.02, 0.02}, {0.02}, {0.04}}
	attackDefHalt         = [][]bool{{true}, {true}, {true}, {false, true}, {true}, {true}}
	attackStrikeTypes     = [][]attacks.StrikeType{
		{attacks.StrikeTypeSlash},
		{attacks.StrikeTypeSlash},
		{attacks.StrikeTypeSlash},
		{attacks.StrikeTypeSlash, attacks.StrikeTypeSlash},
		{attacks.StrikeTypeSpear},
		{attacks.StrikeTypeSlash},
	}
	attackHitboxes = [][][][]float64{
		{
			{{1.9, 3}},     // box
			{{2.6}},        // fan
			{{1.9, 4}},     // box
			{{2.8}, {2.8}}, // circle, circle
			{{2.5}},        // circle
			{{2.5}},        //TODO: circle?
		},
		{
			{{1.9, 3}},     // box
			{{2.6}},        // fan
			{{1.9, 4}},     // box
			{{2.8}, {2.8}}, // circle, circle
			{{2.5}},        // circle
			{{2.5}},        //TODO: circle?
		},
	}
	attackOffsets = [][][]float64{
		{{0, -0.15}},
		{{0, 0.5}},
		{{0, -1.2}},
		{{-0.5, 0.7}, {-0.5, 0.7}},
		{{0, 2.4}},
		{{0, 2.5}},
	}
	attackFanAngles = [][]float64{{360}, {300}, {360}, {360, 360}, {360}, {360}}
)

const naBuffKey = "in-praise-of-shadows"
const normalHitNum = 6

func init() {
	attackFrames = make([][]int, normalHitNum)

	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0][0], 26)
	attackFrames[0][action.ActionAttack] = 25

	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1][0], 27)
	attackFrames[1][action.ActionAttack] = 22

	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2][0], 38)
	attackFrames[2][action.ActionAttack] = 26

	attackFrames[3] = frames.InitNormalCancelSlice(attackHitmarks[3][1], 42)
	attackFrames[3][action.ActionAttack] = 39

	attackFrames[4] = frames.InitNormalCancelSlice(attackHitmarks[4][0], 30)
	attackFrames[4][action.ActionAttack] = 24

	attackFrames[5] = frames.InitNormalCancelSlice(attackHitmarks[5][0], 79)
	attackFrames[5][action.ActionCharge] = 500 //TODO: this action is illegal; need better way to handle it
}

func (c *char) naBuff() {
	c.Core.Events.Subscribe(event.OnHPDebt, func(args ...interface{}) bool {
		target := args[0].(int)
		if target != c.Index {
			return false
		}
		if c.CurrentHPDebt() >= c.MaxHP()*0.3 {
			// can't use negative duration or else `if .arlecchino.status.in-praise-of-shadows` won't work
			c.AddStatus(naBuffKey, 999999, false)
		} else {
			c.DeleteStatus(naBuffKey)
		}
		return false
	}, "arlechinno-bol-hook")

	m := make([]float64, attributes.EndStatType)
	m[attributes.PyroP] = 0.4
	c.AddStatMod(character.StatMod{
		Base:         modifier.NewBase("arlecchino-na", -1),
		AffectedStat: attributes.PyroP,
		Amount: func() ([]float64, bool) {
			if c.StatusIsActive(naBuffKey) {
				return m, true
			}
			return nil, false
		},
	})
}

// Normal attack damage queue generator
// relatively standard with no major differences versus other characters
func (c *char) Attack(p map[string]int) (action.Info, error) {
	for i, mult := range attack[c.NormalCounter] {
		ai := combat.AttackInfo{
			ActorIndex:         c.Index,
			Abil:               fmt.Sprintf("Normal %v", c.NormalCounter),
			AttackTag:          attacks.AttackTagNormal,
			ICDTag:             attacks.ICDTagNormalAttack,
			ICDGroup:           attacks.ICDGroupDefault,
			StrikeType:         attackStrikeTypes[c.NormalCounter][i],
			Element:            attributes.Physical,
			Durability:         25,
			Mult:               mult[c.TalentLvlAttack()],
			HitlagFactor:       0.01,
			HitlagHaltFrames:   attackHitlagHaltFrame[c.NormalCounter][i] * 60,
			CanBeDefenseHalted: attackDefHalt[c.NormalCounter][i],
		}

		naIndex := 0

		if c.StatusIsActive(naBuffKey) {
			naIndex = 1
			ai.Element = attributes.Pyro
			ai.IgnoreInfusion = true
			ai.FlatDmg += blooddebt[c.TalentLvlAttack()] * c.CurrentHPDebt() / c.MaxHP() * c.getTotalAtk()
			c.QueueCharTask(func() {
				c.ModifyHPDebtByAmount(-0.055 * c.CurrentHPDebt())
			}, attackHitmarks[c.NormalCounter][i]+1)
		}

		var ap combat.AttackPattern
		if len(attackHitboxes[naIndex][c.NormalCounter][i]) == 1 { // circle or fan
			ap = combat.NewCircleHitOnTargetFanAngle(
				c.Core.Combat.Player(),
				geometry.Point{X: attackOffsets[c.NormalCounter][i][0], Y: attackOffsets[c.NormalCounter][i][1]},
				attackHitboxes[naIndex][c.NormalCounter][i][0],
				attackFanAngles[c.NormalCounter][i],
			)
		} else { // box
			ap = combat.NewBoxHitOnTarget(
				c.Core.Combat.Player(),
				geometry.Point{X: attackOffsets[c.NormalCounter][i][0], Y: attackOffsets[c.NormalCounter][i][1]},
				attackHitboxes[naIndex][c.NormalCounter][i][0],
				attackHitboxes[naIndex][c.NormalCounter][i][1],
			)
		}

		c.QueueCharTask(func() {
			c.Core.QueueAttack(ai, ap, 0, 0)
		}, attackHitmarks[c.NormalCounter][i])
	}

	defer c.AdvanceNormalIndex()

	return action.Info{
		Frames:          frames.NewAttackFunc(c.Character, attackFrames),
		AnimationLength: attackFrames[c.NormalCounter][action.InvalidAction],
		CanQueueAfter:   attackHitmarks[c.NormalCounter][len(attackHitmarks[c.NormalCounter])-1],
		State:           action.NormalAttackState,
	}, nil
}

func (c *char) getTotalAtk() float64 {
	stats, _ := c.Stats()
	return c.Base.Atk*(1+stats[attributes.ATKP]) + stats[attributes.ATK]
}
