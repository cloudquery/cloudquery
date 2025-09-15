package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"

	"github.com/fatih/color"
)

var (
	aiBold    = color.New(color.Bold)
	aiSuccess = color.New(color.Bold, color.FgGreen)
	aiInfo    = color.New(color.Bold, color.FgCyan)
)

func aiCmd(ctx context.Context, client *cloudquery_api.ClientWithResponses, teamName string) error {

	// Track AI session start
	TrackAISessionStarted(ctx, invocationUUID.UUID)

	fmt.Println()
	aiSuccess.Println("🤖 CloudQuery AI Assistant")
	fmt.Println("I'm here to help you set up CloudQuery syncs!")
	fmt.Println("Type 'exit' or 'quit' to end the conversation.")
	fmt.Println()
	fmt.Println("What are you trying to build with CloudQuery?")
	fmt.Println()

	if _, err := client.AIOnboardingNewConversationWithResponse(ctx, teamName, cloudquery_api.AIOnboardingNewConversationJSONRequestBody{}); err != nil {
		return fmt.Errorf("failed to start new conversation: %w", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(aiInfo.Sprint("You: "))
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}

		// Check for exit commands
		if strings.ToLower(userInput) == "exit" || strings.ToLower(userInput) == "quit" {
			fmt.Println("Goodbye! 👋")
			break
		}

		response, err := api.Chat(ctx, client, teamName, &userInput, nil)
		if err != nil {
			return fmt.Errorf("failed to chat: %w", err)
		}
		for response.FunctionCall != nil {
			switch *response.FunctionCall {
			case "create_spec_file":
				err := createSpecFile(response.FunctionCallArguments["filename_without_extension"].(string), response.FunctionCallArguments["content"].(string))
				if err != nil {
					return fmt.Errorf("failed to create spec file: %w", err)
				}
				response, err = api.Chat(ctx, client, teamName, nil, &[]api.FunctionCallOutput{
					{
						Name:      "create_spec_file",
						CallID:    response.FunctionCallID,
						Arguments: response.FunctionCallArguments,
						Output:    "Spec file created",
					},
				})
				if err != nil {
					return fmt.Errorf("failed to chat: %w", err)
				}
			case "cloudquery_test":
				response, err = api.Chat(ctx, client, teamName, nil, &[]api.FunctionCallOutput{
					{
						Name:      "cloudquery_test",
						CallID:    response.FunctionCallID,
						Arguments: response.FunctionCallArguments,
						Output:    cloudqueryTest(response.FunctionCallArguments["filename"].(string)),
					},
				})
				if err != nil {
					return fmt.Errorf("failed to chat: %w", err)
				}
			default:
				return fmt.Errorf("unsupported function call: %s", *response.FunctionCall)
			}
		}

		fmt.Printf("\n%s %s\n\n", aiBold.Sprint("AI:"), *response.Message)
	}

	// Track AI session end
	TrackAISessionEnded(ctx, invocationUUID.UUID)

	return nil
}

func createSpecFile(filenameWithoutExtension, content string) error {
	return os.WriteFile(filenameWithoutExtension+".yaml", []byte(content), 0644)
}

func cloudqueryTest(filenameWithoutExtension string) string {
	cmd := exec.Command("cloudquery", "test", filenameWithoutExtension+".yaml")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("cloudquery test failed: %v\nOutput:\n%s", err, out.String())
	}
	return fmt.Sprintf("Result of cloudquery test:\n%s", out.String())
}
