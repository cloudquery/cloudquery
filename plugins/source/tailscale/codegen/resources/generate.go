package resources

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/cloudquery/plugin-sdk/codegen"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

func Generate(ctx context.Context) error {
	resources := recipes()

	grp, _ := errgroup.WithContext(ctx)
	for _, resource := range resources {
		grp.Go(resource.generate)
	}

	if err := grp.Wait(); err != nil {
		return err
	}

	return tables(resources)
}

func (r *Resource) generate() error {
	r.sanitize()
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get caller information")
	}

	dir := path.Dir(filename)
	dir = path.Join(dir, "../../resources/services", r.Service)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dir, err)
	}

	name := fmt.Sprintf("tailscale_%s_%s", r.Service, r.SubService)
	if r.Name != "" {
		name = r.Name
	}

	fmt.Println("generate", name)
	for _, child := range r.Children {
		if err := child.generate(); err != nil {
			return err
		}
	}

	var err error
	opts := []codegen.TableOption{
		codegen.WithExtraColumns(append(codegen.ColumnDefinitions{tailnetCol}, r.ExtraColumns...)),
		codegen.WithPKColumns(r.PKColumns...),
		codegen.WithNameTransformer(nameTransformer),
	}

	// All table names must be plural
	if !strings.HasSuffix(name, "s") {
		return fmt.Errorf("invalid table name: %s. must be plural", name)
	}

	r.table, err = codegen.NewTableFromStruct(
		name,
		r.Struct,
		opts...,
	)
	if err != nil {
		return fmt.Errorf("error generating %s: %w", name, err)
	}
	r.table.Description = r.description()
	r.table.Resolver = r.fetcherName()
	if r.PreResolver != "" {
		r.table.PreResourceResolver = r.PreResolver
	}
	if r.PostResolver != "" {
		r.table.PostResourceResolver = r.PostResolver
	}
	for _, child := range r.Children {
		r.table.Relations = append(r.table.Relations, csr.ToCamel(child.SubService)+"()")
	}
	slices.Sort(r.table.Relations)

	if err := r.generateSchema(dir); err != nil {
		return err
	}

	return nil
}
