package iam

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gocarina/gocsv"
	"time"
)

type CredentialReportUser struct {
	User                      string    `csv:"user"`
	Arn                       string    `csv:"arn"`
	UserCreationTime          time.Time `csv:"user_creation_time"`
	PasswordStatus            string    `csv:"password_enabled"`
	PasswordLastChanged       string    `csv:"password_last_changed"`
	PasswordNextRotation      string    `csv:"password_next_rotation"`
	MfaActive                 bool      `csv:"mfa_active"`
	AccessKey1Active          bool      `csv:"access_key_1_active"`
	AccessKey2Active          bool      `csv:"access_key_2_active"`
	AccessKey1LastRotated     string    `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated     string    `csv:"access_key_2_last_rotated"`
	Cert1Active               bool      `csv:"cert_1_active"`
	Cert2Active               bool      `csv:"cert_2_active"`
	Cert1LastRotated          string    `csv:"cert_1_last_rotated"`
	Cert2LastRotated          string    `csv:"cert_2_last_rotated"`
	AccessKey1LastUsedDate    time.Time `csv:"access_key_1_last_used_date"`
	AccessKey1LastUsedRegion  string    `csv:"access_key_1_last_used_region"`
	AccessKey1LastUsedService string    `csv:"access_key_1_last_used_service"`
	AccessKey2LastUsedDate    time.Time `csv:"access_key_2_last_used_date"`
	AccessKey2LastUsedRegion  string    `csv:"access_key_2_last_used_region"`
	AccessKey2LastUsedService string    `csv:"access_key_2_last_used_service"`
	PasswordLastUsed          string    `csv:"password_last_used"`
}

func fetchIamCredentialReportUsers(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	svc := meta.(*client.Client).Services().IAM
	for {
		reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err == nil && reportOutput != nil {
			var users []*CredentialReportUser
			err = gocsv.UnmarshalBytes(reportOutput.Content, &users)
			if err != nil {
				return err
			}
			res <- users
		}
		if !errors.As(err, &apiErr) {
			return err
		}
		switch apiErr.ErrorCode() {
		case "ReportNotPresent", "ReportExpired":
			_, err := svc.GenerateCredentialReport(ctx, &iam.GenerateCredentialReportInput{})
			if err != nil {
				var serviceError smithy.APIError
				if !errors.As(err, &serviceError) {
					return err
				}
				// LimitExceeded is the only specific error that should not stop processing
				// If Limit Exceeded is returned we should try and see if there is a credential report
				// already generated so we want to sleep for 5 seconds then continue
				if serviceError.ErrorCode() != "LimitExceeded" {
					return err
				}
				if err := client.Sleep(ctx, 5*time.Second); err != nil {
					return err
				}
			}
		case "ReportInProgress":
			meta.Logger().Debug().Msg("Waiting for credential report to be generated")
			if err := client.Sleep(ctx, 5*time.Second); err != nil {
				return err
			}
		default:
			return err
		}
	}
}
