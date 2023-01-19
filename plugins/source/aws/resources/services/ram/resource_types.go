package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func ResourceTypes() *schema.Table {
	return &schema.Table{
		Name:        "aws_ram_resource_types",
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_ServiceNameAndResourceType.html`,
		Resolver:    fetchRamResourceTypes,
		Transform: transformers.TransformWithStruct(&types.ServiceNameAndResourceType{},
			transformers.WithPrimaryKeys("ResourceType", "ServiceName")),
		Multiplex: client.ServiceAccountRegionMultiplexer("ram"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
		},
	}
}
