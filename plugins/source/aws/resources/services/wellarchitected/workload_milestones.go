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

func workloadMilestones() *schema.Table {
	name := "aws_wellarchitected_workload_milestones"
	return &schema.Table{
		Name:        name,
		Description: `https://docs.aws.amazon.com/wellarchitected/latest/APIReference/API_MilestoneSummary.html`,
		Transform: transformers.TransformWithStruct(new(types.MilestoneSummary),
			transformers.WithPrimaryKeys("MilestoneName"),
			transformers.WithUnwrapAllEmbeddedStructs(),
			transformers.WithSkipFields("WorkloadSummary"),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:  fetchWorkloadMilestones,
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "workload_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				Name:     "workload_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("workload_id"),
			},
		},
		Relations: schema.Tables{lensReviews()},
	}
}

func fetchWorkloadMilestones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected
	workloadID := parent.Get("workload_id").String()

	p := wellarchitected.NewListMilestonesPaginator(service,
		&wellarchitected.ListMilestonesInput{
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
		res <- output.MilestoneSummaries
	}

	return nil
}
