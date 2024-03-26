package crimsonmoonssemblance

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
	core.RegisterWeaponFunc(keys.CrimsonMoonsSemblance, NewWeapon)
}

type Weapon struct {
	Index int
}

const (
	buffICD    = 840
	buffICDKey = "crimsonmoonssemblance-bol-icd"
)

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

func NewWeapon(c *core.Core, char *character.CharWrapper, p info.WeaponProfile) (info.Weapon, error) {
	w := &Weapon{}
	r := p.Refine

	hasBoL := float64(r)*0.04 + 0.08
	gteBoL := float64(r)*0.08 + 0.16

	char.AddAttackMod(character.AttackMod{
		Base: modifier.NewBase("crimsonmoonssemblance-buff", -1),
		Amount: func(atk *combat.AttackEvent, t combat.Target) ([]float64, bool) {
			buff := make([]float64, attributes.EndStatType)
			if char.CurrentHPDebt() > 0 {
				buff[attributes.DmgP] += hasBoL
			}
			if char.CurrentHPDebt() >= char.MaxHP()*0.2 {
				buff[attributes.DmgP] += gteBoL
			}
			return buff, true
		},
	})

	c.Events.Subscribe(event.OnEnemyDamage, func(args ...interface{}) bool {
		atk := args[1].(*combat.AttackEvent)

		if atk.Info.ActorIndex != char.Index {
			return false
		}
		if c.Player.Active() != char.Index {
			return false
		}
		if atk.Info.AttackTag != attacks.AttackTagExtra {
			return false
		}

		if char.StatusIsActive(buffICDKey) {
			return false
		}

		char.AddStatus(buffICDKey, buffICD, true)

		char.ModifyHPDebtByRatio(0.18)

		return false
	}, fmt.Sprintf("crimsonmoonssemblance-%v", char.Base.Key.String()))
	return w, nil
}
