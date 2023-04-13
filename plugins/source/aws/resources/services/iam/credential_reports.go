package iam

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/gocarina/gocsv"
	"github.com/thoas/go-funk"
)

func CredentialReports() *schema.Table {
	tableName := "aws_iam_credential_reports"
	return &schema.Table{
		Name:     tableName,
		Resolver: fetchIamCredentialReports,
		Transform: transformers.TransformWithStruct(
			&models.CredentialReportEntry{},
			transformers.WithSkipFields(
				"AccessKey1LastRotated",
				"AccessKey1LastUsedDate",
				"Cert1LastRotated",
				"AccessKey2LastRotated",
				"AccessKey2LastUsedDate",
				"Cert2LastRotated",
			),
		),
		Multiplex: client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "user_creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("UserCreationTime"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "password_last_changed",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("PasswordLastChanged"),
			},
			{
				Name:     "password_next_rotation",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("PasswordNextRotation"),
			},
			{
				Name:     "access_key_1_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("AccessKey1LastRotated"),
			},
			{
				Name:     "access_key_2_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("AccessKey2LastRotated"),
			},
			{
				Name:     "cert_1_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("Cert1LastRotated"),
			},
			{
				Name:     "cert_2_last_rotated",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("Cert2LastRotated"),
			},
			{
				Name:     "access_key_1_last_used_date",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("AccessKey1LastUsedDate"),
			},
			{
				Name:     "access_key_2_last_used_date",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("AccessKey2LastUsedDate"),
			},
			{
				Name:     "password_last_used",
				Type:     schema.TypeTimestamp,
				Resolver: timestampPathResolver("PasswordLastUsed"),
			},
			{
				Name:     "password_enabled",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PasswordStatus"),
			},
		},
	}
}

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
