package deploy

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/client"
	"log"
	"os"
)


type Request struct {
	TaskName string `json:"taskName"`
}


func LambdaHandler(ctx context.Context, req Request) (string, error) {
	return TaskExecutor(req.TaskName)
}

func TaskExecutor(taskName string) (string, error) {
	driver := os.Getenv("CQ_DRIVER")
	dsn := os.Getenv("CQ_DSN")
	switch taskName {
	case "fetch":
		Fetch(driver, dsn)
	case "policy":
		Policy(driver, dsn)
	default:
		return fmt.Sprintf("Unknown task: %s", taskName), fmt.Errorf("Unkown task: %s", taskName)
	}
	return fmt.Sprintf("Completed task %s", taskName), nil
}

// Fetches resources from a cloud provider and saves them in the configured database
func Fetch(driver, dsn string) {
	c, err := client.New("config.yml", driver, dsn)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}
	err = c.Run("config.yml")
	if err != nil {
		log.Fatalf("Error fetching resources: %s", err)
	}
}

// Runs a policy SQL statement and returns results
func Policy(driver, dsn string) {
	outputPath := "/tmp/result.json"
	queryPath := os.Getenv("CQ_QUERY_PATH") // TODO: if path is an S3 URI, pull file down
	c, err := client.New("", driver, dsn)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}
	err = c.RunQuery(queryPath, outputPath)
	if err != nil {
		log.Fatalf("Error running query: %s", err)
	}
}
