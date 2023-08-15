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

func buildCloudfrontResponseHeaderPoliciesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockCloudfrontClient(ctrl)
	services := client.Services{
		Cloudfront: m,
	}
	cp := cloudfrontTypes.ResponseHeadersPolicySummary{}
	require.NoError(t, faker.FakeObject(&cp))

	cloudfrontOutput := &cloudfront.ListResponseHeadersPoliciesOutput{
		ResponseHeadersPolicyList: &cloudfrontTypes.ResponseHeadersPolicyList{
			Items: []cloudfrontTypes.ResponseHeadersPolicySummary{cp},
		},
	}
	m.EXPECT().ListResponseHeadersPolicies(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		cloudfrontOutput,
		nil,
	)
	return services
}

func TestCloudfrontResponseHeaderPolicies(t *testing.T) {
	client.AwsMockTestHelper(t, ResponseHeaderPolicies(), buildCloudfrontResponseHeaderPoliciesMock, client.TestOptions{})
}
