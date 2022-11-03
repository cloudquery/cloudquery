package recipes

import (
	"fmt"
	"log"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gertd/go-pluralize"
	"github.com/gofrs/uuid"
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
	MockGetFunctionArgs      []string
	MockRelations            []string
	IsRelation               bool
}

type template struct {
	source            string
	destinationSuffix string
	imports           []string
}

type resourceDefinition struct {
	customColumns            codegen.ColumnDefinitions
	tableName                string
	azureStruct              interface{}
	skipFields               []string
	includeColumns           string
	helpers                  []string
	parent                   string
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
	mockGetFunctionArgs      []string
	subServiceOverride       string
	relations                []resourceDefinition
	singleSubscription       bool
}

type byTemplates struct {
	templates           []template
	definitions         []resourceDefinition
	serviceNameOverride string
}

const (
	pluginName   = "azure"
	valueHandler = `if err != nil {
		return err
	}
	if response.Value == nil {
		return nil
	}
	res <- *response.Value
	`
	mockDirectResponse = `CQ_CODEGEN_DIRECT_RESPONSE`
)

var (
	subscriptionIdColumn = codegen.ColumnDefinition{
		Name:     "subscription_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAzureSubscription",
	}
	defaultSkipFields = []string{"Response", "SubscriptionID"}
	azureCaser        = caser.New(caser.WithCustomInitialisms(map[string]bool{
		"V2": true,
	}))
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

func parseAzureStruct(serviceNameOverride string, definition resourceDefinition) (string, string, string, string) {
	plural := pluralize.NewClient()
	elementTypeParts := strings.Split(reflect.TypeOf(definition.azureStruct).Elem().String(), ".")
	azurePackageName, azureStructName := elementTypeParts[0], elementTypeParts[1]
	azureService := strcase.ToCamel(azurePackageName)
	if serviceNameOverride != "" {
		azureService = serviceNameOverride
	}
	azureSubService := plural.Plural(azureStructName)
	if definition.subServiceOverride != "" {
		azureSubService = definition.subServiceOverride
	}

	return azurePackageName, azureStructName, azureService, azureSubService
}

func initColumns(table *codegen.TableDefinition, definition resourceDefinition) codegen.ColumnDefinitions {
	columns := []codegen.ColumnDefinition{}
	if !definition.singleSubscription {
		// add subscription id if we are not in single subscription mode
		columns = append(columns, subscriptionIdColumn)
	}
	if definition.parent != "" {
		columns = append(columns, codegen.ColumnDefinition{
			Name:     definition.parent,
			Type:     schema.TypeString,
			Resolver: `schema.ParentColumnResolver("id")`,
		})
	}

	columns = append(columns, table.Columns...)
	columns = append(columns, definition.customColumns...)

	for i := range columns {
		if columns[i].Name == "id" || columns[i].Name == "kid" {
			columns[i].Options.PrimaryKey = true
		}
	}

	return columns
}

func getTableName(azureService, azureSubService string, override string) string {
	if override != "" {
		return fmt.Sprintf("%s_%s", pluginName, override)
	}

	return fmt.Sprintf("%s_%s_%s", pluginName, strings.ToLower(azureService), azureCaser.ToSnake(azureSubService))
}

func timeStampTransformer(field reflect.StructField) (schema.ValueType, error) {
	dateTime := date.Time{}
	uuid := uuid.UUID{}
	switch field.Type {
	case reflect.TypeOf(dateTime), reflect.TypeOf(&dateTime):
		return schema.TypeTimestamp, nil
	case reflect.TypeOf(uuid), reflect.TypeOf(&uuid):
		return schema.TypeUUID, nil
	}
	return schema.TypeInvalid, nil
}

func initTable(serviceNameOverride string, definition resourceDefinition, azureService string, azureSubService string, azureStructName string) *codegen.TableDefinition {
	skipFields := append(definition.skipFields, defaultSkipFields...)
	table, err := codegen.NewTableFromStruct(
		getTableName(azureService, azureSubService, definition.tableName),
		definition.azureStruct,
		codegen.WithSkipFields(skipFields),
		codegen.WithUnwrapAllEmbeddedStructs(),                 // Unwrap all embedded structs otherwise all resources will just have `Id, Type, Name, Location, Tags` columns
		codegen.WithUnwrapStructFields([]string{"Properties"}), // Some resources have a `Properties` field which contains the actual resource properties instead of an embedded struct
		codegen.WithTypeTransformer(timeStampTransformer),
	)
	if err != nil {
		log.Fatal(err)
	}
	table.Columns = initColumns(table, definition)

	if definition.parent == "" {
		table.Multiplex = "client.SubscriptionMultiplex"
	}

	if definition.singleSubscription {
		table.Multiplex = "client.SingleSubscriptionMultiplex"
	}

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

	table.Resolver = "fetch" + azureService + azureSubService
	table.Relations = make([]string, 0)
	if definition.relations != nil {
		for _, relation := range definition.relations {
			_, _, _, azureSubService := parseAzureStruct(serviceNameOverride, relation)
			table.Relations = append(table.Relations, strcase.ToLowerCamel(azureSubService)+"()")
		}
	}

	if definition.getFunction != "" {
		table.PreResourceResolver = "get" + azureStructName
	}

	return table
}

func getMockRelations(serviceNameOverride string, azureService string, definition resourceDefinition) []string {
	mockRelations := make([]string, 0)
	if definition.relations != nil {
		for _, relation := range definition.relations {
			_, _, _, azureSubService := parseAzureStruct(serviceNameOverride, relation)
			mockRelations = append(mockRelations, fmt.Sprintf("%s: create%sMock(t,ctrl).%s.%s", azureSubService, azureSubService, azureService, azureSubService))
			mockRelations = append(mockRelations, getMockRelations(serviceNameOverride, azureService, relation)...)
		}
	}
	return mockRelations
}

func generateResources(resourcesByTemplates []byTemplates) []Resource {
	allResources := []Resource{}

	for _, byTemplate := range resourcesByTemplates {
		templates := byTemplate.templates
		definitions := byTemplate.definitions

		for _, template := range templates {
			for _, definition := range definitions {
				azurePackageName, azureStructName, azureService, azureSubService := parseAzureStruct(byTemplate.serviceNameOverride, definition)
				table := initTable(byTemplate.serviceNameOverride, definition, azureService, azureSubService, azureStructName)

				mockRelations := make([]string, 0)
				if definition.parent == "" {
					mockRelations = getMockRelations(byTemplate.serviceNameOverride, azureService, definition)
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
						Destination: path.Join(strings.ToLower(azureService), azureCaser.ToSnake(azureSubService)+template.destinationSuffix),
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
					MockFieldsToIgnore:       append(append(defaultSkipFields, definition.skipFields...), definition.mockFieldsToIgnore...),
					MockValueType:            definition.mockValueType,
					MockDefinitionType:       definition.mockDefinitionType,
					MockGetFunctionArgs:      definition.mockGetFunctionArgs,
					IsRelation:               definition.parent != "",
					MockRelations:            mockRelations,
				}
				allResources = append(allResources, resource)
			}
		}
	}

	return allResources
}

func initParentsForResources(serviceNameOverride string, resources []resourceDefinition) {
	plural := pluralize.NewClient()
	for i := range resources {
		parent := resources[i]
		relations := parent.relations
		if relations != nil {
			_, _, azureService, azureSubService := parseAzureStruct(serviceNameOverride, parent)
			parentColumnName := fmt.Sprintf("%s_%s_id", strings.ToLower(azureService), azureCaser.ToSnake(plural.Singular(azureSubService)))
			for j := range relations {
				relations[j].parent = parentColumnName
			}
			initParentsForResources(serviceNameOverride, relations)
		}
	}
}

func initParents(resourcesByTemplates []byTemplates) {
	for i := range resourcesByTemplates {
		initParentsForResources(resourcesByTemplates[i].serviceNameOverride, resourcesByTemplates[i].definitions)
	}
}
