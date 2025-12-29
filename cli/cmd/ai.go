package cmd

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	gosync "sync"
	"syscall"
	"time"

	cloudquery_api "github.com/cloudquery/cloudquery-api-go"
	"github.com/cloudquery/cloudquery/cli/v6/internal/api"
	"github.com/fatih/color"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
)

var (
	aiBold    = color.New(color.Bold)
	aiSuccess = color.New(color.Bold, color.FgGreen)
	aiInfo    = color.New(color.Bold, color.FgCyan)
)

// Spinner messages for different operations
var (
	generalMessages = []string{
		"🤖 Thinking...",
		"📚 Consulting plugin documentation...",
		"🌐 Calling CloudQuery APIs...",
		"🧠 Processing your request...",
		"⚡ Generating response...",
		"🔍 Analyzing your query...",
		"💭 Crafting the perfect answer...",
		"✨ Almost ready...",
	}

	specFileMessages = []string{
		"📝 Creating spec file...",
		"⚙️  Configuring CloudQuery...",
		"🔧 Setting up sync configuration...",
		"📋 Writing YAML configuration...",
		"🎯 Optimizing table selection...",
		"✨ Finalizing spec file...",
	}

	testMessages = []string{
		"🧪 Running CloudQuery test...",
		"🔍 Validating configuration...",
		"⚡ Testing sync capabilities...",
		"📊 Checking table schemas...",
		"🔧 Verifying connections...",
		"✅ Almost done testing...",
	}
)

// startSpinner displays a spinner with rotating messages and returns a stop function
func startSpinner(ctx context.Context, messages []string) func() {
	var mu gosync.Mutex
	var currentMessage string
	var stopped bool

	// Create a progressbar for the spinner
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription(""),
		progressbar.OptionSetWidth(50),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetRenderBlankState(true),
	)

	// Create a context for the spinner that can be cancelled
	spinnerCtx, spinnerCancel := context.WithCancel(ctx)

	// Channel to signal when cleanup is complete
	cleanupDone := make(chan struct{})

	// Start the message rotation
	go func() {
		ticker := time.NewTicker(4 * time.Second)
		defer ticker.Stop()

		messageIndex := 0
		for {
			select {
			case <-spinnerCtx.Done():
				return
			case <-ticker.C:
				mu.Lock()
				if !stopped {
					currentMessage = messages[messageIndex%len(messages)]
					messageIndex++
				}
				mu.Unlock()
			}
		}
	}()

	// Start the spinner animation
	spinnerChars := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinnerIndex := 0

	go func() {
		defer close(cleanupDone)
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-spinnerCtx.Done():
				mu.Lock()
				stopped = true
				mu.Unlock()
				_ = bar.Clear()
				return
			case <-ticker.C:
				mu.Lock()
				if stopped {
					mu.Unlock()
					continue
				}
				msg := currentMessage
				mu.Unlock()

				if msg == "" {
					msg = messages[0]
				}

				spinner := spinnerChars[spinnerIndex%len(spinnerChars)]
				spinnerIndex++

				description := fmt.Sprintf("%s %s", spinner, msg)
				bar.Describe(description)
				_ = bar.Add(0) // Trigger redraw
			}
		}
	}()

	// Return a stop function
	return func() {
		spinnerCancel()
		<-cleanupDone
	}
}

func aiCmd(ctx context.Context, client *cloudquery_api.ClientWithResponses, teamName string, resumeConversation bool) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	errCh := make(chan error)

	go func() {
		errCh <- aiCmdInner(ctx, client, teamName, resumeConversation)
	}()

	select {
	case err := <-errCh:
		// Regardless of the error, the conversation ends here
		api.EndConversation(context.Background(), client, teamName)
		if err == nil {
			fmt.Println("Goodbye! 👋")
		}
		return err
	case <-ctx.Done():
		// End the conversation when context is cancelled (e.g., Ctrl+C)
		api.EndConversation(context.Background(), client, teamName)
		fmt.Println("\nGoodbye! 👋")
		return nil
	}
}

