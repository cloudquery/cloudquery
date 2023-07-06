package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2InstanceStatuses(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.InstanceStatus{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeInstanceStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstanceStatusOutput{
			InstanceStatuses: []types.InstanceStatus{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}
func TestEc2InstanceStatuses(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceStatuses(), buildEc2InstanceStatuses, client.TestOptions{})
}
