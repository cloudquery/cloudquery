package recipes

import (
	"path"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/v4/caser"
	"github.com/gertd/go-pluralize"
)

type Resource struct {
	// DataStruct that will be used to generate the cloudquery table
	DataStruct any
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields    []string
	PKColumns     []string
	IgnoreInTests []string

	Description string // optional, auto generated from struct name if not provided
	SkipMocks   bool

	Service   string // optional
	TableName string // optional, without "$plugin_" prefix

	Plugin     string // name of plugin, auto generated
	StructName string // name of struct, auto generated

	Single        bool   // true if we're getting a single entity, false if we're getting a list/iterator
	FetchTemplate string // optional, if not provided will use default template, decided by Single

	HasIDPK bool // has "id" column as PK, auto generated, used in template

	Children      []*Resource
	Parent        *Resource // auto calculated from Children
	ExtraChildren []string

	ListParams     string   // optional
	ExpandFields   []string // optional, fields to expand in list call
	StateParamName string   // optional *int64 param name. If empty, no state will be persisted. Only if Single is not true.
}

var (
	pluralizeClient *pluralize.Client
	csr             *caser.Caser
)

func init() {
	pluralizeClient = pluralize.NewClient()
	csr = caser.New()
}

func (r *Resource) Infer(parent *Resource) {
	r.Parent = parent
	if parent != nil {
		r.Service = parent.Service
		r.SkipMocks = true
	}

	r.Plugin = path.Base(strings.TrimSuffix(reflect.TypeOf(r).Elem().PkgPath(), "/codegen/recipes")) // "stripe"
	r.SkipFields = append(r.SkipFields, "APIResource")

	ds := reflect.TypeOf(r.DataStruct)
	if ds.Kind() == reflect.Ptr {
		ds = ds.Elem()
	}
	if r.StructName == "" {
		r.StructName = ds.Name()
	}

	if len(r.PKColumns) == 0 {
		if _, ok := ds.FieldByName("ID"); ok {
			r.PKColumns = []string{"id"}
			r.SkipFields = append(r.SkipFields, "ID")
			r.HasIDPK = true
		}
	}

	var snakeNameBySingularity string
	if r.Single {
		snakeNameBySingularity = pluralizeClient.Singular(csr.ToSnake(ds.Name()))
	} else {
		snakeNameBySingularity = pluralizeClient.Plural(csr.ToSnake(ds.Name()))
	}

	if r.TableName == "" {
		r.TableName = snakeNameBySingularity
	}
	if r.Service == "" {
		r.Service = snakeNameBySingularity
	}
	if r.Description == "" {
		r.Description = "https://stripe.com/docs/api/" + snakeNameBySingularity
	}

	if r.FetchTemplate == "" {
		if r.Single {
			r.FetchTemplate = "get"
		} else {
			r.FetchTemplate = "list"
		}
	}

	for i := range r.Children {
		r.Children[i].Infer(r)
	}
}
