package users

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func UserContactMethods() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_user_contact_methods",
		Description: `https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods`,
		Resolver:    fetchUserContactMethods,
		Transform:   transformers.TransformWithStruct(&pagerduty.ContactMethod{}, transformers.WithSkipFields("HTMLURL", "SendHTMLEmail")),
		Columns: []schema.Column{
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "send_html_email",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("SendHTMLEmail"),
			},
		},
	}
}
