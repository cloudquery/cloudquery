package guardduty

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	gdTypes "github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildDetectors(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGuarddutyClient(ctrl)

	var d guardduty.GetDetectorOutput
	if err := faker.FakeObject(&d); err != nil {
		t.Fatal(err)
	}
	d.CreatedAt = aws.String(time.Now().Format(time.RFC3339))
	d.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))

	m.EXPECT().ListDetectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListDetectorsOutput{
			DetectorIds: []string{""},
		}, nil)

	m.EXPECT().GetDetector(gomock.Any(), gomock.Any(), gomock.Any()).Return(&d, nil)

	var finding gdTypes.Finding
	if err := faker.FakeObject(&finding); err != nil {
		t.Fatal(err)
	}
	finding.CreatedAt = aws.String(time.Now().Format(time.RFC3339))
	finding.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))

	finding.Id = aws.String("test-finding")
	m.EXPECT().ListFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListFindingsOutput{FindingIds: []string{*finding.Id}}, nil,
	)
	m.EXPECT().GetFindings(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.GetFindingsOutput{Findings: []gdTypes.Finding{finding}}, nil,
	)

	var filter guardduty.GetFilterOutput
	if err := faker.FakeObject(&filter); err != nil {
		t.Fatal(err)
	}
	filter.Name = aws.String("test-filter")
	m.EXPECT().ListFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListFiltersOutput{FilterNames: []string{*filter.Name}}, nil,
	)
	m.EXPECT().GetFilter(gomock.Any(), gomock.Any(), gomock.Any()).Return(&filter, nil)

	var ipset guardduty.GetIPSetOutput
	if err := faker.FakeObject(&ipset); err != nil {
		t.Fatal(err)
	}
	ipset.Name = aws.String("test-ipset")
	m.EXPECT().ListIPSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListIPSetsOutput{IpSetIds: []string{*ipset.Name}}, nil,
	)
	m.EXPECT().GetIPSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ipset, nil)

	var dest gdTypes.Destination
	if err := faker.FakeObject(&dest); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListPublishingDestinations(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListPublishingDestinationsOutput{Destinations: []gdTypes.Destination{dest}}, nil,
	)

	var tset guardduty.GetThreatIntelSetOutput
	if err := faker.FakeObject(&tset); err != nil {
		t.Fatal(err)
	}
	tset.Name = aws.String("test-threatset")
	m.EXPECT().ListThreatIntelSets(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListThreatIntelSetsOutput{ThreatIntelSetIds: []string{*tset.Name}}, nil,
	)
	m.EXPECT().GetThreatIntelSet(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tset, nil)

	var member gdTypes.Member
	if err := faker.FakeObject(&member); err != nil {
		t.Fatal(err)
	}
	member.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))
	member.InvitedAt = aws.String(time.Now().Format(time.RFC3339))
	m.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListMembersOutput{Members: []gdTypes.Member{member}}, nil,
	)
	return client.Services{
		Guardduty: m,
	}
}

func TestDetectors(t *testing.T) {
	client.AwsMockTestHelper(t, Detectors(), buildDetectors, client.TestOptions{})
}
