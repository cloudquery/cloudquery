package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildEc2ImagesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	services := client.Services{
		EC2: m,
	}
	g := types.Image{}
	err := faker.FakeData(&g)
	if err != nil {
		t.Fatal(err)
	}

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeImagesOutput{
			Images: []types.Image{g},
		}, nil)

	lastLaunched := "1994-11-05T08:15:30-05:00"
	m.EXPECT().DescribeImageAttribute(
		gomock.Any(),
		&ec2.DescribeImageAttributeInput{
			Attribute: types.ImageAttributeNameLastLaunchedTime,
			ImageId:   g.ImageId,
		},
		gomock.Any(),
	).Return(
		&ec2.DescribeImageAttributeOutput{
			LastLaunchedTime: &types.AttributeValue{Value: &lastLaunched},
		},
		nil,
	)
	return services
}

func TestEc2Images(t *testing.T) {
	client.AwsMockTestHelper(t, Ec2Images(), buildEc2ImagesMock, client.TestOptions{})
}
