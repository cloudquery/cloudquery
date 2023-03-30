package ram

import (
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Resources() *schema.Table {
	tableName := "aws_ram_resources"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/ram/latest/APIReference/API_Resource.html`,
		Resolver:    fetchRamResources,
		Transform:   client.TransformWithStruct(&types.Resource{}, transformers.WithPrimaryKeys("Arn", "ResourceShareArn")),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ram"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
