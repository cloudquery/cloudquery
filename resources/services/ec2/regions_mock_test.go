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

func buildRegionsMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	r := ec2Types.Region{}
	if err := faker.FakeData(&r); err != nil {
		t.Fatal(err)
	}
	r.OptInStatus = aws.String("opted-in")
	m.EXPECT().DescribeRegions(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeRegionsOutput{
			Regions: []ec2Types.Region{r},
		}, nil)

	return client.Services{
		EC2: m,
	}
}

func TestRegions(t *testing.T) {
	client.AwsMockTestHelper(t, AwsRegions(), buildRegionsMock, client.TestOptions{})
}
