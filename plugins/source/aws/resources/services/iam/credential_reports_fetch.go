package iam

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gocarina/gocsv"
	"github.com/thoas/go-funk"
)

type CredentialReportEntry struct {
	User                      string   `csv:"user"`
	Arn                       string   `csv:"arn"`
	UserCreationTime          DateTime `csv:"user_creation_time"`
	PasswordStatus            string   `csv:"password_enabled"`
	PasswordLastChanged       DateTime `csv:"password_last_changed"`
	PasswordNextRotation      DateTime `csv:"password_next_rotation"`
	MfaActive                 bool     `csv:"mfa_active"`
	AccessKey1Active          bool     `csv:"access_key_1_active"`
	AccessKey2Active          bool     `csv:"access_key_2_active"`
	AccessKey1LastRotated     DateTime `csv:"access_key_1_last_rotated"`
	AccessKey2LastRotated     DateTime `csv:"access_key_2_last_rotated"`
	Cert1Active               bool     `csv:"cert_1_active"`
	Cert2Active               bool     `csv:"cert_2_active"`
	Cert1LastRotated          DateTime `csv:"cert_1_last_rotated"`
	Cert2LastRotated          DateTime `csv:"cert_2_last_rotated"`
	AccessKey1LastUsedDate    DateTime `csv:"access_key_1_last_used_date"`
	AccessKey1LastUsedRegion  string   `csv:"access_key_1_last_used_region"`
	AccessKey1LastUsedService string   `csv:"access_key_1_last_used_service"`
	AccessKey2LastUsedDate    DateTime `csv:"access_key_2_last_used_date"`
	AccessKey2LastUsedRegion  string   `csv:"access_key_2_last_used_region"`
	AccessKey2LastUsedService string   `csv:"access_key_2_last_used_service"`
	PasswordLastUsed          DateTime `csv:"password_last_used"`
}

type DateTime struct {
	*time.Time
}

func (d *DateTime) UnmarshalCSV(val string) (err error) {
	switch val {
	case "N/A", "not_supported":
		d.Time = nil
		return nil
	}
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return err
	}
	d.Time = &t
	return nil
}

func fetchIamCredentialReports(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	svc := meta.(*client.Client).Services().IAM
	for {
		reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err == nil && reportOutput != nil {
			var users []*CredentialReportEntry
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
func timestampPathResolver(path string) schema.ColumnResolver {
	return func(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
		t := funk.Get(r.Item, path, funk.WithAllowZero())
		dt := t.(DateTime)
		return r.Set(c.Name, dt.Time)
	}
}
