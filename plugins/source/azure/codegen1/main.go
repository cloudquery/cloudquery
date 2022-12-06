package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen1/recipes"
)

var (
	currentFilename string
	currentDir string
)

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
	Tables []*Table
}


func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	
	// var rr []Recipe
	for _, tt := range recipes.Tables {
		if len(tt) == 0 {
			continue
		}
		var recipe Recipe
		// v := reflect.TypeOf(r[0].NewFunc)
		recipe.PkgPath = tt[0].PkgPath
		recipe.BaseImport = path.Base(recipe.PkgPath)

		for _, table := range tt {
			// var table Table
			v2Table, err := ConvertTableV1ToV2(table)
			if err != nil {
				log.Fatal(err)
			}
			if v2Table == nil {
				continue
			}
			recipe.Tables = append(recipe.Tables, v2Table)
		}

		if len(recipe.Tables) != 0 {
			if err := generateRecipes(recipe); err != nil {
				log.Fatal(err)
			}
		}
	}	
}

// this uses reflection to find the struct type inside Value field in an azure
// responseStruct returned from NewListPager and any other pager
func getStructNameFromResponseStruct(valueFieldType reflect.Type) (string, error) {
	var name string
	var err error

	switch valueFieldType.Kind() {
	case reflect.Struct:
		name = valueFieldType.Name()
	case reflect.Ptr:
		name, err = getStructNameFromResponseStruct(valueFieldType.Elem())
		if err != nil {
			return "", err
		}
	case reflect.Slice:
		name, err = getStructNameFromResponseStruct(valueFieldType.Elem())
		if err != nil {
			return "", err
		}
	case reflect.Interface:
		// this is currently unsupported so we skip this
		return "", nil
	default:
		return "", fmt.Errorf("failed to find struct name for %s", valueFieldType.String())
	}
	return name, nil
}

func ConvertTableV1ToV2(t *recipes.Table) (*Table, error) {
	v := reflect.TypeOf(t.NewFunc)
	clientType := v.Out(0)
	newListPagerMethod, ok := clientType.MethodByName("NewListPager")
	if !ok {
		return nil, fmt.Errorf("failed to find NewListPager method for %s", clientType.String())
	}
	responseStruct, ok := newListPagerMethod.Type.Out(0).Elem().FieldByName("current")
	if !ok {
		return nil, fmt.Errorf("failed to find current field for %s", newListPagerMethod.Type.Out(0).Elem().String())
	}

	st, ok := responseStruct.Type.Elem().FieldByName("Value")
	if !ok {
		return nil, fmt.Errorf("failed to find Value field for %s", responseStruct.Type.Elem().String())
	}
	
	structName, err := getStructNameFromResponseStruct(st.Type)
	if err != nil {
		return nil, err
	}
	if structName == "" {
		log.Printf("skipping %s as generating interface value fields is not supported yet", st.Type)
		return nil, nil
	}
	responseStructName := strings.Split(responseStruct.Type.String(), ".")[1]
	clientName := strings.Split(v.Out(0).String(), ".")[1]
	return &Table{
		Name: strcase.ToSnake(structName),
		Struct: structName,
		ResponseStruct: responseStructName,
		Client: clientName,
		ListFunc: "NewListPager",
		NewFunc: "New" + clientName,
		URL: t.URL,
	}, nil
}

func generateRecipes(s Recipe) error {
	tpl, err := template.New("recipe.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strings.Title,
	}).ParseFS(templateFS, "templates/recipe.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse recipe.go.tpl: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, s); err != nil {
		return fmt.Errorf("failed to execute recipe template: %w", err)
	}
	baseName := strings.TrimPrefix(s.BaseImport, "arm")
	filePath := path.Join(currentDir, "../codegen2/recipes", baseName+".go")
	if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}