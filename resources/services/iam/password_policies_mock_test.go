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

func buildIamPasswordPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.PasswordPolicy{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetAccountPasswordPolicy(gomock.Any(), gomock.Any()).Return(
		&iam.GetAccountPasswordPolicyOutput{
			PasswordPolicy: &g,
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func TestIamPasswordPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, IamPasswordPolicies(), buildIamPasswordPolicies, client.TestOptions{})
}
