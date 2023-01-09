package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MlTransforms() *schema.Table {
	return &schema.Table{
		Name:      "aws_glue_ml_transforms",
		Resolver:  fetchGlueMlTransforms,
		Transform: transformers.TransformWithStruct(&types.MLTransform{}),
		Multiplex: client.ServiceAccountRegionMultiplexer("glue"),
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
				Resolver: resolveGlueMlTransformArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlueMlTransformTags,
			},
			{
				Name:     "schema",
				Type:     schema.TypeJSON,
				Resolver: resolveMlTransformsSchema,
			},
		},

		Relations: []*schema.Table{
			MlTransformTaskRuns(),
		},
	}
}
