package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2NetworkAcls(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	l := types.NetworkAcl{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.IsDefault = aws.Bool(false)
	m.EXPECT().DescribeNetworkAcls(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeNetworkAclsOutput{
			NetworkAcls: []types.NetworkAcl{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2NetworkAclsMockTest(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2NetworkAcls(), buildEc2NetworkAcls, client.TestOptions{})
}
