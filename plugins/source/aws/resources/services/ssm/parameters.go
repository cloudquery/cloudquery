package ssm

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Parameters() *schema.Table {
	tableName := "aws_ssm_parameters"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html`,
		Resolver:    fetchSsmParameters,
		Transform:   transformers.TransformWithStruct(&types.ParameterMetadata{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "ssm"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:        "name",
				Type:        arrow.BinaryTypes.String,
				Description: `The parameter name`,
				PrimaryKey:  true,
			},
		},
	}
}

func fetchSsmParameters(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssm
	paginator := ssm.NewDescribeParametersPaginator(svc, &ssm.DescribeParametersInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *ssm.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Parameters
	}
	return nil
}
