package main

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
)

var currentFilename string
var currentDir string

const azureDir = "azure-sdk-for-go/sdk/resourcemanager"

//go:embed templates/*.go.tpl
var templateFS embed.FS

type ServicePackage struct {
	Name string
	Path string
	Resources []string
}

func generateRecipes(s ServicePackage) {
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
	
		filePath := path.Join(currentDir, "../codegen1/recipes", s.Name+".go")
		if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
			log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
		}
}

func generateServices(s []ServicePackage) {
	tpl, err := template.New("services.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strings.Title,
	}).ParseFS(templateFS, "templates/services.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse recipe.go.tpl: %w", err))
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, s); err != nil {
		log.Fatal(fmt.Errorf("failed to execute recipe template: %w", err))
	}

	filePath := path.Join(currentDir, "../codegen1/recipes/services.go")
	if err := os.WriteFile(filePath, buff.Bytes(), 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", filePath, err))
	}
}

func getARMDirs() ([]string, error) {
	// List all files in the directory
	var res []string
	files, err := os.ReadDir(path.Join(currentDir, azureDir))
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			packagePath := f.Name() + "/arm" + strings.ReplaceAll(f.Name(), "-", "")
			fileinfo, err := os.Stat(path.Join(currentDir, azureDir, packagePath))
			if err != nil {
				if os.IsNotExist(err) {
					log.Printf("package %s does not exist\n", packagePath)
					continue
				}
				return nil, err
			}
			if fileinfo.IsDir() {
				res = append(res, packagePath)
			}
		}
	}
	return res, nil
}


func main() {
	var ok bool
	_, currentFilename, _, ok = runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	currentDir = path.Dir(currentFilename)
	
	res, err := getARMDirs()
	if err != nil {
		log.Fatal(err)
	}

	var services []ServicePackage
	for _, subPackage := range res {
		set := token.NewFileSet()
		pkgPath := path.Join(currentDir, azureDir, subPackage)
		packs, err := parser.ParseDir(set, pkgPath, nil, 0)
		if err != nil {
				fmt.Println("Failed to parse package:", err)
				os.Exit(1)
		}
		service := ServicePackage{
			Name: strings.Split(subPackage, "/")[1],
			Path: subPackage,
		}
		for _, pack := range packs {
				for _, f := range pack.Files {
						for _, d := range f.Decls {
								if fn, isFn := d.(*ast.FuncDecl); isFn {
										if strings.HasPrefix(fn.Name.Name, "New") && strings.HasSuffix(fn.Name.Name, "Client") {
											service.Resources = append(service.Resources, fn.Name.Name)
										}
								}
						}
				}
		}
		fmt.Println(subPackage)
		fmt.Printf("all funcs: %+v\n", service.Resources)
		generateRecipes(service)
		services = append(services, service)
	}
	generateServices(services)
}