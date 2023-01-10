package ec2

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
)

func buildEc2RegionalConfig(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockEc2Client(ctrl)
	m.EXPECT().GetEbsDefaultKmsKeyId(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ec2.GetEbsDefaultKmsKeyIdOutput{KmsKeyId: aws.String("some/key/id")}, nil)
	m.EXPECT().GetEbsEncryptionByDefault(gomock.Any(), gomock.Any(), gomock.Any()).Return(&ec2.GetEbsEncryptionByDefaultOutput{EbsEncryptionByDefault: aws.Bool(true)}, nil)

	return client.Services{
		Ec2: m,
	}
}

func TestEc2RegionalConfig(t *testing.T) {
	client.AwsMockTestHelper(t, RegionalConfigs(), buildEc2RegionalConfig, client.TestOptions{})
}
