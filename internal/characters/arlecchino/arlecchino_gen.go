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
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.4712370038032532,
		0.5095940232276917,
		0.547950029373169,
		0.602744996547699,
		0.6411010026931763,
		0.6849380135536194,
		0.7452120184898376,
		0.8054869771003723,
		0.8657609820365906,
		0.9315149784088135,
		0.9972689747810364,
		1.0630229711532593,
		1.128777027130127,
		1.194530963897705,
		1.2602850198745728,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.5169370174407959,
		0.5590140223503113,
		0.6010900139808655,
		0.6611989736557007,
		0.7032750248908997,
		0.7513629794120789,
		0.8174819946289062,
		0.8836020231246948,
		0.9497219920158386,
		1.021852970123291,
		1.0939840078353882,
		1.1661150455474854,
		1.2382450103759766,
		1.3103760480880737,
		1.3825069665908813,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.6486809849739075,
		0.701479971408844,
		0.7542799711227417,
		0.8297079801559448,
		0.8825079798698425,
		0.9428499937057495,
		1.0258209705352783,
		1.1087919473648071,
		1.1917619705200195,
		1.2822760343551636,
		1.372789978981018,
		1.4633029699325562,
		1.5538170337677002,
		1.6443300247192383,
		1.7348439693450928,
	}
	// attack: attack_4 = [3 3]
	attack_4 = [][]float64{
		{
			0.3563370108604431,
			0.38534098863601685,
			0.41434499621391296,
			0.4557799994945526,
			0.48478400707244873,
			0.5179309844970703,
			0.5635089874267578,
			0.6090869903564453,
			0.6546649932861328,
			0.7043870091438293,
			0.7541080117225647,
			0.8038290143013,
			0.8535509705543518,
			0.9032719731330872,
			0.9529929757118225,
		},
		{
			0.3563370108604431,
			0.38534098863601685,
			0.41434499621391296,
			0.4557799994945526,
			0.48478400707244873,
			0.5179309844970703,
			0.5635089874267578,
			0.6090869903564453,
			0.6546649932861328,
			0.7043870091438293,
			0.7541080117225647,
			0.8038290143013,
			0.8535509705543518,
			0.9032719731330872,
			0.9529929757118225,
		},
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.7080119848251343,
		0.7656409740447998,
		0.8232700228691101,
		0.9055969715118408,
		0.9632260203361511,
		1.029086947441101,
		1.1196470260620117,
		1.2102069854736328,
		1.300766944885254,
		1.3995590209960938,
		1.498350977897644,
		1.5971440076828003,
		1.6959359645843506,
		1.7947289943695068,
		1.8935209512710571,
	}
	// attack: attack_6 = [5]
	attack_6 = []float64{
		0.8485530018806458,
		0.9176220297813416,
		0.9866899847984314,
		1.08535897731781,
		1.1544270515441895,
		1.2333619594573975,
		1.341897964477539,
		1.4504339694976807,
		1.5589699745178223,
		1.6773730516433716,
		1.7957760095596313,
		1.9141789674758911,
		2.032581090927124,
		2.150984048843384,
		2.2693870067596436,
	}
	// attack: blooddebt = [11]
	blooddebt = []float64{
		1.5049999952316284,
		1.627500057220459,
		1.75,
		1.9249999523162842,
		2.047499895095825,
		2.1875,
		2.380000114440918,
		2.572499990463257,
		2.765000104904175,
		2.9749999046325684,
		3.184999942779541,
		3.3949999809265137,
		3.6050000190734863,
		3.815000057220459,
		4.025000095367432,
	}
	// attack: charge = [6]
	charge = []float64{
		1.2968800067901611,
		1.4024399518966675,
		1.5080000162124634,
		1.6588000059127808,
		1.764359951019287,
		1.8849999904632568,
		2.050879955291748,
		2.2167599201202393,
		2.3826398849487305,
		2.5636000633239746,
		2.7445600032806396,
		2.9255199432373047,
		3.1064798831939697,
		3.287440061569214,
		3.468400001525879,
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
	// skill: skillFinal = [1]
	skillFinal = []float64{
		1.3356000185012817,
		1.435770034790039,
		1.5359400510787964,
		1.6694999933242798,
		1.769670009613037,
		1.8698400259017944,
		2.0034000873565674,
		2.136960029602051,
		2.270519971847534,
		2.4040799140930176,
		2.53764009475708,
		2.6712000370025635,
		2.8381500244140625,
		3.0051000118255615,
		3.1720499992370605,
	}
	// skill: skillSigil = [2]
	skillSigil = []float64{
		0.3179999887943268,
		0.34185001254081726,
		0.36570000648498535,
		0.39750000834465027,
		0.42135000228881836,
		0.44519999623298645,
		0.47699999809265137,
		0.5088000297546387,
		0.5406000018119812,
		0.5723999738693237,
		0.604200005531311,
		0.6359999775886536,
		0.6757500171661377,
		0.715499997138977,
		0.7552499771118164,
	}
	// skill: skillSpike = [0]
	skillSpike = []float64{
		0.14839999377727509,
		0.15952999889850616,
		0.17066000401973724,
		0.18549999594688416,
		0.19663000106811523,
		0.2077600061893463,
		0.22259999811649323,
		0.23744000494480133,
		0.25227999687194824,
		0.26712000370025635,
		0.28196001052856445,
		0.29679998755455017,
		0.3153499960899353,
		0.33390000462532043,
		0.35245001316070557,
	}
	// burst: burst = [0]
	burst = []float64{
		3.7039999961853027,
		3.981800079345703,
		4.2596001625061035,
		4.630000114440918,
		4.907800197601318,
		5.1855998039245605,
		5.556000232696533,
		5.926400184631348,
		6.296800136566162,
		6.667200088500977,
		7.037600040435791,
		7.4079999923706055,
		7.870999813079834,
		8.333999633789062,
		8.79699993133545,
	}
)
