package mualani

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
)

func init() {
	core.RegisterCharFunc(keys.Mualani, NewChar)
}

const nightsoulBurstKey = "nightsoul-burst"
const nightsoulBurstCD = 18 * 60 // only mualani right now

type char struct {
	*tmpl.Character
	nightsoulPoints int
	nightsoulSrc    int
	momentumStacks  int
	a4Stacks        int
	c1Done          bool

	a1Count int
}

func NewChar(s *core.Core, w *character.CharWrapper, _ info.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 60
	c.NormalHitNum = normalHitNum
	c.SkillCon = 3
	c.BurstCon = 5
	c.HasArkhe = false

	c.nightsoulPoints = 0
	w.Character = &c

	return nil
}

func (c *char) Init() error {
	if c.Base.Cons >= 1 {
		c.c1()
	}

	if c.Base.Cons >= 2 {
		c.c2()
	}

	if c.Base.Cons >= 4 {
		c.c4()
	}
	c.SetNumCharges(action.ActionAttack, 1)
	return nil
}

func (c *char) NightsoulBurst() {
	c.Core.Events.Subscribe(event.OnEnemyDamage, func(args ...interface{}) bool {
		if c.StatusIsActive(nightsoulBurstKey) {
			return false
		}

		atk := args[1].(*combat.AttackEvent)
		if atk.Info.Element != attributes.Physical {
			c.AddStatus(nightsoulBurstKey, nightsoulBurstCD, false)
		}

		c.a4()
		return false
	}, "nightsoul-burst")
}

func (c *char) ActionReady(a action.Action, p map[string]int) (bool, action.Failure) {
	if a == action.ActionAttack && c.nightsoulPoints > 0 {
		if c.AvailableCDCharge[a] <= 0 {
			// TODO: Implement AttackCD warning
			return false, action.CharacterDeceased
		}
	}

	return c.Character.ActionReady(a, p)
}

func (c *char) Condition(fields []string) (any, error) {
	switch fields[0] {
	case "momentum":
		return c.momentumStacks, nil
	case "nightsoulpoints":
		return c.nightsoulPoints, nil
	default:
		return c.Character.Condition(fields)
	}
}
