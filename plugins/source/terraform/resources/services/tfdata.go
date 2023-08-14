package services

import (
	"context"
	"regexp"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

var providerNameRegex = regexp.MustCompile(`^.*\["(?P<Hostname>.*)/(?P<Namespace>.*)/(?P<Type>.*)"\].*?$`)

func TFData() *schema.Table {
	return &schema.Table{
		Name:        "tf_data",
		Description: "Terraform meta data",
		Resolver:    resolveTerraformMetaData,
		Multiplex:   client.BackendMultiplex,
		Columns: []schema.Column{
			{
				Name:        "backend_type",
				Description: "Terraform backend type",
				Type:        arrow.BinaryTypes.String,
				Resolver:    resolveBackendType,
			},
			{
				Name:        "backend_name",
				Type:        arrow.BinaryTypes.String,
				Description: "Terraform backend name",
				Resolver:    resolveBackendName,
			},
			{
				Name:        "version",
				Type:        arrow.PrimitiveTypes.Int64,
				Description: "Terraform backend version",
			},
			{
				Name:        "terraform_version",
				Type:        arrow.BinaryTypes.String,
				Description: "Terraform version",
			},
			{
				Name:        "serial",
				Type:        arrow.PrimitiveTypes.Int64,
				Description: "Incremental number which describe the state version",
			},
			{
				Name:        "lineage",
				Type:        arrow.BinaryTypes.String,
				Description: "The \"lineage\" is a unique ID assigned to a state when it is created",
			},
		},
		Relations: []*schema.Table{tfResources()},
	}
}

// ====================================================================================================================
//
//	Table Resolver Functions
//
// ====================================================================================================================
func resolveTerraformMetaData(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	backend := c.Backend()
	res <- backend.Data.State
	return nil
}

func resolveBackendType(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	c := meta.(*client.Client)
	backend := c.Backend()
	return resource.Set("backend_type", backend.BackendType)
}

func resolveBackendName(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	c := meta.(*client.Client)
	backend := c.Backend()
	return resource.Set("backend_name", backend.BackendName)
}
