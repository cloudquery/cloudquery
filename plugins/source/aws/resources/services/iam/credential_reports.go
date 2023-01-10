package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/iam/models"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func CredentialReports() *schema.Table {
	return &schema.Table{
		Name:     "aws_iam_credential_reports",
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
		Multiplex: client.AccountMultiplex,
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
