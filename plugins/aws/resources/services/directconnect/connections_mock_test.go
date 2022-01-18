//go:build mock
// +build mock

package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildDirectconnectConnection(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	conn := types.Connection{}
	err := faker.FakeData(&conn)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeConnectionsOutput{
			Connections: []types.Connection{conn},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnectConnection(t *testing.T) {
	client.AwsMockTestHelper(t, DirectconnectConnections(), buildDirectconnectConnection, client.TestOptions{})
}
