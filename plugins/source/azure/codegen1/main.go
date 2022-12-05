package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen1/recipes"
)

var (
	currentFilename string
	currentDir string
)

var responseStructRe = regexp.MustCompile("\\.([a-zA-Z]+)\\]")

//go:embed templates/*.go.tpl
var templateFS embed.FS

type Table struct {
	Name string
	Struct string
	ResponseStruct string
	Client string
	ListFunc string
	NewFunc string
	URL string
}

type Recipe struct {
	PkgPath string
	BaseImport string
	Tables []Table
}

func generateRecipes(s Recipe) {
	tpl, err := template.New("recipe.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strings.Title,
	}).ParseFS(templateFS, "templates/recipe.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipe.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, s); err != nil {
		log.Fatal(fmt.Errorf("failed to execute recipe template: %w", err))
	}
	filePath := path.Join(currentDir, "../codegen2/recipes", s.BaseImport+".go")
	if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	
	// var rr []Recipe
	for _, r := range recipes.Resources {
		if len(r) == 0 {
			continue
		}
		var recipe Recipe
		// v := reflect.TypeOf(r[0].NewFunc)
		recipe.PkgPath = r[0].PkgPath
		recipe.BaseImport = path.Base(recipe.PkgPath)

		for _, table := range r {
			// var table Table
			v := reflect.TypeOf(table.NewFunc)
			results := v.Out(0)
			m, ok := results.MethodByName("NewListPager")
			if !ok {
				continue
			}
			responseStruct, ok := m.Type.Out(0).Elem().FieldByName("current")
			if !ok {
				panic("failed to find current field")
			}

			st, ok := responseStruct.Type.Elem().FieldByName("Value")
			if !ok {
				panic("failed to find Value field")
			}
			stName := strings.Split(st.Type.String(), ".")[1]

			responseStructName := strings.Split(responseStruct.Type.String(), ".")[1]
			clientName := strings.Split(v.Out(0).String(), ".")[1]
			recipe.Tables = append(recipe.Tables, Table{
				Name: strcase.ToSnake(stName),
				Struct: stName,
				ResponseStruct: responseStructName,
				Client: clientName,
				ListFunc: "NewListPager",
				NewFunc: "New" + clientName,
				URL: table.URL,
			})
		}

		if len(recipe.Tables) != 0 {
			generateRecipes(recipe)
		}
	}	
}