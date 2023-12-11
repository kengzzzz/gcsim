package songofdayspast

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

const (
	yearningKey        = "yearning-effect"
	wavesOfDaysPastKey = "waves-of-days-past"
)

func init() {
	core.RegisterSetFunc(keys.SongOfDaysPast, NewSet)
}

type Set struct {
	healStacks float64
	core       *core.Core
	Index      int
}

func (s *Set) SetIndex(idx int) { s.Index = idx }
func (s *Set) Init() error {
	s.wavesOfDaysPastEffect()
	return nil
}

func NewSet(c *core.Core, char *character.CharWrapper, count int, param map[string]int) (info.Set, error) {
	s := Set{
		core: c,
	}

	if count >= 2 {
		m := make([]float64, attributes.EndStatType)
		m[attributes.Heal] = 0.15
		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBase("sodp-2pc", -1),
			AffectedStat: attributes.Heal,
			Amount: func() ([]float64, bool) {
				return m, true
			},
		})
	}

	if count >= 4 {
		c.Events.Subscribe(event.OnHeal, func(args ...interface{}) bool {
			if s.core.Status.Duration(wavesOfDaysPastKey) > 0 {
				return false
			}

			src := args[0].(*player.HealInfo)
			healAmt := args[2].(float64)

			if src.Caller != char.Index {
				return false
			}

			s.healStacks += healAmt
			if s.healStacks >= 15000 {
				s.healStacks = 15000
			}

			if s.core.Status.Duration(yearningKey) == 0 {
				s.core.Status.Add(yearningKey, 6*60)

				c.Tasks.Add(func() {
					s.core.Status.Add(wavesOfDaysPastKey, 10*60)
					s.core.Flags.Custom[wavesOfDaysPastKey] = 5
				}, 6*60)
			}

			return false
		}, fmt.Sprintf("sodp-4pc-heal-accumulation-%v", char.Base.Key.String()))
	}

	return &s, nil
}

func (s *Set) wavesOfDaysPastEffect() {
	s.core.Events.Subscribe(event.OnEnemyHit, func(args ...interface{}) bool {
		atk := args[1].(*combat.AttackEvent)

		switch atk.Info.AttackTag {
		case attacks.AttackTagElementalBurst:
		case attacks.AttackTagElementalArt:
		case attacks.AttackTagElementalArtHold:
		case attacks.AttackTagNormal:
		case attacks.AttackTagExtra:
		case attacks.AttackTagPlunge:
		default:
			return false
		}

		char := s.core.Player.ByIndex(atk.Info.ActorIndex)

		if s.core.Status.Duration(wavesOfDaysPastKey) == 0 {
			return false
		}

		if char.Index != s.core.Player.Active() {
			return false
		}

		if s.core.Flags.Custom[wavesOfDaysPastKey] > 0 {
			s.core.Flags.Custom[wavesOfDaysPastKey]--
			amt := s.healStacks * 0.08
			atk.Info.FlatDmg += amt
		}

		if s.core.Flags.Custom[wavesOfDaysPastKey] == 0 {
			s.healStacks = 0
			s.core.Status.Delete(wavesOfDaysPastKey)
		}

		return false
	}, "waves-of-days-past-hook")
}