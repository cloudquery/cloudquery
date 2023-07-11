package wellarchitected

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildLensesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWellarchitectedClient(ctrl)
	for _, lensType := range types.LensType("").Values() {
		var summary types.LensSummary
		require.NoError(t, faker.FakeObject(&summary))
		summary.LensType = lensType
		m.EXPECT().ListLenses(gomock.Any(),
			&wellarchitected.ListLensesInput{
				LensStatus: types.LensStatusTypeAll,
				LensType:   lensType,
				MaxResults: 50,
			}, gomock.Any()).
			Return(&wellarchitected.ListLensesOutput{LensSummaries: []types.LensSummary{summary}}, nil)

		var lens types.Lens
		require.NoError(t, faker.FakeObject(&lens))
		lens.LensArn = summary.LensAlias
		lens.LensVersion = summary.LensVersion

		getInput := &wellarchitected.GetLensInput{
			LensAlias:   summary.LensAlias,
			LensVersion: summary.LensName,
		}
		if lensType == types.LensTypeAwsOfficial {
			getInput.LensVersion = nil
		}

		m.EXPECT().GetLens(gomock.Any(), getInput, gomock.Any()).
			Return(&wellarchitected.GetLensOutput{Lens: &lens}, nil)
	}
	return client.Services{Wellarchitected: m}
}

func TestLenses(t *testing.T) {
	client.AwsMockTestHelper(t, Lenses(), buildLensesMock, client.TestOptions{})
}
