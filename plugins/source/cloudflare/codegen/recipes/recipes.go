package recipes

import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gertd/go-pluralize"
)

type Resource struct {
	// Table is the table definition that will be used to generate the cloudquery table
	Table *codegen.TableDefinition
	// DataStruct that will be used to generate the cloudquery table
	DataStruct interface{}
	// SkipFields fields in go struct to skip when generating the table from the go struct
	SkipFields []string
	Template   string
	Multiplex  string

	ExtraColumns []codegen.ColumnDefinition
	PKColumns    []string

	PreResourceResolver   string
	PostResourceResolver  string
	Relations             []string
	UnwrapEmbeddedStructs bool

	Service string // Required

	SkipServiceInTableName bool // Don't prepend service name to table name
	SkipParentInTableName  bool // Don't prepend parent name to table name

	// These are inferred with reflection but can be overridden
	SubService string // Inferred from DataStruct name
	TableName  string // singular Service + plural SubService

	// These are auto calculated
	ImportClient     bool   // true if the resource/column resolvers use the client package
	Filename         string // Calculated from TableName
	TableFuncName    string // Calculated from TableName
	ResolverFuncName string // Calculated from TableFuncName

	// Used for generating better table names
	parent   *Resource
	children []*Resource
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

	pluralizeClient *pluralize.Client
	csr             *caser.Caser
)

func init() {
	pluralizeClient = pluralize.NewClient()
	for _, s := range []string{"dns", "waf"} {
		pluralizeClient.AddUncountableRule(s)
	}

	csr = caser.New(
		caser.WithCustomInitialisms(map[string]bool{
			"dns": true,
			"url": true,
			"waf": true,
		}),
		caser.WithCustomExceptions(map[string]string{
			"dns":  "DNS",
			"url":  "URL",
			"urls": "URLs",
			"waf":  "WAF",
		}),
	)
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

	if r.Service == "" {
		log.Fatalf("Service is required for %s", r.SubService)
	}
	if r.SubService == "" {
		r.SubService = csr.ToSnake(pluralizeClient.Singular(ds.Name()))
	}
}

func (r *Resource) GenerateNames() {
	if r.TableName == "" {
		// Table names are always in [<singular>...]<plural> format. Add everything in singular form and pluralize the last word later

		const sep = "_"
		var nParts []string
		if !r.SkipServiceInTableName {
			nParts = strings.Split(pluralizeClient.Singular(r.Service), sep)
		}
		p := r.parent
		for !r.SkipParentInTableName && p != nil {
			nParts = appendNoRepeat(nParts, strings.Split(pluralizeClient.Singular(p.SubService), sep)...)
			p = p.parent
		}
		nParts = appendNoRepeat(nParts, strings.Split(pluralizeClient.Singular(r.SubService), sep)...)
		nParts[len(nParts)-1] = pluralizeClient.Plural(nParts[len(nParts)-1])

		if r.TableName == "" {
			if len(nParts) == 0 {
				log.Fatalf("Could not generate table name for %s.%s", r.Service, r.SubService)
			}

			r.TableName = strings.Join(nParts, sep)
		}
	}

	r.Filename = csr.ToSnake(r.TableName) + ".go"
	r.TableFuncName = csr.ToPascal(r.TableName)
	r.ResolverFuncName = "fetch" + r.TableFuncName
}

// SetParentChildRelationships calculates and sets the parent and children fields on resources.
func SetParentChildRelationships(resources []*Resource) error {
	m := map[string]*Resource{}
	for _, r := range resources {
		key := r.Service + "_" + pluralizeClient.Plural(r.SubService)
		//log.Printf("%s.%s => %s", r.Service, r.SubService, key)
		m[key] = r
	}
	for _, r := range resources {
		for _, ch := range r.Relations {
			name := strings.TrimPrefix(csr.ToSnake(strings.TrimSuffix(ch, "()")), r.Service+"_"+r.SubService+"_")
			name = r.Service + "_" + name
			v, ok := m[name]
			if !ok {
				return fmt.Errorf("child not found for %s.%s: %s missing", r.Service, r.SubService, name)
			}
			r.children = append(r.children, v)
			v.parent = r
		}
	}
	return nil
}

func appendNoRepeat(parts []string, addition ...string) []string {
	// foo + bar = foo_bar
	// foo + foo_bar = foo_bar
	// foo_bar + bar_baz = foo_bar_baz
	// foo_bar + baz_bax = foo_bar_baz_bax
	// foo_bar_baz + bar_baz_bax = foo_bar_baz_bax

	for i := len(addition); i > 0; i-- {
		// ever-increasing from long form to short: foo_bar_baz, foo_bar, foo
		if sliceEndsWith(parts, addition[:i]) {
			return append(parts, addition[i:]...)
		}
	}
	return append(parts, addition...)
}

func sliceEndsWith(haystack, needle []string) bool {
	ln, lh := len(needle), len(haystack)
	if ln > lh {
		return false
	}
	return reflect.DeepEqual(needle, haystack[lh-ln:])
}
