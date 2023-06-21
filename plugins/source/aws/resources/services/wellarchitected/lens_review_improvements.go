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

func lensReviewImprovements() *schema.Table {
	name := "aws_wellarchitected_lens_review_improvements"
	return &schema.Table{
		Name:        name,
		Description: `https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/wellarchitected/types#Lens`,
		Transform: transformers.TransformWithStruct(new(types.ImprovementSummary),
			transformers.WithPrimaryKeys("PillarId", "QuestionId"),
			transformers.WithUnwrapAllEmbeddedStructs(),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(name, "wellarchitected"),
		Resolver:  fetchLensReviewImprovements,
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
				Resolver: schema.ParentColumnResolver("milestone_number"),
			},
			{
				Name:     "lens_alias",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("lens_alias"),
			},
		},
	}
}

func fetchLensReviewImprovements(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	service := cl.Services().Wellarchitected
	milestoneNumber := int32(parent.Get("milestone_number").Get().(int64))
	workloadID := parent.Get("workload_id").String()
	lensAlias := parent.Get("lens_alias").String()

	review, ok := parent.Item.(*types.LensReview)
	if !ok {
		// we need the full resource to get the pillars
		return nil
	}

	for _, pillar := range review.PillarReviewSummaries {
		p := wellarchitected.NewListLensReviewImprovementsPaginator(service,
			&wellarchitected.ListLensReviewImprovementsInput{
				LensAlias:       &lensAlias,
				WorkloadId:      &workloadID,
				MilestoneNumber: milestoneNumber,
				MaxResults:      50,
				PillarId:        pillar.PillarId,
			},
		)
		for p.HasMorePages() {
			output, err := p.NextPage(ctx, func(o *wellarchitected.Options) {
				o.Region = cl.Region
			})
			if err != nil {
				return err
			}

			res <- output.ImprovementSummaries
		}
	}

	return nil
}
