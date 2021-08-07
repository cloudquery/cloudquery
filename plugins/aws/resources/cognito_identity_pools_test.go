package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCognitoIdentityPools(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCognitoIdentityPoolsClient(ctrl)

	var desc types.IdentityPoolShortDescription
	if err := faker.FakeData(&desc); err != nil {
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
	if err := faker.FakeData(&ipo); err != nil {
		t.Fatal(err)
	}
	ipo.IdentityPoolId = desc.IdentityPoolId
	ipo.IdentityPoolId = desc.IdentityPoolName
	m.EXPECT().DescribeIdentityPool(
		gomock.Any(),
		&cognitoidentity.DescribeIdentityPoolInput{IdentityPoolId: desc.IdentityPoolId},
		gomock.Any(),
	).Return(&ipo, nil)

	return client.Services{CognitoIdentityPools: m}
}

func TestCognitoIdentityPools(t *testing.T) {
	awsTestHelper(t, CognitoIdentityPools(), buildCognitoIdentityPools, TestOptions{})
}
