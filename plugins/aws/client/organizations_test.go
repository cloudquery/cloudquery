package client

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type mockOrgClient struct {
	listAccountsForParent func(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error)
	listAccounts          func(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
}

func (m mockOrgClient) ListAccountsForParent(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error) {
	return m.listAccountsForParent(ctx, params, optFns...)
}

func (m mockOrgClient) ListAccounts(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error) {
	return m.listAccounts(ctx, params, optFns...)
}

func Test_Org_Configure(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		listAccountsForParent func(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error)
		listAccounts          func(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error)
		ous                   []string
		accounts              []Account
		err                   error
		config                *Config
	}{
		{
			listAccounts: func(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error) {
				return &organizations.ListAccountsOutput{
					Accounts: []orgTypes.Account{},
				}, nil
			},
			accounts: []Account{},
			err:      nil,
			config: &Config{
				Organization: &AwsOrg{},
			},
		},
		{
			listAccountsForParent: func(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error) {
				return &organizations.ListAccountsForParentOutput{
					Accounts: []orgTypes.Account{
						{
							Id:     aws.String("012345678910"),
							Status: orgTypes.AccountStatusActive,
						},
					},
				}, nil
			},
			config: &Config{
				Organization: &AwsOrg{
					OrganizationUnits:    []string{"test-ou"},
					ChildAccountRoleName: "test",
					AdminAccount:         &Account{},
				},
			},
			accounts: []Account{
				{
					ID:              "012345678910",
					RoleARN:         "arn:aws:iam::012345678910:role/test",
					RoleSessionName: "",
					ExternalID:      "",
					LocalProfile:    "",
					Regions:         []string{},
					source:          "org",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		api := mockOrgClient{
			listAccountsForParent: test.listAccountsForParent,
			listAccounts:          test.listAccounts,
		}
		resp, err := loadAccounts(ctx, test.config, api)
		respDiff := cmp.Diff(resp, test.accounts, cmpopts.IgnoreUnexported(Account{}), cmpopts.EquateEmpty())

		if respDiff != "" {
			t.Fatal(respDiff)
		}
		errDiff := cmp.Diff(err, test.err)

		if errDiff != "" {
			t.Fatal(errDiff)
		}
	}
}
