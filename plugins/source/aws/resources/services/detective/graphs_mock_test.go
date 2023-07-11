package detective

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/detective"
	"github.com/aws/aws-sdk-go-v2/service/detective/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
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

	listGraphOutput := &detective.ListGraphsOutput{
		GraphList: []types.Graph{c},
	}
	m.EXPECT().ListGraphs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		listGraphOutput,
		nil,
	)

	tags := &detective.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)

	memberDetails := types.MemberDetail{}
	require.NoError(t, faker.FakeObject(&memberDetails))
	m.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&detective.ListMembersOutput{
			MemberDetails: []types.MemberDetail{memberDetails},
		},
		nil,
	)

	return services
}

func TestGraphs(t *testing.T) {
	client.AwsMockTestHelper(t, Graphs(), buildGraphs, client.TestOptions{})
}
