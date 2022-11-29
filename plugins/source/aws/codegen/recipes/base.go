package recipes

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"go/format"
	"os"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"text/template"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/recipes/discover"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/iancoleman/strcase"
)

type Resource struct {
	// Name overrides the table name: used only in rare cases for backwards-compatibility.
	Name                  string
	Service               string
	SubService            string
	Struct                interface{}
	SkipFields            []string
	Description           string
	ExtraColumns          []codegen.ColumnDefinition
	PKColumns             []string
	Table                 *codegen.TableDefinition
	Multiplex             string
	PreResourceResolver   string
	PostResourceResolver  string
	Relations             []string
	UnwrapEmbeddedStructs bool

	// NameTransformer custom name transformer for resource
	NameTransformer func(field reflect.StructField) (string, error)

	// Used for generating the resolver and mock tests.
	// --------------------------------
	ShouldGenerateResolverAndMockTest bool        // if true, resolver and mock will be generated using the options below
	ResolverAndMockTestTemplate       string      // required: name of template directory to use
	Client                            interface{} // required: AWS client struct to use, e.g. &ec2.Client{}

	// Applies only to list resources:
	ListMethodName  string // optional: List method on the Client to use. Only required if we need to disambiguate between multiple options.
	CustomListInput string // optional: string to set List input to (otherwise empty input will be used)

	// Applies only to describe resources:
	DescribeMethodName  string // optional: Describe method on the Client to use. Only required if we need to disambiguate between multiple options.
	CustomDescribeInput string // optional: string to set List input to (otherwise empty input will be used)

	// used for generating resolver and mock tests, but set automatically
	parent   *Resource
	children []*Resource
}

//go:embed templates/resolver_and_mock_test/*/*.go.tpl
//go:embed templates/*.go.tpl
var templatesFS embed.FS

var defaultAccountColumns = []codegen.ColumnDefinition{
	{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
	},
}

var defaultRegionalColumns = []codegen.ColumnDefinition{
	{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
	},
	{
		Name:     "region",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSRegion",
	},
}

var defaultRegionalColumnsPK = []codegen.ColumnDefinition{
	{
		Name:     "account_id",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSAccount",
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	},
	{
		Name:     "region",
		Type:     schema.TypeString,
		Resolver: "client.ResolveAWSRegion",
		Options:  schema.ColumnCreationOptions{PrimaryKey: true},
	},
}

func awsNameTransformer(f reflect.StructField) (string, error) {
	c := caser.New(caser.WithCustomInitialisms(map[string]bool{
		"EC2": true,
		"VPC": true,
	}))
	return c.ToSnake(f.Name), nil
}

func awsResolverTransformer(f reflect.StructField, path string) (string, error) {
	if f.Type.String() == "[]types.Tag" {
		if path == "Tags" {
			return "client.ResolveTags", nil
		}
		return `client.ResolveTagField("` + path + `")`, nil
	}

	if path == "Tags" || path == "TagSet" {
		switch f.Type {
		case reflect.TypeOf(map[string]string{}), reflect.TypeOf(map[string]*string{}), reflect.TypeOf(map[string]interface{}{}), reflect.TypeOf([]types.TagDescription{}):
			// valid tag types
		default:
			return "", fmt.Errorf("%q field is not of type []types.Tag or acceptable map: %s", path, f.Type.String())
		}
	}

	return codegen.DefaultResolverTransformer(f, path)
}

func (r *Resource) Generate() error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	dir := path.Dir(filename)
	dir = path.Join(dir, "../../resources/services", r.Service)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	var err error
	opts := []codegen.TableOption{
		codegen.WithSkipFields(r.SkipFields),
		codegen.WithExtraColumns(r.ExtraColumns),
		codegen.WithPKColumns(r.PKColumns...),
		codegen.WithNameTransformer(awsNameTransformer),
		codegen.WithResolverTransformer(awsResolverTransformer),
	}
	if r.UnwrapEmbeddedStructs {
		opts = append(opts, codegen.WithUnwrapAllEmbeddedStructs())
	}
	if r.NameTransformer != nil {
		opts = append(opts, codegen.WithNameTransformer(r.NameTransformer))
	}
	name := fmt.Sprintf("aws_%s_%s", r.Service, r.SubService)
	if r.Name != "" {
		name = r.Name
	}

	// All table names must be plural
	if !strings.HasSuffix(name, "s") {
		return fmt.Errorf("invalid table name: %s. must be plural", name)
	}

	r.Table, err = codegen.NewTableFromStruct(
		name,
		r.Struct,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("error generating %s: %w", name, err)
	}
	r.Table.Description = r.Description
	r.Table.Resolver = "fetch" + strcase.ToCamel(r.Service) + strcase.ToCamel(r.SubService)
	if r.Multiplex != "" {
		r.Table.Multiplex = r.Multiplex
		err = validateServiceMultiplex(r.Multiplex)
		if err != nil {
			return err
		}
	}
	if r.PreResourceResolver != "" {
		r.Table.PreResourceResolver = r.PreResourceResolver
	}
	if r.PostResourceResolver != "" {
		r.Table.PostResourceResolver = r.PostResourceResolver
	}
	if r.Relations != nil {
		r.Table.Relations = r.Relations
	}

	if err := r.generateSchema(dir); err != nil {
		return err
	}

	if r.ShouldGenerateResolverAndMockTest {
		if err := r.generateResolver(dir); err != nil {
			return err
		}
		if err := r.generateMockTest(dir); err != nil {
			return err
		}
	}

	return nil
}

