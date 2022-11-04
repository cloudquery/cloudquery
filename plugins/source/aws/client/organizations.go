package client

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/services"

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
		rawAccounts, err = getOUAccounts(ctx, accountsApi, awsConfig.Organization.OrganizationUnits)
	} else {
		rawAccounts, err = getAllAccounts(ctx, accountsApi)
	}

	if err != nil {
		return []Account{}, err
	}
	accounts := make([]Account, 0)
	for _, account := range rawAccounts {
		// Only load Active accounts
		if account.Status != orgTypes.AccountStatusActive {
			continue
		}
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
func getOUAccounts(ctx context.Context, accountsApi services.OrganizationsClient, ous []string) ([]orgTypes.Account, error) {
	var rawAccounts []orgTypes.Account

	for _, ou := range ous {
		var paginationToken *string
		for {
			resp, err := accountsApi.ListAccountsForParent(ctx, &organizations.ListAccountsForParentInput{
				NextToken: paginationToken,
				ParentId:  aws.String(ou),
			})
			if err != nil {
				return nil, err
			}
			rawAccounts = append(rawAccounts, resp.Accounts...)
			if resp.NextToken == nil {
				break
			}
			paginationToken = resp.NextToken
		}
	}
	return rawAccounts, nil
}

// Get All accounts in a specific organization
func getAllAccounts(ctx context.Context, accountsApi services.OrganizationsClient) ([]orgTypes.Account, error) {
	var rawAccounts []orgTypes.Account
	var paginationToken *string

	for {
		resp, err := accountsApi.ListAccounts(ctx, &organizations.ListAccountsInput{
			NextToken: paginationToken,
		})
		if err != nil {
			return nil, err
		}
		rawAccounts = append(rawAccounts, resp.Accounts...)
		if resp.NextToken == nil {
			break
		}
		paginationToken = resp.NextToken
	}
	return rawAccounts, nil
}
