//go:build mock
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

func buildIamRoles(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	r := iamTypes.Role{}
	err := faker.FakeData(&r)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}

	// generate valid json
	document := `{"stuff": 3}`
	r.AssumeRolePolicyDocument = &document

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
	err = faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRolePolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListRolePoliciesOutput{
			PolicyNames: l,
		}, nil)

	//get policy
	pd := iam.GetRolePolicyOutput{}
	err = faker.FakeData(&pd)
	if err != nil {
		t.Fatal(err)
	}
	pd.PolicyDocument = &document
	m.EXPECT().GetRolePolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&pd, nil)

	//get tags
	tag := iamTypes.Tag{}
	err = faker.FakeData(&tag)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListRoleTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListRoleTagsOutput{
			Tags: []iamTypes.Tag{
				tag,
			},
		}, nil)

	return client.Services{
		IAM: m,
	}
}

func TestIamRoles(t *testing.T) {
	client.AwsMockTestHelper(t, IamRoles(), buildIamRoles, client.TestOptions{})
}
