package wellarchitected

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/wellarchitected"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildLensesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWellarchitectedClient(ctrl)

	m.EXPECT().ListLenses(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context,
			input *wellarchitected.ListLensesInput,
			_ ...func(options *wellarchitected.Options),
		) (*wellarchitected.ListLensesOutput, error) {
			var summary types.LensSummary
			require.NoError(t, faker.FakeObject(&summary))
			summary.LensType = input.LensType
			return &wellarchitected.ListLensesOutput{LensSummaries: []types.LensSummary{summary}}, nil
		}).MinTimes(len(types.LensType("").Values()))

	m.EXPECT().GetLens(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(
		func(_ context.Context,
			input *wellarchitected.GetLensInput,
			_ ...func(options *wellarchitected.Options),
		) (*wellarchitected.GetLensOutput, error) {
			var lens types.Lens
			require.NoError(t, faker.FakeObject(&lens))
			lens.LensArn = input.LensAlias
			lens.LensVersion = input.LensVersion
			return &wellarchitected.GetLensOutput{Lens: &lens}, nil
		}).MinTimes(len(types.LensType("").Values()))

	return client.Services{Wellarchitected: m}
}

func TestLenses(t *testing.T) {
	client.AwsMockTestHelper(t, Lenses(), buildLensesMock, client.TestOptions{})
}
