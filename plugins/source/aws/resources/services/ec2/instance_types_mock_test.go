package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildEc2InstanceTypes(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	info := types.InstanceTypeInfo{}
	err := faker.FakeObject(&info)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInstanceTypes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstanceTypesOutput{
			InstanceTypes:  []types.InstanceTypeInfo{info},
			NextToken:      nil,
			ResultMetadata: middleware.Metadata{},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2InstanceTypes(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceTypes(), buildEc2InstanceTypes, client.TestOptions{})
}
