// Code generated by "pipeline"; DO NOT EDIT.
package ayaka

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
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.45725300908088684,
		0.4944719970226288,
		0.5316900014877319,
		0.5848590135574341,
		0.6220769882202148,
		0.6646130084991455,
		0.7230979800224304,
		0.7815840244293213,
		0.8400700092315674,
		0.9038730263710022,
		0.9676759839057922,
		1.031479001045227,
		1.0952810049057007,
		1.1590839624404907,
		1.2228870391845703,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.486845999956131,
		0.5264729857444763,
		0.566100001335144,
		0.6227099895477295,
		0.6623370051383972,
		0.7076249718666077,
		0.7698959708213806,
		0.8321670293807983,
		0.8944380283355713,
		0.962369978427887,
		1.0303020477294922,
		1.0982340574264526,
		1.1661659479141235,
		1.234097957611084,
		1.3020299673080444,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.6262180209159851,
		0.6771889925003052,
		0.72816002368927,
		0.8009759783744812,
		0.851947009563446,
		0.9101999998092651,
		0.9902979731559753,
		1.0703949928283691,
		1.1504930257797241,
		1.2378720045089722,
		1.3252509832382202,
		1.4126299619674683,
		1.5000100135803223,
		1.5873889923095703,
		1.6747679710388184,
	}
	// attack: attack_4 = [3 3 3]
	attack_4 = [][]float64{
		{
			0.22646400332450867,
			0.2448969930410385,
			0.26333001255989075,
			0.2896629869937897,
			0.3080959916114807,
			0.3291630148887634,
			0.35812899470329285,
			0.38709500432014465,
			0.41606101393699646,
			0.44766101241111755,
			0.47926101088523865,
			0.510860025882721,
			0.542460024356842,
			0.574059009552002,
			0.605659008026123,
		},
		{
			0.22646400332450867,
			0.2448969930410385,
			0.26333001255989075,
			0.2896629869937897,
			0.3080959916114807,
			0.3291630148887634,
			0.35812899470329285,
			0.38709500432014465,
			0.41606101393699646,
			0.44766101241111755,
			0.47926101088523865,
			0.510860025882721,
			0.542460024356842,
			0.574059009552002,
			0.605659008026123,
		},
		{
			0.22646400332450867,
			0.2448969930410385,
			0.26333001255989075,
			0.2896629869937897,
			0.3080959916114807,
			0.3291630148887634,
			0.35812899470329285,
			0.38709500432014465,
			0.41606101393699646,
			0.44766101241111755,
			0.47926101088523865,
			0.510860025882721,
			0.542460024356842,
			0.574059009552002,
			0.605659008026123,
		},
	}
	// attack: attack_5 = [6]
	attack_5 = []float64{
		0.7818170189857483,
		0.8454539775848389,
		0.909089982509613,
		0.9999989867210388,
		1.063634991645813,
		1.1363630294799805,
		1.2363619804382324,
		1.3363620042800903,
		1.4363620281219482,
		1.5454529523849487,
		1.6545439958572388,
		1.7636350393295288,
		1.872725009918213,
		1.981816053390503,
		2.090907096862793,
	}
	// attack: ca = [7]
	ca = []float64{
		0.5512599945068359,
		0.5961300134658813,
		0.640999972820282,
		0.7050999999046326,
		0.749970018863678,
		0.8012499809265137,
		0.8717600107192993,
		0.9422699809074402,
		1.012779951095581,
		1.0896999835968018,
		1.1666200160980225,
		1.2435400485992432,
		1.3204599618911743,
		1.397379994392395,
		1.4743000268936157,
	}
	// skill: skill = [0]
	skill = []float64{
		2.3919999599456787,
		2.5713999271392822,
		2.7507998943328857,
		2.990000009536743,
		3.1693999767303467,
		3.34879994392395,
		3.5880000591278076,
		3.827199935913086,
		4.066400051116943,
		4.305600166320801,
		4.5447998046875,
		4.783999919891357,
		5.083000183105469,
		5.381999969482422,
		5.681000232696533,
	}
	// burst: burstBloom = [1]
	burstBloom = []float64{
		1.684499979019165,
		1.8108370304107666,
		1.9371750354766846,
		2.1056249141693115,
		2.231961965560913,
		2.358299970626831,
		2.526750087738037,
		2.695199966430664,
		2.86365008354187,
		3.032099962234497,
		3.200550079345703,
		3.36899995803833,
		3.579561948776245,
		3.7901248931884766,
		4.000687122344971,
	}
	// burst: burstCut = [0]
	burstCut = []float64{
		1.1230000257492065,
		1.20722496509552,
		1.291450023651123,
		1.403749942779541,
		1.487975001335144,
		1.5721999406814575,
		1.684499979019165,
		1.7968000173568726,
		1.90910005569458,
		2.021399974822998,
		2.133699893951416,
		2.246000051498413,
		2.3863749504089355,
		2.526750087738037,
		2.6671249866485596,
	}
)