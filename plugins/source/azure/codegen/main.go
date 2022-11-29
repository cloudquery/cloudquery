package main

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"text/template"

	codegen "github.com/cloudquery/cloudquery/plugins/source/azure/codegen/recipes"
	sdkgen "github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
)

//go:embed templates/*.go.tpl
var azureTemplatesFS embed.FS

type plugin struct {
	Packages  []string
	Resources []string
}

func main() {
	filter := ""
	if len(os.Args) > 1 {
		filter = os.Args[1]
	}
	pattern := regexp.MustCompile("(?i)" + filter)
	resourcesDir := path.Join(path.Dir(getFilename()), "../resources")
	pluginDir := path.Join(resourcesDir, "plugin")
	servicesDir := path.Join(resourcesDir, "services")

	pluginResources := make(map[string]bool)
	pluginPackages := make(map[string]bool)

	for _, r := range codegen.AllResources() {
		if pattern.MatchString(fmt.Sprintf("%s/%s", r.AzureService, r.AzureSubService)) {
			fmt.Printf("Generating %s\n", r.Template.Destination)
			generateResource(servicesDir, r)
			packageName := strings.ToLower(r.AzureService)
			function := fmt.Sprintf("%s.%s", packageName, r.AzureSubService)

			// Only add top level resources to provider.go
			if !r.IsRelation {
				pluginResources[function] = true
				pluginPackages[packageName] = true
			}
		}
	}

	sortedPackages := maps.Keys(pluginPackages)
	sort.Strings(sortedPackages)

	sortedResources := maps.Keys(pluginResources)
	sort.Strings(sortedResources)

	generatePlugin(pluginDir, plugin{Packages: sortedPackages, Resources: sortedResources})

	err := exec.Command("goimports", "-w", resourcesDir).Run()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to run goimports: %w", err))
	}
}

func getFilename() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get caller information")
	}
	return filename
}

func initTemplate(templateName string) *template.Template {
	pluralizeClient := pluralize.NewClient()
	tpl, err := template.New(templateName).Funcs(template.FuncMap{
		"ToLower":      strings.ToLower,
		"ToSingular":   pluralizeClient.Singular,
		"ToLowerCamel": strcase.ToLowerCamel,
		"TrimEnd":      func(s string, count int) string { return s[:len(s)-count] },
		"ToCamel":      strcase.ToCamel,
	}).ParseFS(azureTemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse azure templates: %w", err))
	}

	tpl, err = tpl.ParseFS(sdkgen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse codegen template: %w", err))
	}

	return tpl
}

func getContent(t string, destination string, r any) []byte {
	tpl := initTemplate(t)
	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		log.Fatal(fmt.Errorf("failed to execute template: %w", err))
	}
	return buff.Bytes()
}

func writeContent(destination string, content []byte) {
	if err := os.MkdirAll(path.Dir(destination), 0755); err != nil {
		log.Fatal(fmt.Errorf("failed to create directory for file %s: %w", destination, err))
	}
	if err := os.WriteFile(destination, content, 0644); err != nil {
		log.Fatal(fmt.Errorf("failed to write file %s: %w", destination, err))
	}
}

func generateResource(destinationDir string, r codegen.Resource) {
	destination := path.Join(destinationDir, r.Template.Destination)
	content := getContent(r.Template.Source, destination, r)
	writeContent(destination, content)
}

func generatePlugin(pluginDir string, p plugin) {
	destination := path.Join(pluginDir, "plugin.go")
	content := getContent("plugin.go.tpl", destination, p)
	writeContent(destination, content)
}
