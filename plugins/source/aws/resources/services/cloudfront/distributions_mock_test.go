package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildCloudfrontDistributionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	ds := cloudfrontTypes.DistributionSummary{}
	if err := faker.FakeObject(&ds); err != nil {
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
	if err := faker.FakeObject(&distribution); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().GetDistribution(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		distribution,
		nil,
	)

	tags := &cloudfront.ListTagsForResourceOutput{}
	if err := faker.FakeObject(&tags); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		tags,
		nil,
	)
	return services
}

func TestCloudfrontDistributions(t *testing.T) {
	client.AwsMockTestHelper(t, Distributions(), buildCloudfrontDistributionsMock, client.TestOptions{})
}
