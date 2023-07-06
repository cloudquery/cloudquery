package route53

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildRoute53HostedZonesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	h := types.HostedZone{}
	require.NoError(t, faker.FakeObject(&h))

	m.EXPECT().ListHostedZones(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListHostedZonesOutput{
			HostedZones: []types.HostedZone{h},
		}, nil)
	tag := types.Tag{}
	require.NoError(t, faker.FakeObject(&tag))

	hzId := *h.Id
	newId := fmt.Sprintf("/%s/%s", types.TagResourceTypeHostedzone, *h.Id)
	h.Id = &newId
	m.EXPECT().ListTagsForResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTagsForResourcesOutput{
			ResourceTagSets: []types.ResourceTagSet{
				{
					ResourceId: &hzId,
					Tags:       []types.Tag{tag},
				},
			},
		}, nil)
	qlc := types.QueryLoggingConfig{}
	require.NoError(t, faker.FakeObject(&qlc))

	m.EXPECT().ListQueryLoggingConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListQueryLoggingConfigsOutput{
			QueryLoggingConfigs: []types.QueryLoggingConfig{qlc},
		}, nil)
	rrs := types.ResourceRecordSet{}
	require.NoError(t, faker.FakeObject(&rrs))

	m.EXPECT().ListResourceRecordSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListResourceRecordSetsOutput{
			ResourceRecordSets: []types.ResourceRecordSet{rrs},
		}, nil)
	tpi := types.TrafficPolicyInstance{}
	require.NoError(t, faker.FakeObject(&tpi))

	m.EXPECT().ListTrafficPolicyInstancesByHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPolicyInstancesByHostedZoneOutput{
			TrafficPolicyInstances: []types.TrafficPolicyInstance{tpi},
		}, nil)
	vpc := types.VPC{}
	require.NoError(t, faker.FakeObject(&vpc))

	ds := types.DelegationSet{}
	require.NoError(t, faker.FakeObject(&ds))

	m.EXPECT().GetHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.GetHostedZoneOutput{
			HostedZone:    &h,
			DelegationSet: &ds,
			VPCs:          []types.VPC{vpc},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func TestRoute53HostedZones(t *testing.T) {
	client.AwsMockTestHelper(t, HostedZones(), buildRoute53HostedZonesMock, client.TestOptions{})
}
