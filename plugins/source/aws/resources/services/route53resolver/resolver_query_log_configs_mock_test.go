package route53resolver

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53resolver"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildResolverQueryLogConfigsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53resolverClient(ctrl)
	rqlc := types.ResolverQueryLogConfig{}
	require.NoError(t, faker.FakeObject(&rqlc))

	m.EXPECT().ListResolverQueryLogConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53resolver.ListResolverQueryLogConfigsOutput{
			ResolverQueryLogConfigs: []types.ResolverQueryLogConfig{rqlc},
		}, nil)

	return client.Services{
		Route53resolver: m,
	}
}
func TestResolverQueryLogConfigs(t *testing.T) {
	client.AwsMockTestHelper(t, ResolverQueryLogConfigs(), buildResolverQueryLogConfigsMock, client.TestOptions{})
}
