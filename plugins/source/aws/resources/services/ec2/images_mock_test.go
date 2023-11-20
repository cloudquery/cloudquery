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
	services := client.Services{Ec2: m}
	g := types.Image{}
	require.NoError(t, faker.FakeObject(&g))

	creationDate := "2019-05-10T13:17:12.000Z" // from docs
	g.OwnerId = aws.String("testAccount")
	g.CreationDate = &creationDate
	deprecationTime := "2050-11-05T08:15:30-05:00"
	g.DeprecationTime = &deprecationTime

	m.EXPECT().
		DescribeImages(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&ec2.DescribeImagesOutput{Images: []types.Image{g}}, nil).
		Times(2)

	lp := types.LaunchPermission{}
	require.NoError(t, faker.FakeObject(&lp))

	m.EXPECT().
		DescribeImageAttribute(gomock.Any(), &ec2.DescribeImageAttributeInput{Attribute: types.ImageAttributeNameLaunchPermission, ImageId: g.ImageId}, gomock.Any()).
		Return(&ec2.DescribeImageAttributeOutput{LaunchPermissions: []types.LaunchPermission{lp}}, nil)

	m.EXPECT().
		DescribeImageAttribute(gomock.Any(), &ec2.DescribeImageAttributeInput{Attribute: types.ImageAttributeNameLastLaunchedTime, ImageId: g.ImageId}, gomock.Any()).
		Return(&ec2.DescribeImageAttributeOutput{LastLaunchedTime: &types.AttributeValue{Value: &creationDate}}, nil)

	return services
}

func TestEc2Images(t *testing.T) {
	client.AwsMockTestHelper(t, Images(), buildEc2ImagesMock, client.TestOptions{})
}
