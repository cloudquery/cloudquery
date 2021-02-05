package cloudtrail

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
)


type EventSelector struct {
	ID uint `gorm:"primarykey"`
	TrailID uint `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region string `gorm:"-"`
	IncludeManagementEvents *bool
	ReadWriteType *string
}

func (EventSelector) TableName() string {
	return "aws_cloudtrail_trail_event_selectors"
}


func (c *Client) transformEventSelectors(values []*cloudtrail.EventSelector) []*EventSelector {
	var tValues []*EventSelector
	for _, value := range values {
		tValue := EventSelector {
			AccountID: c.accountID,
			Region: c.region,
			IncludeManagementEvents: value.IncludeManagementEvents,
			ReadWriteType: value.ReadWriteType,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

