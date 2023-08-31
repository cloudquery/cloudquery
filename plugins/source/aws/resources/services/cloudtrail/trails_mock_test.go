package cloudtrail

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildCloudtrailTrailsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	trail := types.Trail{}
	require.NoError(t, faker.FakeObject(&trail))

	trail.TrailARN = aws.String("arn:aws:cloudtrail:eu-central-1:testAccount:trail/test-trail")
	trail.CloudWatchLogsLogGroupArn = aws.String("arn:aws:logs:eu-central-1:123:log-group:test-group:")

	trailStatus := cloudtrail.GetTrailStatusOutput{}
	require.NoError(t, faker.FakeObject(&trailStatus))
	eventSelector := types.EventSelector{}
	require.NoError(t, faker.FakeObject(&eventSelector))
	m.EXPECT().DescribeTrails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.DescribeTrailsOutput{
			TrailList: []types.Trail{trail},
		},
		nil,
	)
	m.EXPECT().GetTrailStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&trailStatus,
		nil,
	)
	m.EXPECT().GetEventSelectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.GetEventSelectorsOutput{
			EventSelectors: []types.EventSelector{eventSelector},
		},
		nil,
	)
	tags := cloudtrail.ListTagsOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	tags.ResourceTagList[0].ResourceId = trail.TrailARN
	tags.NextToken = nil
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return services
}

func TestCloudtrailTrails(t *testing.T) {
	client.AwsMockTestHelper(t, Trails(), buildCloudtrailTrailsMock, client.TestOptions{})
}
