package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudtrail"
	cloudtrailTypes "github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudtrailTrailsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudtrailClient(ctrl)
	services := client.Services{
		Cloudtrail: m,
	}
	trail := cloudtrailTypes.Trail{}
	err := faker.FakeData(&trail)
	if err != nil {
		t.Fatal(err)
	}
	trailStatus := cloudtrail.GetTrailStatusOutput{}
	err = faker.FakeData(&trailStatus)
	if err != nil {
		t.Fatal(err)
	}
	eventSelector := cloudtrailTypes.EventSelector{}
	err = faker.FakeData(&eventSelector)
	if err != nil {
		t.Fatal(err)
	}
	m.EXPECT().DescribeTrails(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.DescribeTrailsOutput{
			TrailList: []cloudtrailTypes.Trail{trail},
		},
		nil,
	)
	m.EXPECT().GetTrailStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&trailStatus,
		nil,
	)
	m.EXPECT().GetEventSelectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudtrail.GetEventSelectorsOutput{
			EventSelectors: []cloudtrailTypes.EventSelector{eventSelector},
		},
		nil,
	)
	tags := cloudtrail.ListTagsOutput{}
	err = faker.FakeData(&tags)
	if err != nil {
		t.Fatal(err)
	}
	tags.ResourceTagList[0].ResourceId = trail.TrailARN
	tags.NextToken = nil
	m.EXPECT().ListTags(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)

	return services
}

func TestCloudtrailTrails(t *testing.T) {
	awsTestHelper(t, CloudtrailTrails(), buildCloudtrailTrailsMock, TestOptions{})
}