func aiCmdInner(ctx context.Context, client *cloudquery_api.ClientWithResponses, teamName string, resumeConversation bool) error {
	fmt.Println(`Your conversation with the AI may be recorded for quality assurance purposes. If you prefer not to use AI-assisted setup, run cloudquery init --disable-ai.`)
	fmt.Println()
	aiSuccess.Println("🤖 CloudQuery AI Assistant")
	fmt.Println("I'm here to help you set up CloudQuery syncs!")
	fmt.Println("Type 'exit' or 'quit' to end the conversation.")
	fmt.Println()
	if resumeConversation {
		fmt.Println("Your conversation has been resumed. You can now generate the config or ask questions about your sync.")
	} else {
		fmt.Println("What are you trying to build with CloudQuery?")
	}
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(aiInfo.Sprint("You: "))
		if !scanner.Scan() {
			break
		}
		if err := scanner.Err(); err != nil {
			return err
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}

		// Check for exit commands
		if strings.ToLower(userInput) == "exit" || strings.ToLower(userInput) == "quit" {
			break
		}

		// Show spinner while waiting for API response
		stop := startSpinner(ctx, generalMessages)

		response, err := api.Chat(ctx, client, teamName, &userInput, &[]api.FunctionCallOutput{})
		stop() // Stop the spinner

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

				// Show spinner while waiting for API response after creating spec file
				stop := startSpinner(ctx, specFileMessages)

				response, err = api.Chat(ctx, client, teamName, lo.ToPtr(""), &[]api.FunctionCallOutput{
					{
						Name:      "create_spec_file",
						CallID:    response.FunctionCallID,
						Arguments: response.FunctionCallArguments,
						Output:    "Spec file created",
					},
				})
				stop() // Stop the spinner

				if err != nil {
					return fmt.Errorf("failed to chat: %w", err)
				}
			case "create_sql_file":
				err := createSQLFile(response.FunctionCallArguments["filename_without_extension"].(string), response.FunctionCallArguments["content"].(string))
				if err != nil {
					return fmt.Errorf("failed to create SQL file: %w", err)
				}

				// Show spinner while waiting for API response after creating SQL file
				stop := startSpinner(ctx, specFileMessages)

				response, err = api.Chat(ctx, client, teamName, lo.ToPtr(""), &[]api.FunctionCallOutput{
					{
						Name:      "create_sql_file",
						CallID:    response.FunctionCallID,
						Arguments: response.FunctionCallArguments,
						Output:    "SQL file created",
					},
				})
				stop() // Stop the spinner

				if err != nil {
					return fmt.Errorf("failed to chat: %w", err)
				}
			case "cloudquery_test":
				// Show spinner while running the test
				stop := startSpinner(ctx, testMessages)

				testOutput := cloudqueryTest(response.FunctionCallArguments["filename_without_extension"].(string))
				stop() // Stop the spinner

				// Show spinner while waiting for API response after running test
				stop = startSpinner(ctx, testMessages)

				response, err = api.Chat(ctx, client, teamName, lo.ToPtr(""), &[]api.FunctionCallOutput{
					{
						Name:      "cloudquery_test",
						CallID:    response.FunctionCallID,
						Arguments: response.FunctionCallArguments,
						Output:    testOutput,
					},
				})
				stop() // Stop the spinner

				if err != nil {
					return fmt.Errorf("failed to chat: %w", err)
				}
			default:
				return fmt.Errorf("unsupported function call: %s", *response.FunctionCall)
			}
		}

		fmt.Printf("\n%s %s\n\n", aiBold.Sprint("AI:"), *response.Message)
	}

	return nil
}

func createSpecFile(filenameWithoutExtension, content string) error {
	return os.WriteFile(filenameWithoutExtension+".yaml", []byte(content), 0644)
}

func createSQLFile(filenameWithoutExtension, content string) error {
	return os.WriteFile(filenameWithoutExtension+".sql", []byte(content), 0644)
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
