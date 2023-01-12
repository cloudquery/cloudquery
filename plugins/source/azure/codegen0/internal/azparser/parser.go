package azparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sort"
	"strings"
)

// github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources

var newGlobalFuncsToSkip = map[string]bool{
	// We are skipping operationsClient as this just list all operations available and it is quite static
	// so don't think it's of any use and we can always enable it later
	"NewOperationsClient": true,
	// This is mostly not needed and if it is needed we add this resource manually
	"NewSKUsClient": true,
}

var newFuncToSkipPerPackage = map[string]map[string]bool{
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse": {
		"NewKustoOperationsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb": {
		// we migrated this to manual written client as it has childs
		"NewServersClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement": {
		"NewSKUsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql": {
		"NewServersClient": true,
	},
	// seems this api is not working and always returning InvalidResourceType
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql": {
		"NewDeletedServersClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute": {
		// we migrated this to manual written client as it has childs
		"NewVirtualMachinesClient":         true,
		"NewResourceSKUsClient":            true,
		"NewVirtualMachineScaleSetsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos": {
		"NewDatabaseAccountsClient": true,
	},
	// We are skipping this because we already get this info via
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault": {
		"NewVaultsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn": {
		// we skip this because we moved this to manually generated recipes as it has childs
		"NewProfilesClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork": {
		// Seems like a buggy resource that always returns a marshal error. maybe will be fixed in future Azure SDK
		"NewVirtualApplianceSKUsClient": true,
		// Seems like a resource that always return an error. Skipping for now
		"NewExpressRouteCrossConnectionsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maintenance/armmaintenance": {
		// We are getting this data already from NewConfigurationsClient across the whole subscription
		"NewConfigurationsForResourceGroupClient": true,
		// We are getting this data already from NewApplyUpdate across the whole subscription
		"NewApplyUpdateForResourceGroupClient": true,
		// Seems not implemented
		"NewApplyUpdatesClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity": {
		// ManualResource
		"NewSettingsClient": true,
		// Seems like a buggy resource that always returns a marshal error. maybe will be fixed in future Azure SDK
		"NewIngestionSettingsClient": true,
		// Seems like a buggy resource that always returns a marshal error. maybe will be fixed in future Azure SDK
		"NewContactsClient": true,
		// Seems like another buggy resource
		"NewDeletedServersClient": true,
		// Looks like this is a buggy resource. Im always getting marshal error from the Azure SDK.
		// Just skipping this for now
		"NewAccountConnectorsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservice": {
		"NewDeletedAccountsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction": {
		// Too long table name we will handle this with manually written receipe
		"NewMarketplaceRegistrationDefinitionsWithoutScopeClient": true,
		// This is already fetched by subscription level
		"NewAzureTrafficCollectorsByResourceGroupClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights": {
		// Seems like a buggy resource that always returns error. maybe will be fixed in future Azure SDK
		"NewDeletedWorkspacesClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage": {
		// we migrated this to manual written client as it has childs
		"NewAccountsClient":        true,
		"NewDeletedAccountsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory": {
		// requires renaming etag
		"NewFactoriesClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/nginx/armnginx/v2": {
		// requires setting an API version explicitly
		"NewDeploymentsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/windowsesu/armwindowsesu": {
		// Getting "The resource type could not be found in the namespace 'Microsoft.WindowsESU' for api version '2019-09-16-preview'". Seems like the API doesn't work
		"NewMultipleActivationKeysClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuredata/armazuredata": {
		// Requires setting the API version to 2019-05-10-preview
		"NewSQLServerRegistrationsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot": {
		// Requires setting the API version to 2022-08-08
		"NewBotsClient": true,
	},
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering": {
		// The primary key for this resource is Name and not ID
		"NewServiceProvidersClient": true,
	},
}

var reNewClient = regexp.MustCompile(`New[a-zA-Z]+Client`)
var reListRequest = regexp.MustCompile(`listCreateRequest|listAllCreateRequest|listBySubscriptionCreateRequest`)
var rePager = regexp.MustCompile(`NewListPager|NewListAllPager|NewListBySubscriptionPager`)
var reNamespaceFromURL = regexp.MustCompile(`/providers/([a-zA-Z\.]+)/`)

var supportedPagerParams = [][]string{
	{"options"},
}

var supportedNewClientParams = [][]string{
	{"credential", "options"},
	{"subscriptionID", "credential", "options"},
}

type function struct {
	receiver    string
	name        string
	ast         *ast.FuncDecl
	paramNames  []string
	returnTypes []string
}

func parseURLFromFunc(fn *ast.FuncDecl) string {
	for _, stmt := range fn.Body.List {
		if expr, ok := stmt.(*ast.AssignStmt); ok {
			if len(expr.Lhs) == 1 && len(expr.Rhs) == 1 {
				if lhs, ok := expr.Lhs[0].(*ast.Ident); ok {
					if lhs.Name == "urlPath" {
						if rhs, ok := expr.Rhs[0].(*ast.BasicLit); ok {
							return strings.Replace(rhs.Value, "\"", "", -1)
						}
					}
				}
			}
		}
	}
	return ""
}

func compareStrArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isArrayExist(arr [][]string, item []string) bool {
	for _, a := range arr {
		if compareStrArrays(a, item) {
			return true
		}
	}
	return false
}

func getParamNames(fn *ast.FieldList) []string {
	var params []string
	for _, p := range fn.List {
		for _, name := range p.Names {
			params = append(params, name.Name)
		}
	}
	return params
}

func getReturnTypes(fn *ast.FieldList) []string {
	var params []string
	for _, p := range fn.List {
		if ident, ok := p.Type.(*ast.Ident); ok {
			params = append(params, ident.Name)
		} else if star, ok := p.Type.(*ast.StarExpr); ok {
			if index, ok := star.X.(*ast.IndexExpr); ok {
				if ident, ok := index.Index.(*ast.Ident); ok {
					params = append(params, ident.Name)
				}
			}
		}
	}
	return params
}

func getStructValueNameFromResponse(pkgs map[string]*ast.Package, name string) (string, bool, bool) {
	st := getStruct(pkgs, name)
	st1name := st.Fields.List[0].Type.(*ast.Ident).Name
	st1 := getStruct(pkgs, st1name)
	structName := ""
	hasNextField := false
	for _, f := range st1.Fields.List {
		if f.Names[0].Name == "Value" {
			arrType := f.Type.(*ast.ArrayType)
			_, ok := arrType.Elt.(*ast.StarExpr)
			if !ok {
				structName = f.Type.(*ast.ArrayType).Elt.(*ast.Ident).Name
			} else {
				structName = f.Type.(*ast.ArrayType).Elt.(*ast.StarExpr).X.(*ast.Ident).Name
			}
		}
		if f.Names[0].Name == "NextLink" {
			hasNextField = true
		}
	}
	hasId := false
	st2 := getStruct(pkgs, structName)
	for _, f := range st2.Fields.List {
		if f.Names[0].Name == "ID" {
			hasId = true
		}
	}
	return structName, hasNextField, hasId
}

func getStruct(pkgs map[string]*ast.Package, name string) *ast.StructType {
	for _, pack := range pkgs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if gen, isGen := d.(*ast.GenDecl); isGen {
					for _, spec := range gen.Specs {
						if typeSpec, isTypeSpec := spec.(*ast.TypeSpec); isTypeSpec {
							if structType, isStruct := typeSpec.Type.(*ast.StructType); isStruct {
								if typeSpec.Name.Name == name {
									return structType
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

// returns reciever and method name that matches re
func findFunctions(pkgs map[string]*ast.Package, re *regexp.Regexp) map[string]*function {
	var funcs map[string]*function = make(map[string]*function)
	for _, pack := range pkgs {
		for _, f := range pack.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					if re.MatchString(fn.Name.Name) {
						fun := function{
							name: fn.Name.Name,
							ast:  fn,
						}
						// if function is a method extract receiver name
						if fn.Recv != nil && len(fn.Recv.List) == 1 {
							receiver := fn.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
							fun.receiver = receiver
						}
						fun.paramNames = getParamNames(fn.Type.Params)
						if fn.Type != nil && fn.Type.Results != nil {
							fun.returnTypes = getReturnTypes(fn.Type.Results)
						}
						if fun.receiver != "" {
							funcs[fun.receiver+"."+fun.name] = &fun
						} else {
							funcs[fun.name] = &fun
						}
					}
				}
			}
		}
	}
	return funcs
}

// Get a package in format of github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/xxx/yy@v0.1.1
func CreateTablesFromPackage(pkg string) ([]*Table, error) {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		output, err := exec.Command("go", "env", "GOPATH").Output()
		if err != nil {
			return nil, err
		}
		goPath = strings.TrimSpace(string(output))
	}
	pkgWithoutVersion := strings.Split(pkg, "@")[0]
	if packagesToSkip[pkgWithoutVersion] {
		return nil, nil
	}
	cacheDir := goPath + "/pkg/mod"
	// this maps client name to tables
	tables := make(map[string]*Table)
	set := token.NewFileSet()
	pkgPath := path.Join(cacheDir, pkg)
	// thats because azure had to be special with uppercase
	pkgPath = strings.Replace(pkgPath, "A", "!A", 1)
	pkgs, err := parser.ParseDir(set, pkgPath, nil, 0)
	if err != nil {
		return nil, err
	}
	newXClientFuncstions := findFunctions(pkgs, reNewClient)
	for _, fn := range newXClientFuncstions {
		if newGlobalFuncsToSkip[fn.name] || !isArrayExist(supportedNewClientParams, fn.paramNames) {
			continue
		}
		if _, ok := newFuncToSkipPerPackage[pkgWithoutVersion]; ok {
			if newFuncToSkipPerPackage[pkgWithoutVersion][fn.name] {
				continue
			}
		}
		tables[strings.TrimPrefix(fn.name, "New")] = &Table{
			NewFuncName:     fn.name,
			NewClientParams: fn.paramNames,
		}
	}
	listMethods := findFunctions(pkgs, reListRequest)
	pagerMethods := findFunctions(pkgs, rePager)

	var result []*Table
	for client, t := range tables {
		var pagerMethod *function
		var listMethod *function
		if pagerMethods[client+".NewListAllPager"] != nil {
			if listMethods[client+".listAllCreateRequest"] != nil {
				pagerMethod = pagerMethods[client+".NewListAllPager"]
				t.Pager = "NewListAllPager"
				listMethod = listMethods[client+".listAllCreateRequest"]
			} else {
				// this permutation is not supported by codegen
				continue
			}
		} else if pagerMethods[client+".NewListPager"] != nil {
			if listMethods[client+".listCreateRequest"] != nil {
				pagerMethod = pagerMethods[client+".NewListPager"]
				t.Pager = "NewListPager"
				listMethod = listMethods[client+".listCreateRequest"]
			} else {
				// this permutation is not supported by codegen
				continue
			}
		} else if pagerMethods[client+".listBySubscriptionCreateRequest"] != nil {
			if listMethods[client+".listBySubscriptionCreateRequest"] != nil {
				pagerMethod = pagerMethods[client+".listBySubscriptionCreateRequest"]
				t.Pager = "listBySubscriptionCreateRequest"
				listMethod = listMethods[client+".listBySubscriptionCreateRequest"]
			} else {
				// this permutation is not supported by codegen
				continue
			}
		} else {
			// this permutation is not supported by codegen
			continue
		}

		if !isArrayExist(supportedPagerParams, pagerMethod.paramNames) {
			continue
		}
		azURL := parseURLFromFunc(listMethod.ast)
		if azURL == "" {
			continue
		}
		t.URL = azURL
		t.ResponseStruct = pagerMethod.returnTypes[0]
		namespaceMatches := reNamespaceFromURL.FindStringSubmatch(azURL)
		if len(namespaceMatches) == 2 {
			t.Namespace = strings.ToLower(namespaceMatches[1])
			t.Multiplex = fmt.Sprintf("client.SubscriptionMultiplexRegisteredNamespace(client.Namespace%s)", strings.ReplaceAll(t.Namespace, ".", "_"))
		} else {
			t.Multiplex = "client.SubscriptionMultiplex"
		}
		hasId := false
		t.ResponseValueStruct, t.ResponspeStructNextLink, hasId = getStructValueNameFromResponse(pkgs, t.ResponseStruct)
		if hasId {
			result = append(result, t)
		}
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].NewFuncName < result[j].NewFuncName
	})
	return result, nil
}
