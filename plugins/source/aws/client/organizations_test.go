package client

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/golang/mock/gomock"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_Org_Configure(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		listAccountsForParent *organizations.ListAccountsForParentOutput
		listAccounts          *organizations.ListAccountsOutput
		ous                   []string
		accounts              []Account
		err                   error
		config                *Spec
	}{
		{
			listAccounts: &organizations.ListAccountsOutput{
				Accounts: []orgTypes.Account{},
			},
			accounts: []Account{},
			err:      nil,
			config: &Spec{
				Organization: &AwsOrg{},
			},
		},
		{
			listAccountsForParent: &organizations.ListAccountsForParentOutput{
				Accounts: []orgTypes.Account{
					{
						Id:     aws.String("012345678910"),
						Status: orgTypes.AccountStatusActive,
					},
				},
			},
			config: &Spec{
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
		ctrl := gomock.NewController(t)
		api := mocks.NewMockOrganizationsClient(ctrl)
		if test.listAccounts != nil {
			api.EXPECT().ListAccounts(gomock.Any(), gomock.Any()).Return(test.listAccounts, nil)
		}
		if test.listAccountsForParent != nil {
			api.EXPECT().ListAccountsForParent(gomock.Any(), gomock.Any(), gomock.Any()).Return(test.listAccountsForParent, nil)
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
