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

func buildLensReviewImprovements(t *testing.T, m *mocks.MockWellarchitectedClient,
	workload *types.Workload,
	milestone *types.MilestoneSummary,
	review *types.LensReview,
) {
	var summary types.ImprovementSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListLensReviewImprovements(gomock.Any(),
		&wellarchitected.ListLensReviewImprovementsInput{
			LensAlias:       review.LensAlias,
			WorkloadId:      workload.WorkloadId,
			MilestoneNumber: milestone.MilestoneNumber,
			PillarId:        review.PillarReviewSummaries[0].PillarId,
			MaxResults:      50,
		},
		gomock.Any()).
		Return(
			&wellarchitected.ListLensReviewImprovementsOutput{
				ImprovementSummaries: []types.ImprovementSummary{summary},
			},
			nil,
		)
}
