package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/cloudquery/cq-provider-terraform/client"
)

var providerNameRegex = regexp.MustCompile(`^.*\["(?P<Hostname>.*)/(?P<Namespace>.*)/(?P<Type>.*)"\].*?$`)

func TFData() *schema.Table {
	return &schema.Table{
		Name:         "tf_data",
		Description:  "Terraform meta data",
		Resolver:     resolveTerraformMetaData,
		DeleteFilter: client.DeleteLineageSerialFilter,
		Multiplex:    client.BackendMultiplex,
		Columns: []schema.Column{
			{
				Name:        "backend_type",
				Description: "Terraform backend type",
				Type:        schema.TypeString,
				Resolver:    resolveBackendType,
			},
			{
				Name:        "backend_name",
				Type:        schema.TypeString,
				Description: "Terraform backend name",
				Resolver:    resolveBackendName,
			},
			{
				Name:        "version",
				Type:        schema.TypeBigInt,
				Description: "Terraform backend version",
			},
			{
				Name:        "terraform_version",
				Type:        schema.TypeString,
				Description: "Terraform version",
			},
			{
				Name:        "serial",
				Type:        schema.TypeBigInt,
				Description: "Incremental number which describe the state version",
			},
			{
				Name:        "lineage",
				Type:        schema.TypeString,
				Description: "The \"lineage\" is a unique ID assigned to a state when it is created",
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "tf_resources",
				Description: "Terraform resources",
				Resolver:    resolveTerraformResources,
				Columns: []schema.Column{
					{
						Name:        "running_id",
						Description: "Unique fetch operation id",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "module",
						Description: "Resource module if exists",
						Type:        schema.TypeString,
					},
					{
						Name:        "mode",
						Description: "Resource mode, for example: data, managed, etc",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "provider_path",
						Description: "Resource provider full path, for example: provider[\"registry.terraform.io/hashicorp/aws\"]",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ProviderConfig"),
					},
					{
						Name:        "provider",
						Description: "Resource provider name, for example: aws, gcp, etc",
						Type:        schema.TypeString,
						Resolver:    resolveProviderName,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "tf_resource_instances",
						Description: "Terraform resource instances",
						Resolver:    resolveTerraformResourceInstances,
						Columns: []schema.Column{
							{
								Name:        "resource_id",
								Description: "Parent resource id",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "instance_id",
								Description: "Instance id",
								Type:        schema.TypeString,
								Resolver:    resolveInstanceInternalId,
							},
							{
								Name:        "schema_version",
								Description: "Terraform schema version",
								Type:        schema.TypeBigInt,
							},
							{
								Name:        "attributes",
								Description: "Instance attributes",
								Type:        schema.TypeJSON,
								Resolver:    resolveInstanceAttributes,
							},
							{
								Name:        "dependencies",
								Description: "Instance dependencies array",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "create_before_destroy",
								Description: "Should resource should be created before destroying",
								Type:        schema.TypeBool,
							},
						},
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func resolveTerraformMetaData(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
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

func resolveTerraformResources(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	state, ok := parent.Item.(client.State)
	if !ok {
		return fmt.Errorf("not terraform state")
	}
	for _, resource := range state.Resources {
		res <- resource
	}
	return nil
}

func resolveTerraformResourceInstances(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	resource, ok := parent.Item.(client.Resource)
	if !ok {
		return fmt.Errorf("not terraform Resource")
	}
	for _, instance := range resource.Instances {
		res <- instance
	}
	return nil
}

func resolveProviderName(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	res, ok := resource.Item.(client.Resource)
	if !ok {
		return fmt.Errorf("not terraform Resource")
	}

	matches := providerNameRegex.FindStringSubmatch(res.ProviderConfig)
	typeIndex := providerNameRegex.SubexpIndex("Type")
	if len(matches) >= 3 {
		return resource.Set(c.Name, matches[typeIndex])
	}
	return nil
}

func resolveInstanceAttributes(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance, ok := resource.Item.(client.Instance)
	if !ok {
		return fmt.Errorf("not terraform Instance")
	}
	attrs, err := instance.AttributesRaw.MarshalJSON()
	if err != nil {
		return fmt.Errorf("not valid JSON attributes")
	}
	return resource.Set(c.Name, attrs)
}

func resolveInstanceInternalId(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(client.Instance)
	data := make(map[string]interface{})
	if err := json.Unmarshal(instance.AttributesRaw, &data); err != nil {
		return fmt.Errorf("could not parse internal instance id")
	}
	if val, ok := data["id"]; ok {
		return resource.Set(c.Name, val)
	}
	return nil
}
