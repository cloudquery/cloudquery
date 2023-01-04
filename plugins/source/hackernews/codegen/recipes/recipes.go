package recipes

import (
	"reflect"

	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/gertd/go-pluralize"
)

type Resource struct {
	Service     string
	Description string
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// DataStruct that will be used to generate the cloudquery table
	DataStruct any
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields    []string
	Template      string
	Multiplex     string
	IsIncremental bool

	ExtraColumns []codegen.ColumnDefinition
	PKColumns    []string

	PreResourceResolver   string
	PostResourceResolver  string
	Relations             []string
	UnwrapEmbeddedStructs bool

	// These are inferred with reflection but can be overridden
	SubService string // Inferred from DataStruct name, singular
	TableName  string // singular Service + plural SubService

	// These are auto calculated
	ImportClient     bool   // true if the resource/column resolvers use the client package
	Filename         string // Calculated from TableName
	TableFuncName    string // Calculated from TableName
	ResolverFuncName string // Calculated from TableFuncName

	//used for generating better table names
	parent   *Resource
	children []*Resource
}

var (
	pluralizeClient *pluralize.Client
	csr             *caser.Caser
)

func init() {
	pluralizeClient = pluralize.NewClient()
	csr = caser.New()
}

func (r *Resource) Infer() {
	// Set defaults and/or infer fields
	if r.Template == "" {
		r.Template = "resource"
	}

	ds := reflect.TypeOf(r.DataStruct)
	if ds.Kind() == reflect.Ptr {
		ds = ds.Elem()
	}
	if r.SubService == "" {
		r.SubService = csr.ToSnake(pluralizeClient.Singular(ds.Name()))
	}
}

func (r *Resource) GenerateNames() {
	r.Filename = csr.ToSnake(r.TableName) + ".go"
	r.TableFuncName = csr.ToPascal(r.TableName)
	r.ResolverFuncName = "fetch" + r.TableFuncName
}
