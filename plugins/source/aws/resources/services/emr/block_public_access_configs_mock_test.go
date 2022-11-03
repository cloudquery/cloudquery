package emr

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/emr"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
)

func buildEMRClient(t *testing.T, ctrl *gomock.Controller) client.Services {
	emrmock := mocks.NewMockEmrClient(ctrl)

	out := &emr.GetBlockPublicAccessConfigurationOutput{
		BlockPublicAccessConfiguration: &types.BlockPublicAccessConfiguration{
			Classification: aws.String("classification"),
			Configurations: []types.Configuration{},
			PermittedPublicSecurityGroupRuleRanges: []types.PortRange{
				{
					MinRange: aws.Int32(1024),
					MaxRange: aws.Int32(2048),
				},
			},
			Properties: map[string]string{
				"key": "value",
			},
		},
		BlockPublicAccessConfigurationMetadata: &types.BlockPublicAccessConfigurationMetadata{
			CreatedByArn:     aws.String("justsomevalue"),
			CreationDateTime: aws.Time(time.Now()),
		},
	}
	emrmock.EXPECT().GetBlockPublicAccessConfiguration(
		gomock.Any(),
		&emr.GetBlockPublicAccessConfigurationInput{},
		gomock.Any(),
	).Return(out, nil)
	return client.Services{Emr: emrmock}
}

func TestEMRBlockPublicAccessConfigs(t *testing.T) {
	client.AwsMockTestHelper(t, BlockPublicAccessConfigs(), buildEMRClient, client.TestOptions{})
}
