package main

import (
	"context"
	"github.com/mitchellh/mapstructure"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Service struct {
	ID          uint   `gorm:"primarykey"`
	ClusterName string `neo:"unique"`

	Name                         string                 `neo:"unique"`
	Namespace                    string                 `neo:"unique"`
	SpecPorts                    []*ServicePort         `gorm:"constraint:OnDelete:CASCADE;"`
	SpecSelector                 []*ServiceSpecSelector `gorm:"constraint:OnDelete:CASCADE;"`
	SpecClusterIP                string
	SpecType                     string
	SpecExternalIPs              []*ServiceSpecExternalIPs `gorm:"constraint:OnDelete:CASCADE;"`
	SpecSessionAffinity          string
	SpecLoadBalancerIP           string
	SpecLoadBalancerSourceRanges []*ServiceSpecLoadBalancerSourceRanges `gorm:"constraint:OnDelete:CASCADE;"`
	SpecExternalName             string
	SpecExternalTrafficPolicy    string
	SpecHealthCheckNodePort      int32
	SpecPublishNotReadyAddresses bool

	SpecSessionAffinityConfigClientIPTimeoutSeconds *int32
	SpecIPFamily                                    *v1.IPFamily
	SpecTopologyKeys                                []*ServiceSpecTopologyKeys `gorm:"constraint:OnDelete:CASCADE;"`

	StatusLoadBalancerIngress []*ServiceLoadBalancerIngress `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Service) TableName() string {
	return "k8s_services"
}

type ServicePort struct {
	ID          uint   `gorm:"primarykey"`
	ServiceID   uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`

	Name        string
	Protocol    string
	AppProtocol *string
	Port        int32
	TargetPort  string
	NodePort    int32
}

func (ServicePort) TableName() string {
	return "k8s_service_ports"
}

type ServiceSpecSelector struct {
	ID          uint   `gorm:"primarykey"`
	ServiceID   uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`

	Key   string
	Value string
}

func (ServiceSpecSelector) TableName() string {
	return "k8s_service_spec_selectors"
}

type ServiceSpecExternalIPs struct {
	ID          uint   `gorm:"primarykey"`
	ServiceID   uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`

	Value string
}

func (ServiceSpecExternalIPs) TableName() string {
	return "k8s_service_spec_external_ips"
}

type ServiceSpecLoadBalancerSourceRanges struct {
	ID          uint   `gorm:"primarykey"`
	ServiceID   uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`

	Value string
}

func (ServiceSpecLoadBalancerSourceRanges) TableName() string {
	return "k8s_service_spec_load_balancer_source_ranges"
}

type ServiceSpecTopologyKeys struct {
	ID          uint   `gorm:"primarykey"`
	ServiceID   uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`

	Value string
}

func (ServiceSpecTopologyKeys) TableName() string {
	return "k8s_service_spec_topology_keys"
}

type ServiceLoadBalancerIngress struct {
	ID          uint   `gorm:"primarykey"`
	ServiceID   uint   `neo:"ignore"`
	ClusterName string `gorm:"-"`

	IP       string
	Hostname string
}

func (ServiceLoadBalancerIngress) TableName() string {
	return "k8s_service_load_balancer_ingresses"
}

func (p *Provider) transformServices(values []v1.Service) []*Service {
	var tValues []*Service
	for _, value := range values {
		tValue := Service{
			ClusterName:                  p.clusterName,
			Name:                         value.Name,
			Namespace:                    value.Namespace,
			SpecPorts:                    p.transformServicePorts(value.Spec.Ports),
			SpecSelector:                 p.transformServiceSpecSelectors(value.Spec.Selector),
			SpecClusterIP:                value.Spec.ClusterIP,
			SpecType:                     string(value.Spec.Type),
			SpecExternalIPs:              p.transformServiceSpecExternalIPs(value.Spec.ExternalIPs),
			SpecSessionAffinity:          string(value.Spec.SessionAffinity),
			SpecLoadBalancerIP:           value.Spec.LoadBalancerIP,
			SpecLoadBalancerSourceRanges: p.transformServiceSpecLoadBalancerSourceRanges(value.Spec.LoadBalancerSourceRanges),
			SpecExternalName:             value.Spec.ExternalName,
			SpecExternalTrafficPolicy:    string(value.Spec.ExternalTrafficPolicy),
			SpecHealthCheckNodePort:      value.Spec.HealthCheckNodePort,
			SpecPublishNotReadyAddresses: value.Spec.PublishNotReadyAddresses,
//			SpecIPFamily:                 value.Spec.IPFamilies,
			SpecTopologyKeys:             p.transformServiceSpecTopologyKeys(value.Spec.TopologyKeys),

			StatusLoadBalancerIngress: p.transformServiceLoadBalancerIngresses(value.Status.LoadBalancer.Ingress),
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformServicePorts(values []v1.ServicePort) []*ServicePort {
	var tValues []*ServicePort
	for _, value := range values {
		tValue := ServicePort{
			Name:        value.Name,
			ClusterName: p.clusterName,
			Protocol:    string(value.Protocol),
			AppProtocol: value.AppProtocol,
			Port:        value.Port,
			TargetPort:  value.TargetPort.String(),
			NodePort:    value.NodePort,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformServiceSpecSelectors(values map[string]string) []*ServiceSpecSelector {
	var tValues []*ServiceSpecSelector
	for k, v := range values {
		tValue := ServiceSpecSelector{
			ClusterName: p.clusterName,
			Key:         k,
			Value:       v,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (p *Provider) transformServiceSpecExternalIPs(values []string) []*ServiceSpecExternalIPs {
	var tValues []*ServiceSpecExternalIPs
	for _, v := range values {
		tValues = append(tValues, &ServiceSpecExternalIPs{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformServiceSpecLoadBalancerSourceRanges(values []string) []*ServiceSpecLoadBalancerSourceRanges {
	var tValues []*ServiceSpecLoadBalancerSourceRanges
	for _, v := range values {
		tValues = append(tValues, &ServiceSpecLoadBalancerSourceRanges{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformServiceSpecTopologyKeys(values []string) []*ServiceSpecTopologyKeys {
	var tValues []*ServiceSpecTopologyKeys
	for _, v := range values {
		tValues = append(tValues, &ServiceSpecTopologyKeys{
			ClusterName: p.clusterName,
			Value:       v,
		})
	}
	return tValues
}

func (p *Provider) transformServiceLoadBalancerIngresses(values []v1.LoadBalancerIngress) []*ServiceLoadBalancerIngress {
	var tValues []*ServiceLoadBalancerIngress
	for _, value := range values {
		tValue := ServiceLoadBalancerIngress{
			ClusterName: p.clusterName,
			IP:          value.IP,
			Hostname:    value.Hostname,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

type ServiceConfig struct {
	Filter string
}

var serviceTables = []interface{}{
	&Service{},
	&ServicePort{},
	&ServiceSpecSelector{},
	&ServiceSpecExternalIPs{},
	&ServiceSpecLoadBalancerSourceRanges{},
	&ServiceSpecTopologyKeys{},
	&ServiceLoadBalancerIngress{},
}

func (p *Provider) services(gConfig interface{}) error {
	var config ServiceConfig
	ctx := context.Background()
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	output, err := p.client.CoreV1().Services("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return err
	}
	p.db.Where("cluster_name", p.clusterName).Delete(serviceTables...)
	p.db.ChunkedCreate(p.transformServices(output.Items))
	p.Logger.Info("Fetched resources", "resource", "k8s.services", "count", len(output.Items))

	return nil
}
