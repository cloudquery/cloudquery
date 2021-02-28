package sns

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/mitchellh/mapstructure"
)

type Topic struct {
	ID                        uint `gorm:"primarykey"`
	AccountID                 string
	Region                    string
	Owner                     string
	TopicArn                  *string
	Policy                    string
	DeliveryPolicy            string
	DisplayName               string
	SubscriptionsConfirmed    *int
	SubscriptionsDeleted      *int
	SubscriptionsPending      *int
	EffectiveDeliveryPolicy   string
	FifoTopic                 *bool
	ContentBasedDeduplication *bool
}

func (Topic) TableName() string {
	return "aws_sns_topics"
}

func getOrZero(attrs map[string]string, keyName string) string {
	if v, found := attrs[keyName]; found {
		return v
	}
	return "0"
}

func (c *Client) transformTopics(values *[]types.Topic) ([]*Topic, error) {
	var tValues []*Topic
	ctx := context.Background()
	for _, value := range *values {

		// All topic attributes are returned as a string; we have to handle type conversion
		params := sns.GetTopicAttributesInput{
			TopicArn: value.TopicArn,
		}
		output, err := c.svc.GetTopicAttributes(ctx, &params)
		if err != nil {
			return nil, err
		}

		subsConfirmed, err := strconv.Atoi(getOrZero(output.Attributes, "SubscriptionsConfirmed"))
		if err != nil {
			return nil, err
		}

		subsDeleted, err := strconv.Atoi(getOrZero(output.Attributes, "SubscriptionsDeleted"))
		if err != nil {
			return nil, err
		}

		subsPending, err := strconv.Atoi(getOrZero(output.Attributes, "SubscriptionsPending"))
		if err != nil {
			return nil, err
		}

		isFifo, err := strconv.ParseBool(getOrZero(output.Attributes, "FifoTopic"))
		if err != nil {
			return nil, err
		}

		isContentBasedDeduped, err := strconv.ParseBool(getOrZero(output.Attributes, "ContentBasedDeduplication"))
		if err != nil {
			return nil, err
		}

		tValue := Topic{
			AccountID:                 c.accountID,
			Region:                    c.region,
			TopicArn:                  value.TopicArn,
			Policy:                    output.Attributes["Policy"],
			DeliveryPolicy:            output.Attributes["DeliveryPolicy"],
			DisplayName:               output.Attributes["DisplayName"],
			Owner:                     output.Attributes["Owner"],
			EffectiveDeliveryPolicy:   output.Attributes["EffectiveDeliveryPolicy"],
			SubscriptionsConfirmed:    &subsConfirmed,
			SubscriptionsDeleted:      &subsDeleted,
			SubscriptionsPending:      &subsPending,
			FifoTopic:                 &isFifo,
			ContentBasedDeduplication: &isContentBasedDeduped,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues, nil
}

type TopicConfig struct {
	Filter string
}

var TopicTables = []interface{}{
	&Topic{},
}

func (c *Client) topics(gConfig interface{}) error {
	var config sns.ListTopicsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	ctx := context.Background()
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(TopicTables...)

	for {
		output, listErr := c.svc.ListTopics(ctx, &config)
		if listErr != nil {
			return listErr
		}

		topics, transformErr := c.transformTopics(&output.Topics)
		if transformErr != nil {
			return transformErr
		}

		c.db.ChunkedCreate(topics)
		c.log.Info("Fetched resources", "resource", "sns.topics", "count", len(output.Topics))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
