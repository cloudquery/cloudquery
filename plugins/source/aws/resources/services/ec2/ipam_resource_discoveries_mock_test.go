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

func buildIpamResourceDiscoveries(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	ird := types.IpamResourceDiscovery{}
	require.NoError(t, faker.FakeObject(&ird))

	m.EXPECT().DescribeIpamResourceDiscoveries(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeIpamResourceDiscoveriesOutput{
			IpamResourceDiscoveries: []types.IpamResourceDiscovery{ird},
		}, nil)
	return client.Services{
		Ec2: m,
	}
}

func TestIpamResourceDiscoveries(t *testing.T) {
	client.AwsMockTestHelper(t, IpamResourceDiscoveries(), buildIpamResourceDiscoveries, client.TestOptions{})
}
