//go:build mock
// +build mock

package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIamVirtualMfaDevices(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.VirtualMFADevice{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListVirtualMFADevices(gomock.Any(), gomock.Any()).Return(
		&iam.ListVirtualMFADevicesOutput{
			VirtualMFADevices: []iamTypes.VirtualMFADevice{g},
		}, nil)
	return client.Services{
		IAM: m,
	}
}
func TestIAMVirtualMfaDevices(t *testing.T) {
	client.AwsMockTestHelper(t, IamVirtualMfaDevices(), buildIamVirtualMfaDevices, client.TestOptions{})
}
