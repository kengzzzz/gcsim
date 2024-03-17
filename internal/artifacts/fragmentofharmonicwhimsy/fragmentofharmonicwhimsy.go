package fragmentofharmonicwhimsy

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func init() {
	core.RegisterSetFunc(keys.FragmentOfHarmonicWhimsy, NewSet)
}

const buffKey = "fragmentofharmonicwhimsy-4pc"

type Set struct {
	Index  int
	core   *core.Core
	char   *character.CharWrapper
	stacks int
	buff   []float64
}

func (s *Set) SetIndex(idx int) { s.Index = idx }
func (s *Set) Init() error      { return nil }

func NewSet(c *core.Core, char *character.CharWrapper, count int, param map[string]int) (info.Set, error) {
	s := Set{
		core: c,
		char: char,
	}

	if count >= 2 {
		m := make([]float64, attributes.EndStatType)
		m[attributes.ATKP] = 0.18
		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBase("fragmentofharmonicwhimsy-2pc", -1),
			AffectedStat: attributes.ATKP,
			Amount: func() ([]float64, bool) {
				return m, true
			},
		})
	}

	if count >= 4 {
		c.Events.Subscribe(event.OnHPDebt, s.OnHPDept(), fmt.Sprintf("fragmentofharmonicwhimsy-4pc-%v", char.Base.Key.String()))
	}

	return &s, nil
}

func (s *Set) OnHPDept() func(args ...interface{}) bool {
	return func(args ...interface{}) bool {
		target := args[0].(int)
		if target != s.char.Index {
			return false
		}
		if !s.char.StatModIsActive(buffKey) {
			s.stacks = 0
		}
		if s.stacks < 3 {
			s.stacks++
		}

		s.char.AddStatMod(character.StatMod{
			Base:         modifier.NewBaseWithHitlag(buffKey, 360),
			AffectedStat: attributes.DmgP,
			Amount: func() ([]float64, bool) {
				s.buff[attributes.DmgP] = 0.18 * float64(s.stacks)
				return s.buff, true
			},
		})

		return false
	}
}
