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

func buildTemplatesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockQuicksightClient(ctrl)

	var lo quicksight.ListTemplatesOutput
	require.NoError(t, faker.FakeObject(&lo))

	lo.NextToken = nil
	m.EXPECT().ListTemplates(gomock.Any(), gomock.Any(), gomock.Any()).Return(&lo, nil)

	var to quicksight.ListTagsForResourceOutput
	require.NoError(t, faker.FakeObject(&to))

	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&to, nil)

	return client.Services{
		Quicksight: m,
	}
}
func TestQuicksightTemplates(t *testing.T) {
	client.AwsMockTestHelper(t, Templates(), buildTemplatesMock, client.TestOptions{})
}
