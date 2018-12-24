// +build ignore

package template

import (
	"log"

	"github.com/shurcooL/vfsgen"

	"github.com/gofunct/hack/cmd/hack-gen-type/template"
)

func main() {
	err := vfsgen.Generate(template.FS, vfsgen.Options{
		PackageName:  "template",
		BuildTags:    "!vfsgen",
		VariableName: "FS",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
