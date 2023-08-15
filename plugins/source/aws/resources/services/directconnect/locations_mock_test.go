package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildDirectconnectLocations(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	loc := types.Location{}
	require.NoError(t, faker.FakeObject(&loc))
	m.EXPECT().DescribeLocations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeLocationsOutput{
			Locations: []types.Location{loc},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnectLocation(t *testing.T) {
	client.AwsMockTestHelper(t, Locations(), buildDirectconnectLocations, client.TestOptions{})
}
