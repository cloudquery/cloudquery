package cloudtrail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCloutrailImports(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)

	var desc types.ImportsListItem
	require.NoError(t, faker.FakeObject(&desc))

	m.EXPECT().ListImports(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		&cloudtrail.ListImportsOutput{
			Imports: []types.ImportsListItem{desc},
		},
		nil,
	)

	out := &cloudtrail.GetImportOutput{}
	require.NoError(t, faker.FakeObject(out))

	m.EXPECT().GetImport(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(
		out,
		nil,
	)

	return client.Services{Cloudtrail: m}
}

func TestImports(t *testing.T) {
	client.AwsMockTestHelper(t, Imports(), buildCloutrailImports, client.TestOptions{})
}
