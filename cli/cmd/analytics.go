package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"github.com/cloudquery/cloudquery/cli/internal/pb"
	"github.com/cloudquery/plugin-sdk/specs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	analyticsHost = "analyticsv1.cloudquery.io:443"
)

type AnalyticsClient struct {
	client pb.AnalyticsClient
	conn  *grpc.ClientConn
}

func initAnalytics() (*AnalyticsClient, error) {
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
					return nil, err
	}
	cred := credentials.NewTLS(&tls.Config{
					RootCAs: systemRoots,
	})
	conn, err := grpc.Dial(analyticsHost, grpc.WithAuthority(analyticsHost), grpc.WithTransportCredentials(cred))
	// conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &AnalyticsClient{
		client: pb.NewAnalyticsClient(conn),
		conn:  conn,
	}, nil
}

func (c *AnalyticsClient) SendSyncSummary(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, resources uint64, errors uint64, panics uint64) error {
	if c.client != nil {
		summary := &pb.SyncSummary{
			SourceName: sourceSpec.Path,
			SourceVersion: sourceSpec.Version,
			Resources: int64(resources),
			Errors: int64(errors),
			Panics: int64(panics),
		}
		for _, destinationSpec := range destinationsSpecs {
			summary.Dest = append(summary.Dest, &pb.Destination{
				Name: destinationSpec.Name,
				Version: destinationSpec.Version,
			})
		}

		_, err := c.client.SendEvent(ctx, &pb.Event_Request{
			SyncSummary: summary,
		})
		return err
	}
	return nil
}

func (c *AnalyticsClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
