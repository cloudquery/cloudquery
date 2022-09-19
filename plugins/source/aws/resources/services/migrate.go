package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"

	"github.com/ettle/strcase"
)

var directory string
var createStarterTemplate bool

type resource struct {
	Name         string
	ExtraColumns []string
	Relations    []string
	SkipFields   []string
}

func migrate(dir string) []resource {
	items, _ := os.ReadDir(dir)
	resources := make([]resource, 0)
	for _, it := range items {
		if strings.HasSuffix(it.Name(), ".go.bk") {
			fmt.Fprintln(os.Stderr, "Skipping "+it.Name())
			resources = append(resources, resource{Name: strings.TrimSuffix(it.Name(), ".go.bk")})
			continue
		}
		if strings.HasSuffix(it.Name(), "_fetch.go") {
			fmt.Fprintln(os.Stderr, "Skipping "+it.Name())
			continue
		}
		if strings.HasSuffix(it.Name(), ".go") && !strings.HasSuffix(it.Name(), "_test.go") {
			recs := migrateFile(path.Join(dir, it.Name()), dir)
			resources = append(resources, recs...)
		}
	}
	return resources
}

func migrateFile(f, dirname string) []resource {
	data, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	s := string(data)
	stopAt := strings.Index(s, "// ==========")
	if stopAt == -1 {
		fmt.Fprintln(os.Stderr, "Already migrated?")
		return []resource{}
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

	// calculate columns that need to be copied
	resources := []resource{}
	ts := strings.Index(s, "&schema.Table{")
	te := findMatchingBracket(s, ts+len("&schema.Table{"))
	resources = append(resources, extractResources(s[ts:te], dirname)...)
	return resources
}

func extractResources(s, dirname string) []resource {
	name := findName(s, 0, dirname)
	i := strings.Index(s, "Columns: []schema.Column{")
	br := findNextBracket(s, i)
	cl := findMatchingBracket(s, br)
	cols, skip := findExtraColumns(strings.Trim(s[br:cl+1], " {}"))

	res := resource{
		Name:         name,
		ExtraColumns: cols,
		SkipFields:   skip,
	}
	relations := []string{}
	resources := []resource{res}
	substr := "Relations: []*schema.Table{"
	rs := strings.Index(s, substr)
	if rs != -1 {
		index := rs + len(substr) + 1
		for {
			rs = findNextBracket(s, index)
			if rs == -1 {
				break
			}
			re := findMatchingBracket(s, rs)
			rname := findName(s[rs:re], 0, dirname)
			pathTableIndex := strings.Index(s[rs:re], "PathTableResolver(")
			relationsIndex := strings.Index(s[rs:re], "Relations: []*schema.Table{")
			if pathTableIndex != -1 && (pathTableIndex < relationsIndex || relationsIndex == -1) {
				// skip resources with PathTableResolver and their descendants
				return []resource{}
			}
			relations = append(relations, strcase.ToGoPascal(rname))
			resources = append(resources, extractResources(s[rs:re], dirname)...)
			index = re + 1
		}
	}
	resources[0].Relations = relations

	return resources
}

var starterTemplate = `package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/{{.ServiceName}}/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func {{.ServiceName | Title}}Resources() []*Resource {
	resources := []*Resource{
		{{- range $i, $resource := .Resources }}
		{
			SubService: "{{ $resource.Name }}",
			Struct:     &types.{{ $resource.Name | Title }}{},
			SkipFields: []string{
				{{- range $s, $skip := $resource.SkipFields -}}
					"{{ $skip }}",
				{{- end -}}
			},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{{- range $c, $col := $resource.ExtraColumns }}
						{{ $col }},
					{{- end }}
				}...),
			{{- if $resource.Relations }}
			Relations: []string{
 				{{- range $r, $rel := $resource.Relations }}
					"{{ $rel }}()",
				{{- end }}
			},
			{{- end }}
		},
		{{- end }}
	}

	// set default values
	for _, r := range resources {
		r.Service = "{{.ServiceName}}"
		// r.Multiplex = ""
	}
	return resources
}
`

func writeStarterTemplate(dir string, resources []resource) {
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
		"Resources":   resources,
	}
	var buff bytes.Buffer
	err = tpl.Execute(&buff, values)
	if err != nil {
		panic(err)
	}
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", pth, err)
	} else {
		content = formattedContent
	}
	f.Write(content)

	fmt.Println("Starter template created at " + pth)
}

func findName(content string, index int, dirname string) string {
	r := regexp.MustCompile(`Name:\s+\"([^\"]+)\"\,`)
	m := r.FindAllStringSubmatch(content[index:], 1)
	name := m[0][1]
	return strings.TrimPrefix(name, "aws_"+dirname+"_")
}

