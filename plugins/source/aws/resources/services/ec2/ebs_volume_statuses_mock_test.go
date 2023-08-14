package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2EbsVolumeStatuses(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	statusOutput := ec2.DescribeVolumeStatusOutput{}
	require.NoError(t, faker.FakeObject(&statusOutput))

	statusOutput.NextToken = nil
	m.EXPECT().DescribeVolumeStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&statusOutput, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestEc2EbsVolumeStatues(t *testing.T) {
	client.AwsMockTestHelper(t, EbsVolumesStatuses(), buildEc2EbsVolumeStatuses, client.TestOptions{})
}
