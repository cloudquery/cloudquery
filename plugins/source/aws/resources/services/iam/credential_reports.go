// Code generated by codegen; DO NOT EDIT.

package iam

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CredentialReports() *schema.Table {
	return &schema.Table{
		Name:      "aws_iam_credential_reports",
		Resolver:  fetchIamCredentialReports,
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
			{
				Name:     "user",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("User"),
			},
			{
				Name:     "password_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PasswordStatus"),
			},
			{
				Name:     "mfa_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("MfaActive"),
			},
			{
				Name:     "access_key1_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccessKey1Active"),
			},
			{
				Name:     "access_key2_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AccessKey2Active"),
			},
			{
				Name:     "cert1_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Cert1Active"),
			},
			{
				Name:     "cert2_active",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Cert2Active"),
			},
			{
				Name:     "access_key1_last_used_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKey1LastUsedRegion"),
			},
			{
				Name:     "access_key1_last_used_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKey1LastUsedService"),
			},
			{
				Name:     "access_key2_last_used_region",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKey2LastUsedRegion"),
			},
			{
				Name:     "access_key2_last_used_service",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccessKey2LastUsedService"),
			},
		},
	}
}
