package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Triggers() *schema.Table {
	tableName := "aws_glue_triggers"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_Trigger.html`,
		Resolver:            fetchGlueTriggers,
		PreResourceResolver: getTrigger,
		Transform:           client.TransformWithStruct(&types.Trigger{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveGlueTriggerArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueTriggerTags,
			},
		},
	}
}
