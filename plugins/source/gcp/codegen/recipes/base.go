package recipes

import (
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

type Resource struct {
	// PackageName name is the packgename in the source plugin this resource is located
	PackageName string
	// Sets PreResourceResolver
	PreResourceResolver string
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// Struct that will be used to generate the cloudquery table
	Struct interface{}
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// Service is the name of the gcp service the struct/api is residing
	Service string
	// SubService is the name of the gcp subservice the struct/api is residing (gcp is split into service.subservice.list)
	SubService string
	// NewFunction
	NewFunction interface{}
	// NewFunctionName name of the above function via Reflection
	NewFunctionName string
	// ClientName name of the above client via Reflection
	ClientName string
	// ListFunction
	ListFunction interface{}
	// ListFunction name of the above function via Reflection
	ListFunctionName string
	// RequestStruct fills the request struct for google api
	RequestStruct interface{}
	// RequestStructName
	RequestStructName string
	// RequestStructFields is the snippet that fills in the request struct
	RequestStructFields string
	// ResponseStruct
	ResponseStruct interface{}
	// ResponseStructName is reflected name from the ResponseStruct
	ResponseStructName string
	// RegisterServer function to grpc register server
	RegisterServer interface{}
	// RegisterServerName is the name of the above function via Reflection
	RegisterServerName string
	// UnimplementedServer is the unimplemented server for the grpc server
	UnimplementedServer interface{}
	// UnimplementedServerName is the name of the above function via Reflection
	UnimplementedServerName string
	// OutputField is field where the result is located in the output struct
	OutputField string
	// Relations is list of relations functions
	Relations []string
	// Template is the template to use to generate the resource (some services has different template as some services were generated using different original codegen)
	Template string
	// imports to add for this resource
	Imports []string
	// Multiplex
	Multiplex *string
	// ChildTable
	ChildTable bool
	// FakerFieldsToIgnore is a list of fields to ignore when generating faker data
	FakerFieldsToIgnore []string
	// SkipMock is used to skip the mock generation for this resource
	SkipMock bool
	// Pass to MockTemplate
	MockTemplate string
	// MockPostFaker is a code snippet that runs post faker
	MockPostFaker string
	// MockListStruct is the name of the struct that will be used in the mock template
	MockListStruct string
	// MockImports imports used in mock tests
	MockImports []string
	// ProtobufImport path to protobuf struct for this resource/service
	ProtobufImport string
	// Don't generate fetch
	SkipFetch bool
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// ExtraColumns override, override generated columns
	ExtraColumns []codegen.ColumnDefinition
	// NameTransformer custom name transformer for resource
	NameTransformer func(field reflect.StructField) (string, error)
}

var ProjectIdColumn = codegen.ColumnDefinition{
	Name:     "project_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveProject",
}

var ProjectIdColumnPk = codegen.ColumnDefinition{
	Name:     "project_id",
	Type:     schema.TypeString,
	Resolver: "client.ResolveProject",
	Options:  schema.ColumnCreationOptions{PrimaryKey: true},
}

func CreateReplaceTransformer(replace map[string]string) func(field reflect.StructField) (string, error) {
	return func(field reflect.StructField) (string, error) {
		name, err := codegen.DefaultNameTransformer(field)
		if err != nil {
			return "", err
		}
		for k, v := range replace {
			name = strings.ReplaceAll(name, k, v)
		}
		return name, nil
	}
}
