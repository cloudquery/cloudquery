package deploy

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"

	"github.com/spf13/viper"
)

type Request struct {
	TaskName string      `json:"taskName"`
	Config   interface{} `json:"config"`
}

func LambdaHandler(ctx context.Context, req Request) (string, error) {
	return TaskExecutor(ctx, req)
}

func TaskExecutor(ctx context.Context, req Request) (string, error) {
	dsn := os.Getenv("CQ_DSN")
	pluginDir, present := os.LookupEnv("CQ_PLUGIN_DIR")
	if !present {
		pluginDir = "."
	}
	viper.Set("plugin-dir", pluginDir)
	policyDir, present := os.LookupEnv("CQ_POLICY_DIR")
	if !present {
		policyDir = "."
	}
	viper.Set("policy-dir", policyDir)
	b, err := json.Marshal(req.Config)

	if err != nil {
		return "", err
	}
	cfg, diags := config.NewParser(nil).LoadConfigFromJson("config.json", b)
	if diags != nil {
		return "", fmt.Errorf("bad configuration: %s", diags)
	}
	// Override dsn env if set
	if dsn != "" {
		cfg.CloudQuery.Connection.DSN = dsn
	}

	switch req.TaskName {
	case "fetch":
		Fetch(ctx, cfg)
	case "policy":
		Policy(ctx, cfg)
	default:
		return fmt.Sprintf("Unknown task: %s", req.TaskName), fmt.Errorf("unknown task: %s", req.TaskName)
	}
	return fmt.Sprintf("Completed task %s", req.TaskName), nil
}

// Fetch fetches resources from a cloud provider and saves them in the configured database
func Fetch(ctx context.Context, cfg *config.Config) {
	c, err := client.New(cfg, func(c *client.Client) {
		c.Hub = registry.NewRegistryHub(registry.CloudQueryRegistryURl, func(h *registry.Hub) {
			h.PluginDirectory = cfg.CloudQuery.PluginDirectory
		})
	})
	if err != nil {
		log.Fatalf("Unable to create client: %s", err)
	}
	err = c.Initialize(ctx)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}
	err = c.Fetch(ctx, client.FetchRequest{
		Providers: cfg.Providers,
	})
	if err != nil {
		log.Fatalf("Error fetching resources: %s", err)
	}
}

// Policy Runs a policy SQL statement and returns results
func Policy(ctx context.Context, cfg *config.Config) {
	outputPath := "/tmp/result.json"
	queryPath := os.Getenv("CQ_QUERY_PATH") // TODO: if path is an S3 URI, pull file down
	c, err := client.New(cfg)
	if err != nil {
		log.Fatalf("Unable to create client: %s", err)
	}
	_, err = c.ExecutePolicy(ctx, client.ExecutePolicyRequest{
		PolicyPath:    queryPath,
		StopOnFailure: false,
		OutputPath:    outputPath,
	})
	if err != nil {
		log.Fatalf("Error running query: %s", err)
	}
}
