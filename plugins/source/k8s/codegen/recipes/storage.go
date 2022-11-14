package recipes

import (
	resource "k8s.io/api/storage/v1"
	"k8s.io/client-go/kubernetes"
	resourceType "k8s.io/client-go/kubernetes/typed/storage/v1"
)

func Storage() []*Resource {
	resources := []*Resource{
		{
			SubService:     "csi_drivers",
			Struct:         &resource.CSIDriver{},
			ResourceFunc:   resourceType.CSIDriversGetter.CSIDrivers,
			GlobalResource: true,
		},
		{
			SubService:     "csi_nodes",
			Struct:         &resource.CSINode{},
			ResourceFunc:   resourceType.CSINodesGetter.CSINodes,
			GlobalResource: true,
		},
		{
			SubService:   "csi_storage_capacities",
			Struct:       &resource.CSIStorageCapacity{},
			ResourceFunc: resourceType.CSIStorageCapacitiesGetter.CSIStorageCapacities,
		},
		{
			SubService:     "storage_classes",
			Struct:         &resource.StorageClass{},
			ResourceFunc:   resourceType.StorageClassesGetter.StorageClasses,
			GlobalResource: true,
		},
		{
			SubService:     "volume_attachments",
			Struct:         &resource.VolumeAttachment{},
			ResourceFunc:   resourceType.VolumeAttachmentsGetter.VolumeAttachments,
			GlobalResource: true,
		},
	}

	for _, resource := range resources {
		resource.Service = "storage"
		resource.ServiceFunc = kubernetes.Interface.StorageV1
	}

	return resources
}
