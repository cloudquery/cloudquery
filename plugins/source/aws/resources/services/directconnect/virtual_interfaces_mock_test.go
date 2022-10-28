package directconnect

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/directconnect"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDirectconnectVirtualInterfacesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockDirectconnectClient(ctrl)
	l := types.VirtualInterface{}
	err := faker.FakeObject(&l)
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
	client.AwsMockTestHelper(t, VirtualInterfaces(), buildDirectconnectVirtualInterfacesMock, client.TestOptions{})
}
