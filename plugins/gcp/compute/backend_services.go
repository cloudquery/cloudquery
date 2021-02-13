package compute

import (
	"go.uber.org/zap"
	"google.golang.org/api/compute/v1"
)

type BackendService struct {
	ID                   uint `gorm:"primarykey"`
	ProjectID            string
	AffinityCookieTtlSec int64
	Backends             []*BackendServiceBackend `gorm:"constraint:OnDelete:CASCADE;"`

	CdnPolicyCacheKeyPolicyIncludeHost        bool
	CdnPolicyCacheKeyPolicyIncludeProtocol    bool
	CdnPolicyCacheKeyPolicyIncludeQueryString bool

	CdnPolicySignedUrlCacheMaxAgeSec int64

	CircuitBreakersMaxConnections           int64
	CircuitBreakersMaxPendingRequests       int64
	CircuitBreakersMaxRequests              int64
	CircuitBreakersMaxRequestsPerConnection int64
	CircuitBreakersMaxRetries               int64

	ConnectionDrainingDrainingTimeoutSec int64

	ConsistentHashHttpCookieName string
	ConsistentHashHttpCookiePath string

	ConsistentHashHttpCookieTtlNanos   int64
	ConsistentHashHttpCookieTtlSeconds int64

	ConsistentHashHttpHeaderName  string
	ConsistentHashMinimumRingSize int64

	CreationTimestamp    string

	Description          string
	EnableCDN            bool

	FailoverPolicyDisableConnectionDrainOnFailover bool
	FailoverPolicyDropTrafficIfUnhealthy           bool
	FailoverPolicyFailoverRatio                    float64

	Fingerprint  string

	IapEnabled                  bool
	IapOauth2ClientId           string
	IapOauth2ClientSecret       string
	IapOauth2ClientSecretSha256 string

	ResourceID          uint64
	Kind                string
	LoadBalancingScheme string
	LocalityLbPolicy    string

	LogConfigEnable     bool
	LogConfigSampleRate float64

	Name    string
	Network string

	OutlierDetectionBaseEjectionTimeNanos   int64
	OutlierDetectionBaseEjectionTimeSeconds int64

	OutlierDetectionConsecutiveErrors                  int64
	OutlierDetectionConsecutiveGatewayFailure          int64
	OutlierDetectionEnforcingConsecutiveErrors         int64
	OutlierDetectionEnforcingConsecutiveGatewayFailure int64
	OutlierDetectionEnforcingSuccessRate               int64

	OutlierDetectionIntervalNanos   int64
	OutlierDetectionIntervalSeconds int64

	OutlierDetectionMaxEjectionPercent       int64
	OutlierDetectionSuccessRateMinimumHosts  int64
	OutlierDetectionSuccessRateRequestVolume int64
	OutlierDetectionSuccessRateStdevFactor   int64

	Port           int64
	PortName       string
	Protocol       string
	Region         string
	SecurityPolicy string

	SecuritySettingsClientTlsPolicy string

	SelfLink        string
	SessionAffinity string
	TimeoutSec      int64
}

func (BackendService) TableName() string {
	return "gcp_compute_backend_services"
}

type BackendServiceBackend struct {
	ID                        uint   `gorm:"primarykey"`
	ProjectID                 string `gorm:"-"`
	BackendServiceID          uint   `neo:"ignore"`
	BalancingMode             string
	CapacityScaler            float64
	Description               string
	Failover                  bool
	Group                     string
	MaxConnections            int64
	MaxConnectionsPerEndpoint int64
	MaxConnectionsPerInstance int64
	MaxRate                   int64
	MaxRatePerEndpoint        float64
	MaxRatePerInstance        float64
	MaxUtilization            float64
}

func (BackendServiceBackend) TableName() string {
	return "gcp_compute_backend_service_backends"
}


