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

func buildDirectconnectConnection(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	conn := types.Connection{}
	require.NoError(t, faker.FakeObject(&conn))
	m.EXPECT().DescribeConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeConnectionsOutput{
			Connections: []types.Connection{conn},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnectConnection(t *testing.T) {
	client.AwsMockTestHelper(t, Connections(), buildDirectconnectConnection, client.TestOptions{})
}
