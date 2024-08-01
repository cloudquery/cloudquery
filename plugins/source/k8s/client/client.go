package client

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
	v1 "k8s.io/api/core/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	"github.com/cloudquery/cloudquery/plugins/source/k8s/client/spec"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/rs/zerolog"

	// import all k8s auth options
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

type Client struct {
	logger zerolog.Logger

	// map context_name -> namespaces
	namespaces map[string][]v1.Namespace

	// map context_name -> Services struct
	clients map[string]kubernetes.Interface

	// map context_name -> API extensions
	apiExtensions map[string]apiextensionsclientset.Interface

	// map context_name -> Dynamic client
	dynamicClients map[string]dynamic.Interface

	spec     *spec.Spec
	contexts []string
	paths    map[string]struct{}

	Context   string
	Namespace string
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	if c.Namespace != "" {
		return fmt.Sprintf("context:%s:namespace%s", c.Context, c.Namespace)
	}
	return fmt.Sprintf("context:%s", c.Context)
}

func (c *Client) Client() kubernetes.Interface {
	return c.clients[c.Context]
}

func (c *Client) APIExtensions() apiextensionsclientset.Interface {
	return c.apiExtensions[c.Context]
}

func (c *Client) DynamicClient() dynamic.Interface {
	return c.dynamicClients[c.Context]
}

func (c *Client) Namespaces() []v1.Namespace {
	return c.namespaces[c.Context]
}

// Don't confuse `k8sContext` with `context.ctx`! k8s-context is a k8s-term that refers to a k8s cluster.
func (c *Client) WithContext(k8sContext string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("context", k8sContext).Logger()
	newC.Context = k8sContext
	return &newC
}

func (c *Client) WithNamespace(namespace string) *Client {
	newC := *c
	newC.logger = c.logger.With().Str("namespace", namespace).Logger()
	newC.Namespace = namespace
	return &newC
}

func Configure(ctx context.Context, logger zerolog.Logger, s spec.Spec) (schema.ClientMeta, error) {
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)
	rawKubeConfig, err := kubeConfig.RawConfig()
	if err != nil {
		return nil, err
	}

	contexts, err := loadContexts(&s, rawKubeConfig, logger)
	if err != nil {
		return nil, err
	}
	if len(contexts) == 0 {
		return nil, fmt.Errorf("could not find any context. Try to add context, https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-context-and-configuration")
	}

	c := Client{
		logger:         logger,
		clients:        make(map[string]kubernetes.Interface),
		namespaces:     make(map[string][]v1.Namespace),
		apiExtensions:  make(map[string]apiextensionsclientset.Interface),
		dynamicClients: make(map[string]dynamic.Interface),
		spec:           &s,
		contexts:       contexts,
		Context:        contexts[0],
		paths:          make(map[string]struct{}),
	}

	for _, ctxName := range contexts {
		logger.Info().Str("context", ctxName).Msg("creating k8s client for context")

		restConfig, err := buildRESTConfig(logger, rawKubeConfig, ctxName)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s RES config for context %q: %w", ctxName, err)
		}

		kClient, err := kubernetes.NewForConfig(restConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s client for context %q: %w", ctxName, err)
		}

		dClient, err := dynamic.NewForConfig(restConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s dynamic client for context %q: %w", ctxName, err)
		}

		apiExtClient, err := apiextensionsclientset.NewForConfig(restConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to build k8s API Extensions client for context %q: %w", ctxName, err)
		}

		c.paths, err = getAPIsMap(kClient)
		if err != nil {
			logger.Warn().Err(err).Msg("Failed to get OpenAPI schema. It might be not supported in the current version of Kubernetes. OpenAPI has been supported since Kubernetes 1.4")
		}

		namespaces, err := discoverNamespaces(ctx, kClient)
		if err != nil {
			return nil, fmt.Errorf("failed to discover namespaces for context %q: %w", ctxName, err)
		}

		c.clients[ctxName] = kClient
		c.namespaces[ctxName] = namespaces
		c.apiExtensions[ctxName] = apiExtClient
		c.dynamicClients[ctxName] = dClient
	}

	return &c, nil
}

func loadContexts(s *spec.Spec, rawCfg api.Config, logger zerolog.Logger) ([]string, error) {
	if len(s.Contexts) == 0 {
		logger.Debug().Str("context", rawCfg.CurrentContext).Msg("no context set in configuration using current default defined context")
		return []string{rawCfg.CurrentContext}, nil
	}

	if len(s.Contexts) == 1 && s.Contexts[0] == "*" {
		logger.Debug().Msg("loading all available configuration")
		return maps.Keys(rawCfg.Contexts), nil
	}

	// verify contexts
	for _, cName := range s.Contexts {
		if _, ok := rawCfg.Contexts[cName]; !ok {
			return nil, fmt.Errorf("context %q doesn't exist in kube configuration", cName)
		}
	}
	return s.Contexts, nil
}

func discoverNamespaces(ctx context.Context, client kubernetes.Interface) ([]v1.Namespace, error) {
	cl := client.CoreV1().Namespaces()

	opts := metav1.ListOptions{}
	namespaces := make([]v1.Namespace, 0)
	for {
		result, err := cl.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		namespaces = append(namespaces, result.Items...)
		if result.GetContinue() == "" {
			break
		}
		opts.Continue = result.GetContinue()
	}
	return namespaces, nil
}

// buildRESTConfig creates a k8s REST client config from the given config and context name.
func buildRESTConfig(logger zerolog.Logger, kubeConfig api.Config, ctx string) (*rest.Config, error) {
	override := &clientcmd.ConfigOverrides{CurrentContext: ctx}
	clientConfig := clientcmd.NewNonInteractiveClientConfig(
		kubeConfig,
		override.CurrentContext,
		override,
		&clientcmd.ClientConfigLoadingRules{},
	)

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		logger.Warn().Err(err).Msg("Failed to create k8s client, fallback to use the in-cluster config")
		restConfig, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	return restConfig, nil
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
