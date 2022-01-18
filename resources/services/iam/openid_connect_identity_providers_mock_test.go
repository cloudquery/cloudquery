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

func buildIamOpenIDConnectProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.OpenIDConnectProviderListEntry{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListOpenIDConnectProviders(gomock.Any(), gomock.Any()).Return(
		&iam.ListOpenIDConnectProvidersOutput{
			OpenIDConnectProviderList: []iamTypes.OpenIDConnectProviderListEntry{l},
		}, nil)

	p := iam.GetOpenIDConnectProviderOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetOpenIDConnectProvider(gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		IAM: m,
	}
}

func TestIamOpenidConnectIdentityProviders(t *testing.T) {
	client.AwsMockTestHelper(t, IamOpenidConnectIdentityProviders(), buildIamOpenIDConnectProviders, client.TestOptions{})
}
