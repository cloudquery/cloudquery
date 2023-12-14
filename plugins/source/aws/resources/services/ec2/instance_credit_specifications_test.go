package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func buildInstanceCreditSpecifications(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	creditSpecification := types.InstanceCreditSpecification{}
	err := faker.FakeObject(&creditSpecification)
	require.NoError(t, err)

	m.EXPECT().DescribeInstanceCreditSpecifications(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstanceCreditSpecificationsOutput{
			InstanceCreditSpecifications: []types.InstanceCreditSpecification{creditSpecification},
			NextToken:                    nil,
		},
		nil,
	)
	return client.Services{
		Ec2: m,
	}
}

func TestInstanceCreditSpecifications(t *testing.T) {
	client.AwsMockTestHelper(t, InstanceCreditSpecifications(), buildInstanceCreditSpecifications, client.TestOptions{})
}
