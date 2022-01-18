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

func buildDirectconnectVirtualInterfacesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualInterface{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeVirtualInterfaces(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&directconnect.DescribeVirtualInterfacesOutput{
			VirtualInterfaces: []types.VirtualInterface{l},
		}, nil)
	return client.Services{
		Directconnect: m,
	}
}

func TestDirectconnecVirtualInterfaces(t *testing.T) {
	client.AwsMockTestHelper(t, DirectconnectVirtualInterfaces(), buildDirectconnectVirtualInterfacesMock, client.TestOptions{})
}
