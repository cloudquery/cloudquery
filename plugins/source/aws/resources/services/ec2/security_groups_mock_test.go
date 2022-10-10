package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2SecurityGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.SecurityGroup{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSecurityGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeSecurityGroupsOutput{
			SecurityGroups: []types.SecurityGroup{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2SecurityGroups(t *testing.T) {
	client.AwsMockTestHelper(t, SecurityGroups(), buildEc2SecurityGroups, client.TestOptions{})
}
