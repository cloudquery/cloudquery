// Code generated by codegen; DO NOT EDIT.
package services

import (
	"github.com/fastly/go-fastly/v7/fastly"
	"net/http"
)

//go:generate mockgen -package=mocks -destination=../mocks/fastly.go -source=fastly.go FastlyClient
type FastlyClient interface {
	Get(string, *fastly.RequestOptions) (*http.Response, error)
	GetACL(*fastly.GetACLInput) (*fastly.ACL, error)
	GetACLEntry(*fastly.GetACLEntryInput) (*fastly.ACLEntry, error)
	GetAPIEvent(*fastly.GetAPIEventInput) (*fastly.Event, error)
	GetAPIEvents(*fastly.GetAPIEventsFilterInput) (fastly.GetAPIEventsResponse, error)
	GetBackend(*fastly.GetBackendInput) (*fastly.Backend, error)
	GetBigQuery(*fastly.GetBigQueryInput) (*fastly.BigQuery, error)
	GetBilling(*fastly.GetBillingInput) (*fastly.Billing, error)
	GetBlobStorage(*fastly.GetBlobStorageInput) (*fastly.BlobStorage, error)
	GetBulkCertificate(*fastly.GetBulkCertificateInput) (*fastly.BulkCertificate, error)
	GetCacheSetting(*fastly.GetCacheSettingInput) (*fastly.CacheSetting, error)
	GetCloudfiles(*fastly.GetCloudfilesInput) (*fastly.Cloudfiles, error)
	GetCondition(*fastly.GetConditionInput) (*fastly.Condition, error)
	GetCurrentUser() (*fastly.User, error)
	GetCustomTLSCertificate(*fastly.GetCustomTLSCertificateInput) (*fastly.CustomTLSCertificate, error)
	GetCustomTLSConfiguration(*fastly.GetCustomTLSConfigurationInput) (*fastly.CustomTLSConfiguration, error)
	GetDatadog(*fastly.GetDatadogInput) (*fastly.Datadog, error)
	GetDictionary(*fastly.GetDictionaryInput) (*fastly.Dictionary, error)
	GetDictionaryInfo(*fastly.GetDictionaryInfoInput) (*fastly.DictionaryInfo, error)
	GetDictionaryItem(*fastly.GetDictionaryItemInput) (*fastly.DictionaryItem, error)
	GetDiff(*fastly.GetDiffInput) (*fastly.Diff, error)
	GetDigitalOcean(*fastly.GetDigitalOceanInput) (*fastly.DigitalOcean, error)
	GetDirector(*fastly.GetDirectorInput) (*fastly.Director, error)
	GetDirectorBackend(*fastly.GetDirectorBackendInput) (*fastly.DirectorBackend, error)
	GetDomain(*fastly.GetDomainInput) (*fastly.Domain, error)
	GetDynamicSnippet(*fastly.GetDynamicSnippetInput) (*fastly.DynamicSnippet, error)
	GetERL(*fastly.GetERLInput) (*fastly.ERL, error)
	GetElasticsearch(*fastly.GetElasticsearchInput) (*fastly.Elasticsearch, error)
	GetFTP(*fastly.GetFTPInput) (*fastly.FTP, error)
	GetGCS(*fastly.GetGCSInput) (*fastly.GCS, error)
	GetGeneratedVCL(*fastly.GetGeneratedVCLInput) (*fastly.VCL, error)
	GetGzip(*fastly.GetGzipInput) (*fastly.Gzip, error)
	GetHTTPS(*fastly.GetHTTPSInput) (*fastly.HTTPS, error)
	GetHeader(*fastly.GetHeaderInput) (*fastly.Header, error)
	GetHealthCheck(*fastly.GetHealthCheckInput) (*fastly.HealthCheck, error)
	GetHeroku(*fastly.GetHerokuInput) (*fastly.Heroku, error)
	GetHoneycomb(*fastly.GetHoneycombInput) (*fastly.Honeycomb, error)
	GetKafka(*fastly.GetKafkaInput) (*fastly.Kafka, error)
	GetKinesis(*fastly.GetKinesisInput) (*fastly.Kinesis, error)
	GetLogentries(*fastly.GetLogentriesInput) (*fastly.Logentries, error)
	GetLoggly(*fastly.GetLogglyInput) (*fastly.Loggly, error)
	GetLogshuttle(*fastly.GetLogshuttleInput) (*fastly.Logshuttle, error)
	GetNewRelic(*fastly.GetNewRelicInput) (*fastly.NewRelic, error)
	GetObjectStore(*fastly.GetObjectStoreInput) (*fastly.ObjectStore, error)
	GetObjectStoreKey(*fastly.GetObjectStoreKeyInput) (string, error)
	GetOpenstack(*fastly.GetOpenstackInput) (*fastly.Openstack, error)
	GetOriginMetricsForService(*fastly.GetOriginMetricsInput) (*fastly.OriginInspector, error)
	GetOriginMetricsForServiceJSON(*fastly.GetOriginMetricsInput, interface{}) error
	GetPackage(*fastly.GetPackageInput) (*fastly.Package, error)
	GetPapertrail(*fastly.GetPapertrailInput) (*fastly.Papertrail, error)
	GetPool(*fastly.GetPoolInput) (*fastly.Pool, error)
	GetPrivateKey(*fastly.GetPrivateKeyInput) (*fastly.PrivateKey, error)
	GetPubsub(*fastly.GetPubsubInput) (*fastly.Pubsub, error)
	GetRegions() (*fastly.RegionsResponse, error)
	GetRequestSetting(*fastly.GetRequestSettingInput) (*fastly.RequestSetting, error)
	GetResponseObject(*fastly.GetResponseObjectInput) (*fastly.ResponseObject, error)
	GetS3(*fastly.GetS3Input) (*fastly.S3, error)
	GetSFTP(*fastly.GetSFTPInput) (*fastly.SFTP, error)
	GetScalyr(*fastly.GetScalyrInput) (*fastly.Scalyr, error)
	GetSecret(*fastly.GetSecretInput) (*fastly.Secret, error)
	GetSecretStore(*fastly.GetSecretStoreInput) (*fastly.SecretStore, error)
	GetServer(*fastly.GetServerInput) (*fastly.Server, error)
	GetService(*fastly.GetServiceInput) (*fastly.Service, error)
	GetServiceAuthorization(*fastly.GetServiceAuthorizationInput) (*fastly.ServiceAuthorization, error)
	GetServiceDetails(*fastly.GetServiceInput) (*fastly.ServiceDetail, error)
	GetSettings(*fastly.GetSettingsInput) (*fastly.Settings, error)
	GetSnippet(*fastly.GetSnippetInput) (*fastly.Snippet, error)
	GetSplunk(*fastly.GetSplunkInput) (*fastly.Splunk, error)
	GetStats(*fastly.GetStatsInput) (*fastly.StatsResponse, error)
	GetStatsField(*fastly.GetStatsInput) (*fastly.StatsFieldResponse, error)
	GetStatsJSON(*fastly.GetStatsInput, interface{}) error
	GetSumologic(*fastly.GetSumologicInput) (*fastly.Sumologic, error)
	GetSyslog(*fastly.GetSyslogInput) (*fastly.Syslog, error)
	GetTLSActivation(*fastly.GetTLSActivationInput) (*fastly.TLSActivation, error)
	GetTLSSubscription(*fastly.GetTLSSubscriptionInput) (*fastly.TLSSubscription, error)
	GetTokenSelf() (*fastly.Token, error)
	GetUsage(*fastly.GetUsageInput) (*fastly.UsageResponse, error)
	GetUsageByService(*fastly.GetUsageInput) (*fastly.UsageByServiceResponse, error)
	GetUser(*fastly.GetUserInput) (*fastly.User, error)
	GetVCL(*fastly.GetVCLInput) (*fastly.VCL, error)
	GetVersion(*fastly.GetVersionInput) (*fastly.Version, error)
	GetWAF(*fastly.GetWAFInput) (*fastly.WAF, error)
	GetWAFVersion(*fastly.GetWAFVersionInput) (*fastly.WAFVersion, error)
	ListACLEntries(*fastly.ListACLEntriesInput) ([]*fastly.ACLEntry, error)
	ListACLs(*fastly.ListACLsInput) ([]*fastly.ACL, error)
	ListAllWAFActiveRules(*fastly.ListAllWAFActiveRulesInput) (*fastly.WAFActiveRuleResponse, error)
	ListAllWAFRuleExclusions(*fastly.ListAllWAFRuleExclusionsInput) (*fastly.WAFRuleExclusionResponse, error)
	ListAllWAFRules(*fastly.ListAllWAFRulesInput) (*fastly.WAFRuleResponse, error)
	ListAllWAFVersions(*fastly.ListAllWAFVersionsInput) (*fastly.WAFVersionResponse, error)
	ListBackends(*fastly.ListBackendsInput) ([]*fastly.Backend, error)
	ListBigQueries(*fastly.ListBigQueriesInput) ([]*fastly.BigQuery, error)
	ListBlobStorages(*fastly.ListBlobStoragesInput) ([]*fastly.BlobStorage, error)
	ListBulkCertificates(*fastly.ListBulkCertificatesInput) ([]*fastly.BulkCertificate, error)
	ListCacheSettings(*fastly.ListCacheSettingsInput) ([]*fastly.CacheSetting, error)
	ListCloudfiles(*fastly.ListCloudfilesInput) ([]*fastly.Cloudfiles, error)
	ListConditions(*fastly.ListConditionsInput) ([]*fastly.Condition, error)
	ListCustomTLSCertificates(*fastly.ListCustomTLSCertificatesInput) ([]*fastly.CustomTLSCertificate, error)
	ListCustomTLSConfigurations(*fastly.ListCustomTLSConfigurationsInput) ([]*fastly.CustomTLSConfiguration, error)
	ListCustomerTokens(*fastly.ListCustomerTokensInput) ([]*fastly.Token, error)
	ListCustomerUsers(*fastly.ListCustomerUsersInput) ([]*fastly.User, error)
	ListDatadog(*fastly.ListDatadogInput) ([]*fastly.Datadog, error)
	ListDictionaries(*fastly.ListDictionariesInput) ([]*fastly.Dictionary, error)
	ListDictionaryItems(*fastly.ListDictionaryItemsInput) ([]*fastly.DictionaryItem, error)
	ListDigitalOceans(*fastly.ListDigitalOceansInput) ([]*fastly.DigitalOcean, error)
	ListDirectors(*fastly.ListDirectorsInput) ([]*fastly.Director, error)
	ListDomains(*fastly.ListDomainsInput) ([]*fastly.Domain, error)
	ListERLs(*fastly.ListERLsInput) ([]*fastly.ERL, error)
	ListElasticsearch(*fastly.ListElasticsearchInput) ([]*fastly.Elasticsearch, error)
	ListFTPs(*fastly.ListFTPsInput) ([]*fastly.FTP, error)
	ListGCSs(*fastly.ListGCSsInput) ([]*fastly.GCS, error)
	ListGzips(*fastly.ListGzipsInput) ([]*fastly.Gzip, error)
	ListHTTPS(*fastly.ListHTTPSInput) ([]*fastly.HTTPS, error)
	ListHeaders(*fastly.ListHeadersInput) ([]*fastly.Header, error)
	ListHealthChecks(*fastly.ListHealthChecksInput) ([]*fastly.HealthCheck, error)
	ListHerokus(*fastly.ListHerokusInput) ([]*fastly.Heroku, error)
	ListHoneycombs(*fastly.ListHoneycombsInput) ([]*fastly.Honeycomb, error)
	ListKafkas(*fastly.ListKafkasInput) ([]*fastly.Kafka, error)
	ListKinesis(*fastly.ListKinesisInput) ([]*fastly.Kinesis, error)
	ListLogentries(*fastly.ListLogentriesInput) ([]*fastly.Logentries, error)
	ListLoggly(*fastly.ListLogglyInput) ([]*fastly.Loggly, error)
	ListLogshuttles(*fastly.ListLogshuttlesInput) ([]*fastly.Logshuttle, error)
	ListNewRelic(*fastly.ListNewRelicInput) ([]*fastly.NewRelic, error)
	ListObjectStoreKeys(*fastly.ListObjectStoreKeysInput) (*fastly.ListObjectStoreKeysResponse, error)
	ListObjectStores(*fastly.ListObjectStoresInput) (*fastly.ListObjectStoresResponse, error)
	ListOpenstack(*fastly.ListOpenstackInput) ([]*fastly.Openstack, error)
	ListPapertrails(*fastly.ListPapertrailsInput) ([]*fastly.Papertrail, error)
	ListPools(*fastly.ListPoolsInput) ([]*fastly.Pool, error)
	ListPrivateKeys(*fastly.ListPrivateKeysInput) ([]*fastly.PrivateKey, error)
	ListPubsubs(*fastly.ListPubsubsInput) ([]*fastly.Pubsub, error)
	ListRequestSettings(*fastly.ListRequestSettingsInput) ([]*fastly.RequestSetting, error)
	ListResponseObjects(*fastly.ListResponseObjectsInput) ([]*fastly.ResponseObject, error)
	ListS3s(*fastly.ListS3sInput) ([]*fastly.S3, error)
	ListSFTPs(*fastly.ListSFTPsInput) ([]*fastly.SFTP, error)
	ListScalyrs(*fastly.ListScalyrsInput) ([]*fastly.Scalyr, error)
	ListSecretStores(*fastly.ListSecretStoresInput) (*fastly.SecretStores, error)
	ListSecrets(*fastly.ListSecretsInput) (*fastly.Secrets, error)
	ListServers(*fastly.ListServersInput) ([]*fastly.Server, error)
	ListServiceAuthorizations(*fastly.ListServiceAuthorizationsInput) (*fastly.ServiceAuthorizations, error)
	ListServiceDomains(*fastly.ListServiceDomainInput) (fastly.ServiceDomainsList, error)
	ListServices(*fastly.ListServicesInput) ([]*fastly.Service, error)
	ListSnippets(*fastly.ListSnippetsInput) ([]*fastly.Snippet, error)
	ListSplunks(*fastly.ListSplunksInput) ([]*fastly.Splunk, error)
	ListSumologics(*fastly.ListSumologicsInput) ([]*fastly.Sumologic, error)
	ListSyslogs(*fastly.ListSyslogsInput) ([]*fastly.Syslog, error)
	ListTLSActivations(*fastly.ListTLSActivationsInput) ([]*fastly.TLSActivation, error)
	ListTLSDomains(*fastly.ListTLSDomainsInput) ([]*fastly.TLSDomain, error)
	ListTLSSubscriptions(*fastly.ListTLSSubscriptionsInput) ([]*fastly.TLSSubscription, error)
	ListTokens() ([]*fastly.Token, error)
	ListVCLs(*fastly.ListVCLsInput) ([]*fastly.VCL, error)
	ListVersions(*fastly.ListVersionsInput) ([]*fastly.Version, error)
	ListWAFActiveRules(*fastly.ListWAFActiveRulesInput) (*fastly.WAFActiveRuleResponse, error)
	ListWAFRuleExclusions(*fastly.ListWAFRuleExclusionsInput) (*fastly.WAFRuleExclusionResponse, error)
	ListWAFRules(*fastly.ListWAFRulesInput) (*fastly.WAFRuleResponse, error)
	ListWAFVersions(*fastly.ListWAFVersionsInput) (*fastly.WAFVersionResponse, error)
	ListWAFs(*fastly.ListWAFsInput) (*fastly.WAFResponse, error)
	NewListACLEntriesPaginator(*fastly.ListACLEntriesInput) fastly.PaginatorACLEntries
	NewListDictionaryItemsPaginator(*fastly.ListDictionaryItemsInput) fastly.PaginatorDictionaryItems
	NewListObjectStoreKeysPaginator(*fastly.ListObjectStoreKeysInput) *fastly.ListObjectStoreKeysPaginator
	NewListObjectStoresPaginator(*fastly.ListObjectStoresInput) *fastly.ListObjectStoresPaginator
	NewListServicesPaginator(*fastly.ListServicesInput) fastly.PaginatorServices
}
