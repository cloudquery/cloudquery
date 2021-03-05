package eks

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
	"github.com/mitchellh/mapstructure"
)

type Cluster struct {
	ID                       uint `gorm:"primarykey"`
	AccountID                string
	Region                   string
	Name                     *string
	Arn                      *string
	CertificateAuthorityData *string
	CreatedAt                *time.Time
	Endpoint                 *string
	OidcIssuer               *string
	ServiceIpv4Cidr          *string
	PlatformVersion          *string
	VpcID                    *string
	SecurityGroupID          *string
	EndpointPrivateAccess    bool
	EndpointPublicAccess     bool
	RoleArn                  *string
	Status                   string
	Version                  *string
	Tags                     []*ClusterTag                  `gorm:"constraint:OnDelete:CASCADE;"`
	LoggingConfigurations    []*ClusterLoggingConfiguration `gorm:"constraint:OnDelete:CASCADE;"`
	PublicAccessCidrs        []*ClusterPublicAccessCidr     `gorm:"constraint:OnDelete:CASCADE;"`
	SecurityGroups           []*ClusterSecurityGroup        `gorm:"constraint:OnDelete:CASCADE;"`
	Subnets                  []*ClusterSubnet               `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Cluster) TableName() string {
	return "aws_eks_clusters"
}

type ClusterTag struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Key   *string
	Value *string
}

func (ClusterTag) TableName() string {
	return "aws_eks_cluster_tags"
}

type ClusterLoggingConfiguration struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Enabled *bool
	Type    string
}

func (ClusterLoggingConfiguration) TableName() string {
	return "aws_eks_logging_configurations"
}

type ClusterPublicAccessCidr struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	Cidr string
}

func (ClusterPublicAccessCidr) TableName() string {
	return "aws_eks_public_access_cidr"
}

type ClusterSecurityGroup struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	SecurityGroupID string
}

func (ClusterSecurityGroup) TableName() string {
	return "aws_eks_security_groups"
}

type ClusterSubnet struct {
	ID        uint   `gorm:"primarykey"`
	ClusterID uint   `neo:"ignore"`
	AccountID string `gorm:"-"`
	Region    string `gorm:"-"`

	SubnetID string
}

func (ClusterSubnet) TableName() string {
	return "aws_eks_subnets"
}

func (c *Client) transformClusterTags(values *map[string]string) []*ClusterTag {
	var tValues []*ClusterTag
	for key, value := range *values {
		tValues = append(tValues, &ClusterTag{
			AccountID: c.accountID,
			Region:    c.region,
			Key:       &key,
			Value:     &value,
		})
	}
	return tValues
}

func (c *Client) transformClusterLoggingConfigurations(values *[]types.LogSetup) []*ClusterLoggingConfiguration {
	var tValues []*ClusterLoggingConfiguration
	for _, value := range *values {
		for _, logType := range value.Types {
			tValue := ClusterLoggingConfiguration{
				AccountID: c.accountID,
				Region:    c.region,
				Enabled:   value.Enabled,
				Type:      string(logType),
			}
			tValues = append(tValues, &tValue)
		}
	}
	return tValues
}

func (c *Client) transformClusterPublicAccessCidrs(values *[]string) []*ClusterPublicAccessCidr {
	var tValues []*ClusterPublicAccessCidr
	for _, value := range *values {
		tValues = append(tValues, &ClusterPublicAccessCidr{
			AccountID: c.accountID,
			Region:    c.region,
			Cidr:      value,
		})
	}
	return tValues
}

func (c *Client) transformClusterSecurityGroups(values *[]string) []*ClusterSecurityGroup {
	var tValues []*ClusterSecurityGroup
	for _, value := range *values {
		tValues = append(tValues, &ClusterSecurityGroup{
			AccountID:       c.accountID,
			Region:          c.region,
			SecurityGroupID: value,
		})
	}
	return tValues
}

func (c *Client) transformClusterSubnets(values *[]string) []*ClusterSubnet {
	var tValues []*ClusterSubnet
	for _, value := range *values {
		tValues = append(tValues, &ClusterSubnet{
			AccountID: c.accountID,
			Region:    c.region,
			SubnetID:  value,
		})
	}
	return tValues
}

func (c *Client) transformEKSClusters(values *[]types.Cluster) []*Cluster {
	var tValues []*Cluster
	for _, value := range *values {
		tValue := Cluster{
			AccountID:       c.accountID,
			Region:          c.region,
			Name:            value.Name,
			Arn:             value.Arn,
			CreatedAt:       value.CreatedAt,
			Endpoint:        value.Endpoint,
			PlatformVersion: value.PlatformVersion,
			RoleArn:         value.RoleArn,
			Status:          string(value.Status),
			Version:         value.Version,
			Tags:            c.transformClusterTags(&value.Tags),
		}
		if value.CertificateAuthority != nil {
			tValue.CertificateAuthorityData = value.CertificateAuthority.Data
		}
		if value.Identity != nil && value.Identity.Oidc != nil {
			tValue.OidcIssuer = value.Identity.Oidc.Issuer
		}
		if value.KubernetesNetworkConfig != nil {
			tValue.ServiceIpv4Cidr = value.KubernetesNetworkConfig.ServiceIpv4Cidr
		}
		if value.Logging != nil && value.Logging.ClusterLogging != nil {
			tValue.LoggingConfigurations = c.transformClusterLoggingConfigurations(&value.Logging.ClusterLogging)
		}
		if value.ResourcesVpcConfig != nil {
			tValue.PublicAccessCidrs = c.transformClusterPublicAccessCidrs(&value.ResourcesVpcConfig.PublicAccessCidrs)
			tValue.SecurityGroups = c.transformClusterSecurityGroups(&value.ResourcesVpcConfig.SecurityGroupIds)
			tValue.Subnets = c.transformClusterSubnets(&value.ResourcesVpcConfig.SubnetIds)
			tValue.VpcID = value.ResourcesVpcConfig.VpcId
			tValue.SecurityGroupID = value.ResourcesVpcConfig.ClusterSecurityGroupId
			tValue.EndpointPrivateAccess = value.ResourcesVpcConfig.EndpointPrivateAccess
			tValue.EndpointPublicAccess = value.ResourcesVpcConfig.EndpointPublicAccess
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

var ClusterTables = []interface{}{
	&Cluster{},
	&ClusterTag{},
	&ClusterLoggingConfiguration{},
	&ClusterPublicAccessCidr{},
	&ClusterSecurityGroup{},
	&ClusterSubnet{},
}

func (c *Client) clusters(gConfig interface{}) error {
	ctx := context.Background()
	var config eks.ListClustersInput
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	c.db.Where("region", c.region).Where("account_id", c.accountID).Delete(ClusterTables...)
	for {
		var clusters []types.Cluster
		listClustersOutput, err := c.svc.ListClusters(ctx, &config)
		if err != nil {
			return err
		}
		for _, name := range listClustersOutput.Clusters {
			describeClusterOutput, err := c.svc.DescribeCluster(ctx, &eks.DescribeClusterInput{
				Name: &name,
			})
			if err != nil {
				return err
			}
			clusters = append(clusters, *describeClusterOutput.Cluster)
		}

		c.db.ChunkedCreate(c.transformEKSClusters(&clusters))
		c.log.Info("Fetched resources", "resource", "eks.clusters", "count", len(clusters))

		if listClustersOutput.NextToken == nil {
			break
		}
		config.NextToken = listClustersOutput.NextToken
	}
	return nil
}
