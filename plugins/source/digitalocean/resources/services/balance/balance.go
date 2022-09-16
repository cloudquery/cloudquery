// Code generated by codegen; DO NOT EDIT.

package balance

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func Balance() *schema.Table {
	return &schema.Table{
		Name:     "digitalocean_balance",
		Resolver: fetchBalanceBalance,
		Columns: []schema.Column{
			{
				Name:     "month_to_date_balance",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MonthToDateBalance"),
			},
			{
				Name:     "account_balance",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AccountBalance"),
			},
			{
				Name:     "month_to_date_usage",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MonthToDateUsage"),
			},
			{
				Name:     "generated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("GeneratedAt"),
			},
		},
	}
}
