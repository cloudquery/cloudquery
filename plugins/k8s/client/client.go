package client

import (
	"fmt"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"

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

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	kCfg, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
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
				return nil, fmt.Errorf("context %s doesn't in kube configuration", cName)
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
		Context:  "",
	}

	for _, ctxName := range contexts {
		logger.Info("creating k8s client for context", "context", ctxName)
		kClient, err := buildKubeClient(kCfg, kCfg.CurrentContext)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s client for context %s: %w", kCfg.CurrentContext, err)
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
		Client:       client,
		Namespaces:   client.CoreV1().Namespaces(),
		Nodes:        client.CoreV1().Nodes(),
		Pods:         client.CoreV1().Pods(""),
		Services:     client.CoreV1().Services(""),
		Jobs:         client.BatchV1().Jobs(""),
		DaemonSets:   client.AppsV1().DaemonSets(""),
		StatefulSets: client.AppsV1().StatefulSets(""),
		ReplicaSets:  client.AppsV1().ReplicaSets(""),
		Roles:        client.RbacV1().Roles(""),
		RoleBindings: client.RbacV1().RoleBindings(""),
	}
}
