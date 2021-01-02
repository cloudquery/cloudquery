package k8s

import (
	"fmt"
	"github.com/cloudquery/cloudquery/database"
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/cloudquery/cloudquery/providers/provider"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
)

type Provider struct {
	db              *database.Database
	config          Config
	resourceClients map[string]common.ClientInterface
	log             *zap.Logger
	client          *kubernetes.Clientset
	clusterName     string
}

type Config struct {
	KubeConfig string `mapstructure:"domain"`
	Resources  []struct {
		Name  string
		Other map[string]interface{} `mapstructure:",remain"`
	}
}

var tablesArr = [][]interface{}{
	serviceTables,
	podTables,
}

func NewProvider(db *database.Database, log *zap.Logger) (provider.Interface, error) {
	p := Provider{
		db:              db,
		resourceClients: map[string]common.ClientInterface{},
		log:             log,
	}

	for _, tables := range tablesArr {
		err := db.AutoMigrate(tables...)
		if err != nil {
			return nil, err
		}
	}
	return &p, nil
}

func (p *Provider) Run(config interface{}) error {
	err := mapstructure.Decode(config, &p.config)
	if err != nil {
		return err
	}
	if len(p.config.Resources) == 0 {
		p.log.Warn("please specify at least 1 resource in config.yml. see: https://docs.cloudquery.io/k8s/reference")
		return nil
	}

	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//
	//kubeConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	return err
	//}

	clientCfg, err := clientcmd.NewDefaultClientConfigLoadingRules().Load()
	if err != nil {
		return err
	}
	kubeConfig := clientcmd.NewDefaultClientConfig(*clientCfg, &clientcmd.ConfigOverrides{})
	restConfig, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil
	}
	for k, _ := range clientCfg.Clusters {
		p.clusterName = k
		break
	}

	p.client, err = kubernetes.NewForConfig(restConfig)

	if err != nil {
		return err
	}

	for _, resource := range p.config.Resources {
		var err error
		switch resource.Name {
		case "services":
			err = p.services(resource.Other)
		case "pods":
			err = p.pods(resource.Other)
		default:
			return fmt.Errorf("unsupported resource %s", resource)
		}
		if err != nil {
			return err
		}

	}

	return nil
}
