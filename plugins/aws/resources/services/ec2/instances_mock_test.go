// +build mock

package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2Types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2Instances(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	l := ec2Types.Reservation{}
	err := faker.FakeData(&l)
	if err != nil {
		t.Fatal(err)
	}
	l.Instances[0].StateTransitionReason = aws.String("User initiated (2021-11-26 11:33:00 GMT)")
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeInstancesOutput{
			Reservations: []ec2Types.Reservation{l},
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2Instances(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2Instances(), buildEc2Instances, client.TestOptions{})
}
