module github.com/cloudquery/cq-provider-k8s

go 1.15

require (
	github.com/cloudquery/cloudquery v0.8.11
	github.com/gophercloud/gophercloud v0.15.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	go.uber.org/zap v1.16.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
)
