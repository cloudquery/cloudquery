package iam

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/gocarina/gocsv"
	"github.com/thoas/go-funk"
)

func fetchIamCredentialReports(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	var err error
	var apiErr smithy.APIError
	var reportOutput *iam.GetCredentialReportOutput
	svc := meta.(*client.Client).Services().Iam
	for {
		reportOutput, err = svc.GetCredentialReport(ctx, &iam.GetCredentialReportInput{})
		if err == nil && reportOutput != nil {
			var users []*models.CredentialReportEntry
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
			meta.(*client.Client).Logger().Debug().Msg("Waiting for credential report to be generated")
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
		dt := t.(models.DateTime)
		return r.Set(c.Name, dt.Time)
	}
}
