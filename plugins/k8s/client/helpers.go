package client

import (
	"context"
	"errors"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

const (
	ContextFieldName = "k8s_config_context"
	ContextFieldDesc = "Name of the context from k8s configuration."
)

var CommonContextField = schema.Column{
	Name:        ContextFieldName,
	Description: ContextFieldDesc,
	Type:        schema.TypeString,
	Resolver:    ResolveContext,
}

// ContextMultiplex returns a list of clients for each context from the cq config
func ContextMultiplex(meta schema.ClientMeta) []schema.ClientMeta {
	client := meta.(*Client)
	if len(client.config.Contexts) == 0 {
		c, err := buildClient(client, client.kubeConfig.CurrentContext)
		if err != nil {
			client.Log.Warn("failed to create client with default context", "error", err)
			return nil
		}
		return []schema.ClientMeta{c}
	}
	clients := make([]schema.ClientMeta, 0)
	for _, name := range client.config.Contexts {
		client.Log.Error("preparing new context")
		if err := validateContext(name, client.kubeConfig.Contexts); err != nil {
			client.Log.Warn("invalid context", "context", name, "error", err)
			continue
		}

		c, err := buildClient(client, name)
		if err != nil {
			client.Log.Warn("failed to create client with context", "context", name, "error", err)
			continue
		}
		clients = append(clients, c)
	}
	return clients
}

// buildClient creates a client for a given k8s context
func buildClient(from *Client, context string) (*Client, error) {
	clientSet, err := buildKubeClient(from.kubeConfig, context)
	if err != nil {
		return nil, err
	}
	c := *from
	c.Services = initServices(clientSet)
	c.currentContext = context
	c.Log = c.Log.With("context", context)
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

// validateContext checks that context name is not empty and it exists in the provided contexts map.
func validateContext(name string, contexts map[string]*api.Context) error {
	if name == "" {
		return errors.New("empty context name")
	}
	for v := range contexts {
		if v == name {
			return nil
		}
	}
	return errors.New("context not found")
}

// DeleteContextFilter returns a delete filter that cleans up the data belonging to the k8s context.
func DeleteContextFilter(meta schema.ClientMeta, _ *schema.Resource) []interface{} {
	client := meta.(*Client)
	return []interface{}{ContextFieldName, client.currentContext}
}

// ResolveContext is a resolver that fills the k8s context field.
func ResolveContext(_ context.Context, meta schema.ClientMeta, r *schema.Resource, c schema.Column) error {
	client := meta.(*Client)
	return r.Set(c.Name, client.currentContext)
}
