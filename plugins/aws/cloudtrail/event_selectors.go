package cloudtrail

import (
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)


type EventSelector struct {
	ID uint `gorm:"primarykey"`
	TrailID uint `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region string `gorm:"-"`
	IncludeManagementEvents *bool
	ReadWriteType string
}

func (EventSelector) TableName() string {
	return "aws_cloudtrail_trail_event_selectors"
}


func (c *Client) transformEventSelectors(values []types.EventSelector) []*EventSelector {
	var tValues []*EventSelector
	for _, value := range values {
		tValue := EventSelector {
			AccountID: c.accountID,
			Region: c.region,
			IncludeManagementEvents: value.IncludeManagementEvents,
			ReadWriteType: string(value.ReadWriteType),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