func (c *Client) transformBackendServices(values []*compute.BackendService) []*BackendService {
	var tValues []*BackendService
	for _, value := range values {
		tValue := BackendService{
			ProjectID:            c.projectID,
			AffinityCookieTtlSec: value.AffinityCookieTtlSec,
			Backends:             c.transformBackendServiceBackends(value.Backends),
			CreationTimestamp:    value.CreationTimestamp,
			Description:          value.Description,
			EnableCDN:            value.EnableCDN,
			Fingerprint:          value.Fingerprint,
			ResourceID:           value.Id,
			Kind:                 value.Kind,
			LoadBalancingScheme:  value.LoadBalancingScheme,
			LocalityLbPolicy:     value.LocalityLbPolicy,
			Name:                 value.Name,
			Network:              value.Network,
			Port:                 value.Port,
			PortName:             value.PortName,
			Protocol:             value.Protocol,
			Region:               value.Region,
			SecurityPolicy:       value.SecurityPolicy,
			SelfLink:             value.SelfLink,
			SessionAffinity:      value.SessionAffinity,
			TimeoutSec:           value.TimeoutSec,
		}
		if value.CdnPolicy != nil {
			tValue.CdnPolicySignedUrlCacheMaxAgeSec = value.CdnPolicy.SignedUrlCacheMaxAgeSec
		}
		if value.CircuitBreakers != nil {
			tValue.CircuitBreakersMaxConnections = value.CircuitBreakers.MaxConnections
			tValue.CircuitBreakersMaxPendingRequests = value.CircuitBreakers.MaxPendingRequests
			tValue.CircuitBreakersMaxRequests = value.CircuitBreakers.MaxRequests
			tValue.CircuitBreakersMaxRequestsPerConnection = value.CircuitBreakers.MaxRequestsPerConnection
			tValue.CircuitBreakersMaxRetries = value.CircuitBreakers.MaxRetries

		}
		if value.ConnectionDraining != nil {

			tValue.ConnectionDrainingDrainingTimeoutSec = value.ConnectionDraining.DrainingTimeoutSec

		}
		if value.ConsistentHash != nil {

			tValue.ConsistentHashHttpHeaderName = value.ConsistentHash.HttpHeaderName
			tValue.ConsistentHashMinimumRingSize = value.ConsistentHash.MinimumRingSize

		}
		if value.FailoverPolicy != nil {

			tValue.FailoverPolicyDisableConnectionDrainOnFailover = value.FailoverPolicy.DisableConnectionDrainOnFailover
			tValue.FailoverPolicyDropTrafficIfUnhealthy = value.FailoverPolicy.DropTrafficIfUnhealthy
			tValue.FailoverPolicyFailoverRatio = value.FailoverPolicy.FailoverRatio

		}
		if value.Iap != nil {

			tValue.IapEnabled = value.Iap.Enabled
			tValue.IapOauth2ClientId = value.Iap.Oauth2ClientId
			tValue.IapOauth2ClientSecret = value.Iap.Oauth2ClientSecret
			tValue.IapOauth2ClientSecretSha256 = value.Iap.Oauth2ClientSecretSha256

		}
		if value.LogConfig != nil {

			tValue.LogConfigEnable = value.LogConfig.Enable
			tValue.LogConfigSampleRate = value.LogConfig.SampleRate

		}
		if value.OutlierDetection != nil {

			tValue.OutlierDetectionConsecutiveErrors = value.OutlierDetection.ConsecutiveErrors
			tValue.OutlierDetectionConsecutiveGatewayFailure = value.OutlierDetection.ConsecutiveGatewayFailure
			tValue.OutlierDetectionEnforcingConsecutiveErrors = value.OutlierDetection.EnforcingConsecutiveErrors
			tValue.OutlierDetectionEnforcingConsecutiveGatewayFailure = value.OutlierDetection.EnforcingConsecutiveGatewayFailure
			tValue.OutlierDetectionEnforcingSuccessRate = value.OutlierDetection.EnforcingSuccessRate
			tValue.OutlierDetectionMaxEjectionPercent = value.OutlierDetection.MaxEjectionPercent
			tValue.OutlierDetectionSuccessRateMinimumHosts = value.OutlierDetection.SuccessRateMinimumHosts
			tValue.OutlierDetectionSuccessRateRequestVolume = value.OutlierDetection.SuccessRateRequestVolume
			tValue.OutlierDetectionSuccessRateStdevFactor = value.OutlierDetection.SuccessRateStdevFactor

		}
		if value.SecuritySettings != nil {

			tValue.SecuritySettingsClientTlsPolicy = value.SecuritySettings.ClientTlsPolicy

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}

func (c *Client) transformBackendServiceBackends(values []*compute.Backend) []*BackendServiceBackend {
	var tValues []*BackendServiceBackend
	for _, value := range values {
		tValue := BackendServiceBackend{
			ProjectID:                 c.projectID,
			BalancingMode:             value.BalancingMode,
			CapacityScaler:            value.CapacityScaler,
			Description:               value.Description,
			Failover:                  value.Failover,
			Group:                     value.Group,
			MaxConnections:            value.MaxConnections,
			MaxConnectionsPerEndpoint: value.MaxConnectionsPerEndpoint,
			MaxConnectionsPerInstance: value.MaxConnectionsPerInstance,
			MaxRate:                   value.MaxRate,
			MaxRatePerEndpoint:        value.MaxRatePerEndpoint,
			MaxRatePerInstance:        value.MaxRatePerInstance,
			MaxUtilization:            value.MaxUtilization,
		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}


var BackendServiceTables = []interface{}{
	&BackendService{},
	&BackendServiceBackend{},
}

func (c *Client) backendServices(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(BackendServiceTables...)
	for {
		call := c.svc.BackendServices.AggregatedList(c.projectID)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		count := 0
		for _, backendServicesScopedList := range output.Items {
			c.db.ChunkedCreate(c.transformBackendServices(backendServicesScopedList.BackendServices))
			count += len(backendServicesScopedList.BackendServices)
		}
		c.log.Info("populating BackendServices", zap.Int("count", count))

		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
