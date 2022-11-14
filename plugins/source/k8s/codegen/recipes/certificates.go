package recipes

import (
	resource "k8s.io/api/certificates/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/certificates/v1"
)

func Certificates() []*Resource {
	resources := []*Resource{
		{
			SubService:     "signing_requests",
			Struct:         &resource.CertificateSigningRequest{},
			ResourceFunc:   resourceType.CertificateSigningRequestsGetter.CertificateSigningRequests,
			GlobalResource: true,
		},
	}

	for _, resource := range resources {
		resource.Service = "certificates"
		resource.ServiceFunc = kubernetes.Interface.CertificatesV1
	}

	return resources
}
