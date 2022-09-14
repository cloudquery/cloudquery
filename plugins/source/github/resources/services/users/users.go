package users

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

// Users is a utility table to be used as a relation
func Users() *schema.Table {
	return &schema.Table{
		Columns: []schema.Column{
			{
				Name:            "login",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "id",
				Resolver:        schema.PathResolver("ID"),
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "node_id",
				Resolver: schema.PathResolver("NodeID"),
				Type:     schema.TypeString,
			},
			{
				Name:     "avatar_url",
				Resolver: schema.PathResolver("AvatarURL"),
				Type:     schema.TypeString,
			},
			{
				Name:     "html_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HTMLURL"),
			},
			{
				Name:     "gravatar_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GravatarID"),
			},
			{
				Name: "name",
				Type: schema.TypeString,
			},
			{
				Name: "company",
				Type: schema.TypeString,
			},
			{
				Name: "blog",
				Type: schema.TypeString,
			},
			{
				Name: "location",
				Type: schema.TypeString,
			},
			{
				Name: "email",
				Type: schema.TypeString,
			},
			{
				Name: "hireable",
				Type: schema.TypeBool,
			},
			{
				Name: "bio",
				Type: schema.TypeString,
			},
			{
				Name: "twitter_username",
				Type: schema.TypeString,
			},
			{
				Name: "public_repos",
				Type: schema.TypeInt,
			},
			{
				Name: "public_gists",
				Type: schema.TypeInt,
			},
			{
				Name: "followers",
				Type: schema.TypeInt,
			},
			{
				Name: "following",
				Type: schema.TypeInt,
			},
			{
				Name:     "created_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt.Time"),
			},
			{
				Name:     "updated_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt.Time"),
			},
			{
				Name:     "suspended_at_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SuspendedAt.Time"),
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "site_admin",
				Type: schema.TypeBool,
			},
			{
				Name: "total_private_repos",
				Type: schema.TypeInt,
			},
			{
				Name: "owned_private_repos",
				Type: schema.TypeInt,
			},
			{
				Name: "private_gists",
				Type: schema.TypeInt,
			},
			{
				Name: "disk_usage",
				Type: schema.TypeInt,
			},
			{
				Name: "collaborators",
				Type: schema.TypeInt,
			},
			{
				Name: "two_factor_authentication",
				Type: schema.TypeBool,
			},
			{
				Name: "plan",
				Type: schema.TypeJSON,
			},
			{
				Name:        "url",
				Description: "API URLs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("URL"),
			},
			{
				Name:     "events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventsURL"),
			},
			{
				Name:     "following_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FollowingURL"),
			},
			{
				Name:     "followers_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FollowersURL"),
			},
			{
				Name:     "gists_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("GistsURL"),
			},
			{
				Name:     "organizations_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OrganizationsURL"),
			},
			{
				Name:     "received_events_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReceivedEventsURL"),
			},
			{
				Name:     "repos_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReposURL"),
			},
			{
				Name:     "starred_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StarredURL"),
			},
			{
				Name:     "subscriptions_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubscriptionsURL"),
			},
			{
				Name: "text_matches",
				Type: schema.TypeJSON,
			},
			{
				Name: "permissions",
				Type: schema.TypeJSON,
			},
			{
				Name: "role_name",
				Type: schema.TypeString,
			},
		},
	}
}
