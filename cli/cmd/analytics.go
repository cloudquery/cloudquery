package cmd

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/internal/env"
	"github.com/cloudquery/cloudquery/cli/v6/internal/specs/v0"
	"github.com/cloudquery/plugin-pb-go/metrics"
	"github.com/cloudquery/plugin-pb-go/pb/analytics/v0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAnalyticsHost = "analyticsv1.cloudquery.io:443"
)

const (
	ExitReasonUnset     ExitReason = ""
	ExitReasonStopped   ExitReason = "stopped"
	ExitReasonCompleted ExitReason = "completed"
)

type AnalyticsClient struct {
	client analytics.AnalyticsClient
	conn   *grpc.ClientConn
	host   string
}

func initAnalytics() (*AnalyticsClient, error) {
	host := env.GetEnvOrDefault("CQ_ANALYTICS_HOST", defaultAnalyticsHost)
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
	// TODO: Remove once there's a documented migration path per https://github.com/grpc/grpc-go/issues/7244
	// nolint:staticcheck
	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to dial analytics host %v: %w", host, err)
	}
	return &AnalyticsClient{
		client: analytics.NewAnalyticsClient(conn),
		conn:   conn,
		host:   host,
	}, nil
}

func (c *AnalyticsClient) SendSyncMetrics(ctx context.Context, sourceSpec specs.Source, destinationsSpecs []specs.Destination, invocationUUID string, m *metrics.Metrics, exitReason ExitReason) error {
	if m == nil {
		// handle nil metrics
		m = &metrics.Metrics{TableClient: map[string]map[string]*metrics.TableClientMetrics{}}
	}
	if c.client != nil {
		sourcePath := sourceSpec.Path
		if sourceSpec.Registry == specs.RegistryLocal || sourceSpec.Registry == specs.RegistryGRPC {
			_, sourcePath = path.Split(sourceSpec.Path)
		}
		syncSummary := &analytics.SyncSummary{
			Invocation_UUID: invocationUUID,
			SourcePath:      sourcePath,
			SourceVersion:   sourceSpec.Version,
			Destinations:    make([]*analytics.Destination, 0, 1),
			Resources:       int64(m.TotalResources()),
			Errors:          int64(m.TotalErrors()),
			Panics:          int64(m.TotalPanics()),
			ClientVersion:   Version,
			ExitReason:      string(exitReason),
		}
		for _, destinationSpec := range destinationsSpecs {
			destPath := destinationSpec.Path
			if destinationSpec.Registry == specs.RegistryLocal || destinationSpec.Registry == specs.RegistryGRPC {
				_, destPath = path.Split(destinationSpec.Path)
			}
			syncSummary.Destinations = append(syncSummary.Destinations, &analytics.Destination{
				Path:    destPath,
				Version: destinationSpec.Version,
			})
		}

		_, err := c.client.SendEvent(ctx, &analytics.Event_Request{
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
