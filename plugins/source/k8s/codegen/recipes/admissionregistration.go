package recipes

import (
	admissionregistration "k8s.io/api/admissionregistration/v1"
	"k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
)

func AdmissionRegistration() []*Resource {
	resources := []*Resource{
		{
			SubService:     "mutating_webhook_configurations",
			Struct:         &admissionregistration.MutatingWebhookConfiguration{},
			ResourceFunc:   v1.MutatingWebhookConfigurationsGetter.MutatingWebhookConfigurations,
			GlobalResource: true,
		},
		{
			SubService:     "validating_webhook_configurations",
			Struct:         &admissionregistration.ValidatingWebhookConfiguration{},
			ResourceFunc:   v1.ValidatingWebhookConfigurationsGetter.ValidatingWebhookConfigurations,
			GlobalResource: true,
		},
	}

	for _, resource := range resources {
		resource.Service = "admissionregistration"
		resource.ServiceFunc = kubernetes.Interface.AdmissionregistrationV1
	}

	return resources
}
