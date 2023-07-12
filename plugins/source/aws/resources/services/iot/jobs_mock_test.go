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

func buildIotJobs(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIotClient(ctrl)

	lp := iot.ListJobsOutput{}
	require.NoError(t, faker.FakeObject(&lp))
	lp.NextToken = nil
	m.EXPECT().ListJobs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&lp, nil)

	p := iot.DescribeJobOutput{}
	require.NoError(t, faker.FakeObject(&p))
	m.EXPECT().DescribeJob(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&p, nil)

	tags := iot.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.NextToken = nil
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&tags, nil)

	return client.Services{
		Iot: m,
	}
}

func TestIotJobs(t *testing.T) {
	client.AwsMockTestHelper(t, Jobs(), buildIotJobs, client.TestOptions{})
}
