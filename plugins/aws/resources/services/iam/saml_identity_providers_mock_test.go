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

func buildIamSAMLProviders(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	l := iamTypes.SAMLProviderListEntry{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListSAMLProviders(gomock.Any(), gomock.Any()).Return(
		&iam.ListSAMLProvidersOutput{
			SAMLProviderList: []iamTypes.SAMLProviderListEntry{l},
		}, nil)

	p := iam.GetSAMLProviderOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetSAMLProvider(gomock.Any(), gomock.Any()).Return(&p, nil)

	return client.Services{
		IAM: m,
	}
}

func TestIAMSamlIdentityProviders(t *testing.T) {
	client.AwsMockTestHelper(t, IamSamlIdentityProviders(), buildIamSAMLProviders, client.TestOptions{})
}
