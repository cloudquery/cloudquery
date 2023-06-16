package wafv2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	cftypes "github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildWAFV2WebACLMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockWafv2Client(ctrl)
	cfm := mocks.NewMockCloudfrontClient(ctrl)

	tempWebACLSum := types.WebACLSummary{}
	require.NoError(t, faker.FakeObject(&tempWebACLSum))

	var tempResourceArns []string
	require.NoError(t, faker.FakeObject(&tempResourceArns))

	var tempTags []types.Tag
	require.NoError(t, faker.FakeObject(&tempTags))

	var loggingConfiguration types.LoggingConfiguration
	require.NoError(t, faker.FakeObject(&loggingConfiguration))

	rule := types.Rule{}
	require.NoError(t, faker.FakeObject(&rule))

	for _, scope := range []types.Scope{types.ScopeCloudfront, types.ScopeRegional} {
		tempWebACL := types.WebACL{}
		require.NoError(t, faker.FakeObject(&tempWebACL))
		m.EXPECT().ListWebACLs(gomock.Any(), &wafv2.ListWebACLsInput{
			Scope: scope,
			Limit: aws.Int32(100),
		}, gomock.Any()).Return(&wafv2.ListWebACLsOutput{
			WebACLs: []types.WebACLSummary{tempWebACLSum},
		}, nil)
		m.EXPECT().GetWebACL(gomock.Any(), &wafv2.GetWebACLInput{
			Id:    tempWebACLSum.Id,
			Name:  tempWebACLSum.Name,
			Scope: scope,
		}, gomock.Any()).Return(&wafv2.GetWebACLOutput{
			WebACL: &tempWebACL,
		}, nil)
		m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListTagsForResourceOutput{
			TagInfoForResource: &types.TagInfoForResource{TagList: tempTags},
		}, nil)
		m.EXPECT().GetLoggingConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.GetLoggingConfigurationOutput{
			LoggingConfiguration: &loggingConfiguration,
		}, nil)
	}
	m.EXPECT().ListResourcesForWebACL(gomock.Any(), gomock.Any(), gomock.Any()).Return(&wafv2.ListResourcesForWebACLOutput{
		ResourceArns: tempResourceArns,
	}, nil).MinTimes(1)

	distributionList := cftypes.DistributionList{}
	require.NoError(t, faker.FakeObject(&distributionList))

	distributionList.NextMarker = nil
	cfm.EXPECT().ListDistributionsByWebACLId(gomock.Any(), gomock.Any(), gomock.Any()).Return(&cloudfront.ListDistributionsByWebACLIdOutput{
		DistributionList: &distributionList,
	}, nil)

	return client.Services{Wafv2: m, Cloudfront: cfm}
}

func TestWafV2WebACL(t *testing.T) {
	client.AwsMockTestHelper(t, WebAcls(), buildWAFV2WebACLMock, client.TestOptions{})
}
