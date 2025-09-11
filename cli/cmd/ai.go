package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/cloudquery/cloudquery/cli/v6/aionboarding"
	"github.com/cloudquery/cloudquery/cli/v6/aionboarding/openaiapi"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	aiShort   = `Start an AI-powered interactive conversation to help with CloudQuery syncs`
	aiExample = `# Start an AI conversation to get help with CloudQuery
cloudquery ai`
)

var (
	aiBold    = color.New(color.Bold)
	aiSuccess = color.New(color.Bold, color.FgGreen)
	aiInfo    = color.New(color.Bold, color.FgCyan)
)

func newCmdAI() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ai",
		Short:   aiShort,
		Long:    aiShort,
		Example: aiExample,
		Args:    cobra.ExactArgs(0),
		RunE:    aiCmd,
	}
	return cmd
}

func aiCmd(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// Check for OpenAI API key
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		return errors.New("OPENAI_API_KEY environment variable is required")
	}

	// Track AI session start
	TrackAISessionStarted(ctx, invocationUUID.UUID)

	fmt.Println()
	aiSuccess.Println("🤖 CloudQuery AI Assistant")
	fmt.Println("I'm here to help you set up CloudQuery syncs!")
	fmt.Println("Type 'exit' or 'quit' to end the conversation.")
	fmt.Println()
	fmt.Println("What are you trying to build with CloudQuery?")
	fmt.Println()

	aiOnboarding, err := aionboarding.NewAIOnboarding(
		ctx,
		aionboarding.WithOpenAIAPIKey(openaiKey),
		aionboarding.WithPosthogAPIKey("fake"),
		aionboarding.WithDebug(false),
	)
	if err != nil {
		return fmt.Errorf("failed to create AI onboarding: %w", err)
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

		response, err := aiOnboarding.Chat(ctx, "123", "default", &userInput)
		if err != nil {
			return fmt.Errorf("failed to chat: %w", err)
		}

		for response.HasFunctionCalls() {
			switch response.FunctionCall {
			case "create_spec_file":
				err := createSpecFile(response.FunctionCallArguments["filename_without_extension"].(string), response.FunctionCallArguments["content"].(string))
				if err != nil {
					return fmt.Errorf("failed to create spec file: %w", err)
				}
				response, err = aiOnboarding.Chat(
					ctx,
					"123",
					"default",
					nil,
					openaiapi.WithFunctionCallOutput(
						response.FunctionCall,
						response.FunctionCallArguments,
						response.FunctionCallID,
						"Spec file created",
					),
				)
				if err != nil {
					return fmt.Errorf("failed to chat: %w", err)
				}
			default:
				return fmt.Errorf("unsupported function call: %s", response.FunctionCall)
			}
		}

		// Display the AI response
		fmt.Printf("\n%s %s\n\n", aiBold.Sprint("AI:"), response.Message)
	}

	// Track AI session end
	TrackAISessionEnded(ctx, invocationUUID.UUID)

	return nil
}

func createSpecFile(filename, content string) error {
	return os.WriteFile(filename+".yaml", []byte(content), 0644)
}