func (r *Resource) generateSchema(dir string) error {
	tpl, err := template.New("resource.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/resource.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}
	tpl, err = tpl.ParseFS(codegen.TemplatesFS, "templates/*.go.tpl")
	if err != nil {
		return fmt.Errorf("failed to parse sdk template: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	filePath := path.Join(dir, r.SubService+".go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}
	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func (r *Resource) generateResolver(dir string) error {
	tpl, err := template.New("fetch.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS,
		fmt.Sprintf("templates/resolver_and_mock_test/%s/fetch.go.tpl", r.ResolverAndMockTestTemplate))
	if err != nil {
		return fmt.Errorf("failed to parse templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute resolver template for %s.%s: %w", r.Service, r.SubService, err)
	}

	filePath := path.Join(dir, r.SubService+"_fetch.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

func (r *Resource) generateMockTest(dir string) error {
	tpl, err := template.New("mock_test.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS,
		fmt.Sprintf("templates/resolver_and_mock_test/%s/mock_test.go.tpl", r.ResolverAndMockTestTemplate))
	if err != nil {
		return fmt.Errorf("failed to parse gcp templates: %w", err)
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, r); err != nil {
		return fmt.Errorf("failed to execute mock template for %s.%s: %w", r.Service, r.SubService, err)
	}

	filePath := path.Join(dir, r.SubService+"_mock_test.go")
	content := buff.Bytes()
	formattedContent, err := format.Source(buff.Bytes())
	if err != nil {
		fmt.Printf("failed to format source: %s: %v\n", filePath, err)
	} else {
		content = formattedContent
	}

	if err := os.WriteFile(filePath, content, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}

// SetParentChildRelationships calculates and sets the parent and children fields on resources.
func SetParentChildRelationships(resources []*Resource) error {
	m := map[string]*Resource{}
	for _, r := range resources {
		m[r.Service+"."+r.SubService] = r
	}
	csr := caser.New()
	for _, r := range resources {
		for _, ch := range r.Relations {
			name := csr.ToSnake(strings.TrimSuffix(ch, "()"))
			v, ok := m[r.Service+"."+name]
			if !ok {
				return errors.New("child not found for " + r.Service + "." + r.SubService + " : " + name)
			}
			r.children = append(r.children, v)
			v.parent = r
		}
	}
	return nil
}

// These methods are called from the template.
// Because of this, we use a value receiver.
// -------------------------------------------------------------------------------

// StructName returns the name of the resource's Struct field
func (r Resource) StructName() string {
	// because usually the 'Struct' field contains a pointer, we need to dereference with '.Elem()'.
	return reflect.TypeOf(r.Struct).Elem().Name()
}

type ListTagsMethodResponse struct {
	Method reflect.Method
	Found  bool
}

// ListTagsMethod finds a ListTags method for the service, if any
func (r Resource) ListTagsMethod() ListTagsMethodResponse {
	if r.Client == nil || reflect.ValueOf(r.Client).IsNil() {
		panic("Client needs to be set to generate resolvers and mocks")
	}
	m, err := discover.FindListTagsMethod(r.Client)
	if err != nil {
		return ListTagsMethodResponse{
			Found: false,
		}
	}
	return ListTagsMethodResponse{
		Method: m.Method,
		Found:  true,
	}
}

// DescribeMethod finds a describe method for the resource
func (r Resource) DescribeMethod() discover.DiscoveredMethod {
	if r.Client == nil || (reflect.ValueOf(r.Client).Kind() == reflect.Ptr && reflect.ValueOf(r.Client).IsNil()) {
		panic("Client needs to be set to generate resolvers and mocks")
	}
	if r.DescribeMethodName != "" {
		m, err := discover.MethodByName(r.Client, r.Struct, r.DescribeMethodName)
		if err != nil {
			panic(err)
		}
		return m
	}
	m, err := discover.FindDescribeMethod(r.Client, r.Struct)
	if err != nil {
		panic(err)
	}
	return m
}

// ListMethod finds a list method for the resource
func (r Resource) ListMethod() discover.DiscoveredMethod {
	if r.Client == nil {
		panic("Client needs to be set to generate resolvers and mocks")
	}
	if r.ListMethodName != "" {
		m, err := discover.MethodByName(r.Client, r.Struct, r.ListMethodName)
		if err != nil {
			panic(err)
		}
		return m
	}
	m, err := discover.FindListMethod(r.Client, r.Struct)
	if err != nil {
		panic(err)
	}
	return m
}

// Parent returns the parent resource, if any
func (r Resource) Parent() *Resource {
	return r.parent
}

// Children returns the child resources, if any
func (r Resource) Children() []*Resource {
	return r.children
}

// CloudQueryServiceName is used for accessing 'client.Services().{{.CloudqueryServiceName}}' in templates
func (r Resource) CloudQueryServiceName() string {
	csr := caser.New()
	return csr.ToPascal(r.Service)
}

// CreateReplaceTransformer allows overriding column names
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

func validateServiceMultiplex(multiplexerCall string) error {
	re := regexp.MustCompile(`\"(.*?)\"`)
	// Find the value of the service parameter
	submatchAll := re.FindStringSubmatch(multiplexerCall)
	if len(submatchAll) != 2 {
		return nil
	}

	t := client.ReadSupportedServiceRegions()
	services := make(map[string]bool)

	for _, partition := range t.Partitions {
		for service := range partition.Services {
			if _, ok := services[service]; !ok {
				services[service] = true
			}
		}
	}
	if _, ok := services[submatchAll[1]]; !ok {
		return fmt.Errorf("invalid partition: %s", submatchAll[1])
	}
	return nil
}
