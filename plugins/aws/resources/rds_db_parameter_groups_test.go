package resources

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

func buildRDSDBParameterGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockRdsClient(ctrl)
	var g types.DBParameterGroup
	if err := faker.FakeData(&g); err != nil {
		t.Fatal(err)
	}
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
	if err := faker.FakeData(&p); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeDBParameters(
		gomock.Any(),
		&rds.DescribeDBParametersInput{DBParameterGroupName: g.DBParameterGroupName},
		gomock.Any(),
	).Return(
		&rds.DescribeDBParametersOutput{Parameters: []types.Parameter{p}},
		nil,
	)
	return client.Services{RDS: mock}
}

func TestRDSDBParameterGroups(t *testing.T) {
	awsTestHelper(t, RdsDbParameterGroups(), buildRDSDBParameterGroups, TestOptions{})
}
