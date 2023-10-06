package client

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/spec"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	orgTypes "github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/google/go-cmp/cmp"
)

// sets up test accounts with the following structure:
// root
//
//	ou-parent1
//	  ou-child1
//	    id-child1-account
//	  id-parent1-account
//	ou-parent2
//	  ou-child2
//	    id-child2-account
//	  id-parent2-account
//	id-top-level-account
//	id-top-level-account-inactive
func setupTestAccounts(t *testing.T) *mocks.MockOrganizationsClient {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockOrganizationsClient(ctrl)

	api.EXPECT().ListAccounts(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(
		func(ctx context.Context, params *organizations.ListAccountsInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsOutput, error) {
			output := &organizations.ListAccountsOutput{
				Accounts: []orgTypes.Account{
					{Id: aws.String("id-top-level-account"), Status: orgTypes.AccountStatusActive},
					{Id: aws.String("id-top-level-account-inactive"), Status: orgTypes.AccountStatusSuspended},
					{Id: aws.String("id-parent1-account"), Status: orgTypes.AccountStatusActive},
					{Id: aws.String("id-parent2-account"), Status: orgTypes.AccountStatusActive},
					{Id: aws.String("id-child1-account"), Status: orgTypes.AccountStatusActive},
					{Id: aws.String("id-child2-account"), Status: orgTypes.AccountStatusActive},
				},
			}
			return output, nil
		},
	)

	api.EXPECT().ListChildren(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(
		func(ctx context.Context, params *organizations.ListChildrenInput, optFns ...func(*organizations.Options)) (*organizations.ListChildrenOutput, error) {
			output := &organizations.ListChildrenOutput{}
			switch *params.ParentId {
			case "root":
				output.Children = []orgTypes.Child{
					{Id: aws.String("ou-parent1")},
					{Id: aws.String("ou-parent2")},
				}
			case "ou-parent1":
				output.Children = []orgTypes.Child{
					{Id: aws.String("ou-child1")},
				}
			case "ou-parent2":
				output.Children = []orgTypes.Child{
					{Id: aws.String("ou-child2")},
				}
			case "ou-child1", "ou-child2":
			default:
				return nil, fmt.Errorf("no such OU: %v", *params.ParentId)
			}
			return output, nil
		},
	)

	api.EXPECT().ListAccountsForParent(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes().DoAndReturn(
		func(ctx context.Context, params *organizations.ListAccountsForParentInput, optFns ...func(*organizations.Options)) (*organizations.ListAccountsForParentOutput, error) {
			output := &organizations.ListAccountsForParentOutput{}
			switch *params.ParentId {
			case "root":
				output.Accounts = []orgTypes.Account{
					{Id: aws.String("id-top-level-account"), Status: orgTypes.AccountStatusActive},
					{Id: aws.String("id-top-level-account-inactive"), Status: orgTypes.AccountStatusSuspended},
				}
			case "ou-parent1":
				output.Accounts = []orgTypes.Account{
					{Id: aws.String("id-parent1-account"), Status: orgTypes.AccountStatusActive},
				}
			case "ou-parent2":
				output.Accounts = []orgTypes.Account{
					{Id: aws.String("id-parent2-account"), Status: orgTypes.AccountStatusActive},
				}
			case "ou-child1":
				output.Accounts = []orgTypes.Account{
					{Id: aws.String("id-child1-account"), Status: orgTypes.AccountStatusActive},
				}
			case "ou-child2":
				output.Accounts = []orgTypes.Account{
					{Id: aws.String("id-child2-account"), Status: orgTypes.AccountStatusActive},
				}
			default:
				return nil, fmt.Errorf("no such OU: %v", *params.ParentId)
			}
			return output, nil
		},
	)
	return api
}

func Test_loadAccounts(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name    string
		spec    *spec.Spec
		want    []string
		wantErr error
	}{
		{
			name: "all_accounts",
			spec: &spec.Spec{
				Organization: &spec.Org{
					AdminAccount: &spec.Account{},
				},
			},
			want: []string{"id-child1-account", "id-child2-account", "id-parent1-account", "id-parent2-account", "id-top-level-account"},
		},
		{
			name: "all_accounts_with_skip_member_accounts",
			spec: &spec.Spec{
				Organization: &spec.Org{
					AdminAccount:       &spec.Account{},
					SkipMemberAccounts: []string{"id-child2-account", "id-parent1-account", "id-parent2-account", "id-top-level-account"},
				},
			},
			want: []string{"id-child1-account"},
		},
		{
			name: "org_root",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits: []string{"root"},
					AdminAccount:      &spec.Account{},
				},
			},
			want: []string{"id-top-level-account", "id-child1-account", "id-parent1-account", "id-child2-account", "id-parent2-account"},
		},
		{
			name: "ou_parent1",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits: []string{"ou-parent1"},
					AdminAccount:      &spec.Account{},
				},
			},
			want: []string{"id-parent1-account", "id-child1-account"},
		},
		{
			name: "ou_parent1_and_parent2",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits: []string{"ou-parent1", "ou-parent2"},
					AdminAccount:      &spec.Account{},
				},
			},
			want: []string{"id-parent1-account", "id-child1-account", "id-parent2-account", "id-child2-account"},
		},
		{
			name: "ou_parent1_skip_child1",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits:  []string{"ou-parent1"},
					SkipMemberAccounts: []string{"id-child1-account"},
					AdminAccount:       &spec.Account{},
				},
			},
			want: []string{"id-parent1-account"},
		},
		{
			name: "ou_root_skip_parent1",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits:       []string{"root"},
					SkipOrganizationalUnits: []string{"ou-parent1"},
					AdminAccount:            &spec.Account{},
				},
			},
			want: []string{"id-top-level-account", "id-parent2-account", "id-child2-account"},
		},
		{
			name: "ou_root_skip_parent1",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits:       []string{"root"},
					SkipOrganizationalUnits: []string{"ou-parent1"},
					AdminAccount:            &spec.Account{},
				},
			},
			want: []string{"id-top-level-account", "id-parent2-account", "id-child2-account"},
		},
		{
			name: "ou_root_and_parent1",
			spec: &spec.Spec{
				Organization: &spec.Org{
					OrganizationUnits:       []string{"root", "ou-parent1"},
					SkipOrganizationalUnits: []string{},
					AdminAccount:            &spec.Account{},
				},
			},
			want: []string{"id-top-level-account", "id-parent1-account", "id-child1-account", "id-parent2-account", "id-child2-account"},
		},
	}
	api := setupTestAccounts(t)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := loadAccounts(ctx, tc.spec, api, "us-east-1")
			sort.Slice(tc.want, func(i, j int) bool {
				return tc.want[i] < tc.want[j]
			})
			respIDs := make([]string, len(resp))
			for i, a := range resp {
				respIDs[i] = a.ID
			}
			sort.Strings(respIDs)
			respDiff := cmp.Diff(respIDs, tc.want)
			errDiff := cmp.Diff(err, tc.wantErr)
			if errDiff != "" {
				t.Fatal(errDiff)
			}
			if respDiff != "" {
				t.Fatal(respDiff)
			}
		})
	}
}
