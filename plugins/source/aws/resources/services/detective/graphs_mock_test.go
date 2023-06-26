package detective

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildGraphs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDetectiveClient(ctrl)
	services := client.Services{
		Detective: m,
	}
	c := types.Graph{}
	require.NoError(t, faker.FakeObject(&c))
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	listGraphOutput := &detective.ListGraphsOutput{
		GraphList: []types.Graph{c},
	}
	m.EXPECT().ListGraphs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listGraphOutput,
		nil,
	)

	tags := &detective.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestGraphs(t *testing.T) {
	client.AwsMockTestHelper(t, Graphs(), buildGraphs, client.TestOptions{})
}
