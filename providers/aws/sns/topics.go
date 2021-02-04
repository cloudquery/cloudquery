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

func GetTopicAttributes(c *Client, TopicArn *string) map[string]*string {
	params := &sns.GetTopicAttributesInput{
		TopicArn: TopicArn,
	}
	output, _ := c.svc.GetTopicAttributes(params)
	return output.Attributes
}

func getOrDefault(attrs map[string]*string, keyName string) string {
	if v, found := attrs[keyName]; found {
		return *v
	}
	return "0"
}

func (c *Client) transformTopics(values []*sns.Topic) []*Topic {
	var tValues []*Topic
	for _, value := range values {

		output := GetTopicAttributes(c, value.TopicArn)

		subsConfirmed, _ := strconv.Atoi(getOrDefault(output, "SubscriptionsConfirmed"))
		subsDeleted, _ := strconv.Atoi(getOrDefault(output, "SubscriptionsDeleted"))
		subsPending, _ := strconv.Atoi(getOrDefault(output, "SubscriptionsPending"))
		isFifo, _ := strconv.ParseBool(getOrDefault(output, "FifoTopic"))
		isContentBasedDeduped, _ := strconv.ParseBool(getOrDefault(output, "ContentBasedDeduplication"))

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
	return tValues
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
		output, err := c.svc.ListTopics(&config)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(c.transformTopics(output.Topics))
		c.log.Info("Fetched resources", zap.String("resource", "sns.topics"), zap.Int("count", len(output.Topics)))
		if aws.StringValue(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}

	return nil
}
