// Code generated by "pipeline"; DO NOT EDIT.
package mavuika
import (
	_ "embed"
	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)
//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	1: {"hold", "recast"},
	5: {"collision"},
	6: {"collision"},
}
func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Mavuika, ValidateParamKeys)
}
func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}
func (x *char) Data() *model.AvatarData {
	return base
}
var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
	skillAttack = [][]float64{
		skillAttack_1,
		skillAttack_2,
		skillAttack_3,
		skillAttack_4,
		skillAttack_5,
	}
)
var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.80035,
		0.865495,
		0.93064,
		1.023704,
		1.088849,
		1.1633,
		1.26567,
		1.368041,
		1.470411,
		1.582088,
		1.693765,
		1.805442,
		1.917118,
		2.028795,
		2.140472,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.364799,
		0.394492,
		0.424185,
		0.466604,
		0.496296,
		0.530231,
		0.576892,
		0.623552,
		0.670212,
		0.721114,
		0.772017,
		0.822919,
		0.873821,
		0.924723,
		0.975625,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.332232,
		0.359274,
		0.386317,
		0.424948,
		0.45199,
		0.482896,
		0.525391,
		0.567885,
		0.61038,
		0.656738,
		0.703096,
		0.749454,
		0.795812,
		0.84217,
		0.888528,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		1.161929,
		1.256504,
		1.35108,
		1.486188,
		1.580764,
		1.68885,
		1.837469,
		1.986088,
		2.134706,
		2.296836,
		2.458966,
		2.621095,
		2.783225,
		2.945354,
		3.107484,
	}
	// attack: charge = [4]
	charge = []float64{
		1.93844,
		2.09622,
		2.254,
		2.4794,
		2.63718,
		2.8175,
		3.06544,
		3.31338,
		3.56132,
		3.8318,
		4.10228,
		4.37276,
		4.64324,
		4.91372,
		5.1842,
	}
	// attack: collision = [6]
	collision = []float64{
		0.745878,
		0.806589,
		0.8673,
		0.95403,
		1.014741,
		1.084125,
		1.179528,
		1.274931,
		1.370334,
		1.47441,
		1.578486,
		1.682562,
		1.786638,
		1.890714,
		1.99479,
	}
	// attack: highPlunge = [8]
	highPlunge = []float64{
		1.862889,
		2.01452,
		2.16615,
		2.382765,
		2.534396,
		2.707688,
		2.945964,
		3.184241,
		3.422517,
		3.682455,
		3.942393,
		4.202331,
		4.462269,
		4.722207,
		4.982145,
	}
	// attack: lowPlunge = [7]
	lowPlunge = []float64{
		1.49144,
		1.612836,
		1.734233,
		1.907656,
		2.029052,
		2.167791,
		2.358556,
		2.549322,
		2.740087,
		2.948195,
		3.156303,
		3.364411,
		3.572519,
		3.780627,
		3.988735,
	}
	// skill: skill = [0]
	skill = []float64{
		0.744,
		0.7998,
		0.8556,
		0.93,
		0.9858,
		1.0416,
		1.116,
		1.1904,
		1.2648,
		1.3392,
		1.4136,
		1.488,
		1.581,
		1.674,
		1.767,
	}
	// skill: skillAttack_1 = [3]
	skillAttack_1 = []float64{
		0.572648,
		0.619259,
		0.66587,
		0.732457,
		0.779068,
		0.832337,
		0.905583,
		0.978829,
		1.052075,
		1.131979,
		1.211883,
		1.291788,
		1.371692,
		1.451597,
		1.531501,
	}
	// skill: skillAttack_2 = [4]
	skillAttack_2 = []float64{
		0.591327,
		0.639459,
		0.68759,
		0.756349,
		0.80448,
		0.859488,
		0.935122,
		1.010757,
		1.086392,
		1.168903,
		1.251414,
		1.333925,
		1.416435,
		1.498946,
		1.581457,
	}
	// skill: skillAttack_3 = [5]
	skillAttack_3 = []float64{
		0.699868,
		0.756834,
		0.8138,
		0.89518,
		0.952146,
		1.01725,
		1.106768,
		1.196286,
		1.285804,
		1.38346,
		1.481116,
		1.578772,
		1.676428,
		1.774084,
		1.87174,
	}
	// skill: skillAttack_4 = [6]
	skillAttack_4 = []float64{
		0.697047,
		0.753784,
		0.81052,
		0.891572,
		0.948308,
		1.01315,
		1.102307,
		1.191464,
		1.280622,
		1.377884,
		1.475146,
		1.572409,
		1.669671,
		1.766934,
		1.864196,
	}
	// skill: skillAttack_5 = [7]
	skillAttack_5 = []float64{
		0.910035,
		0.984107,
		1.05818,
		1.163998,
		1.238071,
		1.322725,
		1.439125,
		1.555525,
		1.671924,
		1.798906,
		1.925888,
		2.052869,
		2.179851,
		2.306832,
		2.433814,
	}
	// skill: skillCharge = [9]
	skillCharge = []float64{
		0.989,
		1.0695,
		1.15,
		1.265,
		1.3455,
		1.4375,
		1.564,
		1.6905,
		1.817,
		1.955,
		2.093,
		2.231,
		2.369,
		2.507,
		2.645,
	}
	// skill: skillChargeFinal = [10]
	skillChargeFinal = []float64{
		1.376,
		1.488,
		1.6,
		1.76,
		1.872,
		2,
		2.176,
		2.352,
		2.528,
		2.72,
		2.912,
		3.104,
		3.296,
		3.488,
		3.68,
	}
	// skill: skillDash = [8]
	skillDash = []float64{
		0.8084,
		0.8742,
		0.94,
		1.034,
		1.0998,
		1.175,
		1.2784,
		1.3818,
		1.4852,
		1.598,
		1.7108,
		1.8236,
		1.9364,
		2.0492,
		2.162,
	}
	// skill: skillPlunge = [11]
	skillPlunge = []float64{
		1.5996,
		1.7298,
		1.86,
		2.046,
		2.1762,
		2.325,
		2.5296,
		2.7342,
		2.9388,
		3.162,
		3.3852,
		3.6084,
		3.8316,
		4.0548,
		4.278,
	}
	// skill: skillRing = [1]
	skillRing = []float64{
		1.28,
		1.376,
		1.472,
		1.6,
		1.696,
		1.792,
		1.92,
		2.048,
		2.176,
		2.304,
		2.432,
		2.56,
		2.72,
		2.88,
		3.04,
	}
	// burst: burst = [0]
	burst = []float64{
		4.448,
		4.7816,
		5.1152,
		5.56,
		5.8936,
		6.2272,
		6.672,
		7.1168,
		7.5616,
		8.0064,
		8.4512,
		8.896,
		9.452,
		10.008,
		10.564,
	}
	// burst: burstCABonus = [4]
	burstCABonus = []float64{
		0.00516,
		0.00558,
		0.006,
		0.0066,
		0.00702,
		0.0075,
		0.00816,
		0.00882,
		0.00948,
		0.0102,
		0.01092,
		0.01164,
		0.01236,
		0.01308,
		0.0138,
	}
	// burst: burstNABonus = [3]
	burstNABonus = []float64{
		0.00258,
		0.00279,
		0.003,
		0.0033,
		0.00351,
		0.00375,
		0.00408,
		0.00441,
		0.00474,
		0.0051,
		0.00546,
		0.00582,
		0.00618,
		0.00654,
		0.0069,
	}
	// burst: burstQBonus = [2]
	burstQBonus = []float64{
		0.016,
		0.0172,
		0.0184,
		0.02,
		0.0212,
		0.0224,
		0.024,
		0.0256,
		0.0272,
		0.0288,
		0.0304,
		0.032,
		0.034,
		0.036,
		0.038,
	}
)