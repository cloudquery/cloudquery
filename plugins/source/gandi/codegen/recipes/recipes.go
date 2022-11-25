package recipes

import (
	"fmt"
	"log"
	"path"
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

	SkipSubserviceName bool // Skip name of the subservice in auto generated table names (including relations)

	// These are inferred with reflection but can be overridden
	Service    string // Inferred from DataStruct package, pluralized
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
	SharingIDColumn = codegen.ColumnDefinition{
		Name:        "sharing_id",
		Description: "The Sharing ID of the resource.",
		Type:        schema.TypeString,
		Resolver:    "client.ResolveSharingID",
	}

	pluralizeClient *pluralize.Client
	csr             *caser.Caser
)

func init() {
	pluralizeClient = pluralize.NewClient()
	for _, s := range []string{"livedns", "simplehosting"} {
		pluralizeClient.AddUncountableRule(s)
	}

	csr = caser.New(
		caser.WithCustomInitialisms(map[string]bool{
			"DNS":     true,
			"DNSSec":  true,
			"LiveDNS": true,
			//"Simplehosting": true,
			//"Vhost":         true,
		}),
		caser.WithCustomExceptions(map[string]string{
			"dnssec":  "DNSSec",
			"livedns": "LiveDNS",
			"vhost":   "Vhost",
		}),
	)
}

func enforceAbbrevations(word string) string {
	word = strings.ReplaceAll(word, "live_dns", "livedns")
	word = strings.ReplaceAll(word, "dns_sec", "dnssec")
	return word
}

func toSnake(word string) string {
	return enforceAbbrevations(csr.ToSnake(word))
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
	basepkg := strings.ToLower(path.Base(ds.PkgPath()))

	if r.Service == "" {
		if !pluralizeClient.IsPlural(basepkg) {
			basepkg = pluralizeClient.Plural(basepkg)
		}
		r.Service = basepkg
	}
	if r.SubService == "" {
		r.SubService = toSnake(pluralizeClient.Singular(ds.Name()))
	}
}

func (r *Resource) GenerateNames() {
	if r.TableName == "" {
		nParts := []string{pluralizeClient.Singular(r.Service)}
		p := r.parent
		for p != nil {
			if !p.SkipSubserviceName {
				nParts = append(nParts, pluralizeClient.Singular(p.SubService))
			}
			p = p.parent
		}
		if !r.SkipSubserviceName {
			nParts = append(nParts, pluralizeClient.Plural(r.SubService))
		}
		if l := len(nParts); l == 1 {
			nParts[0] = pluralizeClient.Plural(nParts[0])
		} else if l == 0 {
			log.Fatalf("Could not generate table name for %s.%s", r.Service, r.SubService)
		}

		if r.TableName == "" {
			r.TableName = strings.Join(nParts, "_")
		}
	}

	r.Filename = toSnake(r.TableName) + ".go"
	r.TableFuncName = csr.ToPascal(r.TableName)
	r.ResolverFuncName = "fetch" + r.TableFuncName
}

// SetParentChildRelationships calculates and sets the parent and children fields on resources.
func SetParentChildRelationships(resources []*Resource) error {
	m := map[string]*Resource{}
	for _, r := range resources {
		key := enforceAbbrevations(pluralizeClient.Singular(r.Service) + "_" + pluralizeClient.Plural(r.SubService))
		m[key] = r
	}
	for _, r := range resources {
		for _, ch := range r.Relations {
			name := strings.TrimPrefix(toSnake(strings.TrimSuffix(ch, "()")), r.Service+"_"+r.SubService+"_")
			if !r.SkipSubserviceName {
				name = r.Service + "_" + name
			}
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
