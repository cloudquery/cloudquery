package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

type logEntry struct {
	invocationID string
	key          string
	ts           string
	typ          string
}

type tableData struct {
	times []time.Time
	types []string
}

const (
	analyzeShort   = "Analyze CloudQuery log files to detect stalled tables and calculate execution times"
	analyzeExample = `# Analyze a CloudQuery log file to find stalled tables
cloudquery analyze --file path/to/cloudquery.log

# Analyze data for a specific invocation ID only
cloudquery analyze --file path/to/cloudquery.log --invocation-id abc123`
)

func NewCmdAnalyze() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "analyze",
		Short:   analyzeShort,
		Long:    analyzeShort,
		Example: analyzeExample,
		RunE:    analyze,
	}
	cmd.Flags().StringP("file", "f", "", "Path to the CloudQuery log file")
	cmd.Flags().StringP("invocation-id", "i", "", "Only analyze data for a specific invocation ID")
	_ = cmd.MarkFlagRequired("file")

	return cmd
}

func analyze(cmd *cobra.Command, args []string) error {
	logFile, err := cmd.Flags().GetString("file")
	if err != nil {
		return err
	}

	invocationID, err := cmd.Flags().GetString("invocation-id")
	if err != nil {
		return err
	}

	return analyzeLogFile(logFile, invocationID)
}

func extractInvocationID(line string) string {
	invocationPattern := regexp.MustCompile(`invocation_id=([^\s]+)`)
	if matches := invocationPattern.FindStringSubmatch(line); matches != nil {
		return matches[1]
	}
	return ""
}

func analyzeLogFile(filePath string, filterInvocationID string) error {
	// Define the regular expression patterns
	patternEnd := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z)\s+(INF|ERR|WRN)\s+(.*)\s+client=(.*)\s+(.*)\s+errors=(\d+)?\s+(.*)+\s+module=([a-zA-Z-]+)?\s+resources=(\d+)?\s+table=(\w+)?`)
	patternStart := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z)\s+(INF|ERR|WRN)\s+(.*)\s+client=(.*)\s+(.*)\s+module=([a-zA-Z-]+)?\s+table=(\w+)?`)

	fmt.Printf("Analyzing log file: %s\n", filePath)

	// Open the log file
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	// Process log entries
	allData := []logEntry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Extract invocation ID for filtering
		extractedInvocationID := extractInvocationID(line)

		// Skip if we're filtering by invocation ID and this line doesn't match
		if filterInvocationID != "" && extractedInvocationID != filterInvocationID {
			continue
		}

		// Try to match end pattern
		if matches := patternEnd.FindStringSubmatch(line); matches != nil {
			allData = append(allData, logEntry{
				invocationID: extractedInvocationID,
				key:          matches[4] + matches[10],
				ts:           matches[1],
				typ:          "end",
			})
			continue
		}

		// Try to match start pattern
		if matches := patternStart.FindStringSubmatch(line); matches != nil {
			allData = append(allData, logEntry{
				invocationID: extractedInvocationID,
				key:          matches[4] + matches[7],
				ts:           matches[1],
				typ:          "start",
			})
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading log file: %w", err)
	}

	// Normalize data
	normalizedData := make(map[string]map[string]*tableData)
	for _, data := range allData {
		clientPair := data.key
		ts := data.ts

		if _, exists := normalizedData[data.invocationID]; !exists {
			normalizedData[data.invocationID] = make(map[string]*tableData)
		}
		if _, exists := normalizedData[data.invocationID][data.key]; !exists {
			normalizedData[data.invocationID][clientPair] = &tableData{
				times: []time.Time{},
				types: []string{},
			}
		}

		// Parse timestamp
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z", ts)
		if err != nil {
			fmt.Printf("Warning: could not parse timestamp %s: %v\n", ts, err)
			continue
		}

		normalizedData[data.invocationID][clientPair].times = append(normalizedData[data.invocationID][clientPair].times, parsedTime)
		normalizedData[data.invocationID][clientPair].types = append(normalizedData[data.invocationID][clientPair].types, data.typ)
	}

	// Sort times for each entry
	for invocationID := range normalizedData {
		for _, data := range normalizedData[invocationID] {
			sort.Slice(data.times, func(i, j int) bool {
				return data.times[i].Before(data.times[j])
			})
		}
	}

	// Calculate time differences and find tables that never completed
	type timeDiffKeyPair struct {
		timeDiff int
		key      string
	}

	timeDiffKeyPairs := make(map[string][]timeDiffKeyPair)

	for invocationID := range normalizedData {
		for key, data := range normalizedData[invocationID] {
			if len(data.times) > 1 {
				// Calculate time difference in minutes
				timeDiff := int(data.times[len(data.times)-1].Sub(data.times[0]).Seconds() / 60)

				if _, exists := timeDiffKeyPairs[invocationID]; !exists {
					timeDiffKeyPairs[invocationID] = []timeDiffKeyPair{}
				}
				timeDiffKeyPairs[invocationID] = append(timeDiffKeyPairs[invocationID], timeDiffKeyPair{
					timeDiff: timeDiff,
					key:      key,
				})

			} else if len(data.types) == 1 && data.types[0] == "start" {
				fmt.Printf("Table never completed: %s for invocation %s\n", key, invocationID)
			}
		}
	}

	for invocationID := range timeDiffKeyPairs {
		// Sort time differences in descending order
		sort.Slice(timeDiffKeyPairs[invocationID], func(i, j int) bool {
			return timeDiffKeyPairs[invocationID][i].timeDiff > timeDiffKeyPairs[invocationID][j].timeDiff
		})
	}

	// Display results for all invocation IDs
	for invocationID := range timeDiffKeyPairs {
		fmt.Printf("Invocation ID: %s\n", invocationID)
		// Print time differences
		for _, pair := range timeDiffKeyPairs[invocationID] {
			fmt.Printf("   %d minutes - %s\n", pair.timeDiff, pair.key)
		}
	}

	return nil
}
