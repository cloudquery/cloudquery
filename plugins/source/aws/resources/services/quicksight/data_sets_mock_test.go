package quicksight

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/quicksight"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDataSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var ld quicksight.ListDataSetsOutput
	require.NoError(t, faker.FakeObject(&ld))

	ld.NextToken = nil
	m.EXPECT().ListDataSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ld, nil)

	var to quicksight.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&to))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	var ig quicksight.ListIngestionsOutput
	require.NoError(t, faker.FakeObject(&ig))

	ig.NextToken = nil
	m.EXPECT().ListIngestions(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ig, nil)

	var to2 quicksight.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&to2))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to2, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightDataSets(t *testing.T) {
	client.AwsMockTestHelper(t, DataSets(), buildDataSetsMock, client.TestOptions{})
}
