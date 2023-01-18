package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Triggers() *schema.Table {
	return &schema.Table{
		Name:                "aws_glue_triggers",
		Description:         `https://docs.aws.amazon.com/glue/latest/webapi/API_Trigger.html`,
		Resolver:            fetchGlueTriggers,
		PreResourceResolver: getTrigger,
		Transform:           transformers.TransformWithStruct(&types.Trigger{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
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
