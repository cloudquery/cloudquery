package cloudfront

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cloudfrontTypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildOriginAccessIdentitiesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	coai := cloudfrontTypes.CloudFrontOriginAccessIdentitySummary{}
	require.NoError(t, faker.FakeObject(&coai))

	m.EXPECT().ListCloudFrontOriginAccessIdentities(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&cloudfront.ListCloudFrontOriginAccessIdentitiesOutput{
			CloudFrontOriginAccessIdentityList: &cloudfrontTypes.CloudFrontOriginAccessIdentityList{
				Items: []cloudfrontTypes.CloudFrontOriginAccessIdentitySummary{coai},
			},
		},
		nil,
	)
	return services
}

func TestCloudFrontOriginAccessIdentities(t *testing.T) {
	client.AwsMockTestHelper(t, OriginAccessIdentities(), buildOriginAccessIdentitiesMock, client.TestOptions{})
}
