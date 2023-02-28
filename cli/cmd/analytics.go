package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/cli/internal/pb"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAnalyticsHost = "analyticsv1.cloudquery.io:443"
)

type AnalyticsClient struct {
	client pb.AnalyticsClient
	conn   *grpc.ClientConn
	host   string
}

func initAnalytics() (*AnalyticsClient, error) {
	host := getEnvOrDefault("CQ_ANALYTICS_HOST", defaultAnalyticsHost)
	var opts []grpc.DialOption
	if strings.HasSuffix(host, ":443") {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = []grpc.DialOption{grpc.WithAuthority(host), grpc.WithTransportCredentials(cred)}
	} else {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial analytics host %v: %w", host, err)
	}
	return &AnalyticsClient{
		client: pb.NewAnalyticsClient(conn),
		conn:   conn,
		host:   host,
	}, nil
}

func (c *AnalyticsClient) SendSyncMetrics(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, invocationUUID string, metrics *source.Metrics, exitReason string) error {
	if metrics == nil {
		// handle nil metrics
		metrics = &source.Metrics{TableClient: map[string]map[string]*source.TableClientMetrics{}}
	}
	if c.client != nil {
		sourcePath := sourceSpec.Path
		if sourceSpec.Registry == specs.RegistryLocal || sourceSpec.Registry == specs.RegistryGrpc {
			_, sourcePath = path.Split(sourceSpec.Path)
		}
		syncSummary := &pb.SyncSummary{
			Invocation_UUID: invocationUUID,
			SourcePath:      sourcePath,
			SourceVersion:   sourceSpec.Version,
			Destinations:    make([]*pb.Destination, 0, 1),
			Resources:       int64(metrics.TotalResources()),
			Errors:          int64(metrics.TotalErrors()),
			Panics:          int64(metrics.TotalPanics()),
			ClientVersion:   Version,
			ExitReason:      exitReason,
		}
		for _, destinationSpec := range destinationsSpecs {
			destPath := destinationSpec.Path
			if destinationSpec.Registry == specs.RegistryLocal || destinationSpec.Registry == specs.RegistryGrpc {
				_, destPath = path.Split(destinationSpec.Path)
			}
			syncSummary.Destinations = append(syncSummary.Destinations, &pb.Destination{
				Path:    destPath,
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

func (c *AnalyticsClient) Host() string {
	return c.host
}

func (c *AnalyticsClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
