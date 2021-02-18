package iam
import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)


type UserAttachedPolicy struct {
	ID uint `gorm:"primarykey"`
	UserID uint `neo:"ignore"`
	AccountID string `gorm:"-"`
	PolicyArn *string
	PolicyName *string
}

func (UserAttachedPolicy) TableName() string {
	return "aws_iam_user_attached_policies"
}

func (c *Client) transformAttachedPolicies(values *[]types.AttachedPolicy) []*UserAttachedPolicy {
	var tValues []*UserAttachedPolicy
	for _, value := range *values {
		tValue := UserAttachedPolicy{
			AccountID: c.accountID,
			PolicyArn: value.PolicyArn,
			PolicyName: value.PolicyName,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

