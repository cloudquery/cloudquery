package route53

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildRoute53HostedZonesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	h := types.HostedZone{}
	if err := faker.FakeObject(&h); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListHostedZones(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListHostedZonesOutput{
			HostedZones: []types.HostedZone{h},
		}, nil)
	tag := types.Tag{}
	if err := faker.FakeObject(&tag); err != nil {
		t.Fatal(err)
	}
	//create id that is usually returned by aws
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
	if err := faker.FakeObject(&qlc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListQueryLoggingConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListQueryLoggingConfigsOutput{
			QueryLoggingConfigs: []types.QueryLoggingConfig{qlc},
		}, nil)
	rrs := types.ResourceRecordSet{}
	if err := faker.FakeObject(&rrs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListResourceRecordSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListResourceRecordSetsOutput{
			ResourceRecordSets: []types.ResourceRecordSet{rrs},
		}, nil)
	tpi := types.TrafficPolicyInstance{}
	if err := faker.FakeObject(&tpi); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTrafficPolicyInstancesByHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPolicyInstancesByHostedZoneOutput{
			TrafficPolicyInstances: []types.TrafficPolicyInstance{tpi},
		}, nil)
	vpc := types.VPC{}
	if err := faker.FakeObject(&vpc); err != nil {
		t.Fatal(err)
	}
	ds := types.DelegationSet{}
	if err := faker.FakeObject(&ds); err != nil {
		t.Fatal(err)
	}

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
