package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/rs/zerolog"
)

func findEnabledRegions(ctx context.Context, logger zerolog.Logger, accountName string, ec2Client services.Ec2Client, localRegions []string, accountDefaultRegion string) []string {
	// By default we should use the default region (us-east-1)
	regionsToCheck := []string{defaultRegion}
	// If user specifies a Default Region we should use it
	if accountDefaultRegion != "" {
		regionsToCheck = []string{accountDefaultRegion}
		// If no default region and * is not specified we should use all specified regions
	} else if len(localRegions) > 0 && !isAllRegions(localRegions) {
		regionsToCheck = localRegions
	}

	for _, region := range regionsToCheck {
		enabledRegions, err := getEnabledRegions(ctx, ec2Client, region)
		if err != nil {
			logger.Warn().Str("account", accountName).Err(err).Msgf("Failed to find disabled regions for account when checking: %s", region)
			continue
		}
		filteredRegions := filterDisabledRegions(localRegions, enabledRegions)
		if len(filteredRegions) > 0 {
			return filteredRegions
		}
	}
	return []string{}
}

func getEnabledRegions(ctx context.Context, ec2Client services.Ec2Client, region string) ([]types.Region, error) {
	res, err := ec2Client.DescribeRegions(ctx,
		&ec2.DescribeRegionsInput{AllRegions: aws.Bool(false)},
		func(o *ec2.Options) {
			o.Region = region
		})
	if err != nil {
		return nil, err
	}
	return res.Regions, nil
}

func filterDisabledRegions(regions []string, enabledRegions []types.Region) []string {
	regionsMap := map[string]bool{}
	for _, r := range enabledRegions {
		if r.RegionName != nil && r.OptInStatus != nil && *r.OptInStatus != "not-opted-in" {
			regionsMap[*r.RegionName] = true
		}
	}

	var filteredRegions []string
	// Our list of regions might not always be the latest and most up to date list
	// if a user specifies all regions via a "*" then they should get the most broad list possible
	if isAllRegions(regions) {
		for region := range regionsMap {
			filteredRegions = append(filteredRegions, region)
		}
	} else {
		for _, r := range regions {
			if regionsMap[r] {
				filteredRegions = append(filteredRegions, r)
			}
		}
	}
	return filteredRegions
}

func getAccountId(ctx context.Context, awsCfg aws.Config) (*sts.GetCallerIdentityOutput, error) {
	svc := sts.NewFromConfig(awsCfg)
	return svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
}

func verifyRegions(regions []string) error {
	availableRegions, err := getAvailableRegions()
	if err != nil {
		return err
	}

	// validate regions values
	var hasWildcard bool
	for i, region := range regions {
		if region == "*" {
			hasWildcard = true
		}
		if i != 0 && region == "*" {
			return errInvalidRegion
		}
		if i > 0 && hasWildcard {
			return errInvalidRegion
		}
		regionExist := availableRegions[region]
		if !hasWildcard && !regionExist {
			return errUnknownRegion(region)
		}
	}
	return nil
}
func isAllRegions(regions []string) bool {
	// if regions array is not valid return false
	err := verifyRegions(regions)
	if err != nil {
		return false
	}

	wildcardAllRegions := false
	if (len(regions) == 1 && regions[0] == "*") || (len(regions) == 0) {
		wildcardAllRegions = true
	}
	return wildcardAllRegions
}
