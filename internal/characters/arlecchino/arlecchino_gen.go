// Code generated by "pipeline"; DO NOT EDIT.
package arlecchino

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		{attack_3},
		attack_4,
		{attack_5},
		{attack_6},
	}
)

var (
	// attack: attack_1 = [0 0]
	attack_1 = []float64{
		0.4712,
		0.5096,
		0.5480,
		0.6027,
		0.6411,
		0.6849,
		0.7452,
		0.8055,
		0.8658,
		0.9315,
		0.9973,
		1.0630,
		1.1288,
		1.1945,
		1.2603,
	}
	// attack: attack_2 = [2]
	attack_2 = []float64{
		0.5169,
		0.559,
		0.6011,
		0.6612,
		0.7033,
		0.7514,
		0.8175,
		0.8836,
		0.9497,
		1.0219,
		1.094,
		1.1661,
		1.2382,
		1.3104,
		1.3825,
	}
	// attack: attack_3 = [3]
	attack_3 = []float64{
		0.6487,
		0.7015,
		0.7543,
		0.8297,
		0.8825,
		0.9428,
		1.0258,
		1.1088,
		1.1918,
		1.2823,
		1.3728,
		1.4633,
		1.5538,
		1.6443,
		1.7348,
	}
	// attack: attack_4 = [4 4]
	attack_4 = [][]float64{
		{
			0.3563,
			0.3853,
			0.4143,
			0.4558,
			0.4848,
			0.5179,
			0.5635,
			0.6091,
			0.6547,
			0.7044,
			0.7541,
			0.8038,
			0.8536,
			0.9033,
			0.953,
		},
		{
			0.3563,
			0.3853,
			0.4143,
			0.4558,
			0.4848,
			0.5179,
			0.5635,
			0.6091,
			0.6547,
			0.7044,
			0.7541,
			0.8038,
			0.8536,
			0.9033,
			0.953,
		},
	}
	// attack: attack_5 = [6]
	attack_5 = []float64{
		0.708,
		0.7656,
		0.8233,
		0.9056,
		0.9632,
		1.0291,
		1.1196,
		1.2102,
		1.3008,
		1.3996,
		1.4984,
		1.5971,
		1.6959,
		1.7947,
		1.8935,
	}
	// attack: attack_6 = [7]
	attack_6 = []float64{
		0.708,
		0.7656,
		0.8233,
		0.9056,
		0.9632,
		1.0291,
		1.1196,
		1.2102,
		1.3996,
		1.4984,
		1.5971,
		1.6959,
		1.7947,
		1.8935,
	}
	// attack: charge = [8]
	charge = []float64{
		1.2969,
		1.4024,
		1.508,
		1.6588,
		1.7644,
		1.885,
		2.0509,
		2.2168,
		2.3826,
		2.5636,
		2.7446,
		2.9255,
		3.1065,
		3.2874,
		3.4684,
	}
	// attack: collision = [8]
	collision = []float64{
		0.6393240094184875,
		0.6913620233535767,
		0.743399977684021,
		0.8177400231361389,
		0.8697779774665833,
		0.9292500019073486,
		1.011023998260498,
		1.0927979946136475,
		1.1745719909667969,
		1.2637799978256226,
		1.3529880046844482,
		1.442196011543274,
		1.5314040184020996,
		1.6206120252609253,
		1.709820032119751,
	}
	// attack: highPlunge = [10]
	highPlunge = []float64{
		1.59676194190979,
		1.7267309427261353,
		1.8566999435424805,
		2.042370080947876,
		2.1723389625549316,
		2.3208749294281006,
		2.5251119136810303,
		2.72934889793396,
		2.9335858821868896,
		3.1563899517059326,
		3.3791940212249756,
		3.6019980907440186,
		3.8248019218444824,
		4.047605991363525,
		4.270410060882568,
	}
	// attack: lowPlunge = [9]
	lowPlunge = []float64{
		1.2783770561218262,
		1.3824310302734375,
		1.4864850044250488,
		1.635133981704712,
		1.7391870021820068,
		1.858106017112732,
		2.021620035171509,
		2.1851329803466797,
		2.3486459255218506,
		2.527024984359741,
		2.7054030895233154,
		2.8837809562683105,
		3.0621590614318848,
		3.24053692817688,
		3.418915033340454,
	}
	blooddebt = []float64{
		1.204,
		1.302,
		1.400,
		1.540,
		1.638,
		1.750,
		1.904,
		2.058,
		2.212,
		2.380,
		2.548,
		2.716,
		2.884,
		3.052,
		3.220,
	}
	// skill: skill = [0]
	skillSpike = []float64{
		0.1484,
		0.1595,
		0.1707,
		0.1855,
		0.1966,
		0.2078,
		0.2226,
		0.2374,
		0.2523,
		0.2671,
		0.282,
		0.2968,
		0.3153,
		0.3339,
		0.3525,
	}
	// skill: skill = [0]
	skillFinal = []float64{
		1.3356,
		1.4358,
		1.5359,
		1.6695,
		1.7697,
		1.8698,
		2.0034,
		2.137,
		2.2705,
		2.4041,
		2.5376,
		2.6712,
		2.8382,
		3.0051,
		3.172,
	}
	// skill: skill = [0]
	skillSigil = []float64{
		0.212,
		0.2279,
		0.2438,
		0.265,
		0.2809,
		0.2968,
		0.318,
		0.3392,
		0.3604,
		0.3816,
		0.4028,
		0.424,
		0.4505,
		0.477,
		0.5035,
	}
	// burst: burstDrain = [1]
	burst = []float64{
		3.704,
		3.9818,
		4.2596,
		4.63,
		4.9078,
		5.1856,
		5.556,
		5.9264,
		6.2968,
		6.6672,
		7.0376,
		7.408,
		7.871,
		8.334,
		8.797,
	}
)