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

func buildIamPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.ManagedPolicyDetail{}
	err := faker.FakeObject(&g)
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

	//get tags
	tag := iamTypes.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListPolicyTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListPolicyTagsOutput{
			Tags: []iamTypes.Tag{
				tag,
			},
		}, nil)
	return client.Services{
		Iam: m,
	}
}

func TestIamPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, Policies(), buildIamPolicies, client.TestOptions{})
}
