package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotSecurityProfilesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	sp := iot.ListSecurityProfilesOutput{}
	err := faker.FakeData(&sp)
	if err != nil {
		t.Fatal(err)
	}
	sp.NextToken = nil
	m.EXPECT().ListSecurityProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sp, nil)

	profileOutput := iot.DescribeSecurityProfileOutput{}
	err = faker.FakeData(&profileOutput)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeSecurityProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&profileOutput, nil)

	targets := iot.ListTargetsForSecurityProfileOutput{}
	err = faker.FakeData(&targets)
	if err != nil {
		t.Fatal(err)
	}
	targets.NextToken = nil

	m.EXPECT().ListTargetsForSecurityProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&targets, nil)

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

func TestIotSecurityProfiles(t *testing.T) {
	client.AwsMockTestHelper(t, IotSecurityProfiles(), buildIotSecurityProfilesMock, client.TestOptions{})
}
