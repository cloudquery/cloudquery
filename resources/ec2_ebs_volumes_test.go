package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2EbsVolumes(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	volumesOutput := ec2.DescribeVolumesOutput{}
	err := faker.FakeData(&volumesOutput)
	if err != nil {
		t.Fatal(err)
	}
	volumesOutput.NextToken = nil
	m.EXPECT().DescribeVolumes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&volumesOutput, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2EbsVolumes(t *testing.T) {
	awsTestHelper(t, Ec2EbsVolumes(), buildEc2EbsVolumes, TestOptions{})
}
