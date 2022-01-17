// +build mock

package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotPolicies(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	lp := iot.ListPoliciesOutput{}
	err := faker.FakeData(&lp)
	if err != nil {
		t.Fatal(err)
	}
	lp.NextMarker = nil
	m.EXPECT().ListPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	p := iot.GetPolicyOutput{}
	err = faker.FakeData(&p)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetPolicy(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		IOT: m,
	}
}

func TestIotPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, IotPolicies(), buildIotPolicies, client.TestOptions{})
}
