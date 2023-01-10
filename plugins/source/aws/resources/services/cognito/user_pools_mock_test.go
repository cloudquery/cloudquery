package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCognitoUserPools(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCognitoidentityproviderClient(ctrl)

	var desc types.UserPoolDescriptionType
	if err := faker.FakeObject(&desc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPools(
		gomock.Any(),
		&cognitoidentityprovider.ListUserPoolsInput{MaxResults: 60},
		gomock.Any(),
	).Return(
		&cognitoidentityprovider.ListUserPoolsOutput{UserPools: []types.UserPoolDescriptionType{desc}},
		nil,
	)

	var pool types.UserPoolType
	if err := faker.FakeObject(&pool); err != nil {
		t.Fatal(err)
	}
	pool.Id = desc.Id
	m.EXPECT().DescribeUserPool(
		gomock.Any(),
		&cognitoidentityprovider.DescribeUserPoolInput{UserPoolId: desc.Id},
		gomock.Any(),
	).Return(
		&cognitoidentityprovider.DescribeUserPoolOutput{UserPool: &pool},
		nil,
	)

	var providerDesc types.ProviderDescription
	if err := faker.FakeObject(&providerDesc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListIdentityProviders(
		gomock.Any(),
		&cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id},
		gomock.Any(),
	).Return(
		&cognitoidentityprovider.ListIdentityProvidersOutput{Providers: []types.ProviderDescription{providerDesc}},
		nil,
	)

	var provider types.IdentityProviderType
	if err := faker.FakeObject(&provider); err != nil {
		t.Fatal(err)
	}
	provider.ProviderName = providerDesc.ProviderName
	provider.UserPoolId = pool.Id
	m.EXPECT().DescribeIdentityProvider(
		gomock.Any(),
		&cognitoidentityprovider.DescribeIdentityProviderInput{
			ProviderName: providerDesc.ProviderName,
			UserPoolId:   pool.Id,
		},
		gomock.Any(),
	).Return(
		&cognitoidentityprovider.DescribeIdentityProviderOutput{IdentityProvider: &provider},
		nil,
	)

	return client.Services{Cognitoidentityprovider: m}
}

func TestCognitoUserPools(t *testing.T) {
	client.AwsMockTestHelper(t, UserPools(), buildCognitoUserPools, client.TestOptions{})
}
