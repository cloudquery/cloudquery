package iam

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v4/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAccount(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)

	acc := struct {
		Users                             int32
		UsersQuota                        int32
		Groups                            int32
		GroupsQuota                       int32
		ServerCertificates                int32
		ServerCertificatesQuota           int32
		UserPolicySizeQuota               int32
		GroupPolicySizeQuota              int32
		GroupsPerUserQuota                int32
		SigningCertificatesPerUserQuota   int32
		AccessKeysPerUserQuota            int32
		MFADevices                        int32
		MFADevicesInUse                   int32
		AccountMFAEnabled                 int32
		AccountAccessKeysPresent          int32
		AccountSigningCertificatesPresent int32
		AttachedPoliciesPerGroupQuota     int32
		AttachedPoliciesPerRoleQuota      int32
		AttachedPoliciesPerUserQuota      int32
		Policies                          int32
		PoliciesQuota                     int32
		PolicySizeQuota                   int32
		PolicyVersionsInUse               int32
		PolicyVersionsInUseQuota          int32
		VersionsPerPolicyQuota            int32
		GlobalEndpointTokenVersion        int32
	}{}

	require.NoError(t, faker.FakeObject(&acc))

	data, err := json.Marshal(acc)
	if err != nil {
		t.Fatal(err)
	}
	summaryData := make(map[string]int32)
	if err := json.Unmarshal(data, &summaryData); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetAccountSummary(gomock.Any(), gomock.Any(), gomock.Any()).Return(&iam.GetAccountSummaryOutput{SummaryMap: summaryData}, nil)
	m.EXPECT().ListAccountAliases(gomock.Any(), gomock.Any(), gomock.Any()).Return(&iam.ListAccountAliasesOutput{AccountAliases: []string{"testAccount"}}, nil)

	return client.Services{
		Iam: m,
	}
}

func TestAccounts(t *testing.T) {
	client.AwsMockTestHelper(t, Accounts(), buildAccount, client.TestOptions{})
}
