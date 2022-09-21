package client

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/rs/zerolog"
)

// Parses org configuration and grabs the appropriate accounts
func loadOrgAccounts(ctx context.Context, logger zerolog.Logger, spec *Spec) ([]Account, AssumeRoleAPIClient, error) {
	// If user doesn't specify any configs for admin account instantiate default values
	if spec.Organization.AdminAccount == nil {
		spec.Organization.AdminAccount = &Account{
			Name:  "Default-Admin-Account",
			LocalProfile: "",
		}
	}
	awsCfg, err := configureAwsClient(ctx, logger, spec, *spec.Organization.AdminAccount, nil)
	if err != nil {
		return nil, nil, err
	}
	svc := organizations.NewFromConfig(awsCfg)
	accounts, err := loadAccounts(ctx, spec, svc)
	if err != nil {
		return nil, nil, err
	}
	if spec.Organization.MemberCredentials != nil {
		awsCfg, err = configureAwsClient(ctx, logger, spec, *spec.Organization.MemberCredentials, nil)
		if err != nil {
			return nil, nil, err
		}
	}
	return accounts, sts.NewFromConfig(awsCfg), err
}

// Load accounts from the appropriate endpoint as well as normalizing response
func loadAccounts(ctx context.Context, spec *Spec, accountsApi OrganizationsClient) ([]Account, error) {
	var rawAccounts []orgTypes.Account
	var err error
	if len(spec.Organization.OrganizationUnits) > 0 {
		rawAccounts, err = getOUAccounts(ctx, accountsApi, spec.Organization.OrganizationUnits)
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
			Resource:  "role/" + spec.Organization.ChildAccountRoleName,
		}
		if parsed, err := arn.Parse(aws.ToString(account.Arn)); err == nil {
			roleArn.Partition = parsed.Partition
		}

		accounts = append(accounts, Account{
			Name:                    *account.Id,
			AssumeRoleARN:         roleArn.String(),
			AssumeRoleSessionName: spec.Organization.ChildAccountRoleSessionName,
			AssumeRoleExternalID:  spec.Organization.ChildAccountExternalID,
			LocalProfile:          spec.Organization.AdminAccount.LocalProfile,
			Regions:               spec.Organization.ChildAccountRegions,
			source:                "org",
		})
	}
	return accounts, err
}

// Get Accounts for specific Organizational Units
func getOUAccounts(ctx context.Context, accountsApi OrganizationsClient, ous []string) ([]orgTypes.Account, error) {
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
func getAllAccounts(ctx context.Context, accountsApi OrganizationsClient) ([]orgTypes.Account, error) {
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
