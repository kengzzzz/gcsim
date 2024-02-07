// Code generated by "pipeline"; DO NOT EDIT.
package sucrose

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
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.3346399962902069,
		0.35973799228668213,
		0.38483598828315735,
		0.41830000281333923,
		0.44339799880981445,
		0.4684959948062897,
		0.5019599795341492,
		0.535423994064331,
		0.5688880085945129,
		0.6023520231246948,
		0.6358159780502319,
		0.6692799925804138,
		0.71110999584198,
		0.7529399991035461,
		0.7947700023651123,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.30616000294685364,
		0.32912200689315796,
		0.3520840108394623,
		0.38269999623298645,
		0.40566200017929077,
		0.4286240041255951,
		0.45923998951911926,
		0.4898560047149658,
		0.52047199010849,
		0.5510879755020142,
		0.5817040205001831,
		0.6123200058937073,
		0.6505900025367737,
		0.6888599991798401,
		0.7271299958229065,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.38447999954223633,
		0.413316011428833,
		0.4421519935131073,
		0.4805999994277954,
		0.5094360113143921,
		0.5382720232009888,
		0.5767199993133545,
		0.6151679754257202,
		0.6536160111427307,
		0.6920639872550964,
		0.7305120229721069,
		0.7689599990844727,
		0.8170199990272522,
		0.8650799989700317,
		0.9131399989128113,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.47917601466178894,
		0.5151140093803406,
		0.5510519742965698,
		0.5989699959754944,
		0.6349080204963684,
		0.6708459854125977,
		0.7187640070915222,
		0.7666820287704468,
		0.8145989775657654,
		0.8625169992446899,
		0.9104340076446533,
		0.9583520293235779,
		1.0182490348815918,
		1.078145980834961,
		1.1380430459976196,
	}
	// attack: charge = [4]
	charge = []float64{
		1.2015999555587769,
		1.2917200326919556,
		1.3818399906158447,
		1.5019999742507935,
		1.5921200513839722,
		1.6822400093078613,
		1.80239999294281,
		1.9225599765777588,
		2.042720079421997,
		2.1628799438476562,
		2.2830400466918945,
		2.4031999111175537,
		2.5534000396728516,
		2.7035999298095703,
		2.853800058364868,
	}
	// skill: skill = [0]
	skill = []float64{
		2.111999988555908,
		2.270400047302246,
		2.428800106048584,
		2.640000104904175,
		2.7983999252319336,
		2.9567999839782715,
		3.1679999828338623,
		3.379199981689453,
		3.590399980545044,
		3.8015999794006348,
		4.012800216674805,
		4.223999977111816,
		4.48799991607666,
		4.751999855041504,
		5.015999794006348,
	}
	// burst: burstAbsorb = [1]
	burstAbsorb = []float64{
		0.4399999976158142,
		0.4729999899864197,
		0.5059999823570251,
		0.550000011920929,
		0.5830000042915344,
		0.6159999966621399,
		0.6600000262260437,
		0.7039999961853027,
		0.7480000257492065,
		0.7919999957084656,
		0.8360000252723694,
		0.8799999952316284,
		0.9350000023841858,
		0.9900000095367432,
		1.0449999570846558,
	}
	// burst: burstDot = [0]
	burstDot = []float64{
		1.4800000190734863,
		1.590999960899353,
		1.7020000219345093,
		1.850000023841858,
		1.9609999656677246,
		2.072000026702881,
		2.2200000286102295,
		2.368000030517578,
		2.5160000324249268,
		2.6640000343322754,
		2.812000036239624,
		2.9600000381469727,
		3.1449999809265137,
		3.3299999237060547,
		3.515000104904175,
	}
)