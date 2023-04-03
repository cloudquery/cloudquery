package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildIotPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	lp := iot.ListPoliciesOutput{}
	err := faker.FakeObject(&lp)
	if err != nil {
		t.Fatal(err)
	}
	lp.NextMarker = nil
	m.EXPECT().ListPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	p := iot.GetPolicyOutput{}
	err = faker.FakeObject(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	err = faker.FakeObject(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, Policies(), buildIotPolicies, client.TestOptions{})
}
