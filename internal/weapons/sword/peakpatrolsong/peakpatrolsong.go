package peakpatrolsong

import (
	"fmt"

	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/event"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/modifier"
)

func init() {
	core.RegisterWeaponFunc(keys.PeakPatrolSong, NewWeapon)
}

type Weapon struct {
	Index  int
	stacks int
}

const (
	buffKey     = "peakpatrolsong-buff"
	buffDur     = 6 * 60
	teamBuffKey = "peakpatrolsong-team-buff"
	teamBuffDur = 15 * 60
	icdKey      = "peakpatrolsong-buff-icd"
	icdDur      = 0.1 * 60
)

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

func NewWeapon(c *core.Core, char *character.CharWrapper, p info.WeaponProfile) (info.Weapon, error) {
	w := &Weapon{}
	r := float64(p.Refine)

	d := make([]float64, attributes.EndStatType)
	b := make([]float64, attributes.EndStatType)

	def := 0.045 + 0.015*r
	bonusMulti := 0.09 + 0.03*r
	maxBonus := 0.27 + 0.09*r

	c.Events.Subscribe(event.OnEnemyDamage, func(args ...interface{}) bool {
		atk := args[1].(*combat.AttackEvent)
		if atk.Info.ActorIndex != char.Index {
			return false
		}
		if atk.Info.AttackTag != attacks.AttackTagNormal && atk.Info.AttackTag != attacks.AttackTagPlunge {
			return false
		}
		if char.StatusIsActive(icdKey) {
			return false
		}

		if !char.StatModIsActive(buffKey) {
			w.stacks = 0
		}
		if w.stacks < 2 {
			w.stacks++
		}

		d[attributes.DEFP] = def * float64(w.stacks)
		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBaseWithHitlag(buffKey, buffDur),
			AffectedStat: attributes.DEFP,
			Amount: func() ([]float64, bool) {
				return d, true
			},
		})

		if w.stacks == 2 {
			bonus := bonusMulti * char.TotalDef() / 1000.0
			bonus = min(bonus, maxBonus)
			for i := attributes.PyroP; i <= attributes.DendroP; i++ {
				b[i] = bonus
			}
			for _, this := range c.Player.Chars() {
				this.AddStatMod(character.StatMod{
					Base: modifier.NewBaseWithHitlag(teamBuffKey, teamBuffDur),
					Amount: func() ([]float64, bool) {
						return b, true
					},
				})
			}
		}

		char.AddStatus(icdKey, icdDur, true)
		return false
	}, fmt.Sprintf("peakpatrolsong-hit-%v", char.Base.Key.String()))

	return w, nil
}
