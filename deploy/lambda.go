package deploy

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudquery/cloudquery/pkg/core"
	"github.com/cloudquery/cloudquery/pkg/plugin"
	"github.com/cloudquery/cloudquery/pkg/plugin/registry"

	"github.com/cloudquery/cloudquery/pkg/config"
	"github.com/cloudquery/cloudquery/pkg/core/database"
	"github.com/cloudquery/cloudquery/pkg/policy"
	"github.com/spf13/viper"
)

type Request struct {
	TaskName string `json:"taskName"`
	HCL      string `json:"hcl"`
}

func LambdaHandler(ctx context.Context, req Request) (string, error) {
	return TaskExecutor(ctx, req)
}

func TaskExecutor(ctx context.Context, req Request) (string, error) {
	dsn := os.Getenv("CQ_DSN")
	dataDir, present := os.LookupEnv("CQ_DATA_DIR")
	if !present {
		dataDir = ".cq"
	}
	viper.Set("data-dir", dataDir)
	pluginDir, present := os.LookupEnv("CQ_PLUGIN_DIR")
	if !present {
		pluginDir = "."
	}
	viper.Set("plugin-dir", pluginDir)
	policyDir, present := os.LookupEnv("CQ_POLICY_DIR")
	if !present {
		policyDir = "."
	}
	viper.Set("policy-dir", policyDir)

	cfg, diags := config.NewParser(
		config.WithEnvironmentVariables(config.EnvVarPrefix, os.Environ()),
	).LoadConfigFromSource("config.hcl", []byte(req.HCL))
	if diags != nil {
		return "", fmt.Errorf("bad configuration: %s", diags)
	}
	// Override dsn env if set
	if dsn != "" {
		cfg.CloudQuery.Connection.DSN = dsn
	}

	completedMsg := fmt.Sprintf("Completed task %s", req.TaskName)
	switch req.TaskName {
	case "fetch":
		return completedMsg, Fetch(ctx, cfg)
	case "policy":
		return completedMsg, Policy(ctx, cfg)
	default:
		return fmt.Sprintf("Unknown task: %s", req.TaskName), fmt.Errorf("unknown task: %s", req.TaskName)
	}
}

// Fetch fetches resources from a cloud provider and saves them in the configured database
func Fetch(ctx context.Context, cfg *config.Config) error {

	pm, err := plugin.NewManager(registry.NewRegistryHub(registry.CloudQueryRegistryURL, registry.WithPluginDirectory(cfg.CloudQuery.PluginDirectory)))
	if err != nil {
		return err
	}
	defer pm.Shutdown()

	_, dialect, err := database.GetExecutor(cfg.CloudQuery.Connection.DSN, cfg.CloudQuery.History)
	if err != nil {
		return err
	}
	storage := database.NewStorage(cfg.CloudQuery.Connection.DSN, dialect)

	providers := make([]core.ProviderInfo, len(cfg.Providers))
	for i, p := range cfg.Providers {
		rp, _ := cfg.CloudQuery.GetRequiredProvider(p.Name)
		src, _, _ := core.ParseProviderSource(rp)

		if _, err := core.Download(ctx, pm, &core.DownloadOptions{
			Providers: []registry.Provider{
				{Name: p.Name, Version: rp.Version, Source: src},
			},
			NoVerify: false,
		}); err != nil {
			return err
		}
		providers[i] = core.ProviderInfo{
			Provider: registry.Provider{Name: p.Name, Version: rp.Version, Source: src},
			Config:   p,
		}
	}
	_, diags := core.Fetch(ctx, storage, pm, &core.FetchOptions{
		ProvidersInfo: providers,
		History:       cfg.CloudQuery.History,
	})
	if diags.HasErrors() {
		return fmt.Errorf("failed to fetch, check logs for more details")
	}
	return nil
}

// Policy Runs a policy SQL statement and returns results
func Policy(ctx context.Context, cfg *config.Config) error {
	outputPath := "/tmp/"
	storage := database.NewStorage(cfg.CloudQuery.Connection.DSN, nil)
	_, err := policy.Run(ctx, storage, &policy.RunRequest{
		Policies:  cfg.Policies,
		Directory: cfg.CloudQuery.PolicyDirectory,
		OutputDir: outputPath,
	})
	if err != nil {
		return fmt.Errorf("error running query: %s", err)
	}
	return nil
}
