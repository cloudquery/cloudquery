package client

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	Log      hclog.Logger
	Services Services
}

func (c *Client) Logger() hclog.Logger {
	return c.Log
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	restConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	client, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	return &Client{
		Log: logger,
		Services: Services{
			Nodes: client.CoreV1().Nodes(),
		},
	}, nil
}
