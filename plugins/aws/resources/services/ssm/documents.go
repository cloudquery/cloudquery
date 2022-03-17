package ssm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SsmDocuments() *schema.Table {
	return &schema.Table{
		Name:                 "aws_ssm_documents",
		Description:          "Describes a Amazon Web Services Systems Manager document (SSM document).",
		Resolver:             fetchSsmDocuments,
		Multiplex:            client.ServiceAccountRegionMultiplexer("ssm"),
		IgnoreError:          client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:         client.DeleteAccountRegionFilter,
		PostResourceResolver: ssmDocumentPostResolver,
		Options:              schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the managed instance.",
				Type:        schema.TypeString,
				Resolver:    resolveSSMDocumentARN,
			},
			{
				Name:          "approved_version",
				Description:   "The version of the document currently approved for use in the organization.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "attachments_information",
				Description:   "Details about the document attachments, including names, locations, sizes, and so on.",
				Type:          schema.TypeJSON,
				Resolver:      resolveSSMDocumentJSONField(func(d *types.DocumentDescription) interface{} { return d.AttachmentsInformation }),
				IgnoreInTests: true,
			},
			{
				Name:          "author",
				Description:   "The user in your organization who created the document.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "created_date",
				Description: "The date when the document was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "default_version",
				Description: "The default version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "description",
				Description: "A description of the document.",
				Type:        schema.TypeString,
			},
			{
				Name:          "display_name",
				Description:   "The friendly name of the SSM document",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "document_format",
				Description: "The document format, either JSON or YAML.",
				Type:        schema.TypeString,
			},
			{
				Name:        "document_type",
				Description: "The type of document.",
				Type:        schema.TypeString,
			},
			{
				Name:        "document_version",
				Description: "The document version.",
				Type:        schema.TypeString,
			},
			{
				Name:        "hash",
				Description: "The Sha256 or Sha1 hash created by the system when the document was created. Sha1 hashes have been deprecated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "hash_type",
				Description: "The hash type of the document",
				Type:        schema.TypeString,
			},
			{
				Name:        "latest_version",
				Description: "The latest version of the document.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "The name of the SSM document.",
				Type:        schema.TypeString,
			},
			{
				Name:        "owner",
				Description: "The Amazon Web Services user account that created the document.",
				Type:        schema.TypeString,
			},
			{
				Name:        "parameters",
				Description: "A description of the parameters for a document.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSSMDocumentJSONField(func(d *types.DocumentDescription) interface{} { return d.Parameters }),
			},
			{
				Name:          "pending_review_version",
				Description:   "The version of the document that is currently under review.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "platform_types",
				Description: "The list of OS platforms compatible with this SSM document.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:          "requires",
				Description:   "A list of SSM documents required by a document.",
				Type:          schema.TypeJSON,
				Resolver:      resolveSSMDocumentJSONField(func(d *types.DocumentDescription) interface{} { return d.Requires }),
				IgnoreInTests: true,
			},
			{
				Name:        "review_status",
				Description: "The current status of the review.",
				Type:        schema.TypeString,
			},
			{
				Name:        "schema_version",
				Description: "The schema version.",
				Type:        schema.TypeString,
			},
			{
				Name:          "sha1",
				Description:   "The SHA1 hash of the document, which you can use for verification.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "status",
				Description: "The status of the SSM document.",
				Type:        schema.TypeString,
			},
			{
				Name:          "status_information",
				Description:   "A message returned by Amazon Web Services Systems Manager that explains the Status value",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "target_type",
				Description:   "The target type which defines the kinds of resources the document can run on. For example, /AWS::EC2::Instance",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "version_name",
				Description:   "The version of the artifact associated with the document.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "review_information",
				Description:   "Details about the review of a document.",
				Type:          schema.TypeJSON,
				Resolver:      resolveSSMDocumentJSONField(func(d *types.DocumentDescription) interface{} { return d.ReviewInformation }),
				IgnoreInTests: true,
			},
			{
				Name:        "tags",
				Description: "The tags, or metadata, that have been applied to the document.",
				Type:        schema.TypeJSON,
				Resolver:    resolveSSMDocumentTags,
			},
			{
				Name:          "account_ids",
				Description:   "The account IDs that have permission to use this document",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:          "account_sharing_info_list",
				Description:   "A list of Amazon Web Services accounts where the current document is shared and the version shared with each account.",
				Type:          schema.TypeJSON,
				IgnoreInTests: true,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSsmDocuments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	client := meta.(*client.Client)
	svc := client.Services().SSM
	optsFn := func(o *ssm.Options) {
		o.Region = client.Region
	}
	params := ssm.ListDocumentsInput{
		Filters: []types.DocumentKeyValuesFilter{{Key: aws.String("Owner"), Values: []string{"Self"}}},
	}
	for {
		output, err := svc.ListDocuments(ctx, &params, optsFn)
		if err != nil {
			return diag.WrapError(err)
		}

		for _, d := range output.DocumentIdentifiers {
			dd, err := svc.DescribeDocument(ctx, &ssm.DescribeDocumentInput{Name: d.Name}, optsFn)
			if err != nil {
				return diag.WrapError(err)
			}
			res <- dd.Document
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		params.NextToken = output.NextToken
	}
	return nil
}

func resolveSSMDocumentJSONField(getter func(d *types.DocumentDescription) interface{}) func(context.Context, schema.ClientMeta, *schema.Resource, schema.Column) error {
	return func(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
		d, ok := resource.Item.(*types.DocumentDescription)
		if !ok {
			return fmt.Errorf("not a %T instance: %T", d, resource.Item)
		}
		b, err := json.Marshal(getter(d))
		if err != nil {
			return diag.WrapError(err)
		}
		return resource.Set(c.Name, b)
	}
}

func resolveSSMDocumentTags(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	d, ok := resource.Item.(*types.DocumentDescription)
	if !ok {
		return fmt.Errorf("not a %T instance: %T", d, resource.Item)
	}
	tags := make(map[string]string)
	for _, t := range d.Tags {
		tags[aws.ToString(t.Key)] = aws.ToString(t.Value)
	}
	return resource.Set(c.Name, tags)
}

func ssmDocumentPostResolver(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) (exitErr error) {
	d, ok := resource.Item.(*types.DocumentDescription)
	if !ok {
		return fmt.Errorf("not a %T instance: %T", d, resource.Item)
	}
	client := meta.(*client.Client)
	svc := client.Services().SSM
	optsFn := func(o *ssm.Options) {
		o.Region = client.Region
	}
	input := ssm.DescribeDocumentPermissionInput{
		Name:           d.Name,
		PermissionType: types.DocumentPermissionTypeShare,
	}
	var accountIDs []string
	var infoList []types.AccountSharingInfo
	for {
		output, err := svc.DescribeDocumentPermission(ctx, &input, optsFn)
		if err != nil {
			return diag.WrapError(err)
		}
		accountIDs = append(accountIDs, output.AccountIds...)
		infoList = append(infoList, output.AccountSharingInfoList...)
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	if err := resource.Set("account_ids", accountIDs); err != nil {
		return err
	}
	b, err := json.Marshal(infoList)
	if err != nil {
		return diag.WrapError(err)
	}
	return resource.Set("account_sharing_info_list", b)
}

func resolveSSMDocumentARN(_ context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	d, ok := resource.Item.(*types.DocumentDescription)
	if !ok {
		return fmt.Errorf("not a %T instance: %T", d, resource.Item)
	}
	cl := meta.(*client.Client)
	return resource.Set(c.Name, client.GenerateResourceARN("ssm", "document", *d.Name, cl.Region, cl.AccountID))
}
