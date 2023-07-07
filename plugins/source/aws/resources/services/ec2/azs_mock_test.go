package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAzsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	r := types.AvailabilityZone{}
	require.NoError(t, faker.FakeObject(&r))

	r.OptInStatus = types.AvailabilityZoneOptInStatusNotOptedIn
	r.RegionName = aws.String("us-east-1")
	m.EXPECT().DescribeAvailabilityZones(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeAvailabilityZonesOutput{
			AvailabilityZones: []types.AvailabilityZone{r},
		}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestAZs(t *testing.T) {
	client.AwsMockTestHelper(t, AvailabilityZones(), buildAzsMock, client.TestOptions{})
}
