package glue

import (
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func MlTransforms() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_ml_transforms",
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_MLTransform.html`,
		Resolver:    fetchGlueMlTransforms,
		Transform:   transformers.TransformWithStruct(&types.MLTransform{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
