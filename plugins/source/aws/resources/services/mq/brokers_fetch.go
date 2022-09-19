package mq

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mq"
	xj "github.com/basgys/goxml2json"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchMqBrokers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config mq.ListBrokersInput
	c := meta.(*client.Client)
	svc := c.Services().MQ
	for {
		response, err := svc.ListBrokers(ctx, &config)
		if err != nil {
			return err
		}
		for _, bs := range response.BrokerSummaries {
			output, err := svc.DescribeBroker(ctx, &mq.DescribeBrokerInput{BrokerId: bs.BrokerId}, func(options *mq.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- output
		}
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

func fetchMqBrokerConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ
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
func fetchMqBrokerConfigurationRevisions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cfg := parent.Item.(mq.DescribeConfigurationOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ

	input := mq.ListConfigurationRevisionsInput{ConfigurationId: cfg.Id}
	for {
		output, err := svc.ListConfigurationRevisions(ctx, &input)
		if err != nil {
			return err
		}
		for _, rev := range output.Revisions {
			revId := strconv.Itoa(int(rev.Revision))
			output, err := svc.DescribeConfigurationRevision(ctx, &mq.DescribeConfigurationRevisionInput{ConfigurationId: cfg.Id, ConfigurationRevision: &revId}, func(options *mq.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- output
		}
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
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
	unmarshalledJson := map[string]interface{}{}
	err = json.Unmarshal(marshalledJson.Bytes(), &unmarshalledJson)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, unmarshalledJson)
}

func fetchMqBrokerUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	broker := parent.Item.(*mq.DescribeBrokerOutput)
	c := meta.(*client.Client)
	svc := c.Services().MQ
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
