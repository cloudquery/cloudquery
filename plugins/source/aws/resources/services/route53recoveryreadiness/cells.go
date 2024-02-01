package route53recoveryreadiness

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func Cells() *schema.Table {
	tableName := "aws_route53recoveryreadiness_cells"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/recovery-readiness/latest/api/cells.html`,
		Resolver:    fetchCells,
		Transform:   transformers.TransformWithStruct(&types.CellOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "route53-recovery-control-config"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("CellArn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchCells(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceRoute53recoveryreadiness).Route53recoveryreadiness
	paginator := route53recoveryreadiness.NewListCellsPaginator(svc, &route53recoveryreadiness.ListCellsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(o *route53recoveryreadiness.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.Cells
	}
	return nil
}
