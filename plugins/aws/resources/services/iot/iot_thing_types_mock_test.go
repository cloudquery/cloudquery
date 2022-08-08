package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildIotThingTypesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIOTClient(ctrl)

	groupsOutput := iot.ListThingTypesOutput{}
	err := faker.FakeData(&groupsOutput)
	if err != nil {
		t.Fatal(err)
	}
	groupsOutput.NextToken = nil
	m.EXPECT().ListThingTypes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

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

func TestIotThingTypes(t *testing.T) {
	client.AwsMockTestHelper(t, IotThingTypes(), buildIotThingTypesMock, client.TestOptions{})
}