func findExtraColumns(content string) ([]string, []string) {
	cols := []string{}
	skipFields := []string{}
	blocks := findBlocks(content)
	for _, block := range blocks {
		block = removeDescription(block)

		// put quotes around Resolver
		r := regexp.MustCompile(`Resolver:\s+([^\,]+)\,`)
		m := r.FindAllStringSubmatchIndex(block, -1)
		if len(m) > 0 && len(m[0]) > 2 {
			s := m[0][2]
			e := m[0][3]
			block = block[:s] + "`" + block[s:e] + "`" + block[e:]
		}

		if strings.Contains(block, "CreationOptions") {
			block = strings.ReplaceAll(block, "CreationOptions:", "Options:")
			cols = append(cols, block)
			if !strings.Contains(block, "Resolver:") {
				skipFields = append(skipFields, strcase.ToPascal(findName(block, 0, "")))
			} else if strings.Contains(block, "PathResolver(") {
				r := regexp.MustCompile(`PathResolver\(\"([^\"]+)\"`)
				name := r.FindAllStringSubmatch(block, 1)
				skipFields = append(skipFields, name[0][1])
			}
			continue
		}
		if strings.Contains(block, "client.ResolveAWSRegion") {
			continue
		}
		if strings.Contains(block, "client.ResolveAWSAccount") {
			continue
		}
		if strings.Contains(block, "PathResolver(") {
			continue
		}
		if strings.Contains(block, "Resolver:") {
			block = strings.ReplaceAll(block, "_cq_id", "_arn")
			block = strings.ReplaceAll(block, "schema.TypeUUID", "schema.TypeString")
			block = strings.ReplaceAll(block, "schema.ParentIdResolver", `schema.ParentResourceFieldResolver("arn")`)
			if strings.Contains(block, "client.ResolveTags") {
				skipFields = append(skipFields, "Tags")
			}
			cols = append(cols, block)
			continue
		}
	}
	return cols, skipFields
}

func removeDescription(b string) string {
	lines := strings.Split(b, "\n")
	final := []string{}
	for _, l := range lines {
		if strings.Contains(l, "Description:") {
			continue
		}
		final = append(final, l)
	}
	return strings.Join(final, "\n")
}

func findBlocks(content string) []string {
	blocks := []string{}
	index := 0
	for {
		s := findNextBracket(content, index)
		if s == -1 {
			return blocks
		}
		e := findMatchingBracket(content, s)
		blocks = append(blocks, content[s:e+1])
		index = e + 1
	}
}

func findNextBracket(content string, index int) int {
	i := strings.Index(content[index:], "{")
	if i == -1 {
		return -1
	}
	return i + index
}

func findMatchingBracket(content string, bracketIndex int) int {
	stack := 0
	index := bracketIndex
	for {
		nextOpening := strings.Index(content[index+1:], "{")
		nextClosing := strings.Index(content[index+1:], "}")
		if nextOpening == -1 || nextClosing < nextOpening {
			if stack == 0 {
				return nextClosing + index + 1
			}
			stack--
			index = nextClosing + index + 1
		} else {
			stack++
			index = nextOpening + index + 1
		}
	}
}

func calcProgress() {
	items, _ := os.ReadDir(".")
	migratedDirs := 0
	migratedFiles := 0
	unmigratedFiles := 0
	for _, it := range items {
		if it.IsDir() {
			dirItems, _ := os.ReadDir(it.Name())
			migrated := false
			n := 0
			for _, dit := range dirItems {
				if strings.HasSuffix(dit.Name(), ".go") && !strings.HasSuffix(dit.Name(), "_test.go") {
					n++
					if strings.HasSuffix(dit.Name(), "_fetch.go") {
						if !migrated {
							migratedDirs++
						}
						migratedFiles++
						migrated = true
					}
				}
			}
			if !migrated {
				unmigratedFiles += n
			}
		}
	}
	fmt.Println("==============================")
	fmt.Printf("Directories migrated: %d/%d\n", migratedDirs, len(items))
	fmt.Printf("Files left to migrate: %d/%d\n", unmigratedFiles, unmigratedFiles+migratedFiles)
	fmt.Println("==============================")
}

func main() {
	flag.StringVar(&directory, "d", "", "directory to migrate")
	flag.BoolVar(&createStarterTemplate, "s", false, "whether to create starter template")
	flag.Parse()

	if directory == "" {
		fmt.Fprintln(os.Stderr, "Usage: -d <directory_to_migrate>")
		return
	}
	resources := migrate(directory)
	if createStarterTemplate {
		writeStarterTemplate(directory, resources)
	}

	calcProgress()
}
