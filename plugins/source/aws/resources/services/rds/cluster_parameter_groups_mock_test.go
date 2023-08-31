package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRdsClusterParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBClusterParameterGroup
	require.NoError(t, faker.FakeObject(&g))

	mock.EXPECT().DescribeDBClusterParameterGroups(
		gomock.Any(),
		&rds.DescribeDBClusterParameterGroupsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterParameterGroupsOutput{DBClusterParameterGroups: []types.DBClusterParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: g.DBClusterParameterGroupArn},
		gomock.Any(),
	).Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	var p types.Parameter
	require.NoError(t, faker.FakeObject(&p))

	mock.EXPECT().DescribeDBClusterParameters(
		gomock.Any(),
		&rds.DescribeDBClusterParametersInput{DBClusterParameterGroupName: g.DBClusterParameterGroupName},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRdsClusterParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, ClusterParameterGroups(), buildRdsClusterParameterGroups, client.TestOptions{})
}
