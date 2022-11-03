package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2NetworkAcls(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)

	l := types.NetworkAcl{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.IsDefault = aws.Bool(false)
	m.EXPECT().DescribeNetworkAcls(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeNetworkAclsOutput{
			NetworkAcls: []types.NetworkAcl{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2NetworkAclsMockTest(t *testing.T) {
	client.AwsMockTestHelper(t, NetworkAcls(), buildEc2NetworkAcls, client.TestOptions{})
}
