package wellarchitected

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v3/schema"
	"github.com/cloudquery/plugin-sdk/v3/transformers"
)

func lensReviews() *schema.Table {
	name := "aws_wellarchitected_lens_reviews"
	return &schema.Table{
		Name:        name,
		Description: `https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/wellarchitected/types#Lens`,
		Transform: transformers.TransformWithStruct(new(types.LensReview),
			transformers.WithPrimaryKeys("LensAlias"),
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:  fetchLensReviews,
		Columns: schema.ColumnList{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
			{
				Name:     "workload_id",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("workload_id"),
			},
			{
				Name:     "milestone_number",
				Type:     arrow.PrimitiveTypes.Int64,
				Resolver: schema.ParentColumnResolver("number"),
			},
		},
		Relations: nil,
	}
}

func fetchLensReviews(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected
	milestoneNumber := int32(parent.Get("number").Get().(int64))
	workloadID := parent.Get("workload_id").String()

	p := wellarchitected.NewListLensReviewsPaginator(service,
		&wellarchitected.ListLensReviewsInput{
			WorkloadId:      &workloadID,
			MilestoneNumber: milestoneNumber,
			MaxResults:      50,
		},
	)
	for p.HasMorePages() {
		output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			return err
		}

		// we also need workload summary for Get call, so we do it here
		for _, summary := range output.LensReviewSummaries {
			out, err := service.GetLensReview(ctx,
				&wellarchitected.GetLensReviewInput{
					LensAlias:       summary.LensAlias,
					WorkloadId:      &workloadID,
					MilestoneNumber: milestoneNumber,
				},
				func(o *wellarchitected.Options) { o.Region = cl.Region },
			)
			if err != nil {
				cl.Logger().Err(err).Str("table", "aws_wellarchitected_lens_reviews").Msg("Failed to perform get, ignoring...")
				// At the very least we want the summary data to be filled in
				res <- summary
				continue
			}

			res <- out.LensReview
		}
	}

	return nil
}
