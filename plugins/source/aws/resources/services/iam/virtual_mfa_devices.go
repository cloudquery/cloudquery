package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualMfaDevices() *schema.Table {
	tableName := "aws_iam_virtual_mfa_devices"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html`,
		Resolver:    fetchIamVirtualMfaDevices,
		Transform:   transformers.TransformWithStruct(&types.VirtualMFADevice{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "iam"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
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

func fetchIamVirtualMfaDevices(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config iam.ListVirtualMFADevicesInput
	svc := meta.(*client.Client).Services().Iam
	for {
		response, err := svc.ListVirtualMFADevices(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.VirtualMFADevices
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}

	return nil
}
