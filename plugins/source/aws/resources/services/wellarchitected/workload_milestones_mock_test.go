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

func buildWorkloadMilestones(t *testing.T, m *mocks.MockWellarchitectedClient, workload *types.Workload) {
	var summary types.MilestoneSummary
	require.NoError(t, faker.FakeObject(&summary))

	m.EXPECT().ListMilestones(gomock.Any(),
		&wellarchitected.ListMilestonesInput{
			WorkloadId: workload.WorkloadId,
			MaxResults: 50,
		},
		gomock.Any()).
		Return(
			&wellarchitected.ListMilestonesOutput{
				MilestoneSummaries: []types.MilestoneSummary{summary},
			},
			nil,
		)

	buildLensReviews(t, m, workload, &summary)
}
