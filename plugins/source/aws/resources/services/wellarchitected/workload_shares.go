package wellarchitected

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func workloadShares() *schema.Table {
	name := "aws_wellarchitected_workload_shares"
	return &schema.Table{
		Name:        name,
		Description: `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_WorkloadShareSummary.html`,
		Transform: transformers.TransformWithStruct(new(types.WorkloadShareSummary),
			transformers.WithPrimaryKeys("ShareId"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:  fetchWorkloadShares,
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "workload_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
		Relations: nil,
	}
}

func fetchWorkloadShares(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected
	workloadID := parent.Get("workload_id").String()

	p := wellarchitected.NewListWorkloadSharesPaginator(service,
		&wellarchitected.ListWorkloadSharesInput{
			WorkloadId: &workloadID,
			MaxResults: 50,
		},
	)
	for p.HasMorePages() {
		output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.WorkloadShareSummaries
	}

	return nil
}
