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

func buildRDSDBParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBParameterGroup
	require.NoError(t, faker.FakeObject(&g))

	mock.EXPECT().DescribeDBParameterGroups(
		gomock.Any(),
		&rds.DescribeDBParameterGroupsInput{},
		gomock.Any(),
	).Return(
		&rds.DescribeDBParameterGroupsOutput{DBParameterGroups: []types.DBParameterGroup{g}},
		nil,
	)

	mock.EXPECT().ListTagsForResource(
		gomock.Any(),
		&rds.ListTagsForResourceInput{ResourceName: g.DBParameterGroupArn},
		gomock.Any(),
	).Return(
		&rds.ListTagsForResourceOutput{
			TagList: []types.Tag{{Key: aws.String("key"), Value: aws.String("value")}},
		},
		nil,
	)

	var p types.Parameter
	require.NoError(t, faker.FakeObject(&p))

	mock.EXPECT().DescribeDBParameters(
		gomock.Any(),
		&rds.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName},
		gomock.Any(),
	).Return(
		&rds.DescribeDBParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return client.Services{Rds: mock}
}

func TestRDSDBParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, DbParameterGroups(), buildRDSDBParameterGroups, client.TestOptions{})
}
