package artifact

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"
)

type assetData struct {
	Data []assetMapping
}

type assetMapping struct {
	Key      string
	IconName string
}

func (g *Generator) GenerateAssetsKey(dir string) error {
	t, err := template.New("artifacts_asset_tmpl").Parse(assetTmpl)
	if err != nil {
		return fmt.Errorf("failed to build template: %w", err)
	}

	var data assetData
	for i := range g.artifacts {
		v := g.artifacts[i]
		dm, ok := g.data[v.Key]
		if !ok {
			log.Printf("No data found for %v; skipping", v.Key)
			continue
		}
		data.Data = append(data.Data,
			assetMapping{
				Key:      v.Key + "_circlet",
				IconName: dm.ImageNames.Circlet,
			},
			assetMapping{
				Key:      v.Key + "_flower",
				IconName: dm.ImageNames.Flower,
			},
			assetMapping{
				Key:      v.Key + "_goblet",
				IconName: dm.ImageNames.Goblet,
			},
			assetMapping{
				Key:      v.Key + "_plume",
				IconName: dm.ImageNames.Plume,
			},
			assetMapping{
				Key:      v.Key + "_sands",
				IconName: dm.ImageNames.Sands,
			},
		)
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)
	if err != nil {
		return fmt.Errorf("failed to execute artifacts assets template: %w", err)
	}
	src := buf.Bytes()
	dst, err := format.Source(src)
	if err != nil {
		fmt.Println(string(src))
		return fmt.Errorf("failed to gofmt on generated artifacts assets data: %w", err)
	}
	err = os.WriteFile(fmt.Sprintf("%v/artifacts_gen.go", dir), dst, 0o644)
	if err != nil {
		return fmt.Errorf("failed to write artifacts_gen.go: %w", err)
	}

	return nil
}

const assetTmpl = `// Code generated by "pipeline"; DO NOT EDIT.
package assets

var artfactMap = map[string]string{
{{ range $index, $ele := .Data -}}
"{{$ele.Key}}": "{{$ele.IconName}}",
{{ end }}
}
`