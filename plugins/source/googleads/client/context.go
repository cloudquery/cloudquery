package client

import (
	"context"

	"google.golang.org/grpc/metadata"
)

const (
	mdDeveloperToken  = "developer-token"   // https://developers.google.com/google-ads/api/docs/concepts/call-structure#developer-token
	mdLoginCustomerID = "login-customer-id" // https://developers.google.com/google-ads/api/docs/concepts/call-structure#cid
)

func (c *Client) OutgoingContext(ctx context.Context) context.Context {
	ctx = c.addDeveloperToken(ctx)
	if len(c.ManagerID) == 0 {
		return ctx
	}
	return addLoginCustomerID(ctx, c.ManagerID)
}

func (c *Client) addDeveloperToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, mdDeveloperToken, c.developerToken)
}

// addLoginCustomerID is useful separately for hierarchy scan as well
func addLoginCustomerID(ctx context.Context, id string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, mdLoginCustomerID, id)
}
