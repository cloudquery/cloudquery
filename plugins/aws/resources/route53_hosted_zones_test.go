package resources

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53Types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildRoute53HostedZonesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockRoute53Client(ctrl)
	h := route53Types.HostedZone{}
	if err := faker.FakeData(&h); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListHostedZones(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListHostedZonesOutput{
			HostedZones: []route53Types.HostedZone{h},
		}, nil)
	tag := route53Types.Tag{}
	if err := faker.FakeData(&tag); err != nil {
		t.Fatal(err)
	}
	//create id that is usually returned by aws
	hzId := *h.Id
	newId := fmt.Sprintf("/%s/%s", route53Types.TagResourceTypeHostedzone, *h.Id)
	h.Id = &newId
	m.EXPECT().ListTagsForResources(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTagsForResourcesOutput{
			ResourceTagSets: []route53Types.ResourceTagSet{
				{
					ResourceId: &hzId,
					Tags:       []route53Types.Tag{tag},
				},
			},
		}, nil)
	qlc := route53Types.QueryLoggingConfig{}
	if err := faker.FakeData(&qlc); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListQueryLoggingConfigs(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListQueryLoggingConfigsOutput{
			QueryLoggingConfigs: []route53Types.QueryLoggingConfig{qlc},
		}, nil)
	rrs := route53Types.ResourceRecordSet{}
	if err := faker.FakeData(&rrs); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListResourceRecordSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListResourceRecordSetsOutput{
			ResourceRecordSets: []route53Types.ResourceRecordSet{rrs},
		}, nil)
	tpi := route53Types.TrafficPolicyInstance{}
	if err := faker.FakeData(&tpi); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListTrafficPolicyInstancesByHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.ListTrafficPolicyInstancesByHostedZoneOutput{
			TrafficPolicyInstances: []route53Types.TrafficPolicyInstance{tpi},
		}, nil)
	vpc := route53Types.VPC{}
	if err := faker.FakeData(&vpc); err != nil {
		t.Fatal(err)
	}
	ds := route53Types.DelegationSet{}
	if err := faker.FakeData(&ds); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetHostedZone(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&route53.GetHostedZoneOutput{
			HostedZone:    &h,
			DelegationSet: &ds,
			VPCs:          []route53Types.VPC{vpc},
		}, nil)
	return client.Services{
		Route53: m,
	}
}

func TestRoute53HostedZones(t *testing.T) {
	awsTestHelper(t, Route53HostedZones(), buildRoute53HostedZonesMock)
}
