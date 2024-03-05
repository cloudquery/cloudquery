package client

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog"
)

func Test_findEnabledRegions(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		regionsReply           *ec2.DescribeRegionsOutput
		regionsReplyError      error
		requestedRegions       []string
		requestedDefaultRegion string
		expectedRegions        []string
	}{
		{
			// order should not matter
			regionsReply: &ec2.DescribeRegionsOutput{
				Regions: []types.Region{
					{
						OptInStatus: aws.String("opt-in-not-required"),
						RegionName:  aws.String("us-east-1"),
					}, {
						OptInStatus: aws.String("not-opted-in"),
						RegionName:  aws.String("us-east-5"),
					},
				},
			},
			regionsReplyError:      nil,
			requestedRegions:       []string{"us-east-5", "us-east-1"},
			requestedDefaultRegion: "",
			expectedRegions:        []string{"us-east-1"},
		},
		{
			// order should not matter
			regionsReply: &ec2.DescribeRegionsOutput{
				Regions: []types.Region{
					{
						OptInStatus: aws.String("opt-in-not-required"),
						RegionName:  aws.String("us-east-1"),
					}, {
						OptInStatus: aws.String("not-opted-in"),
						RegionName:  aws.String("us-east-5"),
					},
				},
			},
			regionsReplyError:      nil,
			requestedRegions:       []string{"us-east-1", "us-east-5"},
			requestedDefaultRegion: "",
			expectedRegions:        []string{"us-east-1"},
		},
		{
			// User does not have the permissions to call DescribeRegions
			regionsReply:           nil,
			regionsReplyError:      &smithy.GenericAPIError{Code: "UnauthorizedOperation", Message: "You are not authorized to perform this operation. User: arn:aws:sts::012345678910:assumed-role/RoleName is not authorized to perform: ec2:DescribeRegions with an explicit deny in an identity-based policy"},
			requestedRegions:       []string{"us-east-1"},
			requestedDefaultRegion: "",
			expectedRegions:        []string{"us-east-1"},
		},
		{
			// User does not have the permissions to call DescribeRegions
			// If user is using "*" as requested region or a DefaultRegion we should still error out
			regionsReply:           nil,
			regionsReplyError:      &smithy.GenericAPIError{Code: "UnauthorizedOperation", Message: "You are not authorized to perform this operation. User: arn:aws:sts::012345678910:assumed-role/RoleName is not authorized to perform: ec2:DescribeRegions with an explicit deny in an identity-based policy"},
			requestedRegions:       []string{"*"},
			requestedDefaultRegion: "us-east-2",
			expectedRegions:        []string{},
		},
		{
			// User is making a request to a disabled region that returns an error
			regionsReply:           nil,
			regionsReplyError:      errors.New("region not enabled"),
			requestedRegions:       []string{"us-east-5"},
			requestedDefaultRegion: "",
			expectedRegions:        []string{},
		},
		{
			// User is making a request to a disabled region that returns an error (in DefaultRegion)
			regionsReply:           nil,
			regionsReplyError:      errors.New("region not enabled"),
			requestedRegions:       []string{},
			requestedDefaultRegion: "us-east-5",
			expectedRegions:        []string{},
		},
	}

	for _, test := range tests {
		ctrl := gomock.NewController(t)
		api := mocks.NewMockEc2Client(ctrl)
		api.EXPECT().DescribeRegions(gomock.Any(), gomock.Any(), gomock.Any()).Return(test.regionsReply, test.regionsReplyError)
		enabledRegions := findEnabledRegions(ctx, zerolog.New(os.Stderr).With().Logger(), "test", api, test.requestedRegions, test.requestedDefaultRegion)
		respDiff := cmp.Diff(test.expectedRegions, enabledRegions)

		if respDiff != "" {
			t.Fatal(respDiff)
		}
	}
}
