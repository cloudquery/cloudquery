package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildEc2ImagesMock(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	services := client.Services{
		Ec2: m,
	}
	image := types.Image{}
	require.NoError(t, faker.FakeObject(&image))

	creationDate := "1994-11-05T08:15:30-05:00"
	image.OwnerId = aws.String("testAccount")
	image.CreationDate = &creationDate
	deprecationTime := "2050-11-05T08:15:30-05:00"
	image.DeprecationTime = &deprecationTime

	m.EXPECT().DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&ec2.DescribeImagesOutput{
			Images: []types.Image{image},
		}, nil).Times(2)

	lp := types.LaunchPermission{}
	require.NoError(t, faker.FakeObject(&lp))

	m.EXPECT().DescribeImageAttribute(
		gomock.Any(),
		&ec2.DescribeImageAttributeInput{
			Attribute: types.ImageAttributeNameLaunchPermission,
			ImageId:   image.ImageId,
		},
		gomock.Any(),
	).Return(
		&ec2.DescribeImageAttributeOutput{
			LaunchPermissions: []types.LaunchPermission{lp},
		},
		nil,
	)
	m.EXPECT().DescribeImageAttribute(
		gomock.Any(),
		&ec2.DescribeImageAttributeInput{
			Attribute: types.ImageAttributeNameLastLaunchedTime,
			ImageId:   image.ImageId,
		},
		gomock.Any(),
	).Return(
		&ec2.DescribeImageAttributeOutput{
			LastLaunchedTime: &types.AttributeValue{Value: &creationDate},
		},
		nil,
	)

	return services
}

func TestEc2Images(t *testing.T) {
	client.AwsMockTestHelper(t, Images(), buildEc2ImagesMock, client.TestOptions{})
}
