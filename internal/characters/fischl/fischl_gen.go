// Code generated by "pipeline"; DO NOT EDIT.
package fischl

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
	auto = [][]float64{
		auto_1,
		auto_2,
		auto_3,
		auto_4,
		auto_5,
	}
)

var (
	// attack: aim = [5]
	aim = []float64{
		0.43860000371932983,
		0.47429999709129333,
		0.5099999904632568,
		0.5609999895095825,
		0.5967000126838684,
		0.637499988079071,
		0.6935999989509583,
		0.7497000098228455,
		0.8058000206947327,
		0.8669999837875366,
		0.9282000064849854,
		0.9894000291824341,
		1.0506000518798828,
		1.111799955368042,
		1.1729999780654907,
	}
	// attack: auto_1 = [0]
	auto_1 = []float64{
		0.4411799907684326,
		0.4770900011062622,
		0.5130000114440918,
		0.564300000667572,
		0.6002100110054016,
		0.6412500143051147,
		0.6976799964904785,
		0.7541099786758423,
		0.8105400204658508,
		0.8720999956130981,
		0.9336599707603455,
		0.9952200055122375,
		1.0567799806594849,
		1.118340015411377,
		1.179900050163269,
	}
	// attack: auto_2 = [1]
	auto_2 = []float64{
		0.4678399860858917,
		0.5059199929237366,
		0.5440000295639038,
		0.5983999967575073,
		0.6364799737930298,
		0.6800000071525574,
		0.7398399710655212,
		0.7996799945831299,
		0.8595200181007385,
		0.9247999787330627,
		0.9900799989700317,
		1.055359959602356,
		1.1206400394439697,
		1.185920000076294,
		1.2511999607086182,
	}
	// attack: auto_3 = [2]
	auto_3 = []float64{
		0.5813599824905396,
		0.6286799907684326,
		0.6759999990463257,
		0.7436000108718872,
		0.7909200191497803,
		0.8450000286102295,
		0.9193599820137024,
		0.9937199950218201,
		1.068079948425293,
		1.1491999626159668,
		1.2303199768066406,
		1.3114399909973145,
		1.3925600051879883,
		1.473680019378662,
		1.554800033569336,
	}
	// attack: auto_4 = [3]
	auto_4 = []float64{
		0.5770599842071533,
		0.6240299940109253,
		0.6710000038146973,
		0.738099992275238,
		0.78507000207901,
		0.8387500047683716,
		0.912559986114502,
		0.9863700270652771,
		1.0601799488067627,
		1.1406999826431274,
		1.2212200164794922,
		1.301740050315857,
		1.3822599649429321,
		1.4627799987792969,
		1.5433000326156616,
	}
	// attack: auto_5 = [4]
	auto_5 = []float64{
		0.7206799983978271,
		0.7793400287628174,
		0.8379999995231628,
		0.9218000173568726,
		0.980459988117218,
		1.0475000143051147,
		1.139680027961731,
		1.2318600416183472,
		1.3240400552749634,
		1.4246000051498413,
		1.5251599550247192,
		1.6257200241088867,
		1.7262799739837646,
		1.8268400430679321,
		1.92739999294281,
	}
	// attack: fullaim = [6]
	fullaim = []float64{
		1.2400000095367432,
		1.3329999446868896,
		1.4259999990463257,
		1.5499999523162842,
		1.6430000066757202,
		1.7359999418258667,
		1.8600000143051147,
		1.9839999675750732,
		2.1080000400543213,
		2.2320001125335693,
		2.3559999465942383,
		2.4800000190734863,
		2.634999990463257,
		2.7899999618530273,
		2.944999933242798,
	}
	// skill: birdAtk = [0]
	birdAtk = []float64{
		0.8880000114440918,
		0.9545999765396118,
		1.0211999416351318,
		1.1100000143051147,
		1.1765999794006348,
		1.2431999444961548,
		1.3320000171661377,
		1.420799970626831,
		1.509600043296814,
		1.5983999967575073,
		1.6871999502182007,
		1.7760000228881836,
		1.8869999647140503,
		1.9980000257492065,
		2.1089999675750732,
	}
	// skill: birdSum = [1]
	birdSum = []float64{
		1.1543999910354614,
		1.2409800291061401,
		1.3275599479675293,
		1.4429999589920044,
		1.529579997062683,
		1.6161600351333618,
		1.731600046157837,
		1.847040057182312,
		1.9624799489974976,
		2.0779199600219727,
		2.1933600902557373,
		2.308799982070923,
		2.4530999660491943,
		2.597399950027466,
		2.7416999340057373,
	}
	// burst: burst = [0]
	burst = []float64{
		2.0799999237060547,
		2.2360000610351562,
		2.3919999599456787,
		2.5999999046325684,
		2.75600004196167,
		2.9119999408721924,
		3.119999885559082,
		3.328000068664551,
		3.5360000133514404,
		3.74399995803833,
		3.9519999027252197,
		4.159999847412109,
		4.420000076293945,
		4.679999828338623,
		4.940000057220459,
	}
)