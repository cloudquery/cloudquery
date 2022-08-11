package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2InstanceTypes(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	info := ec2Types.InstanceTypeInfo{}
	err := faker.FakeData(&info)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeInstanceTypes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstanceTypesOutput{
			InstanceTypes:  []ec2Types.InstanceTypeInfo{info},
			NextToken:      nil,
			ResultMetadata: middleware.Metadata{},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2InstanceTypes(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceTypes(), buildEc2InstanceTypes, client.TestOptions{})
}
