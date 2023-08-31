package route53

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRoute53DelegationSetsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	ds := route53Types.DelegationSet{}
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().ListReusableDelegationSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListReusableDelegationSetsOutput{
			DelegationSets: []route53Types.DelegationSet{ds},
		}, nil)
	return client.Services{
		Route53: m,
	}
}
func TestRoute53DelegationSets(t *testing.T) {
	client.AwsMockTestHelper(t, DelegationSets(), buildRoute53DelegationSetsMock, client.TestOptions{})
}
