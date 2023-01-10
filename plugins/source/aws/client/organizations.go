package client

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"
	"github.com/thoas/go-funk"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/rs/zerolog"
)

// Parses org configuration and grabs the appropriate accounts
func loadOrgAccounts(ctx context.Context, logger zerolog.Logger, awsConfig *Spec) ([]Account, AssumeRoleAPIClient, error) {
	// If user doesn't specify any configs for admin account instantiate default values
	if awsConfig.Organization.AdminAccount == nil {
		awsConfig.Organization.AdminAccount = &Account{
			AccountName:  "Default-Admin-Account",
			LocalProfile: "",
		}
	}
	awsCfg, err := configureAwsClient(ctx, logger, awsConfig, *awsConfig.Organization.AdminAccount, nil)
	if err != nil {
		return nil, nil, err
	}
	svc := organizations.NewFromConfig(awsCfg)
	accounts, err := loadAccounts(ctx, awsConfig, svc)
	if err != nil {
		return nil, nil, err
	}
	if awsConfig.Organization.MemberCredentials != nil {
		awsCfg, err = configureAwsClient(ctx, logger, awsConfig, *awsConfig.Organization.MemberCredentials, nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return accounts, sts.NewFromConfig(awsCfg), err
}

// Load accounts from the appropriate endpoint as well as normalizing response
func loadAccounts(ctx context.Context, awsConfig *Spec, accountsApi services.OrganizationsClient) ([]Account, error) {
	var rawAccounts []orgTypes.Account
	var err error
	if len(awsConfig.Organization.OrganizationUnits) > 0 {
		rawAccounts, err = getOUAccounts(ctx, accountsApi, awsConfig.Organization)
	} else {
		rawAccounts, err = getAllAccounts(ctx, accountsApi, awsConfig.Organization)
	}

	if err != nil {
		return []Account{}, err
	}
	seen := map[string]struct{}{}
	accounts := make([]Account, 0)
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
			Resource:  "role/" + awsConfig.Organization.ChildAccountRoleName,
		}
		if parsed, err := arn.Parse(aws.ToString(account.Arn)); err == nil {
			roleArn.Partition = parsed.Partition
		}

		accounts = append(accounts, Account{
			ID:              *account.Id,
			RoleARN:         roleArn.String(),
			RoleSessionName: awsConfig.Organization.ChildAccountRoleSessionName,
			ExternalID:      awsConfig.Organization.ChildAccountExternalID,
			LocalProfile:    awsConfig.Organization.AdminAccount.LocalProfile,
			Regions:         awsConfig.Organization.ChildAccountRegions,
			source:          "org",
		})
	}
	return accounts, err
}

// Get Accounts for specific Organizational Units
func getOUAccounts(ctx context.Context, accountsApi services.OrganizationsClient, awsOrg *AwsOrg) ([]orgTypes.Account, error) {
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
			output, err := accountsPaginator.NextPage(ctx)
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
			output, err := ouPaginator.NextPage(ctx)
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
func getAllAccounts(ctx context.Context, accountsApi services.OrganizationsClient, org *AwsOrg) ([]orgTypes.Account, error) {
	var rawAccounts []orgTypes.Account
	accountsPaginator := organizations.NewListAccountsPaginator(accountsApi, &organizations.ListAccountsInput{})
	for accountsPaginator.HasMorePages() {
		output, err := accountsPaginator.NextPage(ctx)
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
