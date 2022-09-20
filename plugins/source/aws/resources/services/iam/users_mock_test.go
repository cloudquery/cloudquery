package iam

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIamUsers(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	u := iamTypes.User{}
	err := faker.FakeData(&u)
	if err != nil {
		t.Fatal(err)
	}
	g := iamTypes.Group{}
	err = faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}
	km := iamTypes.AccessKeyMetadata{}
	err = faker.FakeData(&km)
	if err != nil {
		t.Fatal(err)
	}
	aup := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&aup)
	if err != nil {
		t.Fatal(err)
	}
	akl := iam.GetAccessKeyLastUsedOutput{}
	err = faker.FakeData(&akl)
	if err != nil {
		t.Fatal(err)
	}

	var tags []iamTypes.Tag
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListUsers(gomock.Any(), gomock.Any()).Return(
		&iam.ListUsersOutput{
			Users: []iamTypes.User{u},
		}, nil)
	m.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(
		&iam.GetUserOutput{
			User: &u,
		}, nil)
	m.EXPECT().ListGroupsForUser(gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsForUserOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAccessKeys(gomock.Any(), gomock.Any()).Return(
		&iam.ListAccessKeysOutput{
			AccessKeyMetadata: []iamTypes.AccessKeyMetadata{km},
		}, nil)
	m.EXPECT().ListAttachedUserPolicies(gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedUserPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{aup},
		}, nil)
	m.EXPECT().GetAccessKeyLastUsed(gomock.Any(), gomock.Any()).Return(
		&akl, nil)

	//list user inline policies
	var l []string
	err = faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListUserPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListUserPoliciesOutput{
			PolicyNames: l,
		}, nil)

	//get policy
	p := iam.GetUserPolicyOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	p.PolicyDocument = &document
	m.EXPECT().GetUserPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	return client.Services{
		IAM: m,
	}
}

func TestIamUsers(t *testing.T) {
	client.AwsMockTestHelper(t, Users(), buildIamUsers, client.TestOptions{})
}
