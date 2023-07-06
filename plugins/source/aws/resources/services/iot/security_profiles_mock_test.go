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

func buildIotSecurityProfilesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	sp := iot.ListSecurityProfilesOutput{}
	require.NoError(t, faker.FakeObject(&sp))
	sp.NextToken = nil
	m.EXPECT().ListSecurityProfiles(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&sp, nil)

	profileOutput := iot.DescribeSecurityProfileOutput{}
	require.NoError(t, faker.FakeObject(&profileOutput))
	m.EXPECT().DescribeSecurityProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&profileOutput, nil)

	targets := iot.ListTargetsForSecurityProfileOutput{}
	require.NoError(t, faker.FakeObject(&targets))
	targets.NextToken = nil

	m.EXPECT().ListTargetsForSecurityProfile(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&targets, nil)

	tags := iot.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotSecurityProfiles(t *testing.T) {
	client.AwsMockTestHelper(t, SecurityProfiles(), buildIotSecurityProfilesMock, client.TestOptions{})
}
