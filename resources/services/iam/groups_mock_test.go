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

func buildIamGroups(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	g := iamTypes.Group{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	p := iamTypes.AttachedPolicy{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().ListGroups(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupsOutput{
			Groups: []iamTypes.Group{g},
		}, nil)
	m.EXPECT().ListAttachedGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListAttachedGroupPoliciesOutput{
			AttachedPolicies: []iamTypes.AttachedPolicy{p},
		}, nil)

	//list policies
	var l []string
	err = faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListGroupPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.ListGroupPoliciesOutput{
			PolicyNames: l,
		}, nil)

	//get policy
	gp := iam.GetGroupPolicyOutput{}
	err = faker.FakeData(&gp)
	if err != nil {
		t.Fatal(err)
	}
	document := "{\"test\": {\"t1\":1}}"
	gp.PolicyDocument = &document
	m.EXPECT().GetGroupPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&gp, nil)
	return client.Services{
		IAM: m,
	}
}

func TestIamGroups(t *testing.T) {
	client.AwsMockTestHelper(t, IamGroups(), buildIamGroups, client.TestOptions{})
}
