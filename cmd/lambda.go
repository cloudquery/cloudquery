package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cloudquery/cloudquery/cloudqueryclient"
	"github.com/spf13/cobra"
)

var DRIVER string
var DSN string

type Request struct {
	TaskName string `json:"taskName"`
}

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "Runs cloudquery compatibly with AWS Lambda",
	RunE: func(cmd *cobra.Command, args []string) error {
		if env := os.Getenv("AWS_LAMBDA_RUNTIME_API"); env != "" {
			lambda.Start(LambdaHandler)
		} else if len(args) > 0 {
			TaskExecutor(args[0])
		} else {
			return fmt.Errorf("No AWS_LAMBDA_RUNTIME_API environment variable detected, or no argument was passed.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(lambdaCmd)
	DRIVER = os.Getenv("CLOUDQUERY_DRIVER")
	DSN = os.Getenv("CLOUDQUERY_DATABASE_STRING")
}

func LambdaHandler(ctx context.Context, req Request) (string, error) {
	return TaskExecutor(req.TaskName)
}

func TaskExecutor(taskName string) (string, error) {
	switch taskName {
	case "fetch":
		Fetch(DRIVER, DSN, false)
	case "policy":
		Policy(DRIVER, DSN, false)
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
