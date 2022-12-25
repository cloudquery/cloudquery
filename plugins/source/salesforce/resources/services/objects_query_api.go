package services

import (
	"context"
	"strings"

	"github.com/cloudquery/cloudquery/plugins/source/salesforce/client"
)

type queryResponse struct {
	TotalSize      int              `json:"totalSize"`
	Records        []map[string]any `json:"records"`
	Done           bool             `json:"done"`
	NextRecordsUrl string           `json:"nextRecordsUrl"`
}

func fetchQueryApi(ctx context.Context, c *client.Client, fields []string, res chan<- any) error {
	var url strings.Builder
	url.WriteString(c.HTTPDataEndpoint)
	url.WriteString("/query?q=SELECT+")
	url.WriteString(strings.Join(fields, ","))
	url.WriteString("+FROM+")
	url.WriteString(c.Object)

	var queryRes queryResponse
	if err := c.Get(ctx, url.String(), &queryRes); err != nil {
		return err
	}
	res <- queryRes.Records
	for !queryRes.Done {
		url := c.HTTPDataEndpoint + "/query/" + queryRes.NextRecordsUrl
		if err := c.Get(ctx, url, &queryRes); err != nil {
			return err
		}
		res <- queryRes.Records
	}
	return nil
}
