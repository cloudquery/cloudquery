package ssoadmin

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func permissionsBoundaries() *schema.Table {
	tableName := "aws_ssoadmin_permission_set_permissions_boundaries"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_GetPermissionsBoundaryForPermissionSet.html`,
		Resolver:    fetchPermissionBoundaries,
		Transform:   transformers.TransformWithStruct(&types.PermissionsBoundary{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			{
				Name:                "permission_set_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("permission_set_arn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:                "instance_arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.ParentColumnResolver("instance_arn"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchPermissionBoundaries(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceSsoadmin).Ssoadmin

	permissionSetARN := parent.Item.(*types.PermissionSet).PermissionSetArn
	instanceARN := parent.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.GetPermissionsBoundaryForPermissionSetInput{
		InstanceArn:      instanceARN,
		PermissionSetArn: permissionSetARN,
	}

	response, err := svc.GetPermissionsBoundaryForPermissionSet(ctx, &config, func(o *ssoadmin.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response.PermissionsBoundary
	return nil
}
