package client

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
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
