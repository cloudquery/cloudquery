package iam

import (
	"context"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

var exampleReport = `user,arn,user_creation_time,password_enabled,password_last_used,password_last_changed,password_next_rotation,mfa_active,access_key_1_active,access_key_1_last_rotated,access_key_1_last_used_date,access_key_1_last_used_region,access_key_1_last_used_service,access_key_2_active,access_key_2_last_rotated,access_key_2_last_used_date,access_key_2_last_used_region,access_key_2_last_used_service,cert_1_active,cert_1_last_rotated,cert_2_active,cert_2_last_rotated
user-readonly,arn:aws:iam::123456789012:user/user-readonly,2022-08-31T11:10:33+00:00,false,2022-08-30T11:10:33+00:00,2022-08-31T11:10:33+00:00,2023-08-31T11:10:33+00:00,false,true,2022-08-31T11:10:34+00:00,2022-08-31T11:23:00+00:00,us-east-1,iam,true,2022-08-31T11:10:33+00:00,2022-08-31T11:10:33+00:00,N/A,N/A,false,2022-08-31T11:10:33+00:00,false,2022-08-31T11:10:33+00:00`

var exampleReportWithNilValues = `user,arn,user_creation_time,password_enabled,password_last_used,password_last_changed,password_next_rotation,mfa_active,access_key_1_active,access_key_1_last_rotated,access_key_1_last_used_date,access_key_1_last_used_region,access_key_1_last_used_service,access_key_2_active,access_key_2_last_rotated,access_key_2_last_used_date,access_key_2_last_used_region,access_key_2_last_used_service,cert_1_active,cert_1_last_rotated,cert_2_active,cert_2_last_rotated
<root_account>,arn:aws:iam::123456789012:root,2022-03-23T10:21:07+00:00,not_supported,2022-08-26T15:26:38+00:00,not_supported,not_supported,false,false,N/A,N/A,N/A,N/A,false,N/A,N/A,N/A,N/A,false,N/A,false,N/A
user-cli,arn:aws:iam::123456789012:user/user-cli,2022-07-18T09:03:38+00:00,false,N/A,N/A,N/A,false,true,2022-08-01T13:51:50+00:00,2022-08-05T08:49:00+00:00,ap-northeast-3,glue,true,2022-08-29T08:39:55+00:00,2022-09-01T15:41:00+00:00,us-east-1,logs,false,N/A,false,N/A
user-readonly,arn:aws:iam::123456789012:user/user-readonly,2022-08-31T11:10:33+00:00,false,N/A,N/A,N/A,false,true,2022-08-31T11:10:34+00:00,2022-08-31T11:23:00+00:00,us-east-1,iam,false,N/A,N/A,N/A,N/A,false,N/A,false,N/A`

func buildCredentialReports(_ *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	m.EXPECT().GetCredentialReport(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetCredentialReportOutput{
			Content: []byte(exampleReport),
		}, nil)

	return client.Services{
		Iam: m,
	}
}

func buildCredentialReportsWithNilValues(ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockIamClient(ctrl)
	m.EXPECT().GetCredentialReport(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		&iam.GetCredentialReportOutput{
			Content: []byte(exampleReportWithNilValues),
		}, nil)

	return client.Services{
		Iam: m,
	}
}

// this is an additional custom test to check that we can also handle N/A and not_supported values in CSV
// (these columns should be set to nil)
func testCredentialReportsWithNilValues(t *testing.T) {
	t.Helper()

	t.Run("test with nil values", func(t *testing.T) {
		ctx := context.Background()
		ctrl := gomock.NewController(t)
		services := buildCredentialReportsWithNilValues(ctrl)
		services.Regions = []string{"us-east-1"}
		cl := client.NewAwsClient(zerolog.Logger{}, nil)
		cl.ServicesManager.InitServicesForPartitionAccount("aws", "testAccount", services)
		cl.Partition = "aws"
		cl.Region = "us-east-1"
		cl.AccountID = "testAccount"
		res := make(chan any, 1)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			got := make([]*models.CredentialReportEntry, 0, 3)
			for v := range res {
				vals := v.([]*models.CredentialReportEntry)
				got = append(got, vals...)
			}
			if len(got) != 3 {
				t.Errorf("got %d credential report entries, want %d", len(got), 3)
			}
			wg.Done()
		}()
		err := fetchIamCredentialReports(ctx, &cl, nil, res)
		if err != nil {
			t.Fatalf("unexpected error calling fetchIamCredentialReports: %v", err)
		}
		close(res)
		wg.Wait()
	})
}

func TestCredentialReports(t *testing.T) {
	client.AwsMockTestHelper(t, CredentialReports(), buildCredentialReports, client.TestOptions{})
	testCredentialReportsWithNilValues(t)
}
