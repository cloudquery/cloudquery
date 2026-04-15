package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

// headingReplacer promotes cobra/doc's heading levels so each page has a
// proper H1. cobra/doc generates:
//
//	## <command>          (should be H1 — the page title)
//	### Synopsis          (should be H2)
//	### Options           (should be H2)
//	### SEE ALSO          (should be H2)
//	#### <flag-detail>    (should be H3)
//
// We promote deeper levels first (#### → ###, then ### → ##), then promote the
// first remaining ## to # so that only the command-name heading becomes H1.
var headingReplacer = strings.NewReplacer(
	"\n#### ", "\n### ",
	"\n### ", "\n## ",
)

const (
	docShort = "Generate CLI documentation markdown files"
)

// seeAlsoSections maps generated filename to the "## See Also" content to append.
var seeAlsoSections = map[string]string{
	"cloudquery.md": `## See Also

- [Getting Started](/cli/getting-started) - Install and run your first sync
- [Configuration Guide](/cli/core-concepts/configuration) - Configure source and destination integrations
`,
	"cloudquery_addon.md": `## See Also

- [Transformations](/cli/core-concepts/transformations) - Official dbt and SQL transformations
- [Dashboards & Visualizations](/cli/core-concepts/dashboards) - Grafana dashboards from the hub
`,
	"cloudquery_addon_download.md": `## See Also

- [Transformations](/cli/core-concepts/transformations) - Available transformations and policies
- [Dashboards & Visualizations](/cli/core-concepts/dashboards) - Grafana dashboards from the hub
`,
	"cloudquery_addon_publish.md": `## See Also

- [Publishing an Addon](/cli/advanced/publishing-an-addon-to-the-hub) - Full addon publishing guide
- [Transformations](/cli/core-concepts/transformations) - Official transformations and policies
`,
	"cloudquery_init.md": `## See Also

- [Getting Started](/cli/getting-started) - Full quickstart guide using the init command
- [Configuration Guide](/cli/core-concepts/configuration) - Understand the generated configuration files
`,
	"cloudquery_login.md": `## See Also

- [Generate API Key](/cli/managing-cloudquery/deployments/generate-api-key) - Create API keys for headless authentication
- [Getting Started](/cli/getting-started) - Install and run your first sync
`,
	"cloudquery_logout.md": `## See Also

- [Generate API Key](/cli/managing-cloudquery/deployments/generate-api-key) - Manage API keys for authentication
- [Security](/cli/managing-cloudquery/security) - CloudQuery security best practices
`,
	"cloudquery_migrate.md": `## See Also

- [Schema Migrations](/cli/managing-cloudquery/migrations) - How CloudQuery handles schema changes
- [Destination Integrations](/cli/integrations/destinations) - Configure migration modes
`,
	"cloudquery_plugin.md": `## See Also

- [Integration Concepts](/cli/core-concepts/integrations) - How integrations work
- [Managing Versions](/cli/advanced/managing-versions) - Integration versioning
`,
	"cloudquery_plugin_install.md": `## See Also

- [Managing Versions](/cli/advanced/managing-versions) - Understand version management
- [Source Integrations](/cli/integrations/sources) - Available source integrations
`,
	"cloudquery_plugin_publish.md": `## See Also

- [Publishing an Integration](/cli/integrations/creating-new-integration/publishing) - Full publishing guide
- [Creating a New Integration](/cli/integrations/creating-new-integration) - Build an integration first
`,
	"cloudquery_switch.md": `## See Also

- [Managing Versions](/cli/advanced/managing-versions) - Understand integration versioning
- [Source Integrations](/cli/integrations/sources) - Configure source integration versions
`,
	"cloudquery_sync.md": `## See Also

- [Syncs](/cli/core-concepts/syncs) - Understand full and incremental sync modes
- [Configuration Guide](/cli/core-concepts/configuration) - Set up sync configurations
- [Performance Tuning](/cli/advanced/performance-tuning) - Optimize sync performance
`,
	"cloudquery_tables.md": `## See Also

- [Source Integrations](/cli/integrations/sources) - Configure which tables to sync
- [Integration Concepts](/cli/core-concepts/integrations) - How source integrations define tables
`,
	"cloudquery_test-connection.md": `## See Also

- [Source Integrations](/cli/integrations/sources) - Configure source connections
- [Destination Integrations](/cli/integrations/destinations) - Configure destination connections
- [Troubleshooting](/cli/getting-support/troubleshooting) - Debug connection issues
`,
	"cloudquery_validate-config.md": `## See Also

- [Configuration Guide](/cli/core-concepts/configuration) - Configuration format and options
- [Environment Variables](/cli/managing-cloudquery/environment-variables) - Variable substitution in configuration files
`,
}

