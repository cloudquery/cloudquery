package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/types"
)

func tfResourceInstances() *schema.Table {
	return &schema.Table{
		Name:        "tf_resource_instances",
		Description: "Terraform resource instances",
		Resolver:    resolveTerraformResourceInstances,
		Columns: []schema.Column{
			{
				Name:     "resource_name",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("name"),
			},
			{
				Name:        "instance_id",
				Description: "Instance id",
				Type:        arrow.BinaryTypes.String,
				Resolver:    resolveInstanceInternalId,
			},
			{
				Name:        "schema_version",
				Description: "Terraform schema version",
				Type:        arrow.PrimitiveTypes.Int64,
			},
			{
				Name:        "attributes",
				Description: "Instance attributes",
				Type:        types.ExtensionTypes.JSON,
				Resolver:    resolveInstanceAttributes,
			},
			{
				Name:        "dependencies",
				Description: "Instance dependencies array",
				Type:        arrow.ListOf(arrow.BinaryTypes.String),
			},
			{
				Name:        "create_before_destroy",
				Description: "Should resource should be created before destroying",
				Type:        arrow.FixedWidthTypes.Boolean,
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
