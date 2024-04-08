package crimsonmoonssemblance

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
	core.RegisterWeaponFunc(keys.CrimsonMoonsSemblance, NewWeapon)
}

type Weapon struct {
	Index int
}

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

// Grants a Bond of Life equal to 18% of Max HP when a Charged Attack hits an opponent.
// This effect can be triggered up to once every 14s.
// In addition, when the equipping character has a Bond of Life,
// they gain a 12/16/20/24/28% DMG Bonus;
// if the value of the Bond of Life is greater than or equal to 30% of Max HP,
// then gain an additional 24/32/40/48/56% DMG Bonus.

func NewWeapon(c *core.Core, char *character.CharWrapper, p info.WeaponProfile) (info.Weapon, error) {
	w := &Weapon{}
	r := p.Refine

	const icdKey = "cms-icd"
	const bondKey = "cms-bond"
	hp := 0.25
	bondDMG := 0.08 + float64(r)*0.04
	bondDMG2 := 0.16 + float64(r)*0.08
	duration := 14 * 60
	maxhp := char.MaxHP()
	dmg := make([]float64, attributes.EndStatType)

	c.Events.Subscribe(event.OnChargeAttack, func(args ...interface{}) bool {
		if c.Player.Active() != char.Index {
			return false
		}
		if char.StatusIsActive(icdKey) {
			return false
		}
		char.AddStatus(icdKey, duration, true)
		char.AddStatus(bondKey, -1, false)

		char.ModifyHPDebtByRatio(hp)
		if char.CurrentHPDebt() > 0 {
			if char.CurrentHPDebt() >= 0.3*maxhp {
				dmg[attributes.DmgP] = bondDMG + bondDMG2
				char.AddStatMod(character.StatMod{
					Base:         modifier.NewBaseWithHitlag("cms-dmg-bonus", -1),
					AffectedStat: attributes.DmgP,
					Amount: func() ([]float64, bool) {
						return dmg, true
					},
				})
			} else {
				dmg[attributes.DmgP] = bondDMG
				char.AddStatMod(character.StatMod{
					Base:         modifier.NewBaseWithHitlag("cms-dmg-bonus", -1),
					AffectedStat: attributes.DmgP,
					Amount: func() ([]float64, bool) {
						return dmg, true
					},
				})
			}
		}
		return false
	}, fmt.Sprintf("cms-dmg-%v", char.Base.Key.String()))

	// Remove BOL status when healed
	c.Events.Subscribe(event.OnHeal, func(args ...interface{}) bool {
		index := args[1].(int)
		if index != char.Index {
			return false
		}
		if char.CurrentHPDebt() > 0 {
			return false
		}
		if !char.StatusIsActive(bondKey) {
			return false
		}
		char.DeleteStatus(bondKey)
		return false
	}, fmt.Sprintf("cms-remove-bond-%v", char.Base.Key.String()))

	return w, nil
}
