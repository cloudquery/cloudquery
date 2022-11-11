package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIamOpenIDConnectProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.OpenIDConnectProviderListEntry{}
	err := faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListOpenIDConnectProviders(gomock.Any(), gomock.Any()).Return(
		&iam.ListOpenIDConnectProvidersOutput{
			OpenIDConnectProviderList: []iamTypes.OpenIDConnectProviderListEntry{l},
		}, nil)

	p := iam.GetOpenIDConnectProviderOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetOpenIDConnectProvider(gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamOpenidConnectIdentityProviders(t *testing.T) {
	client.AwsMockTestHelper(t, OpenidConnectIdentityProviders(), buildIamOpenIDConnectProviders, client.TestOptions{})
}
