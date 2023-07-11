package ssoadmin

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
)

func inlinePolicies() *schema.Table {
	tableName := "aws_ssoadmin_permission_set_inline_policies"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_GetInlinePolicyForPermissionSet.html`,
		Resolver:    fetchInlinePolicies,
		Transform:   transformers.TransformWithStruct(&ssoadmin.GetInlinePolicyForPermissionSetOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			{
				Name:       "permission_set_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("permission_set_arn"),
				PrimaryKey: true,
			},
			{
				Name:       "instance_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("instance_arn"),
				PrimaryKey: true,
			},
			{
				Name:     "inline_policy_json",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveInlinePolicyJSON,
			},
		},
	}
}

func fetchInlinePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Ssoadmin

	permissionSetARN := parent.Item.(*types.PermissionSet).PermissionSetArn
	instanceARN := parent.Parent.Item.(types.InstanceMetadata).InstanceArn
	config := ssoadmin.GetInlinePolicyForPermissionSetInput{
		InstanceArn:      instanceARN,
		PermissionSetArn: permissionSetARN,
	}

	response, err := svc.GetInlinePolicyForPermissionSet(ctx, &config, func(o *ssoadmin.Options) {
		o.Region = cl.Region
	})
	if err != nil {
		return err
	}

	res <- response
	return nil
}

func resolveInlinePolicyJSON(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(*ssoadmin.GetInlinePolicyForPermissionSetOutput)

	decodedDocument, err := url.QueryUnescape(*r.InlinePolicy)
	if err != nil {
		return err
	}

	var document map[string]any
	err = json.Unmarshal([]byte(decodedDocument), &document)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, document)
}
