package services

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"text/template"

	"github.com/cloudquery/cloudquery/plugins/source/azure/codegen/util"
	"github.com/cloudquery/plugin-sdk/caser"
	"github.com/iancoleman/strcase"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

//go:embed templates/*.go.tpl
var templatesFS embed.FS

// Generate generates a services.go file and individual service files from the constructors defined in constructors.go
func Generate() error {
	services := make([]*serviceInfo, 0, len(constructors))
	for _, constructor := range constructors {
		services = append(services, getServiceInfo(constructor))
	}
	svcMap := make(map[string]*servicePackage)
	for _, constructor := range constructors {
		svc := getServiceInfo(constructor)
		pkg, ok := svcMap[svc.PackageName]
		if !ok {
			pkg = new(servicePackage)
			svcMap[svc.PackageName] = pkg
		}
		pkg.addService(svc)
	}

	pkgs := maps.Keys(svcMap)
	slices.Sort(pkgs)

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	// write services.go file
	servicesTpl, err := template.New("services.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/services.go.tpl")
	if err != nil {
		return err
	}

	buff := new(bytes.Buffer)
	if err := servicesTpl.Execute(buff, pkgs); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	filePath := path.Join(path.Dir(filename), "../../client/services.go")
	if err := util.WriteAndFormat(filePath, buff.Bytes()); err != nil {
		return err
	}

	dir := path.Dir(filename)
	grp, _ := errgroup.WithContext(context.Background())
	for _, pkg := range svcMap {
		pkg := pkg
		grp.Go(func() error {
			return pkg.generateServicePackage(dir)
		})
	}
	return grp.Wait()
}

type servicePackage struct {
	PackageName string
	Services    []*serviceInfo
}

func (sp *servicePackage) addService(s *serviceInfo) {
	if len(sp.Services) == 0 {
		sp.PackageName = s.PackageName
	}
	sp.Services = append(sp.Services, s)
}

func (sp *servicePackage) generateServicePackage(dir string) error {
	// write encompassing clients
	clientTpl, err := template.New("client.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/client.go.tpl")
	if err != nil {
		return err
	}

	buff := new(bytes.Buffer)
	if err := clientTpl.Execute(buff, sp); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	dirPath := path.Join(dir, fmt.Sprintf("../../client/services/%s", sp.PackageName))
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}

	csr := caser.New()
	filePath := path.Join(dirPath, csr.ToSnake(sp.PackageName)+"_client.go")
	if err := util.WriteAndFormat(filePath, buff.Bytes()); err != nil {
		return err
	}

	// write individual service files
	grp, _ := errgroup.WithContext(context.Background())
	for _, service := range sp.Services {
		service := service
		grp.Go(func() error {
			return service.generate(dir)
		})
	}
	return grp.Wait()
}

func (s *serviceInfo) generate(dir string) error {
	serviceTpl, err := template.New("service.go.tpl").Funcs(template.FuncMap{
		"ToCamel": strcase.ToCamel,
		"ToLower": strings.ToLower,
	}).ParseFS(templatesFS, "templates/service.go.tpl")
	if err != nil {
		return err
	}

	buff := new(bytes.Buffer)
	if err := serviceTpl.Execute(buff, s); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	dirPath := path.Join(dir, fmt.Sprintf("../../client/services/%s", s.PackageName))
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dirPath, err)
	}

	filePath := path.Join(dirPath, s.SourceFile+".go")
	return util.WriteAndFormat(filePath, buff.Bytes())
}
