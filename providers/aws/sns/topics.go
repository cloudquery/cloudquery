package sns

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
)

type Topic struct {
	ID                        uint `gorm:"primarykey"`
	AccountID                 string
	Region                    string
	Owner                     *string
	TopicArn                  *string
	Policy                    *string
	DeliveryPolicy            *string
	DisplayName               *string
	SubscriptionsConfirmed    *int
	SubscriptionsDeleted      *int
	SubscriptionsPending      *int
	EffectiveDeliveryPolicy   *string
	FifoTopic                 *bool
	ContentBasedDeduplication *bool
}

func (Topic) TableName() string {
	return "aws_sns_topics"
}

func GetTopicAttributes(c *Client, TopicArn *string) (map[string]*string, error) {
	params := &sns.GetTopicAttributesInput{
		TopicArn: TopicArn,
	}
	output, err := c.svc.GetTopicAttributes(params)
	if err != nil {
		return nil, err
	}
	return output.Attributes, nil
}

func getOrZero(attrs map[string]*string, keyName string) string {
	if v, found := attrs[keyName]; found {
		return *v
	}
	return "0"
}

func (c *Client) transformTopics(values []*sns.Topic) ([]*Topic, error) {
	var tValues []*Topic
	for _, value := range values {

		// All topic attributes are returned as a string; we have to handle type conversion
		output, err := GetTopicAttributes(c, value.TopicArn)
		if err != nil {
			return nil, err
		}

		subsConfirmed, subsConfirmedErr := strconv.Atoi(getOrZero(output, "SubscriptionsConfirmed"))
		if subsConfirmedErr != nil {
			return nil, subsConfirmedErr
		}

		subsDeleted, subsDeletedErr := strconv.Atoi(getOrZero(output, "SubscriptionsDeleted"))
		if subsDeletedErr != nil {
			return nil, subsDeletedErr
		}

		subsPending, subsPendingErr := strconv.Atoi(getOrZero(output, "SubscriptionsPending"))
		if subsPendingErr != nil {
			return nil, subsPendingErr
		}

		isFifo, isFifoErr := strconv.ParseBool(getOrZero(output, "FifoTopic"))
		if isFifoErr != nil {
			return nil, isFifoErr
		}

		isContentBasedDeduped, isContentBasedDedupedErr := strconv.ParseBool(getOrZero(output, "ContentBasedDeduplication"))
		if isContentBasedDedupedErr != nil {
			return nil, isContentBasedDedupedErr
		}

		tValue := Topic{
			AccountID:                 c.accountID,
			Region:                    c.region,
			TopicArn:                  value.TopicArn,
			Policy:                    output["Policy"],
			DeliveryPolicy:            output["DeliveryPolicy"],
			DisplayName:               output["DisplayName"],
			Owner:                     output["Owner"],
			EffectiveDeliveryPolicy:   output["EffectiveDeliveryPolicy"],
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
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(TopicTables...)

	for {
		output, listErr := c.svc.ListTopics(&config)
		if listErr != nil {
			return listErr
		}

		topics, transformErr := c.transformTopics(output.Topics)
		if transformErr != nil {
			return transformErr
		}

		c.db.ChunkedCreate(topics)
		c.log.Info("Fetched resources", zap.String("resource", "sns.topics"), zap.Int("count", len(output.Topics)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
