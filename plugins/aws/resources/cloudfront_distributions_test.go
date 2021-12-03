package resources

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildCloudfrontDistributionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	ds := cloudfrontTypes.DistributionSummary{}
	if err := faker.FakeData(&ds); err != nil {
		t.Fatal(err)
	}
	cloudfrontOutput := &cloudfront.ListDistributionsOutput{
		DistributionList: &cloudfrontTypes.DistributionList{
			Items: []cloudfrontTypes.DistributionSummary{ds},
		},
	}
	m.EXPECT().ListDistributions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)

	distribution := &cloudfront.GetDistributionOutput{}
	if err := faker.FakeData(&distribution); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDistribution(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		distribution,
		nil,
	)

	tags := &cloudfront.ListTagsForResourceOutput{}
	if err := faker.FakeData(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestCloudfrontDistributions(t *testing.T) {
	awsTestHelper(t, CloudfrontDistributions(), buildCloudfrontDistributionsMock, TestOptions{})
}
