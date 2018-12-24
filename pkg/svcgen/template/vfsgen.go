// +build ignore

package template

import (
	"log"

	"github.com/gofunct/hack/pkg/svcgen/template"
	"github.com/shurcooL/vfsgen"
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
