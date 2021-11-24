package resources

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/guardduty"
	gdTypes "github.com/aws/aws-sdk-go-v2/service/guardduty/types"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildGuardDutyDetectors(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockGuardDutyClient(ctrl)

	var d guardduty.GetDetectorOutput
	if err := faker.FakeData(&d); err != nil {
		t.Fatal(err)
	}
	d.CreatedAt = aws.String(time.Now().Format(time.RFC3339))
	d.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))

	m.EXPECT().ListDetectors(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListDetectorsOutput{
			DetectorIds: []string{""},
		}, nil)

	m.EXPECT().GetDetector(gomock.Any(), gomock.Any(), gomock.Any()).Return(&d, nil)

	var member gdTypes.Member
	if err := faker.FakeData(&member); err != nil {
		t.Fatal(err)
	}
	member.UpdatedAt = aws.String(time.Now().Format(time.RFC3339))
	member.InvitedAt = aws.String(time.Now().Format(time.RFC3339))

	m.EXPECT().ListMembers(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&guardduty.ListMembersOutput{Members: []gdTypes.Member{member}}, nil,
	)
	return client.Services{
		GuardDuty: m,
	}
}

func TestGuarddutyDetectors(t *testing.T) {
	awsTestHelper(t, GuarddutyDetectors(), buildGuardDutyDetectors, TestOptions{})
}
