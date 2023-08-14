package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIamVirtualMfaDevices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.VirtualMFADevice{}
	require.NoError(t, faker.FakeObject(&g))

	m.EXPECT().ListVirtualMFADevices(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListVirtualMFADevicesOutput{
			VirtualMFADevices: []iamTypes.VirtualMFADevice{g},
		}, nil)
	return client.Services{
		Iam: m,
	}
}
func TestIAMVirtualMfaDevices(t *testing.T) {
	client.AwsMockTestHelper(t, VirtualMfaDevices(), buildIamVirtualMfaDevices, client.TestOptions{})
}
