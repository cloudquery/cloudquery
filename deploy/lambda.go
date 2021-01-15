package deploy

import (
	"context"
	"fmt"
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/viper"
	"log"
)


type Request struct {
	TaskName string `json:"taskName"`
}


func LambdaHandler(ctx context.Context, req Request) (string, error) {
	return TaskExecutor(req.TaskName)
}

func TaskExecutor(taskName string) (string, error) {
	driver := viper.GetString("driver")
	dsn := viper.GetString("dsn")
	switch taskName {
	case "fetch":
		Fetch(driver, dsn, false)
	case "policy":
		Policy(driver, dsn, false)
	default:
		return fmt.Sprintf("Unknown task: %s", taskName), fmt.Errorf("Unkown task: %s", taskName)
	}
	return fmt.Sprintf("Completed task %s", taskName), nil
}

// Fetches resources from a cloud provider and saves them in the configured database
func Fetch(driver, dsn string, verbose bool) {
	client, err := cloudqueryclient.New(driver, dsn, verbose)
	if err != nil {
		log.Fatalf("Unable to initialize client: %s", err)
	}
	err = client.Run("config.yml")
	if err != nil {
		log.Fatalf("Error fetching resources: %s", err)
	}
}

// Runs a policy SQL statement and returns results
func Policy(driver, dsn string, verbose bool) {
	fmt.Println("Running policy queries")
}
