package mq

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"strconv"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	xj "github.com/basgys/goxml2json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func brokerConfigurationRevisions() *schema.Table {
	tableName := "aws_mq_broker_configuration_revisions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/amazon-mq/latest/api-reference/configurations-configuration-id-revisions.html`,
		Resolver:            fetchMqBrokerConfigurationRevisions,
		PreResourceResolver: getMqBrokerConfigurationRevision,
		Transform:           transformers.TransformWithStruct(&mq.DescribeConfigurationRevisionOutput{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "mq"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "broker_configuration_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "data",
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveBrokerConfigurationRevisionsData,
			},
		},
	}
}

func fetchMqBrokerConfigurationRevisions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cfg := parent.Item.(mq.DescribeConfigurationOutput)
	cl := meta.(*client.Client)
	svc := cl.Services().Mq

	input := mq.ListConfigurationRevisionsInput{ConfigurationId: cfg.Id}
	// No paginator available
	for {
		output, err := svc.ListConfigurationRevisions(ctx, &input, func(options *mq.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- output.Revisions

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}

func getMqBrokerConfigurationRevision(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Mq
	rev := resource.Item.(types.ConfigurationRevision)
	cfg := resource.Parent.Item.(mq.DescribeConfigurationOutput)

	revId := strconv.Itoa(int(rev.Revision))
	output, err := svc.DescribeConfigurationRevision(ctx, &mq.DescribeConfigurationRevisionInput{ConfigurationId: cfg.Id, ConfigurationRevision: &revId}, func(options *mq.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

func resolveBrokerConfigurationRevisionsData(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	revision := resource.Item.(*mq.DescribeConfigurationRevisionOutput)
	rawDecodedText, err := base64.StdEncoding.DecodeString(*revision.Data)
	if err != nil {
		return err
	}
	xml := bytes.NewReader(rawDecodedText)
	marshalledJson, err := xj.Convert(xml)
	if err != nil {
		return err
	}
	unmarshalledJson := map[string]any{}
	err = json.Unmarshal(marshalledJson.Bytes(), &unmarshalledJson)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, unmarshalledJson)
}
