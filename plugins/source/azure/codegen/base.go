package codegen

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

type Template struct {
	Source      string
	Destination string
}

type Resource struct {
	// DefaultColumns columns that will be appended to the main table
	DefaultColumns []codegen.ColumnDefinition
	// Table is the table definition that will be used to generate the CloudQuery table
	Table *codegen.TableDefinition
	// AzureStruct that will be used to generate the CloudQuery table
	AzureStruct          interface{}
	AzureStructName      string
	AzurePackageName     string
	AzureService         string
	AzureSubService      string
	Imports              []string
	SkipFields           []string
	CreateTableOptions   schema.TableCreationOptions
	Template             Template
	TemplateParams       []string
	ListFunction         string
	ListFunctionArgs     []string
	ListHandler          string
	MockListResult       string
	MockListFunctionArgs []string
}

type template struct {
	source            string
	destinationSuffix string
	imports           []string
}

type resourceDefinition struct {
	azureStruct          interface{}
	listFunction         string
	listFunctionArgs     []string
	listHandler          string
	mockListResult       string
	mockListFunctionArgs []string
}

type byTemplates struct {
	templates           []template
	definitions         []resourceDefinition
	serviceNameOverride string
}

const pluginName = "azure"

const valueHandler = `if err != nil {
	return errors.WithStack(err)
}
if response.Value == nil {
	return nil
}
res <- *response.Value
`

var SubscriptionIdColumn = codegen.ColumnDefinition{
	Name:     "subscription_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveAzureSubscription",
}

var IdColumn = codegen.ColumnDefinition{
	Name:     "id",
	Type:     schema.TypeString,
	Resolver: "schema.PathResolver(\"ID\")",
}

func initResourceTable(resource *Resource) {
	var err error
	resource.Table, err = codegen.NewTableFromStruct(fmt.Sprintf("%s_%s_%s", pluginName, resource.AzurePackageName, strcase.ToSnake(resource.AzureSubService)), resource.AzureStruct, codegen.WithSkipFields(resource.SkipFields...))
	if err != nil {
		log.Fatal(err)
	}
	resource.Table.Columns = append(resource.DefaultColumns, resource.Table.Columns...)
	resource.Table.Multiplex = "client.SubscriptionMultiplex"
	resource.Table.Resolver = "fetch" + resource.AzureService + resource.AzureSubService
	resource.Table.Options.PrimaryKeys = resource.CreateTableOptions.PrimaryKeys
}

func generateResources(resourcesByTemplates []byTemplates) []Resource {
	plural := pluralize.NewClient()
	allResources := []Resource{}

	for _, byTemplate := range resourcesByTemplates {
		templates := byTemplate.templates
		definitions := byTemplate.definitions

		for _, template := range templates {
			for _, definition := range definitions {
				elementTypeParts := strings.Split(reflect.TypeOf(definition.azureStruct).Elem().String(), ".")
				azurePackageName, azureStructName := elementTypeParts[0], elementTypeParts[1]
				azureService := strcase.ToCamel(azurePackageName)
				if byTemplate.serviceNameOverride != "" {
					azureService = byTemplate.serviceNameOverride
				}
				resource := Resource{
					AzurePackageName:   azurePackageName,
					AzureStructName:    azureStructName,
					AzureStruct:        definition.azureStruct,
					AzureService:       azureService,
					AzureSubService:    plural.Plural(azureStructName),
					DefaultColumns:     []codegen.ColumnDefinition{SubscriptionIdColumn, IdColumn},
					SkipFields:         []string{"ID"},
					Imports:            template.imports,
					CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
					Template: Template{
						Source:      template.source,
						Destination: path.Join(azurePackageName, strcase.ToSnake(azureStructName)+template.destinationSuffix),
					},
					ListFunction:         definition.listFunction,
					ListHandler:          definition.listHandler,
					ListFunctionArgs:     definition.listFunctionArgs,
					MockListResult:       definition.mockListResult,
					MockListFunctionArgs: definition.mockListFunctionArgs,
				}
				initResourceTable(&resource)
				allResources = append(allResources, resource)
			}
		}
	}

	return allResources
}
