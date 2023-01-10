package cognito

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCognitoIdentityPools(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCognitoidentityClient(ctrl)

	var desc types.IdentityPoolShortDescription
	if err := faker.FakeObject(&desc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListIdentityPools(
		gomock.Any(),
		&cognitoidentity.ListIdentityPoolsInput{MaxResults: 60},
		gomock.Any(),
	).Return(
		&cognitoidentity.ListIdentityPoolsOutput{IdentityPools: []types.IdentityPoolShortDescription{desc}},
		nil,
	)

	var ipo cognitoidentity.DescribeIdentityPoolOutput
	if err := faker.FakeObject(&ipo); err != nil {
		t.Fatal(err)
	}
	ipo.IdentityPoolId = desc.IdentityPoolId
	ipo.IdentityPoolId = desc.IdentityPoolName
	m.EXPECT().DescribeIdentityPool(
		gomock.Any(),
		&cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: desc.IdentityPoolId},
		gomock.Any(),
	).Return(&ipo, nil)

	return client.Services{Cognitoidentity: m}
}

func TestCognitoIdentityPools(t *testing.T) {
	client.AwsMockTestHelper(t, IdentityPools(), buildCognitoIdentityPools, client.TestOptions{})
}
