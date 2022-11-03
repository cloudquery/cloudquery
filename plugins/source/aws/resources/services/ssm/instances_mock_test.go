package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
)

func buildSSMInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.InstanceInformation
	if err := faker.FakeObject(&i); err != nil {
		t.Fatal(err)
	}
	i.IPAddress = aws.String("192.168.1.1")
	mock.EXPECT().DescribeInstanceInformation(
		gomock.Any(),
		&ssm.DescribeInstanceInformationInput{},
		gomock.Any(),
	).Return(
		&ssm.DescribeInstanceInformationOutput{InstanceInformationList: []types.InstanceInformation{i}},
		nil,
	)

	var c types.ComplianceItem
	if err := faker.FakeObject(&c); err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListComplianceItems(gomock.Any(),
		&ssm.ListComplianceItemsInput{ResourceIds: []string{*i.InstanceId}},
		gomock.Any(),
	).Return(
		&ssm.ListComplianceItemsOutput{ComplianceItems: []types.ComplianceItem{c}},
		nil,
	)
	return client.Services{Ssm: mock}
}

func TestSSMInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildSSMInstances, client.TestOptions{})
}
