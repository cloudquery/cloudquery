package deploy

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/cloudquery/cloudquery/client"
	"github.com/cloudquery/cloudquery/config"
)

type Request struct {
	TaskName string        `json:"taskName"`
	Config   config.Config `json:"config"`
}

func LambdaHandler(ctx context.Context, req Request) (string, error) {
	return TaskExecutor(req)
}

func TaskExecutor(req Request) (string, error) {
	driver := os.Getenv("CQ_DRIVER")
	dsn := os.Getenv("CQ_DSN")
	pluginDir, present := os.LookupEnv("CQ_PLUGIN_DIR")
	if !present {
		pluginDir = "."
	}
	viper.Set("plugin-dir", pluginDir)
	switch req.TaskName {
	case "fetch":
		Fetch(driver, dsn, req.Config)
	case "policy":
		Policy(driver, dsn)
	default:
		return fmt.Sprintf("Unknown task: %s", req.TaskName), fmt.Errorf("unknown task: %s", req.TaskName)
	}
	return fmt.Sprintf("Completed task %s", req.TaskName), nil
}

// Fetches resources from a cloud provider and saves them in the configured database
func Fetch(driver, dsn string, cfg config.Config) {
	c, err := client.New(driver, dsn)
	if err != nil {
		log.Fatalf("Unable to create client: %s", err)
	}
	err = c.Initialize(&cfg)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}
	err = c.Run(&cfg)
	if err != nil {
		log.Fatalf("Error fetching resources: %s", err)
	}
}

// Runs a policy SQL statement and returns results
func Policy(driver, dsn string) {
	outputPath := "/tmp/result.json"
	queryPath := os.Getenv("CQ_QUERY_PATH") // TODO: if path is an S3 URI, pull file down
	c, err := client.New(driver, dsn)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}
	err = c.RunQuery(queryPath, outputPath)
	if err != nil {
		log.Fatalf("Error running query: %s", err)
	}
}
