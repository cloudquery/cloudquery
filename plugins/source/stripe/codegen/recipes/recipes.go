package recipes

import (
	"path"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/gertd/go-pluralize"
)

type Resource struct {
	// DataStruct that will be used to generate the cloudquery table
	DataStruct any
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	PKColumns  []string

	Description string // optional, auto generated from struct name if not provided
	SkipMocks   bool

	Service   string // optional
	TableName string // optional, without "$plugin_" prefix

	Plugin     string // name of plugin, auto generated
	StructName string // name of struct, auto generated
}

var (
	AllResources []*Resource

	pluralizeClient *pluralize.Client
	csr             *caser.Caser
)

func init() {
	pluralizeClient = pluralize.NewClient()
	csr = caser.New()
}

func (r *Resource) Infer() {
	r.Plugin = path.Base(strings.TrimSuffix(reflect.TypeOf(r).Elem().PkgPath(), "/codegen/recipes")) // "stripe"
	r.SkipFields = append(r.SkipFields, "ID", "APIResource")

	ds := reflect.TypeOf(r.DataStruct)
	if ds.Kind() == reflect.Ptr {
		ds = ds.Elem()
	}
	r.StructName = ds.Name()

	if r.TableName == "" {
		r.TableName = pluralizeClient.Plural(csr.ToSnake(ds.Name()))
	}
	if r.Service == "" {
		r.Service = pluralizeClient.Plural(csr.ToSnake(ds.Name()))
	}
	if r.Description == "" {
		r.Description = "https://stripe.com/docs/api/" + pluralizeClient.Plural(csr.ToSnake(ds.Name()))
	}
}
