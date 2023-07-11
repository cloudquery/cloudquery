package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCognitoUserPools(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCognitoidentityproviderClient(ctrl)

	var desc types.UserPoolDescriptionType
	require.NoError(t, faker.FakeObject(&desc))

	m.EXPECT().ListUserPools(
		gomock.Any(),
		&cognitoidentityprovider.ListUserPoolsInput{MaxResults: 60},
		gomock.Any(),
	).Return(
		&cognitoidentityprovider.ListUserPoolsOutput{UserPools: []types.UserPoolDescriptionType{desc}},
		nil,
	)

	var pool types.UserPoolType
	require.NoError(t, faker.FakeObject(&pool))

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
	require.NoError(t, faker.FakeObject(&providerDesc))

	m.EXPECT().ListIdentityProviders(
		gomock.Any(),
		&cognitoidentityprovider.ListIdentityProvidersInput{UserPoolId: pool.Id},
		gomock.Any(),
	).Return(
		&cognitoidentityprovider.ListIdentityProvidersOutput{Providers: []types.ProviderDescription{providerDesc}},
		nil,
	)

	var provider types.IdentityProviderType
	require.NoError(t, faker.FakeObject(&provider))

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
