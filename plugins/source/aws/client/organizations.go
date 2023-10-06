package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/spec"
	"github.com/thoas/go-funk"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/rs/zerolog"
)

// Parses org configuration and grabs the appropriate accounts
func loadOrgAccounts(ctx context.Context, logger zerolog.Logger, awsPluginSpec *spec.Spec) ([]spec.Account, AssumeRoleAPIClient, error) {
	// If user doesn't specify any configs for admin account instantiate default values
	if awsPluginSpec.Organization.AdminAccount == nil {
		awsPluginSpec.Organization.AdminAccount = &spec.Account{
			AccountName:  "Default-Admin-Account",
			LocalProfile: "",
		}
	}
	awsCfg, err := ConfigureAwsSDK(ctx, logger, awsPluginSpec, *awsPluginSpec.Organization.AdminAccount, nil)
	if err != nil {
		return nil, nil, err
	}
	svc := organizations.NewFromConfig(awsCfg)
	region := awsCfg.Region
	if region == "" {
		region = defaultRegion
	}
	accounts, err := loadAccounts(ctx, awsPluginSpec, svc, region)
	if err != nil {
		return nil, nil, err
	}
	if awsPluginSpec.Organization.MemberCredentials != nil {
		awsCfg, err = ConfigureAwsSDK(ctx, logger, awsPluginSpec, *awsPluginSpec.Organization.MemberCredentials, nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return accounts, sts.NewFromConfig(awsCfg), err
}

// Load accounts from the appropriate endpoint as well as normalizing response
func loadAccounts(ctx context.Context, awsPluginSpec *spec.Spec, accountsApi services.OrganizationsClient, region string) ([]spec.Account, error) {
	var rawAccounts []orgTypes.Account
	var err error
	if len(awsPluginSpec.Organization.OrganizationUnits) > 0 {
		rawAccounts, err = getOUAccounts(ctx, accountsApi, awsPluginSpec.Organization, region)
	} else {
		rawAccounts, err = getAllAccounts(ctx, accountsApi, awsPluginSpec.Organization, region)
	}

	if err != nil {
		return []spec.Account{}, err
	}
	seen := map[string]struct{}{}
	accounts := make([]spec.Account, 0)
	for _, account := range rawAccounts {
		// Only load Active accounts
		if account.Status != orgTypes.AccountStatusActive || account.Id == nil {
			continue
		}

		// Skip duplicates
		if _, found := seen[*account.Id]; found {
			continue
		}
		seen[*account.Id] = struct{}{}

		roleArn := arn.ARN{
			Partition: "aws",
			Service:   "iam",
			Region:    "",
			AccountID: *account.Id,
			Resource:  "role/" + awsPluginSpec.Organization.ChildAccountRoleName,
		}
		if parsed, err := arn.Parse(aws.ToString(account.Arn)); err == nil {
			roleArn.Partition = parsed.Partition
		}

		accounts = append(accounts, spec.Account{
			ID:              *account.Id,
			RoleARN:         roleArn.String(),
			RoleSessionName: awsPluginSpec.Organization.ChildAccountRoleSessionName,
			ExternalID:      awsPluginSpec.Organization.ChildAccountExternalID,
			LocalProfile:    awsPluginSpec.Organization.AdminAccount.LocalProfile,
			Regions:         awsPluginSpec.Organization.ChildAccountRegions,
			Source:          spec.AccountSourceOrg,
		})
	}
	return accounts, err
}

// Get Accounts for specific Organizational Units
func getOUAccounts(ctx context.Context, accountsApi services.OrganizationsClient, awsOrg *spec.Org, region string) ([]orgTypes.Account, error) {
	q := awsOrg.OrganizationUnits
	var ou string
	var rawAccounts []orgTypes.Account
	seenOUs := map[string]struct{}{}
	for len(q) > 0 {
		ou, q = q[0], q[1:]

		// Skip duplicates to avoid making duplicate API calls
		if _, found := seenOUs[ou]; found {
			continue
		}
		seenOUs[ou] = struct{}{}

		// Skip any OUs that user has asked to skip
		if funk.ContainsString(awsOrg.SkipOrganizationalUnits, ou) {
			continue
		}

		// get accounts directly under this OU
		accountsPaginator := organizations.NewListAccountsForParentPaginator(accountsApi, &organizations.ListAccountsForParentInput{
			ParentId: aws.String(ou),
		})
		for accountsPaginator.HasMorePages() {
			output, err := accountsPaginator.NextPage(ctx, func(options *organizations.Options) {
				options.Region = region
			})
			if err != nil {
				return nil, err
			}
			for _, account := range output.Accounts {
				// Skip any accounts that user has asked to skip
				if funk.ContainsString(awsOrg.SkipMemberAccounts, *account.Id) {
					continue
				}
				rawAccounts = append(rawAccounts, account)
			}
		}

		// get OUs directly under this OU, and add them to the queue
		ouPaginator := organizations.NewListChildrenPaginator(accountsApi, &organizations.ListChildrenInput{
			ChildType: orgTypes.ChildTypeOrganizationalUnit,
			ParentId:  aws.String(ou),
		})
		for ouPaginator.HasMorePages() {
			output, err := ouPaginator.NextPage(ctx, func(options *organizations.Options) {
				options.Region = region
			})
			if err != nil {
				return nil, err
			}
			for _, child := range output.Children {
				q = append(q, *child.Id)
			}
		}
	}

	return rawAccounts, nil
}

// Get All accounts in a specific organization
func getAllAccounts(ctx context.Context, accountsApi services.OrganizationsClient, org *spec.Org, region string) ([]orgTypes.Account, error) {
	var rawAccounts []orgTypes.Account
	accountsPaginator := organizations.NewListAccountsPaginator(accountsApi, &organizations.ListAccountsInput{})
	for accountsPaginator.HasMorePages() {
		output, err := accountsPaginator.NextPage(ctx, func(options *organizations.Options) {
			options.Region = region
		})
		if err != nil {
			return nil, err
		}
		for _, account := range output.Accounts {
			// Skip any accounts that user has asked to skip
			if funk.ContainsString(org.SkipMemberAccounts, *account.Id) {
				continue
			}
			rawAccounts = append(rawAccounts, account)
		}
	}
	return rawAccounts, nil
}
