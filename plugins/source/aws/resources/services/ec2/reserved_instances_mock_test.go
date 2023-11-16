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

func buildReservedEc2Instances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := types.ReservedInstances{}
	require.NoError(t, faker.FakeObject(&l))

	m.EXPECT().DescribeReservedInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(
		&ec2.DescribeReservedInstancesOutput{
			ReservedInstances: []types.ReservedInstances{l},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestReservedEc2Instances(t *testing.T) {
	client.AwsMockTestHelper(t, ReservedInstances(), buildReservedEc2Instances, client.TestOptions{})
}
