package fragmentofharmonicwhimsy

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func init() {
	core.RegisterSetFunc(keys.FragmentOfHarmonicWhimsy, NewSet)
}

type Set struct {
	stacks int
	core   *core.Core
	char   *character.CharWrapper
	buff   []float64
	Index  int
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
			Base:         modifier.NewBase("harmonicwhimsy-2pc", -1),
			AffectedStat: attributes.ATKP,
			Amount: func() ([]float64, bool) {
				return m, true
			},
		})
	}

	// When the value of a Bond of Life increases or decreases, t
	// his character deals 18% increased DMG for 6s. Max 3 stacks.
	if count < 4 {
		return &s, nil
	}

	m := make([]float64, attributes.EndStatType)
	m[attributes.DmgP] = 0.18
	c.Events.Subscribe(event.OnHPDebt, func(args ...interface{}) bool {
		amount := args[1].(float64)
		if amount != 0 {
			char.AddStatMod(character.StatMod{
				Base: modifier.NewBaseWithHitlag(fmt.Sprintf("harmonic-whimny-%v-stack", s.stacks+1), 6*60),
				Amount: func() ([]float64, bool) {
					return m, true
				},
			})
			s.stacks = (s.stacks + 1) % 3
			c.Log.NewEvent("Harmonic Whimsy stack gained", glog.LogArtifactEvent, char.Index).Write("stacks", s.stacks)
		}
		return false
	}, "Checking for HPDebt")

	return &s, nil
}
