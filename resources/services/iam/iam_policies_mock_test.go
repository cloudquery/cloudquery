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

func buildIamPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.ManagedPolicyDetail{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}
	document := `{"stuff": 3}`
	// generate valid json
	for i := range g.PolicyVersionList {
		g.PolicyVersionList[i].Document = &document
	}

	m.EXPECT().GetAccountAuthorizationDetails(gomock.Any(), gomock.Any()).Return(
		&iam.GetAccountAuthorizationDetailsOutput{
			Policies: []iamTypes.ManagedPolicyDetail{g},
		}, nil)
	return client.Services{
		IAM: m,
	}
}

func TestIamPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, IamPolicies(), buildIamPolicies, client.TestOptions{})
}
