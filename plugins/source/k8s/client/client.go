package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"k8s.io/client-go/kubernetes"

	// import all k8s auth options
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	logger zerolog.Logger
	// map context_name -> Services struct
	clients  map[string]kubernetes.Interface
	spec     *Spec
	contexts []string
	paths    map[string]struct{}

	Context string
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return c.Context
}

func (c *Client) Client() kubernetes.Interface {
	return c.clients[c.Context]
}

// Don't confuse `k8sContext` with `context.ctx`! k8s-context is a k8s-term that refers to a k8s cluster.
func (c Client) WithContext(k8sContext string) *Client {
	return &Client{
		logger:   c.logger.With().Str("context", k8sContext).Logger(),
		clients:  c.clients,
		spec:     c.spec,
		contexts: c.contexts,
		paths:    c.paths,
		Context:  k8sContext,
	}
}

func Configure(ctx context.Context, logger zerolog.Logger, s specs.Source, _ source.Options) (schema.ClientMeta, error) {
	var k8sSpec Spec

	if err := s.UnmarshalSpec(&k8sSpec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal k8s spec: %w", err)
	}

	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	rawKubeConfig, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	var contexts []string
	switch len(k8sSpec.Contexts) {
	case 0:
		logger.Debug().Str("context", rawKubeConfig.CurrentContext).Msg("no context set in configuration using current default defined context")
		contexts = []string{rawKubeConfig.CurrentContext}
	case 1:
		if k8sSpec.Contexts[0] == "*" {
			logger.Debug().Msg("loading all available configuration")
			for cName := range rawKubeConfig.Contexts {
				contexts = append(contexts, cName)
			}
		} else {
			if _, ok := rawKubeConfig.Contexts[k8sSpec.Contexts[0]]; !ok {
				return nil, fmt.Errorf("context %q doesn't exist in kube configuration", k8sSpec.Contexts[0])
			}
			contexts = []string{k8sSpec.Contexts[0]}
		}
	default:
		for _, cName := range k8sSpec.Contexts {
			if _, ok := rawKubeConfig.Contexts[cName]; !ok {
				return nil, fmt.Errorf("context %q doesn't exist in kube configuration", cName)
			}
			contexts = append(contexts, cName)
		}
	}

	if len(contexts) == 0 {
		return nil, fmt.Errorf("could not find any context. Try to add context, https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-context-and-configuration")
	}

	c := Client{
		logger:   logger,
		clients:  make(map[string]kubernetes.Interface),
		spec:     &k8sSpec,
		contexts: contexts,
		Context:  contexts[0],
		paths:    make(map[string]struct{}),
	}

	for _, ctxName := range contexts {
		logger.Info().Str("context", ctxName).Msg("creating k8s client for context")
		kClient, err := buildKubeClient(rawKubeConfig, ctxName)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s client for context %q: %w", ctxName, err)
		}
		c.paths, err = getAPIsMap(kClient)
		if err != nil {
			logger.Warn().Err(err).Msg("Failed to get OpenAPI schema. It might be not supported in the current version of Kubernetes. OpenAPI has been supported since Kubernetes 1.4")
		}
		c.clients[ctxName] = kClient
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

func getAPIsMap(client *kubernetes.Clientset) (map[string]struct{}, error) {
	doc, err := client.OpenAPISchema()
	if err != nil {
		return nil, err
	}
	paths := make(map[string]struct{})
	for _, p := range doc.Paths.Path {
		path := p.Name
		if strings.HasPrefix(path, "/apis/") {
			paths[path] = struct{}{}
		}
	}
	return paths, nil
}
