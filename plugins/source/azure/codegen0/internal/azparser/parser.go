package azparser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"regexp"
	"strings"
)

var newFuncsToSkip = map[string]bool{
	// We are skipping operationsClient as this just list all operations available and it is quite static
	// so don't think it's of anyuse and we can always enable it later
	"NewOperationsClient": true,
}

const (
	cacheDir = "/Users/yevgenyp/go/pkg/mod"
)

var reNewClient = regexp.MustCompile(`New[a-zA-Z]+Client`)
var reListCreateRequest = regexp.MustCompile(`listCreateRequest`)
var reNewListPager = regexp.MustCompile(`NewListPager`)
var supportedNewListPagerParams = [][]string{
	{"options"},
	{"resourceGroupName", "options"},
	{"subscriptionID", "options"},
}

var supportedNewClientParams = [][]string{
	{"credential", "options"},
	{"subscriptionID", "credential", "options"},
}

type function struct {
	receiver string
	name     string
	ast 		*ast.FuncDecl
	paramNames []string
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

// returns reciever and method name that matches re
func findFunctions(pkgs map[string]*ast.Package, re *regexp.Regexp) []function {
	var funcs []function
	for _, pack := range pkgs {
		for _, f := range pack.Files {
				for _, d := range f.Decls {
						if fn, isFn := d.(*ast.FuncDecl); isFn {
							if re.MatchString(fn.Name.Name) {
								fun := function{
									name: fn.Name.Name,
									ast: fn,
								}								
								// if function is a method extract receiver name
								if fn.Recv != nil && len(fn.Recv.List) == 1 {
									receiver := fn.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
									fun.receiver = receiver
								}
								fun.paramNames = getParamNames(fn.Type.Params)
								funcs = append(funcs, fun)
							}
						}
				}
		}
	}
	return funcs
}

// Get a package in format of github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/xxx/yy@v0.1.1
func CreateTablesFromPackage(pkg string) ([]*Table, error) {
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
		if newFuncsToSkip[fn.name] || !isArrayExist(supportedNewClientParams, fn.paramNames) {
			continue
		}
		tables[strings.TrimPrefix(fn.name, "New")] = &Table{
			NewFuncName: fn.name,
		}
	}

	listMethods := findFunctions(pkgs, reListCreateRequest)
	for _, fn := range listMethods {
		if fn.receiver == "" {
			continue
		}
		if _, ok := tables[fn.receiver]; !ok {
			continue
		}
		azURL := parseURLFromFunc(fn.ast)
		if azURL == "" {
			return nil, fmt.Errorf("could not find url for %s", fn.name)
		}
		tables[fn.receiver].URL = azURL
	}

	listNewPagerMethods := findFunctions(pkgs, reNewListPager)
	for _, fn := range listNewPagerMethods {
		if fn.receiver == "" {
			continue
		}
		if _, ok := tables[fn.receiver]; !ok {
			continue
		}
		tables[fn.receiver].HasListPager = true
		tables[fn.receiver].NewListPagerParams = fn.paramNames
	}

	var result []*Table
	for _, t := range tables {
		// skip tables witout URL (or at least that we didn't find one)
		// not NewListPager struct and more than 3 params
		if t.URL == "" || !t.HasListPager || !isArrayExist(supportedNewListPagerParams, t.NewListPagerParams) {
			continue
		}
		result = append(result, t)
	}
	return result, nil
}