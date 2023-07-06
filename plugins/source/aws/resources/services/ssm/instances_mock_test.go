package ssm

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildSSMInstances(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockSsmClient(ctrl)

	var i types.InstanceInformation
	require.NoError(t, faker.FakeObject(&i))

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
	require.NoError(t, faker.FakeObject(&c))

	mock.EXPECT().ListComplianceItems(gomock.Any(),
		&ssm.ListComplianceItemsInput{ResourceIds: []string{*i.InstanceId}},
		gomock.Any(),
	).Return(
		&ssm.ListComplianceItemsOutput{ComplianceItems: []types.ComplianceItem{c}},
		nil,
	)

	var p types.PatchComplianceData
	require.NoError(t, faker.FakeObject(&p))

	mock.EXPECT().DescribeInstancePatches(gomock.Any(),
		&ssm.DescribeInstancePatchesInput{InstanceId: i.InstanceId},
		gomock.Any(),
	).Return(
		&ssm.DescribeInstancePatchesOutput{Patches: []types.PatchComplianceData{p}},
		nil,
	)

	return client.Services{Ssm: mock}
}

func TestSSMInstances(t *testing.T) {
	client.AwsMockTestHelper(t, Instances(), buildSSMInstances, client.TestOptions{})
}
