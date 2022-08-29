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
	// Table is the table definition that will be used to generate the CloudQuery table
	Table                    *codegen.TableDefinition
	AzureStructName          string
	AzurePackageName         string
	AzureService             string
	AzureSubService          string
	Imports                  []string
	Template                 Template
	ListFunction             string
	ListFunctionArgs         []string
	ListFunctionArgsInit     []string
	ListHandler              string
	MockHelpers              []string
	MockListResult           string
	MockListFunctionArgs     []string
	MockListFunctionArgsInit []string
	MockFieldsToIgnore       []string
}

type template struct {
	source            string
	destinationSuffix string
	imports           []string
}

type resourceDefinition struct {
	azureStruct              interface{}
	skipFields               []string
	listFunction             string
	listFunctionArgs         []string
	listFunctionArgsInit     []string
	listHandler              string
	mockListResult           string
	mockHelpers              []string
	mockListFunctionArgs     []string
	mockListFunctionArgsInit []string
	mockFieldsToIgnore       []string
	subServiceOverride       string
}

type byTemplates struct {
	templates           []template
	definitions         []resourceDefinition
	serviceNameOverride string
}

const (
	pluginName   = "azure"
	valueHandler = `if err != nil {
		return errors.WithStack(err)
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	`
)

var (
	SubscriptionIdColumn = codegen.ColumnDefinition{
		Name:     "subscription_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAzureSubscription",
	}
)

func AllResources() []Resource {
	var resources = []Resource{}
	resources = append(resources, Authorization()...)
	resources = append(resources, Batch()...)
	resources = append(resources, CDN()...)
	resources = append(resources, Compute()...)
	resources = append(resources, Container()...)
	resources = append(resources, CosmosDB()...)
	resources = append(resources, EventHub()...)
	resources = append(resources, FrontDoor()...)
	resources = append(resources, IotHub()...)
	resources = append(resources, Network()...)
	resources = append(resources, KeyValue()...)
	resources = append(resources, Logic()...)
	resources = append(resources, MariaDB()...)
	resources = append(resources, Monitor()...)
	resources = append(resources, MySQL()...)
	resources = append(resources, PostgresSQL()...)
	resources = append(resources, Resources()...)
	resources = append(resources, Search()...)
	resources = append(resources, Security()...)
	resources = append(resources, ServiceBus()...)
	resources = append(resources, SQL()...)
	resources = append(resources, Storage()...)
	resources = append(resources, StreamAnalytics()...)
	resources = append(resources, Web()...)
	return resources
}

func needsSubscriptionId(table *codegen.TableDefinition) bool {
	for _, column := range table.Columns {
		if column.Name == "subscription_id" {
			return false
		}
	}
	return true
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
				azureSubService := plural.Plural(azureStructName)
				if definition.subServiceOverride != "" {
					azureSubService = definition.subServiceOverride
				}

				skipFields := append(definition.skipFields, "Response")
				table, err := codegen.NewTableFromStruct(fmt.Sprintf("%s_%s_%s", pluginName, azurePackageName, strcase.ToSnake(azureSubService)), definition.azureStruct, codegen.WithSkipFields(skipFields...))
				if err != nil {
					log.Fatal(err)
				}

				defaultColumns := []codegen.ColumnDefinition{}
				if needsSubscriptionId(table) {
					defaultColumns = append(defaultColumns, SubscriptionIdColumn)
				}

				table.Columns = append(defaultColumns, table.Columns...)
				table.Multiplex = "client.SubscriptionMultiplex"
				table.Resolver = "fetch" + azureService + azureSubService
				table.Options.PrimaryKeys = []string{"subscription_id", "id"}

				resource := Resource{
					Table:            table,
					AzurePackageName: azurePackageName,
					AzureStructName:  azureStructName,
					AzureService:     azureService,
					AzureSubService:  azureSubService,
					Imports:          template.imports,
					Template: Template{
						Source:      template.source,
						Destination: path.Join(strings.ToLower(azureService), strcase.ToSnake(azureSubService)+template.destinationSuffix),
					},
					ListFunction:             definition.listFunction,
					ListHandler:              definition.listHandler,
					ListFunctionArgs:         definition.listFunctionArgs,
					ListFunctionArgsInit:     definition.listFunctionArgsInit,
					MockHelpers:              definition.mockHelpers,
					MockListResult:           definition.mockListResult,
					MockListFunctionArgs:     definition.mockListFunctionArgs,
					MockListFunctionArgsInit: definition.mockListFunctionArgsInit,
					MockFieldsToIgnore:       append(skipFields, definition.mockFieldsToIgnore...),
				}
				allResources = append(allResources, resource)
			}
		}
	}

	return allResources
}
