package codegen

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"regexp"
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
	GetFunction              string
	GetFunctionArgs          []string
	GetFunctionArgsInit      []string
	Helpers                  []string
	MockHelpers              []string
	MockListResult           string
	MockListFunctionArgs     []string
	MockListFunctionArgsInit []string
	MockFieldsToIgnore       []string
	MockValueType            string
	MockDefinitionType       string
}

type template struct {
	source            string
	destinationSuffix string
	imports           []string
}

type resourceDefinition struct {
	customColumns            codegen.ColumnDefinitions
	azureStruct              interface{}
	skipFields               []string
	includeColumns           string
	helpers                  []string
	listFunction             string
	listFunctionArgs         []string
	listFunctionArgsInit     []string
	listHandler              string
	getFunction              string
	getFunctionArgs          []string
	getFunctionArgsInit      []string
	mockListResult           string
	mockHelpers              []string
	mockListFunctionArgs     []string
	mockListFunctionArgsInit []string
	mockFieldsToIgnore       []string
	mockValueType            string
	mockDefinitionType       string
	subServiceOverride       string
	relations                []string
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
	resources = append(resources, Datalake()...)
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
	resources = append(resources, Redis()...)
	resources = append(resources, Resources()...)
	resources = append(resources, Search()...)
	resources = append(resources, Security()...)
	resources = append(resources, ServiceBus()...)
	resources = append(resources, SQL()...)
	resources = append(resources, Storage()...)
	resources = append(resources, StreamAnalytics()...)
	resources = append(resources, Subscriptions()...)
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
				table, err := codegen.NewTableFromStruct(fmt.Sprintf("%s_%s_%s", pluginName, strings.ToLower(azureService), strcase.ToSnake(azureSubService)), definition.azureStruct, codegen.WithSkipFields(skipFields))
				if err != nil {
					log.Fatal(err)
				}

				defaultColumns := []codegen.ColumnDefinition{}
				if needsSubscriptionId(table) {
					defaultColumns = append(defaultColumns, SubscriptionIdColumn)
				}

				table.Columns = append(defaultColumns, table.Columns...)
				table.Columns = append(table.Columns, definition.customColumns...)

				if definition.includeColumns != "" {
					regex := regexp.MustCompile(definition.includeColumns)
					newColumns := make(codegen.ColumnDefinitions, 0)
					for _, column := range table.Columns {
						if regex.MatchString(column.Name) {
							newColumns = append(newColumns, column)
						}
					}
					table.Columns = newColumns
				}

				table.Multiplex = "client.SubscriptionMultiplex"
				table.Resolver = "fetch" + azureService + azureSubService
				table.Options.PrimaryKeys = []string{"id"}
				table.Relations = definition.relations

				if definition.getFunction != "" {
					table.PreResourceResolver = "get" + azureStructName
				}

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
					Helpers:                  definition.helpers,
					ListFunction:             definition.listFunction,
					ListHandler:              definition.listHandler,
					ListFunctionArgs:         definition.listFunctionArgs,
					ListFunctionArgsInit:     definition.listFunctionArgsInit,
					GetFunction:              definition.getFunction,
					GetFunctionArgs:          definition.getFunctionArgs,
					GetFunctionArgsInit:      definition.getFunctionArgsInit,
					MockHelpers:              definition.mockHelpers,
					MockListResult:           definition.mockListResult,
					MockListFunctionArgs:     definition.mockListFunctionArgs,
					MockListFunctionArgsInit: definition.mockListFunctionArgsInit,
					MockFieldsToIgnore:       append(skipFields, definition.mockFieldsToIgnore...),
					MockValueType:            definition.mockValueType,
					MockDefinitionType:       definition.mockDefinitionType,
				}
				allResources = append(allResources, resource)
			}
		}
	}

	return allResources
}
