package console

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/cloudquery/cloudquery/internal/getter"

	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func FilterPolicies(policyPath string, policies policy.Policies) (policy.Policies, error) {
	if policyPath == "" && len(policies) == 0 {
		return nil, errors.New("no policies defined in configuration")
	}
	policyName, subPath := getter.ParseSourceSubPolicy(policyPath)
	// run them all
	if policyName == "" {
		return policies, nil
	}
	// select policies to run
	for _, p := range policies {
		// request to run only specific policy
		if policyName == p.Name {
			return policy.Policies{p}, nil
		}
	}
	// run hub detector
	p, found, err := policy.DetectPolicy(policyName, subPath)
	if err != nil {
		return nil, err
	}
	if found {
		return policy.Policies{p}, nil
	}
	return nil, fmt.Errorf("no policy with name %s found in configuration or remote. Available in config: %s", policyName, policies.All())
}

func buildDescribePolicyTable(t ui.Table, pp policy.Policies, policyRootPath string) {
	for _, p := range pp {
		policyPath := buildPolicyPath(policyRootPath, p.Name)
		t.Append(policyPath, p.Title)
		buildDescribePolicyTable(t, p.Policies, policyPath)
	}
}

// buildPolicyPath separates policy root path from in policy path with `//`
func buildPolicyPath(rootPath, name string) string {
	policyPath := fmt.Sprintf("%s//%s", rootPath, strings.ToLower(name))
	if strings.Contains(rootPath, "/") {
		policyPath = fmt.Sprintf("%s/%s", rootPath, strings.ToLower(name))
	}
	if rootPath == "" {
		policyPath = strings.ToLower(name)
	}
	return policyPath
}

func getNestedPolicyExample(p *policy.Policy, policyPath string) string {
	if len(p.Policies) > 0 {
		return getNestedPolicyExample(p.Policies[0], path.Join(policyPath, strings.ToLower(p.Name)))
	}
	return policyPath
}

func printPolicyResponse(results []*policy.ExecutionResult) {
	if len(results) == 0 {
		return
	}
	for _, execResult := range results {
		ui.ColorizedOutput(ui.ColorUnderline, "%s %s Results:\n\n", emojiStatus[ui.StatusInfo], execResult.PolicyName)

		if !execResult.Passed {
			if execResult.Error != "" {
				ui.ColorizedOutput(ui.ColorHeader, ui.ColorErrorBold.Sprintf("%s Policy failed to run\nError: %s\n\n", emojiStatus[ui.StatusError], execResult.Error))
			} else {
				ui.ColorizedOutput(ui.ColorHeader, ui.ColorErrorBold.Sprintf("%s Policy finished with warnings\n\n", emojiStatus[ui.StatusWarn]))
			}
		}
		fmtString := defineResultColumnWidths(execResult.Results)
		for _, res := range execResult.Results {
			switch {
			case res.Passed:
				ui.ColorizedOutput(ui.ColorInfo, fmtString, emojiStatus[ui.StatusOK]+" ", res.Name, res.Description, color.GreenString("passed"))
				ui.ColorizedOutput(ui.ColorInfo, "\n")
			case res.Type == policy.ManualQuery:
				ui.ColorizedOutput(ui.ColorInfo, fmtString, emojiStatus[ui.StatusWarn], res.Name, res.Description, color.YellowString("manual"))
				ui.ColorizedOutput(ui.ColorInfo, "\n")
				outputTable := createOutputTable(res)
				for _, row := range strings.Split(outputTable, "\n") {
					ui.ColorizedOutput(ui.ColorInfo, "\t\t  %-10s \n", row)
				}

			default:
				ui.ColorizedOutput(ui.ColorInfo, fmtString, emojiStatus[ui.StatusError], res.Name, res.Description, color.RedString("failed"))
				ui.ColorizedOutput(ui.ColorWarning, "\n")
				queryOutput := findOutput(res.Columns, res.Data)
				if len(queryOutput) > 0 {
					for _, output := range queryOutput {
						ui.ColorizedOutput(ui.ColorInfo, "\t\t%s  %-10s \n\n", emojiStatus[ui.StatusError], output)
					}

				}

			}
			ui.ColorizedOutput(ui.ColorWarning, "\n")
		}
	}
}

func findOutput(columnNames []string, data [][]interface{}) []string {
	outputKeys := []string{"id", "identifier", "resource_identifier", "uid", "uuid", "arn"}
	outputKey := ""
	outputResources := make([]string, 0)
	for _, key := range outputKeys {
		for _, column := range columnNames {
			if key == column {
				outputKey = key
			}
		}
		if outputKey != "" {
			break
		}
	}
	if outputKey == "" {
		return []string{}
	}
	for index, column := range columnNames {
		if column != outputKey {
			continue
		}
		for _, row := range data {
			outputResources = append(outputResources, fmt.Sprintf("%v", row[index]))
		}
	}
	return outputResources
}

func createOutputTable(res *policy.QueryResult) string {
	data := make([][]string, 0)
	for i := range res.Data {
		rowData := make([]string, len(res.Data[i]))
		for j, value := range res.Data[i] {
			rowData[j] = fmt.Sprintf("%v", value)
		}
		data = append(data, rowData)
	}
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader(res.Columns)
	table.SetRowLine(true)
	table.SetAutoFormatHeaders(false)
	table.AppendBulk(data)
	table.Render()
	return tableString.String()
}

func defineResultColumnWidths(execResult []*policy.QueryResult) string {
	maxNameLength := 0
	maxDescrLength := 0
	for _, res := range execResult {
		if len(res.Name) > maxNameLength {
			maxNameLength = len(res.Name) + 1
		}
		if len(res.Description) > maxDescrLength {
			maxDescrLength = len(res.Description) + 1
		}
	}
	return fmt.Sprintf("\t%%s  %%-%ds %%-%ds %%%ds", maxNameLength, maxDescrLength, 10)
}
