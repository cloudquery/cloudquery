package resourcemanager

import (
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/apache/arrow/go/v14/arrow"
	"github.com/cloudquery/cloudquery/plugins/source/gcp/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func ProjectTagKeys() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_project_tag_keys",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagKeys/list`,
		Resolver:    fetchProjectTagKeys,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TagKey{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
		Relations: []*schema.Table{
			projectTagValues(),
		},
	}
}

func projectTagValues() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_project_tag_values",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagValues/list`,
		Resolver:    fetchProjectTagValues,
		Multiplex:   client.ProjectMultiplexEnabledServices("resourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TagValue{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}

func ProjectTagBindings() *schema.Table {
	return &schema.Table{
		Name:        "gcp_resourcemanager_project_tag_bindings",
		Description: `https://cloud.google.com/resource-manager/reference/rest/v3/tagBindings`,
		Resolver:    fetchProjectTagBindings,
		Multiplex:   client.ProjectMultiplexEnabledServices("cloudresourcemanager.googleapis.com"),
		Transform:   client.TransformWithStruct(&pb.TagValue{}, transformers.WithPrimaryKeys("Name")),
		Columns: []schema.Column{
			{
				Name:       "project_id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   client.ResolveProject,
				PrimaryKey: true,
			},
		},
	}
}
