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
	creationDate := "1994-11-05T08:15:30-05:00"
	l.Instances[0].ElasticGpuAssociations[0].ElasticGpuAssociationTime = &creationDate
	nextToken := "test"
	// this test ensures that pagination works by returning a token once, but not the second time
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(
		&ec2.DescribeInstancesOutput{
			Reservations: []ec2Types.Reservation{},
			NextToken:    &nextToken,
		}, nil)
	m.EXPECT().DescribeInstances(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(
		&ec2.DescribeInstancesOutput{
			Reservations: []ec2Types.Reservation{l},
			NextToken:    nil,
		}, nil)
	return client.Services{
		EC2: m,
	}
}

func TestEc2Instances(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2Instances(), buildEc2Instances, client.TestOptions{})
}
