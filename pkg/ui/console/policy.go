package console

import (
	"errors"
	"fmt"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/cloudquery/cloudquery/internal/getter"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/cloudquery/cloudquery/pkg/ui"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
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
		if policyName != p.Name {
			continue
		}
		if subPath != "" && p.SubPolicy() == "" {
			return policy.Policies{
				&policy.Policy{
					Name:   p.Name,
					Source: p.Source + "//" + subPath,
				},
			}, nil
		}
		return policy.Policies{p}, nil
	}
	// run hub detector. We got here if we couldn't find the policy specified by the command argument in the configuration
	p, found, err := policy.DetectPolicy(policyName, subPath)
	if err != nil {
		return nil, err
	}
	if found {
		return policy.Policies{p}, nil
	}

	configPolicies := policies.All()
	if len(configPolicies) == 0 {
		return nil, diag.FromError(fmt.Errorf("no valid policy with name %q found. If using a local policy directory, ensure the path is correct and the directory exists", policyName), diag.USER)
	}

	return nil, diag.FromError(fmt.Errorf("no valid policy with name %q found in configuration or remote. Available in config: %s", policyName, configPolicies), diag.USER)

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
				ui.ColorizedOutput(ui.ColorHeader, ui.ColorErrorBold.Sprintf("%s Policy finished with violations\n\n", emojiStatus[ui.StatusWarn]))
			}
		}
		for _, res := range execResult.Results {
			switch res.Type {
			case policy.ManualQuery:
				ui.ColorizedOutput(ui.ColorInfo, "%s: Policy %s - %s\n\n", color.YellowString("Manual"), res.Name, res.Description)
				ui.ColorizedOutput(ui.ColorInfo, "\n")
			case policy.AutomaticQuery:
				if res.Passed {
					ui.ColorizedOutput(ui.ColorInfo, "%s: Policy %s - %s\n\n", color.GreenString("Passed"), res.Name, res.Description)
				} else {
					ui.ColorizedOutput(ui.ColorInfo, "%s: Policy %s - %s\n\n", color.RedString("Failed"), res.Name, res.Description)
				}
			}
			if len(res.Rows) > 0 {
				createOutputTable(res)
				ui.ColorizedOutput(ui.ColorInfo, "\n\n")
			}

		}
	}
}

func createOutputTable(res *policy.QueryResult) {
	table := tablewriter.NewWriter(os.Stdout)

	if len(res.Rows[0].Identifiers) == 0 {
		table.SetHeader(append([]string{"status", "reason"}, res.QueryColumns...))
		table.SetFooter(append(makeStringArrayOfLength(len(res.QueryColumns)), "Total:", strconv.Itoa(len(res.Rows))))
	} else {
		table.SetHeader(res.Columns)
		table.SetFooter(append(makeStringArrayOfLength(len(res.Columns)-2), "Total:", strconv.Itoa(len(res.Rows))))
	}
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetAutoWrapText(true)
	table.SetReflowDuringAutoWrap(true)
	table.SetRowLine(false)
	table.SetBorder(false)
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)
	sort.Sort(res.Rows)
	for _, row := range res.Rows {
		data := make([]string, 0)
		data = append(data, color.HiRedString(row.Status))
		if len(row.Identifiers) > 0 {
			data = append(data, cast.ToStringSlice(row.Identifiers)...)
		}
		data = append(data, row.Reason)
		data = append(data, cast.ToStringSlice(row.AdditionalData)...)
		table.Append(data)

	}
	table.Render()
}

func makeStringArrayOfLength(length int) []string {
	s := make([]string, length)
	for i := 0; i < length; i++ {
		s[i] = ""
	}
	return s
}
