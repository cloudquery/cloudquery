package deploy

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudquery/cloudquery/pkg/client"
	"github.com/cloudquery/cloudquery/pkg/config"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Request struct {
	TaskName string                 `json:"taskName"`
	Config   map[string]interface{} `json:"config"`
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
	if dsn != "" {

	}
	data, err := json.Marshal(req.Config)
	if err != nil {
		return "", fmt.Errorf("failed to parse request config: %w", err)
	}
	cfg, diags := config.NewParser(nil).LoadConfigFromSource("lambda_config.json", data)
	if diags != nil {
		return "", fmt.Errorf("bad configuration: %s", diags)
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
	c, err := client.New(cfg)
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
	//outputPath := "/tmp/result.json"
	//queryPath := os.Getenv("CQ_QUERY_PATH") // TODO: if path is an S3 URI, pull file down
	//c, err := client.New(cfg)
	//if err != nil {
	//	log.Fatalf("Unable to initialize client: %s", err)
	//}
	////err = c.ExecutePolicy(context.Background(), queryPath, outputPath)
	//if err != nil {
	//	log.Fatalf("Error running query: %s", err)
	//}
}
