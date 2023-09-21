package services

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/notion/client"
	"github.com/cloudquery/cloudquery/plugins/source/notion/internal/databases"
	"github.com/cloudquery/cloudquery/plugins/source/notion/internal/pages"
	"github.com/cloudquery/cloudquery/plugins/source/notion/internal/users"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func NotionUsersTable() *schema.Table {
	return &schema.Table{
		Name:      "notion_users_table",
		Resolver:  fetchUsersTable,
		Transform: transformers.TransformWithStruct(&users.User{}, transformers.WithPrimaryKeys("Id")),
	}
}

func NotionPagesTable() *schema.Table {
	return &schema.Table{
		Name:      "notion_pages_table",
		Resolver:  fetchPagesTable,
		Transform: transformers.TransformWithStruct(&pages.Page{}, transformers.WithPrimaryKeys("Id")),
	}
}

func NotionDatabasesTable() *schema.Table {
	return &schema.Table{
		Name:      "notion_databases_table",
		Resolver:  fetchDatabasesTable,
		Transform: transformers.TransformWithStruct(&databases.Database{}, transformers.WithPrimaryKeys("Id")),
	}
}

func fetchUsersTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	data, err := c.Notion.GetUsers("", false)
	if err != nil {
		return err
	}
	sendUserDataToChan(data.Results, res)
	for data.HasMore {
		data, err = c.Notion.GetUsers(data.NextCursor, data.HasMore)
		if err != nil {
			return nil
		}
		sendUserDataToChan(data.Results, res)
	}
	return nil
}

func sendUserDataToChan(u []users.User, res chan<- any) {
	for _, user := range u {
		res <- user
	}
}

func fetchPagesTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	data, err := c.Notion.GetPages("", false)
	if err != nil {
		return err
	}
	sendPagesDataToChan(data.Results, res)
	for data.HasMore {
		data, err = c.Notion.GetPages(data.NextCursor, data.HasMore)
		if err != nil {
			return nil
		}
		sendPagesDataToChan(data.Results, res)
	}
	return nil
}

func sendPagesDataToChan(p []pages.Page, res chan<- any) {
	for _, page := range p {
		res <- page
	}
}

func fetchDatabasesTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	data, err := c.Notion.GetDatabases("", false)
	if err != nil {
		return err
	}
	sendDatabasesDataToChan(data.Results, res)
	for data.HasMore {
		data, err = c.Notion.GetDatabases(data.NextCursor, data.HasMore)
		if err != nil {
			return nil
		}
		sendDatabasesDataToChan(data.Results, res)
	}
	return nil
}

func sendDatabasesDataToChan(d []databases.Database, res chan<- any) {
	for _, database := range d {
		res <- database
	}
}
