package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualMfaDevices() *schema.Table {
	return &schema.Table{
		Name:        "aws_iam_virtual_mfa_devices",
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html`,
		Resolver:    fetchIamVirtualMfaDevices,
		Transform:   transformers.TransformWithStruct(&types.VirtualMFADevice{}),
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name: "serial_number",
				Type: schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
