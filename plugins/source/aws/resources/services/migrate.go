package main

import (
	"flag"
	"fmt"
	"github.com/ettle/strcase"
	"os"
	"path"
	"strings"
	"text/template"
)

var directory string
var createStarterTemplate bool

func migrate(dir string) []string {
	items, _ := os.ReadDir(dir)
	names := make([]string, 0)
	for _, it := range items {
		if strings.HasSuffix(it.Name(), ".go.bk") {
			fmt.Fprintln(os.Stderr, "Skipping "+it.Name())
			names = append(names, strings.TrimSuffix(it.Name(), ".go.bk"))
			continue
		}
		if strings.HasSuffix(it.Name(), "_fetch.go") {
			fmt.Fprintln(os.Stderr, "Skipping "+it.Name())
			continue
		}
		if strings.HasSuffix(it.Name(), ".go") && !strings.HasSuffix(it.Name(), "_test.go") {
			migrateFile(path.Join(dir, it.Name()))
			names = append(names, strings.TrimSuffix(it.Name(), ".go"))
		}
	}
	return names
}

func migrateFile(f string) {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	s := string(data)
	stopAt := strings.Index(s, "// ==========")
	if stopAt == -1 {
		fmt.Fprintln(os.Stderr, "Already migrated?")
		return
	}

	importEnd := strings.Index(s, ")")

	d := path.Dir(f)
	b := path.Base(f)
	old := path.Join(d, b+".bk")
	fetch := path.Join(d, b[:len(b)-3]+"_fetch.go")

	// write backup file
	err = os.WriteFile(old, []byte(s[:stopAt]), 0755)
	if err != nil {
		panic(err)
	}

	// write resolvers
	fetchContent := s[stopAt:]
	lines := strings.Split(fetchContent, "\n")
	firstLine := 0 // skip comment lines
	for _, line := range lines {
		if strings.HasPrefix(line, "//") {
			firstLine++
			continue
		}
		break
	}
	lines = lines[firstLine:]
	fetchContent = strings.Join(lines, "\n")
	fetchContent = s[:importEnd+1] + "\n" + fetchContent
	err = os.WriteFile(fetch, []byte(fetchContent), 0755)
	if err != nil {
		panic(err)
	}

	// delete original file (this will need to be regenerated)
	os.RemoveAll(f)
}

var starterTemplate = `package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/{{.ServiceName}}/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func {{.ServiceName | Title}}Resources() []*Resource {
	resources := []*Resource{
		{{ range $i, $name := .Names }}
		{
			SubService: "{{ $name }}",
			Struct:     &types.{{ $name | Title }}{},
			SkipFields: []string{"ARN"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: ` + "`" + `schema.PathResolver("ARN")` + "`" + `,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{{ end }}
	}

	// set default values
	for _, r := range resources {
		r.Service = "{{.ServiceName}}"
		// r.Multiplex = ""
	}
	return resources
}
`

func writeStarterTemplate(dir string, names []string) {
	tpl, err := template.New("starter").Funcs(template.FuncMap{
		"Title": strcase.ToGoPascal,
	}).Parse(starterTemplate)
	if err != nil {
		panic(err)
	}
	pth := "../../codegen/recipes/" + dir + ".go"
	f, err := os.Create(pth)
	if err != nil {
		panic(err)
	}
	values := map[string]interface{}{
		"ServiceName": dir,
		"Names":       names,
	}
	err = tpl.Execute(f, values)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starter template created at " + pth)
}

func main() {
	flag.StringVar(&directory, "d", "", "directory to migrate")
	flag.BoolVar(&createStarterTemplate, "s", false, "whether to create starter template")
	flag.Parse()

	if directory == "" {
		fmt.Fprintln(os.Stderr, "Usage: -d <directory_to_migrate>")
		return
	}
	names := migrate(directory)
	if createStarterTemplate {
		writeStarterTemplate(directory, names)
	}
}
