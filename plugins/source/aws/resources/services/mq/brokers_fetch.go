package mq

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	xj "github.com/basgys/goxml2json"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchMqBrokers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config mq.ListBrokersInput
	c := meta.(*client.Client)
	svc := c.Services().Mq
	for {
		response, err := svc.ListBrokers(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.BrokerSummaries

		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func getMqBroker(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Mq
	bs := resource.Item.(types.BrokerSummary)

	output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId})
	if err != nil {
		return err
	}
	resource.Item = output
	return nil
}

func fetchMqBrokerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().Mq
	// Ensure Configurations is not nil
	// This *might* occur during initial creation of broker
	if broker.Configurations == nil {
		return nil
	}

	list := broker.Configurations.History
	if broker.Configurations.Current != nil {
		list = append(list, *broker.Configurations.Current)
	}

	// History might contain same Id multiple times (maybe with different revisions) but we're only interested in the latest revision of each
	dupes := make(map[string]struct{}, len(list))
	configurations := make([]mq.DescribeConfigurationOutput, 0, len(list))
	for _, cfg := range list {
		if cfg.Id == nil {
			continue
		}

		if _, ok := dupes[*cfg.Id]; ok {
			continue
		}
		dupes[*cfg.Id] = struct{}{}

		input := mq.DescribeConfigurationInput{ConfigurationId: cfg.Id}
		output, err := svc.DescribeConfiguration(ctx, &input)
		if err != nil {
			return err
		}
		configurations = append(configurations, *output)
	}
	res <- configurations
	return nil
}

func fetchMqBrokerConfigurationRevisions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cfg := parent.Item.(mq.DescribeConfigurationOutput)
	c := meta.(*client.Client)
	svc := c.Services().Mq

	input := mq.ListConfigurationRevisionsInput{ConfigurationId: cfg.Id}
	for {
		output, err := svc.ListConfigurationRevisions(ctx, &input)
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
	c := meta.(*client.Client)
	svc := c.Services().Mq
	rev := resource.Item.(types.ConfigurationRevision)
	cfg := resource.Parent.Item.(mq.DescribeConfigurationOutput)

	revId := strconv.Itoa(int(rev.Revision))
	output, err := svc.DescribeConfigurationRevision(ctx, &mq.DescribeConfigurationRevisionInput{ConfigurationId: cfg.Id, ConfigurationRevision: &revId})
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

func fetchMqBrokerUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().Mq
	for _, us := range broker.Users {
		input := mq.DescribeUserInput{
			BrokerId: broker.BrokerId,
			Username: us.Username,
		}
		output, err := svc.DescribeUser(ctx, &input)
		if err != nil {
			return err
		}
		res <- output
	}
	return nil
}
