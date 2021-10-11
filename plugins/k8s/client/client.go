package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	Log      hclog.Logger
	Services Services

	kubeConfig     api.Config
	config         *Config
	currentContext string
}

func (c *Client) Logger() hclog.Logger {
	return c.Log
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	cfg, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}
	return &Client{
		Log:        logger,
		kubeConfig: cfg,
		config:     config.(*Config),
	}, nil
}

func initServices(client *kubernetes.Clientset) Services {
	return Services{
		Nodes:    client.CoreV1().Nodes(),
		Pods:     client.CoreV1().Pods(""),
		Services: client.CoreV1().Services(""),
	}
}
