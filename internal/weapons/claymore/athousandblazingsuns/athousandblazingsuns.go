package athousandblazingsuns

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
	core.RegisterWeaponFunc(keys.AThousandBlazingSuns, NewWeapon)
}

type Weapon struct {
	Index    int
	extended int
}

func (w *Weapon) SetIndex(idx int) { w.Index = idx }
func (w *Weapon) Init() error      { return nil }

const (
	BuffICDKey   = "athousandblazingsuns-buff-icd"
	ExtendICDKey = "athousandblazingsuns-extend-icd"
	BuffATKPKey  = "athousandblazingsuns-atkp"
	BuffCDKey    = "athousandblazingsuns-cd"
	BuffICDDur   = 10 * 60
	ExtendICDDur = 60
	BuffDur      = 6 * 60
	ExtendDur    = 2 * 60
	MaxExtendDur = 6 * 60
)

func NewWeapon(c *core.Core, char *character.CharWrapper, p info.WeaponProfile) (info.Weapon, error) {
	w := &Weapon{}
	r := float64(p.Refine)

	mAtk := make([]float64, attributes.EndStatType)
	mCD := make([]float64, attributes.EndStatType)

	buff := func(duration int) {
		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBaseWithHitlag(BuffATKPKey, duration),
			AffectedStat: attributes.ATKP,
			Amount: func() ([]float64, bool) {
				mAtk[attributes.ATKP] = 0.21 + 0.07*r
				if char.StatusIsActive("nightsoul-blessing") { // FIXME
					mAtk[attributes.ATKP] *= 1.75
				}
				return mAtk, true
			},
		})

		char.AddStatMod(character.StatMod{
			Base:         modifier.NewBaseWithHitlag(BuffCDKey, duration),
			AffectedStat: attributes.CD,
			Amount: func() ([]float64, bool) {
				mCD[attributes.CD] = 0.15 + 0.05*r
				if char.StatusIsActive("nightsoul-blessing") { // FIXME
					mCD[attributes.CD] *= 1.75
				}
				return mCD, true
			},
		})
	}

	scorchingBrilliance := func(args ...interface{}) bool {
		if c.Player.Active() != char.Index {
			return false
		}
		if char.StatModIsActive(BuffICDKey) {
			return false
		}
		char.AddStatus(BuffICDKey, BuffICDDur, true)
		w.extended = 0
		buff(BuffDur)
		return false
	}

	c.Events.Subscribe(event.OnSkill, scorchingBrilliance, fmt.Sprintf("%v-athousandblazingsuns-skill", char.Base.Key.String()))
	c.Events.Subscribe(event.OnBurst, scorchingBrilliance, fmt.Sprintf("%v-athousandblazingsuns-burst", char.Base.Key.String()))
	c.Events.Subscribe(event.OnEnemyHit, func(args ...interface{}) bool {
		if c.Player.Active() != char.Index {
			return false
		}
		atk := args[1].(*combat.AttackEvent)
		if atk.Info.ActorIndex != char.Index {
			return false
		}
		if atk.Info.AttackTag != attacks.AttackTagNormal && atk.Info.AttackTag != attacks.AttackTagExtra {
			return false
		}
		if atk.Info.Element == attributes.Physical || atk.Info.Element == attributes.NoElement {
			return false
		}
		if char.StatModIsActive(ExtendICDKey) || !char.StatModIsActive(BuffATKPKey) {
			return false
		}
		if w.extended == MaxExtendDur {
			return false
		}

		char.AddStatus(ExtendICDKey, ExtendICDDur, true)
		w.extended += 2 * 60
		buff(char.StatusDuration(BuffATKPKey) + 2*60)

		return false
	}, fmt.Sprintf("%v-athousandblazingsuns-hit", char.Base.Key.String()))
	c.Events.Subscribe(event.OnTick, func(args ...interface{}) bool {
		if !char.StatModIsActive(BuffATKPKey) {
			return false
		}
		if c.Player.Active() == char.Index {
			return false
		}
		if !char.StatusIsActive("nightsoul-blessing") { // FIXME
			return false
		}

		buff(char.StatusDuration(BuffATKPKey) + 1)

		return false
	}, fmt.Sprintf("%v-athousandblazingsuns-tick", char.Base.Key.String()))

	return w, nil
}
