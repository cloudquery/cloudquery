package rds

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRdsClusterParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBClusterParameterGroup
	if err := faker.FakeData(&g); err != nil {
		t.Fatal(err)
	}
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
	if err := faker.FakeData(&p); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBClusterParameters(
		gomock.Any(),
		&rds.DescribeDBClusterParametersInput{DBClusterParameterGroupName: g.DBClusterParameterGroupName},
		gomock.Any(),
	).Return(
		&rds.DescribeDBClusterParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return client.Services{RDS: mock}
}

func TestRdsClusterParameterGroups(t *testing.T) {
	client.AwsMockTestHelper(t, RdsClusterParameterGroups(), buildRdsClusterParameterGroups, client.TestOptions{})
}
