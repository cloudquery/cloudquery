package codegen

import (
	"github.com/cloudquery/plugin-sdk/codegen"
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
	// MockStruct that will be used to generate the mock
	IsStructPointer bool
	MockStruct      interface{}
	// ParentStruct that will be used to generate the mock
	ParentStruct interface{}
	// ParentStructName that will be used to generate the mock
	IsParentPointer  bool
	Args             string
	ParentStructName string
	// MockWrapper
	MockWrapper bool
	// MockStructName is the name of the Struct because it can't be inferred by reflection
	MockStructName string
	// StructName is the name of the Struct because it can't be inferred by reflection
	StructName string
	// Service is the name of the gcp service the struct/api is residing
	Service string
	// SubService is the name of the subservice
	SubService string
	// SubService is the name of the subservice
	SubServiceName string
	// NewFunction function
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
	// RequestStructFields is the snippet that fills in in the request struct
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
	// SkipMock is used to skip the mock generation for this resource
	SkipMock bool
	// Pass to MockTemplate
	MockTemplate string
	// MockFieldName is a name of a struct that is used for mocking
	ResponsePath string
	// ProtobufImport path to protobuf struct for this resource/service
	ProtobufImport string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	// Columns override, override generated columns
	OverrideColumns []codegen.ColumnDefinition
	FunctionName    string
	SkipFetch       bool
}
