package recipes

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cloudflare/cloudflare-go"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// TableName can be used to override the default generated table name
	TableName string
	// CFStruct that will be used to generate the cloudquery table
	CFStruct interface{}
	// CFStructName is name of CFStruct
	CFStructName string
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	PrimaryKey string
	Template   string
	Multiplex  string

	DefaultColumns []codegen.ColumnDefinition
	ExtraColumns   []codegen.ColumnDefinition
	RenameColumns  map[string]string

	Parent        *Resource
	TableFuncName string
	Filename      string
	ImportClient  bool

	setupDone bool
}

var (
	AccountIDColumn = codegen.ColumnDefinition{
		Name:        "account_id",
		Description: "The Account ID of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveAccountID",
	}

	ZoneIDColumn = codegen.ColumnDefinition{
		Name:        "zone_id",
		Description: "Zone identifier tag.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveZoneID",
	}
)

var listResources = combine(
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
		Multiplex:      "client.ZoneMultiplex",
		CFStruct:       &cloudflare.AccessGroup{},
	},

	parentize(&Resource{
		CFStruct: &cloudflare.Account{},
	},
		&Resource{
			CFStruct: &cloudflare.AccountMember{},
		},
	),

	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
		Multiplex:      "client.ZoneMultiplex",
		CFStruct:       &cloudflare.CertificatePack{},
	},
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn}, // ZoneIDColumn is already in the response
		Multiplex:      "client.ZoneMultiplex",

		CFStruct:   &cloudflare.DNSRecord{},
		SkipFields: []string{"Meta", "Data"},
		ExtraColumns: []codegen.ColumnDefinition{
			{
				Name:        "meta",
				Description: "Extra Cloudflare-specific information about the record.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "data",
				Description: "Metadata about the record.",
				Type:        schema.TypeJSON,
			},
		},
	},
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn},
		Multiplex:      "client.AccountMultiplex",
		CFStruct:       &cloudflare.Image{},
	},
	parentize(
		&Resource{
			DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn}, // ZoneIDColumn is already in the response
			Multiplex:      "client.ZoneMultiplex",
			CFStruct:       &cloudflare.WAFPackage{},
		},
		&Resource{
			CFStruct: &cloudflare.WAFGroup{},
		},
		&Resource{
			CFStruct: &cloudflare.WAFRule{},
		},
	),
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
		Multiplex:      "client.ZoneMultiplex",
		CFStruct:       &cloudflare.WAFOverride{},
		RenameColumns:  map[string]string{"ur_ls": "urls"},
	},
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn, ZoneIDColumn},
		Multiplex:      "client.ZoneMultiplex",
		CFStruct:       &cloudflare.WorkerRoute{},
	},
	parentize(
		&Resource{
			DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn},
			Multiplex:      "client.AccountMultiplex",
			CFStruct:       &cloudflare.WorkerMetaData{},
		},
		&Resource{
			CFStruct: &cloudflare.WorkerCronTrigger{},
		},
		&Resource{
			CFStruct: &cloudflare.WorkersSecret{},
		},
	),
	&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIDColumn},
		Multiplex:      "client.AccountMultiplex",
		CFStruct:       &cloudflare.Zone{},
	},
)

// Setup sets defaults for the given resource
func Setup(r *Resource) {
	if r.setupDone {
		return
	}

	r.CFStructName = reflect.TypeOf(r.CFStruct).Elem().Name()

	if r.Parent != nil {
		r.DefaultColumns = append([]codegen.ColumnDefinition{
			{
				Name:     "parent_cq_id",
				Type:     schema.TypeUUID,
				Resolver: "schema.ParentIDResolver",
			},
		}, r.DefaultColumns...)
	}

	r.TableFuncName = inflection.Plural(r.CFStructName)
	if r.Parent != nil {
		r.TableFuncName = strings.ToLower(string(r.TableFuncName[0])) + r.TableFuncName[1:]
	}

	r.Template = "resource_manual"
	if r.Parent == nil && r.PrimaryKey == "" {
		r.PrimaryKey = "id"
	}

	r.Filename = strcase.ToSnake(inflection.Plural(r.CFStructName))
	if r.Parent != nil {
		r.Filename = fmt.Sprintf("%s_%s", r.Parent.Filename, r.Filename)
	}
	r.setupDone = true
}

func All() []*Resource {
	for i := range listResources {
		Setup(listResources[i])
		//p := listResources[i].Parent
		//for p != nil {
		//	Setup(p)
		//	p = p.Parent
		//}
	}
	return listResources
}

// parentize adds the given parent to each resource (in subs) and returns the combined list
func parentize(parent *Resource, subs ...*Resource) []*Resource {
	Setup(parent)
	ret := make([]*Resource, len(subs)+1)
	ret[0] = parent
	for i := range subs {
		if subs[i].Parent == nil {
			subs[i].Parent = parent
		}
		ret[i+1] = subs[i]
	}
	return ret
}

// combine the given *Resource or []*Resource into a single []*Resource
// if the given argument is of another type, combine will panic
func combine(list ...interface{}) []*Resource {
	res := make([]*Resource, 0, len(list))
	for i := range list {
		switch v := list[i].(type) {
		case *Resource:
			res = append(res, v)
		case []*Resource:
			res = append(res, v...)
		default:
			panic(fmt.Sprintf("combine: unhandled type %T", list[i]))
		}
	}
	return res
}
