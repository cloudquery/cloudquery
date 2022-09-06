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

type providerResource struct {
	ServiceName string
	Function    string
}

type provider struct {
	Packages  []string
	Resources []providerResource
}

func main() {
	filter := ""
	if len(os.Args) > 1 {
		filter = os.Args[1]
	}
	pattern := regexp.MustCompile("(?i)" + filter)
	resourcesDir := path.Join(path.Dir(getFilename()), "../resources")
	providerDir := path.Join(resourcesDir, "provider")
	servicesDir := path.Join(resourcesDir, "services")

	providerResources := make(map[string]providerResource)
	providerPackages := make(map[string]bool)

	for _, r := range codegen.AllResources() {
		if pattern.MatchString(fmt.Sprintf("%s/%s", r.AzureService, r.AzureSubService)) {
			fmt.Printf("Generating %s\n", r.Template.Destination)
			generateResource(servicesDir, r)
			packageName := strings.ToLower(r.AzureService)
			serviceName := fmt.Sprintf("%s.%s", packageName, strcase.ToSnake(r.AzureSubService))
			function := fmt.Sprintf("%s.%s", packageName, r.AzureSubService)
			providerResources[serviceName] = providerResource{
				ServiceName: serviceName,
				Function:    function,
			}
			providerPackages[packageName] = true
		}
	}

	sortedPackages := maps.Keys(providerPackages)
	sort.Strings(sortedPackages)

	sortedResources := maps.Values(providerResources)
	sort.SliceStable(sortedResources, func(i, j int) bool { return sortedResources[i].ServiceName < sortedResources[j].ServiceName })

	generateProvider(providerDir, provider{Packages: sortedPackages, Resources: sortedResources})

	exec.Command("goimports", "-w", resourcesDir).Run()
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
	tpl, err := template.New(templateName).Funcs(template.FuncMap{"ToLower": strings.ToLower, "ToSingular": pluralizeClient.Singular}).ParseFS(azureTemplatesFS, "templates/*.go.tpl")
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

func generateProvider(providerDir string, p provider) {
	destination := path.Join(providerDir, "provider.go")
	content := getContent("provider.go.tpl", destination, p)
	writeContent(destination, content)
}
