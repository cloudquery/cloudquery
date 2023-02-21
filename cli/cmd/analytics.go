package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"github.com/cloudquery/cloudquery/cli/internal/pb"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// default analytics host
	analyticsHost = "analyticsv1.cloudquery.io:443"
)

type AnalyticsClient struct {
	client pb.AnalyticsClient
	conn   *grpc.ClientConn
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
	if err != nil {
		return nil, err
	}

	return &AnalyticsClient{
		client: pb.NewAnalyticsClient(conn),
		conn:   conn,
	}, nil
}

func (c *AnalyticsClient) SendSyncMetrics(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, invocationUUID string, metrics *source.Metrics) error {
	if c.client != nil {
		syncSummary := &pb.SyncSummary{
			Invocation_UUID: invocationUUID,
			SourcePath:      sourceSpec.Path,
			SourceVersion:   sourceSpec.Version,
			Destinations:    make([]*pb.Destination, 0, 1),
			Resources:       int64(metrics.TotalResources()),
			Errors:          int64(metrics.TotalErrors()),
			Panics:          int64(metrics.TotalPanics()),
			ClientVersion:   Version,
		}
		for _, destinationSpec := range destinationsSpecs {
			syncSummary.Destinations = append(syncSummary.Destinations, &pb.Destination{
				Path:    destinationSpec.Path,
				Version: destinationSpec.Version,
			})
		}

		_, err := c.client.SendEvent(ctx, &pb.Event_Request{
			SyncSummary: syncSummary,
		})
		return err
	}
	return nil
}

func (c *AnalyticsClient) SendSyncSummary(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, invocationUUID string, summary schema.SyncSummary) error {
	if c.client != nil {
		syncSummary := &pb.SyncSummary{
			Invocation_UUID: invocationUUID,
			SourcePath:      sourceSpec.Path,
			SourceVersion:   sourceSpec.Version,
			Destinations:    make([]*pb.Destination, 0, 1),
			Resources:       int64(summary.Resources),
			Errors:          int64(summary.Errors),
			Panics:          int64(summary.Panics),
			ClientVersion:   Version,
		}
		for _, destinationSpec := range destinationsSpecs {
			syncSummary.Destinations = append(syncSummary.Destinations, &pb.Destination{
				Path:    destinationSpec.Path,
				Version: destinationSpec.Version,
			})
		}

		_, err := c.client.SendEvent(ctx, &pb.Event_Request{
			SyncSummary: syncSummary,
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
