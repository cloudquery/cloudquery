package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCognitoUserPools(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCognitoUserPoolsClient(ctrl)

	var desc types.UserPoolDescriptionType
	if err := faker.FakeData(&desc); err != nil {
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
	if err := faker.FakeData(&pool); err != nil {
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
	if err := faker.FakeData(&providerDesc); err != nil {
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
	if err := faker.FakeData(&provider); err != nil {
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

	return client.Services{CognitoUserPools: m}
}

func TestCognitoUserPools(t *testing.T) {
	client.AwsMockTestHelper(t, CognitoUserPools(), buildCognitoUserPools, client.TestOptions{})
}