func newCmdDoc() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "doc [directory_path]",
		Short:  docShort,
		Args:   cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := doc.GenMarkdownTreeCustom(cmd.Parent(), args[0], filePrepender, linkHandler); err != nil {
				return err
			}
			if err := fixHeadingLevels(args[0]); err != nil {
				return err
			}
			return appendSeeAlsoSections(args[0])
		},
	}
	return cmd
}

// fixHeadingLevels corrects the heading hierarchy in all generated markdown files.
// cobra/doc emits ## for the command name and ### for sub-sections; this promotes
// them to # and ## so every page has a proper H1.
func fixHeadingLevels(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("reading dir %s: %w", dir, err)
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		fpath := filepath.Join(dir, e.Name())
		data, err := os.ReadFile(fpath)
		if err != nil {
			return fmt.Errorf("reading %s: %w", fpath, err)
		}
		// Promote #### → ### and ### → ## (via headingReplacer), then promote
		// the first ## → # so the command name becomes the page's H1.
		// Also normalize cobra's all-caps ## SEE ALSO to ## See Also for
		// consistency with the custom links appended by appendSeeAlsoSections.
		fixed := headingReplacer.Replace(string(data))
		fixed = strings.Replace(fixed, "\n## ", "\n# ", 1)
		fixed = strings.ReplaceAll(fixed, "\n## SEE ALSO\n", "\n## See Also\n")
		if err := os.WriteFile(fpath, []byte(fixed), 0644); err != nil {
			return fmt.Errorf("writing %s: %w", fpath, err)
		}
	}
	return nil
}

// appendSeeAlsoSections appends custom links into the "## See Also" section of
// each generated file that has entries defined. The cobra-generated markdown
// already contains a "## See Also" heading (normalized by fixHeadingLevels), so
// only the link lines are appended — not a new heading — to avoid a duplicate
// H2 section on every page.
func appendSeeAlsoSections(dir string) error {
	for filename, content := range seeAlsoSections {
		fpath := filepath.Join(dir, filename)
		f, err := os.OpenFile(fpath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return fmt.Errorf("opening %s: %w", fpath, err)
		}
		// Strip the "## See Also\n\n" heading from the content: cobra/doc has
		// already emitted a "## See Also" section (normalized by fixHeadingLevels)
		// and we only want to append the link lines underneath it.
		links := strings.TrimPrefix(content, "## See Also\n\n")
		_, writeErr := fmt.Fprintf(f, "%s", links)
		closeErr := f.Close()
		if writeErr != nil {
			return fmt.Errorf("writing to %s: %w", fpath, writeErr)
		}
		if closeErr != nil {
			return fmt.Errorf("closing %s: %w", fpath, closeErr)
		}
	}
	return nil
}

func linkHandler(s string) string {
	if strings.HasSuffix(s, ".md") {
		fileName := strings.TrimSuffix(s, ".md")
		fullPath := "/cli/cli-reference/" + fileName

		return fullPath
	}

	return s
}

func filePrepender(filename string) string {
	const fmTemplate = `---
title: "%s"
---
`
	name := filepath.Base(filename)
	base := strings.TrimSuffix(name, path.Ext(name))
	id := strings.TrimPrefix(base, "cloudquery_")
	return fmt.Sprintf(fmTemplate, id)
}
