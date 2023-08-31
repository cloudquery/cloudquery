package iot

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildIotThingTypesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	groupsOutput := iot.ListThingTypesOutput{}
	require.NoError(t, faker.FakeObject(&groupsOutput))
	groupsOutput.NextToken = nil
	m.EXPECT().ListThingTypes(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&groupsOutput, nil)

	tags := iot.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotThingTypes(t *testing.T) {
	client.AwsMockTestHelper(t, ThingTypes(), buildIotThingTypesMock, client.TestOptions{})
}
