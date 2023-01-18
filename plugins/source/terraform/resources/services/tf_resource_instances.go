package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TFResourceInstances() *schema.Table {
	return &schema.Table{
		Name:        "tf_resource_instances",
		Description: "Terraform resource instances",
		Resolver:    resolveTerraformResourceInstances,
		Columns: []schema.Column{
			{
				Name:     "resource_name",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("name"),
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
				Type:        schema.TypeInt,
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
	}
}

func resolveTerraformResourceInstances(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	resource := parent.Item.(client.Resource)
	for _, instance := range resource.Instances {
		res <- instance
	}
	return nil
}

func resolveInstanceInternalId(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(client.Instance)
	var data map[string]any
	if err := json.Unmarshal(instance.AttributesRaw, &data); err != nil {
		return fmt.Errorf("could not parse internal instance id")
	}
	if val, ok := data["id"]; ok {
		return resource.Set(c.Name, val)
	}
	return nil
}

func resolveInstanceAttributes(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(client.Instance)
	var attrs map[string]any
	if err := json.Unmarshal(instance.AttributesRaw, &attrs); err != nil {
		return fmt.Errorf("not valid JSON attributes")
	}
	return resource.Set(c.Name, attrs)
}
