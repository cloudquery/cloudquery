package cloudformation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	cqtypes "github.com/cloudquery/plugin-sdk/v4/types"
	"github.com/ghodss/yaml"
)

func stackTemplates() *schema.Table {
	tableName := "aws_cloudformation_stack_templates"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_GetTemplate.html`,
		Resolver:    fetchCloudformationStackTemplates,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "cloudformation"),
		Transform:   transformers.TransformWithStruct(&cloudformation.GetTemplateOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:       "stack_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
			{
				// Might be deprecated in a future release.
				// Contains the template converted to JSON.
				Name:     "template_body",
				Type:     cqtypes.ExtensionTypes.JSON,
				Resolver: resolveTemplateBody,
			},
			{
				// raw template body: could be either YAML or JSON
				Name:     "template_body_text",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("TemplateBody"),
			},
		},
	}
}

func fetchCloudformationStackTemplates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	stack := parent.Item.(types.Stack)
	config := cloudformation.GetTemplateInput{
		StackName: stack.StackName,
	}
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudformation
	resp, err := svc.GetTemplate(ctx, &config, func(options *cloudformation.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- resp
	return nil
}

func resolveTemplateBody(_ context.Context, _ schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	resp := r.Item.(*cloudformation.GetTemplateOutput)
	if resp.TemplateBody == nil {
		return nil
	}
	// this column was originally released as a JSON column, but it turns out that
	// the API can also return YAML. To maintain backwards-compatibility, we attempt
	// to parse the template body as JSON first, and if that fails, we try to parse
	// it as YAML. We return an error if both attempts fail.
	m := map[string]any{}
	err := json.Unmarshal([]byte(*resp.TemplateBody), &m)
	if err != nil {
		// this template might be YAML
		err = yaml.Unmarshal([]byte(*resp.TemplateBody), &m)
		if err != nil {
			return fmt.Errorf("failed to parse Cloudformation template body as either JSON or yaml: %w", err)
		}
	}
	return r.Set(c.Name, m)
}
