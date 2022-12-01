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

func buildRoles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	r := iamTypes.Role{}
	err := faker.FakeObject(&r)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}

	// generate valid json
	document := `{"stuff": 3}`
	r.AssumeRolePolicyDocument = &document

	m.EXPECT().GetRole(gomock.Any(), gomock.Any()).Return(
		&iam.GetRoleOutput{
			Role: &r,
		}, nil)

	m.EXPECT().ListRoles(gomock.Any(), gomock.Any()).Return(
		&iam.ListRolesOutput{
			Roles: []iamTypes.Role{r},
		}, nil)
	m.EXPECT().ListAttachedRolePolicies(gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedRolePoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	// list policies by role
	var l []string
	err = faker.FakeObject(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRolePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListRolePoliciesOutput{
			PolicyNames: l,
		}, nil)

	//get policy
	pd := iam.GetRolePolicyOutput{}
	err = faker.FakeObject(&pd)
	if err != nil {
		t.Fatal(err)
	}
	pd.PolicyDocument = &document
	m.EXPECT().GetRolePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&pd, nil)

	//get tags
	tag := iamTypes.Tag{}
	err = faker.FakeObject(&tag)
	if err != nil {
		t.Fatal(err)
	}

	return client.Services{
		Iam: m,
	}
}

func TestIamRoles(t *testing.T) {
	client.AwsMockTestHelper(t, Roles(), buildRoles, client.TestOptions{})
}
