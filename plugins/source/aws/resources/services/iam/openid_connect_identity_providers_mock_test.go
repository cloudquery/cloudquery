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

func buildIamOpenIDConnectProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.OpenIDConnectProviderListEntry{}
	require.NoError(t, faker.FakeObject(&l))
	m.EXPECT().ListOpenIDConnectProviders(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListOpenIDConnectProvidersOutput{
			OpenIDConnectProviderList: []iamTypes.OpenIDConnectProviderListEntry{l},
		}, nil)

	p := iam.GetOpenIDConnectProviderOutput{}
	require.NoError(t, faker.FakeObject(&p))
	m.EXPECT().GetOpenIDConnectProvider(gomock.Any(), gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		Iam: m,
	}
}

func TestIamOpenidConnectIdentityProviders(t *testing.T) {
	client.AwsMockTestHelper(t, OpenidConnectIdentityProviders(), buildIamOpenIDConnectProviders, client.TestOptions{})
}
