package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

// create cobra subcommand
type goPackage struct {
	Mod module.Version
	NewFuncs []string
	BaseName string
}

const (
	pkgGoDevURL = "https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk"
	cacheDir = "/Users/yevgenyp/go/pkg/mod"
)

//go:embed templates/*.go.tpl
var templateFS embed.FS

var (
	currentFilename string
	currentDir string
)

func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	
	res, err := findAllAzureSdkSubPackages()
	if err != nil {
		log.Fatal(err)
	}
	for _, subPackage := range res {
		set := token.NewFileSet()
		pkgPath := path.Join(cacheDir, subPackage.Mod.String())
		// thats because azure had to be special with uppercase
		pkgPath = strings.Replace(pkgPath, "A", "!A", 1)
		packs, err := parser.ParseDir(set, pkgPath, nil, 0)
		if err != nil {
				os.Exit(1)
		}
		for _, pack := range packs {
				for _, f := range pack.Files {
						for _, d := range f.Decls {
								if fn, isFn := d.(*ast.FuncDecl); isFn {
										if strings.HasPrefix(fn.Name.Name, "New") && strings.HasSuffix(fn.Name.Name, "Client") {
											subPackage.NewFuncs = append(subPackage.NewFuncs, fn.Name.Name)
										}
								}
						}
				}
		}
		generateRecipes(subPackage)
	}
}



func generateRecipes(s goPackage) {
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
	baseName := path.Base(s.Mod.Path)
	filePath := path.Join(currentDir, "../codegen1/recipes", baseName+".go")
	if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}


func findAllAzureSdkSubPackages() ([]goPackage, error) {
	// Open the go.mod file.
	var packages []goPackage
	content, err := os.ReadFile(path.Join(currentDir, "../go.mod"))
	if err != nil {
			return nil, err
	}
	// Parse the go.mod file.
	// modfile
	mod, err := modfile.Parse("go.mod", content, nil)
	if err != nil {
			return nil, err
	}

	for _, req := range mod.Require {
		if strings.HasPrefix(req.Mod.Path, "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager") {
			packages = append(packages, goPackage{
				Mod: req.Mod,
				BaseName: path.Base(req.Mod.Path),
			})
		}
}

return packages, nil
}

