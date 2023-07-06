package wellarchitected

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildLensReviews(t *testing.T, m *mocks.MockWellarchitectedClient,
	workload *types.Workload,
	milestone *types.MilestoneSummary,
) {
	var summary types.LensReviewSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListLensReviews(gomock.Any(),
		&wellarchitected.ListLensReviewsInput{
			WorkloadId:      workload.WorkloadId,
			MilestoneNumber: milestone.MilestoneNumber,
			MaxResults:      50,
		},
		gomock.Any()).
		Return(
			&wellarchitected.ListLensReviewsOutput{
				LensReviewSummaries: []types.LensReviewSummary{summary},
			},
			nil,
		)

	var review types.LensReview
	require.NoError(t, faker.FakeObject(&review))

	m.EXPECT().GetLensReview(gomock.Any(),
		&wellarchitected.GetLensReviewInput{
			WorkloadId:      workload.WorkloadId,
			MilestoneNumber: milestone.MilestoneNumber,
			LensAlias:       summary.LensAlias,
		},
		gomock.Any()).
		Return(
			&wellarchitected.GetLensReviewOutput{
				LensReview:      &review,
				MilestoneNumber: milestone.MilestoneNumber,
				WorkloadId:      workload.WorkloadId,
			},
			nil,
		)

	buildLensReviewImprovements(t, m, workload, milestone, &review)
}
