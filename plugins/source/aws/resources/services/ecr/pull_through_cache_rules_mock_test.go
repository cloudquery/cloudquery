package ecr

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildPullThroughCacheRules(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEcrClient(ctrl)

	ptcr := types.PullThroughCacheRule{}
	require.NoError(t, faker.FakeObject(&ptcr))

	m.EXPECT().DescribePullThroughCacheRules(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ecr.DescribePullThroughCacheRulesOutput{
			PullThroughCacheRules: []types.PullThroughCacheRule{ptcr},
		}, nil)

	return client.Services{
		Ecr: m,
	}
}

func TestPullThroughCacheRules(t *testing.T) {
	client.AwsMockTestHelper(t, PullThroughCacheRules(), buildPullThroughCacheRules, client.TestOptions{})
}
