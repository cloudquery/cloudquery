package organizations

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/mitchellh/mapstructure"
)

type Account struct {
	ID              uint `gorm:"primarykey"`
	CallerAccountID string
	AccountID       *string `neo:"unique"`
	Arn             *string `neo:"unique" gorm:"unique"`
	Email           *string
	JoinedMethod    *string
	JoinedTimestamp *time.Time
	Name            *string
	Status          *string
}

func (Account) TableName() string {
	return "aws_organizations_accounts"
}

func (c *Client) transformAccounts(values *[]types.Account) ([]*Account, error) {
	var tValues []*Account
	for _, value := range *values {
		tValues = append(tValues, &Account{
			CallerAccountID: c.accountID,
			AccountID:       value.Id,
			Arn:             value.Arn,
			Email:           value.Email,
			JoinedMethod:    aws.String(string(value.JoinedMethod)),
			JoinedTimestamp: value.JoinedTimestamp,
			Name:            value.Name,
			Status:          aws.String(string(value.Status)),
		})
	}
	return tValues, nil
}

var AccountTables = []interface{}{
	&Account{},
}

func (c *Client) accounts(gConfig interface{}) error {
	var config organizations.ListAccountsInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	ctx := context.Background()
	// TODO: This doesn't work, since the account ids are not coming from the client but from the sdk call
	c.db.Where("caller_account_id", c.accountID).Delete(AccountTables...)

	for {
		output, err := c.svc.ListAccounts(ctx, &config)
		if err != nil {
			return err
		}
		tValues, err := c.transformAccounts(&output.Accounts)
		if err != nil {
			return err
		}
		c.db.ChunkedUpsert(tValues)
		c.log.Info("Fetched resources", "resource", "organizations.accounts", "count", len(output.Accounts))
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
