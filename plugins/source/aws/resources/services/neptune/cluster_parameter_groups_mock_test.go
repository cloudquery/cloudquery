package neptune

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNeptuneClusterParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockNeptuneClient(ctrl)
	var g types.DBClusterParameterGroup
	require.NoError(t, faker.FakeObject(&g))

	mock.EXPECT().DescribeDBClusterParameterGroups(
		gomock.Any(),
		&neptune.DescribeDBClusterParameterGroupsInput{
			Filters: []types.Filter{{Name: aws.String("engine"), Values: []string{"neptune"}}},
		},
		gomock.Any(),
	).Return(
		&neptune.DescribeDBClusterParameterGroupsOutput{DBClusterParameterGroups: []types.DBClusterParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&neptune.ListTagsForResourceInput{ResourceName: g.DBClusterParameterGroupArn},
		gomock.Any(),
	).Return(
		&neptune.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	var p types.Parameter
	require.NoError(t, faker.FakeObject(&p))

	mock.EXPECT().DescribeDBClusterParameters(
		gomock.Any(),
		&neptune.DescribeDBClusterParametersInput{DBClusterParameterGroupName: g.DBClusterParameterGroupName},
		gomock.Any(),
	).Return(
		&neptune.DescribeDBClusterParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return client.Services{Neptune: mock}
}

func TestNeptuneClusterParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ClusterParameterGroups(), buildNeptuneClusterParameterGroups, client.TestOptions{})
}
