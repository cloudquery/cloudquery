package users

import (
	"github.com/PagerDuty/go-pagerduty"
	"github.com/apache/arrow/go/v13/arrow"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func UserContactMethods() *schema.Table {
	return &schema.Table{
		Name:        "pagerduty_user_contact_methods",
		Description: `https://developer.pagerduty.com/api-reference/50d46c0eb020d-list-a-user-s-contact-methods`,
		Resolver:    fetchUserContactMethods,
		Transform:   transformers.TransformWithStruct(&pagerduty.ContactMethod{}, transformers.WithSkipFields("HTMLURL", "SendHTMLEmail")),
		Columns: []schema.Column{
			{
				Name:       "id",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("ID"),
				PrimaryKey: true,
			},
			{
				Name:     "html_url",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "send_html_email",
				Type:     arrow.FixedWidthTypes.Boolean,
				Resolver: schema.PathResolver("SendHTMLEmail"),
			},
		},
	}
}
