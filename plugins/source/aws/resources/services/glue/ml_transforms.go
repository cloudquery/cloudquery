package glue

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func MlTransforms() *schema.Table {
	tableName := "aws_glue_ml_transforms"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/glue/latest/webapi/API_MLTransform.html`,
		Resolver:    fetchGlueMlTransforms,
		Transform:   transformers.TransformWithStruct(&types.MLTransform{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glue"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            resolveGlueMlTransformArn,
				PrimaryKeyComponent: true,
			},
			{
				Name:     "schema",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveMlTransformsSchema,
			},
			tagsCol(func(cl *client.Client, resource *schema.Resource) string {
				return mlTransformARN(cl, aws.ToString(resource.Item.(types.MLTransform).TransformId))
			}),
		},

		Relations: []*schema.Table{
			mlTransformTaskRuns(),
		},
	}
}

func fetchGlueMlTransforms(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceGlue).Glue
	paginator := glue.NewGetMLTransformsPaginator(svc, &glue.GetMLTransformsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glue.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Transforms
	}
	return nil
}
func resolveGlueMlTransformArn(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	return resource.Set(c.Name, mlTransformARN(cl, aws.ToString(resource.Item.(types.MLTransform).TransformId)))
}

func mlTransformARN(cl *client.Client, transformID string) string {
	return arn.ARN{
		Partition: cl.Partition,
		Service:   string(client.GlueService),
		Region:    cl.Region,
		AccountID: cl.AccountID,
		Resource:  "mlTransform/" + transformID,
	}.String()
}

func resolveMlTransformsSchema(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.MLTransform)
	j := make(map[string]string)
	for _, sCol := range r.Schema {
		j[*sCol.Name] = *sCol.DataType
	}
	return resource.Set(c.Name, j)
}
