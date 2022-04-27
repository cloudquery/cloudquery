package client

import (
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"

	// import all k8s auth options
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

type Client struct {
	Log      hclog.Logger
	services map[string]Services
	kConfig  api.Config
	config   *Config
	contexts []string

	Context string
}

func (c *Client) Logger() hclog.Logger {
	return c.Log
}

func (c *Client) Services() Services {
	return c.services[c.Context]
}

func (c Client) WithContext(context string) *Client {
	return &Client{
		Log:      c.Log.With("context", context),
		services: c.services,
		kConfig:  c.kConfig,
		config:   c.config,
		Context:  context,
	}
}

func (c *Client) SetServices(s map[string]Services) {
	c.services = s
	contexts := make([]string, 0, len(s))
	for k := range s {
		contexts = append(contexts, k)
	}
	c.contexts = contexts
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	kCfg, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, diag.FromError(err, diag.USER)
	}

	cfg := config.(*Config)

	var contexts []string
	switch len(cfg.Contexts) {
	case 0:
		logger.Debug("no context set in configuration using current default defined context", "context", kCfg.CurrentContext)
		contexts = []string{kCfg.CurrentContext}
	case 1:
		if cfg.Contexts[0] != "*" {
			logger.Debug("loading all available configuration")
			for cName := range kCfg.Contexts {
				contexts = append(contexts, cName)
			}
		} else {
			contexts = []string{cfg.Contexts[0]}
		}
	default:
		for _, cName := range cfg.Contexts {
			if _, ok := kCfg.Contexts[cName]; !ok {
				return nil, diag.FromError(fmt.Errorf("context %q doesn't exist in kube configuration", cName), diag.USER)
			}
			contexts = append(contexts, cName)
		}
	}
	c := Client{
		Log:      logger,
		services: make(map[string]Services),
		kConfig:  kCfg,
		config:   cfg,
		contexts: contexts,
		Context:  contexts[0],
	}

	for _, ctxName := range contexts {
		logger.Info("creating k8s client for context", "context", ctxName)
		kClient, err := buildKubeClient(kCfg, kCfg.CurrentContext)
		if err != nil {
			return nil, diag.FromError(fmt.Errorf("failed to build k8s client for context %q: %w", kCfg.CurrentContext, err), diag.INTERNAL)
		}
		c.services[kCfg.CurrentContext] = initServices(kClient)
	}

	return &c, nil
}

// buildKubeClient creates a k8s client from the given config and context name.
func buildKubeClient(kubeConfig api.Config, ctx string) (*kubernetes.Clientset, error) {
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(
		kubeConfig,
		override.CurrentContext,
		override,
		&clientcmd.ClientConfigLoadingRules{},
	)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(restConfig)
}

func initServices(client *kubernetes.Clientset) Services {
	return Services{
		Client:          client,
		CronJobs:        client.BatchV1().CronJobs(""),
		DaemonSets:      client.AppsV1().DaemonSets(""),
		Deployments:     client.AppsV1().Deployments(""),
		Endpoints:       client.CoreV1().Endpoints(""),
		Jobs:            client.BatchV1().Jobs(""),
		LimitRanges:     client.CoreV1().LimitRanges(""),
		Namespaces:      client.CoreV1().Namespaces(),
		NetworkPolicies: client.NetworkingV1().NetworkPolicies(""),
		Nodes:           client.CoreV1().Nodes(),
		Pods:            client.CoreV1().Pods(""),
		ReplicaSets:     client.AppsV1().ReplicaSets(""),
		ResourceQuotas:  client.CoreV1().ResourceQuotas(""),
		RoleBindings:    client.RbacV1().RoleBindings(""),
		Roles:           client.RbacV1().Roles(""),
		ServiceAccounts: client.CoreV1().ServiceAccounts(""),
		Services:        client.CoreV1().Services(""),
		StatefulSets:    client.AppsV1().StatefulSets(""),
	}
}
