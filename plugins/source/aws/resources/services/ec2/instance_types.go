package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func InstanceTypes() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_instance_types",
		Description: "Describes the instance type.",
		Resolver:    fetchEc2InstanceTypes,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "auto_recovery_supported",
				Description: "Indicates whether auto recovery is supported.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "bare_metal",
				Description: "Indicates whether the instance is a bare metal instance type.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "burstable_performance_supported",
				Description: "Indicates whether the instance type is a burstable performance instance type.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "current_generation",
				Description: "Indicates whether the instance type is current generation.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "dedicated_hosts_supported",
				Description: "Indicates whether Dedicated Hosts are supported on the instance type.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ebs_info_ebs_optimized_info_baseline_bandwidth_in_mbps",
				Description: "The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedInfo.BaselineBandwidthInMbps"),
			},
			{
				Name:        "ebs_info_ebs_optimized_info_baseline_iops",
				Description: "The baseline input/output storage operations per seconds for an EBS-optimized instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedInfo.BaselineIops"),
			},
			{
				Name:        "ebs_info_ebs_optimized_info_baseline_throughput_in_mb_ps",
				Description: "The baseline throughput performance for an EBS-optimized instance type, in MB/s.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedInfo.BaselineThroughputInMBps"),
			},
			{
				Name:        "ebs_info_ebs_optimized_info_maximum_bandwidth_in_mbps",
				Description: "The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedInfo.MaximumBandwidthInMbps"),
			},
			{
				Name:        "ebs_info_ebs_optimized_info_maximum_iops",
				Description: "The maximum input/output storage operations per second for an EBS-optimized instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedInfo.MaximumIops"),
			},
			{
				Name:        "ebs_info_ebs_optimized_info_maximum_throughput_in_mb_ps",
				Description: "The maximum throughput performance for an EBS-optimized instance type, in MB/s.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedInfo.MaximumThroughputInMBps"),
			},
			{
				Name:        "ebs_info_ebs_optimized_support",
				Description: "Indicates whether the instance type is Amazon EBS-optimized",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EbsInfo.EbsOptimizedSupport"),
			},
			{
				Name:        "ebs_info_encryption_support",
				Description: "Indicates whether Amazon EBS encryption is supported.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EbsInfo.EncryptionSupport"),
			},
			{
				Name:        "ebs_info_nvme_support",
				Description: "Indicates whether non-volatile memory express (NVMe) is supported.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EbsInfo.NvmeSupport"),
			},
			{
				Name:        "fpga_info_total_fpga_memory_in_mi_b",
				Description: "The total memory of all FPGA accelerators for the instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("FpgaInfo.TotalFpgaMemoryInMiB"),
			},
			{
				Name:        "free_tier_eligible",
				Description: "Indicates whether the instance type is eligible for the free tier.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "gpu_info_total_gpu_memory_in_mi_b",
				Description: "The total size of the memory for the GPU accelerators for the instance type, in MiB.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("GpuInfo.TotalGpuMemoryInMiB"),
			},
			{
				Name:        "hibernation_supported",
				Description: "Indicates whether On-Demand hibernation is supported.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor for the instance type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_storage_info_encryption_support",
				Description: "Indicates whether data is encrypted at rest.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceStorageInfo.EncryptionSupport"),
			},
			{
				Name:        "instance_storage_info_nvme_support",
				Description: "Indicates whether non-volatile memory express (NVMe) is supported.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceStorageInfo.NvmeSupport"),
			},
			{
				Name:        "instance_storage_info_total_size_in_gb",
				Description: "The total size of the disks, in GB.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("InstanceStorageInfo.TotalSizeInGB"),
			},
			{
				Name:        "instance_storage_supported",
				Description: "Indicates whether instance storage is supported.",
				Type:        schema.TypeBool,
			},
			{
				Name:            "instance_type",
				Description:     "The instance type",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "memory_info_size_in_mi_b",
				Description: "The size of the memory, in MiB.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MemoryInfo.SizeInMiB"),
			},
			{
				Name:        "network_info_default_network_card_index",
				Description: "The index of the default network card, starting at 0.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("NetworkInfo.DefaultNetworkCardIndex"),
			},
			{
				Name:        "network_info_efa_info_maximum_efa_interfaces",
				Description: "The maximum number of Elastic Fabric Adapters for the instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("NetworkInfo.EfaInfo.MaximumEfaInterfaces"),
			},
			{
				Name:        "network_info_efa_supported",
				Description: "Indicates whether Elastic Fabric Adapter (EFA) is supported.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkInfo.EfaSupported"),
			},
			{
				Name:        "network_info_ena_support",
				Description: "Indicates whether Elastic Network Adapter (ENA) is supported.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkInfo.EnaSupport"),
			},
			{
				Name:        "network_info_encryption_in_transit_supported",
				Description: "Indicates whether the instance type automatically encrypts in-transit traffic between instances.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkInfo.EncryptionInTransitSupported"),
			},
			{
				Name:        "network_info_ipv4_addresses_per_interface",
				Description: "The maximum number of IPv4 addresses per network interface.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("NetworkInfo.Ipv4AddressesPerInterface"),
			},
			{
				Name:        "network_info_ipv6_addresses_per_interface",
				Description: "The maximum number of IPv6 addresses per network interface.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("NetworkInfo.Ipv6AddressesPerInterface"),
			},
			{
				Name:        "network_info_ipv6_supported",
				Description: "Indicates whether IPv6 is supported.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("NetworkInfo.Ipv6Supported"),
			},
			{
				Name:        "network_info_maximum_network_cards",
				Description: "The maximum number of physical network cards that can be allocated to the instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("NetworkInfo.MaximumNetworkCards"),
			},
			{
				Name:        "network_info_maximum_network_interfaces",
				Description: "The maximum number of network interfaces for the instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("NetworkInfo.MaximumNetworkInterfaces"),
			},
			{
				Name:        "network_info_network_performance",
				Description: "The network performance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("NetworkInfo.NetworkPerformance"),
			},
			{
				Name:        "placement_group_info_supported_strategies",
				Description: "The supported placement group types.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("PlacementGroupInfo.SupportedStrategies"),
			},
			{
				Name:        "processor_info_supported_architectures",
				Description: "The architectures supported by the instance type.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("ProcessorInfo.SupportedArchitectures"),
			},
			{
				Name:        "processor_info_sustained_clock_speed_in_ghz",
				Description: "The speed of the processor, in GHz.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("ProcessorInfo.SustainedClockSpeedInGhz"),
			},
			{
				Name:        "supported_boot_modes",
				Description: "The supported boot modes",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "supported_root_device_types",
				Description: "The supported root device types.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "supported_usage_classes",
				Description: "Indicates whether the instance type is offered for spot or On-Demand.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "supported_virtualization_types",
				Description: "The supported virtualization types.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "v_cpu_info_default_cores",
				Description: "The default number of cores for the instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VCpuInfo.DefaultCores"),
			},
			{
				Name:        "v_cpu_info_default_threads_per_core",
				Description: "The default number of threads per core for the instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VCpuInfo.DefaultThreadsPerCore"),
			},
			{
				Name:        "v_cpu_info_default_v_cpus",
				Description: "The default number of vCPUs for the instance type.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("VCpuInfo.DefaultVCpus"),
			},
			{
				Name:        "v_cpu_info_valid_cores",
				Description: "The valid number of cores that can be configured for the instance type.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("VCpuInfo.ValidCores"),
			},
			{
				Name:        "v_cpu_info_valid_threads_per_core",
				Description: "The valid number of threads per core that can be configured for the instance type.",
				Type:        schema.TypeIntArray,
				Resolver:    schema.PathResolver("VCpuInfo.ValidThreadsPerCore"),
			},
			{
				Name:        "fpga_info",
				Description: "Describes the FPGA accelerator for the instance type.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("FpgaInfo"),
			},
			{
				Name:        "gpu_info",
				Description: "Describes the GPU accelerators for the instance type.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("GpuInfo"),
			},
			{
				Name:        "inference_accelerator_info",
				Description: "Describes the Inference accelerators for the instance type.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("InferenceAcceleratorInfo"),
			},
			{
				Name:        "instance_storage_info",
				Description: "Describes instance storage for the instance type.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("InstanceStorageInfo"),
			},
			{
				Name:        "network_info",
				Description: "Describes the network card support of the instance type.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("NetworkInfo"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2InstanceTypes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeInstanceTypesInput
	c := meta.(*client.Client)
	svc := c.Services().EC2

	for {
		response, err := svc.DescribeInstanceTypes(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.InstanceTypes
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}

	return nil
}
