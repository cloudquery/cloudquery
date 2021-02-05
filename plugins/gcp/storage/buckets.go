package storage

import (
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/storage/v1"
)

type Bucket struct {
	_                           interface{}   `neo:"raw:MERGE (a:GCPProject {project_id: $project_id}) MERGE (a) - [:Resource] -> (n)"`
	ID                          uint          `gorm:"primarykey"`
	ProjectID                   string        `neo:"unique"`
	Acl                         []*BucketACLs `gorm:"constraint:OnDelete:CASCADE;"`
	BillingRequesterPays        bool
	Cors                        []*BucketCors `gorm:"constraint:OnDelete:CASCADE;"`
	DefaultEventBasedHold       bool
	DefaultObjectAcl            []*BucketObjectACLs `gorm:"constraint:OnDelete:CASCADE;"`
	EncryptionDefaultKmsKeyName string
	Etag                        string

	BucketPolicyOnlyEnabled            bool
	BucketPolicyOnlyLockedTime         string
	UniformBucketLevelAccessEnabled    bool
	UniformBucketLevelAccessLockedTime string

	ResourceID                     string `neo:"unique"`
	Kind                           string
	Labels                         []*BucketLabel         `gorm:"constraint:OnDelete:CASCADE;"`
	LifecycleRules                 []*BucketLifecycleRule `gorm:"constraint:OnDelete:CASCADE;"`
	Location                       string
	LocationType                   string
	LoggingLogBucket               string
	LoggingLogObjectPrefix         string
	Metageneration                 int64
	Name                           string
	OwnerEntity                    string
	OwnerEntityId                  string
	ProjectNumber                  uint64
	RetentionPolicyEffectiveTime   string
	RetentionPolicyIsLocked        bool
	RetentionPolicyRetentionPeriod int64
	SelfLink                       string
	StorageClass                   string
	TimeCreated                    string
	Updated                        string
	VersioningEnabled              bool
	WebsiteMainPageSuffix          string
	WebsiteNotFoundPage            string
	ZoneAffinity                   []*BucketZoneAffinity  `gorm:"constraint:OnDelete:CASCADE;"`
	PolicyBindings                 []*BucketPolicyBinding `gorm:"constraint:OnDelete:CASCADE;"`
}

func (Bucket) TableName() string {
	return "gcp_storage_buckets"
}

type BucketPolicyBinding struct {
	ID                   uint   `gorm:"primarykey"`
	BucketID             uint   `neo:"ignore"`
	ProjectID            string `gorm:"-"`
	ConditionDescription string
	ConditionExpression  string
	ConditionLocation    string
	ConditionTitle       string
	Members              []*BucketPolicyBindingsMember `gorm:"constraint:OnDelete:CASCADE;"`
	Role                 string
}

func (BucketPolicyBinding) TableName() string {
	return "gcp_storage_bucket_policy_bindings"
}

type BucketPolicyBindingsMember struct {
	ID                    uint   `gorm:"primarykey"`
	BucketPolicyBindingID uint   `neo:"ignore"`
	ProjectID             string `gorm:"-"`
	Name                  string
}

func (BucketPolicyBindingsMember) TableName() string {
	return "gcp_storage_bucket_policy_binding_members"
}

type BucketACLs struct {
	ID                       uint   `gorm:"primarykey"`
	BucketID                 uint   `neo:"ignore"`
	ProjectID                string `gorm:"-"`
	Bucket                   string
	Domain                   string
	Email                    string
	Entity                   string
	EntityId                 string
	Etag                     string
	ResourceID               string
	Kind                     string
	ProjectTeamProjectNumber string
	ProjectTeamTeam          string
	Role                     string
	SelfLink                 string
}

func (BucketACLs) TableName() string {
	return "gcp_storage_bucket_acls"
}

type BucketCors struct {
	ID             uint   `gorm:"primarykey"`
	BucketID       uint   `neo:"ignore"`
	ProjectID      string `gorm:"-"`
	MaxAgeSeconds  int64
	Method         []*BucketCorsMethod         `gorm:"constraint:OnDelete:CASCADE;"`
	Origin         []*BucketCorsOrigin         `gorm:"constraint:OnDelete:CASCADE;"`
	ResponseHeader []*BucketCorsResponseHeader `gorm:"constraint:OnDelete:CASCADE;"`
}

func (BucketCors) TableName() string {
	return "gcp_storage_bucket_cors"
}

type BucketCorsMethod struct {
	ID           uint   `gorm:"primarykey"`
	BucketCorsID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`
	Value        string
}

func (BucketCorsMethod) TableName() string {
	return "gcp_storage_bucket_cors_methods"
}

type BucketCorsOrigin struct {
	ID           uint   `gorm:"primarykey"`
	BucketCorsID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`
	Value        string
}

func (BucketCorsOrigin) TableName() string {
	return "gcp_storage_bucket_cors_origins"
}

type BucketCorsResponseHeader struct {
	ID           uint   `gorm:"primarykey"`
	BucketCorsID uint   `neo:"ignore"`
	ProjectID    string `gorm:"-"`
	Value        string
}

func (BucketCorsResponseHeader) TableName() string {
	return "gcp_storage_bucket_cors_response_headers"
}

type BucketObjectACLs struct {
	ID                       uint   `gorm:"primarykey"`
	BucketID                 uint   `neo:"ignore"`
	ProjectID                string `gorm:"-"`
	Bucket                   string
	Domain                   string
	Email                    string
	Entity                   string
	EntityId                 string
	Etag                     string
	Generation               int64
	ResourceID               string
	Kind                     string
	Object                   string
	ProjectTeamProjectNumber string
	ProjectTeamTeam          string
	Role                     string
	SelfLink                 string
}

func (BucketObjectACLs) TableName() string {
	return "gcp_storage_bucket_object_acls"
}

type BucketLifecycleRule struct {
	ID                      uint   `gorm:"primarykey"`
	BucketID                uint   `neo:"ignore"`
	ProjectID               string `gorm:"-"`
	ActionStorageClass      string
	ActionType              string
	Age                     int64
	CreatedBefore           string
	CustomTimeBefore        string
	DaysSinceCustomTime     int64
	DaysSinceNoncurrentTime int64
	IsLive                  *bool
	MatchesPattern          string
	MatchesStorageClass     []*BucketLifecycleRuleConditionMatchesStorageClass `gorm:"constraint:OnDelete:CASCADE;"`
	NoncurrentTimeBefore    string
	NumNewerVersions        int64
}

func (BucketLifecycleRule) TableName() string {
	return "gcp_storage_bucket_lifecycle_rules"
}

type BucketLifecycleRuleConditionMatchesStorageClass struct {
	ID                    uint   `gorm:"primarykey"`
	BucketLifecycleRuleID uint   `neo:"ignore"`
	ProjectID             string `gorm:"-"`
	Value                 string
}

func (BucketLifecycleRuleConditionMatchesStorageClass) TableName() string {
	return "gcp_storage_bucket_lifecycle_rule_condition_matches_storage_class"
}

type BucketZoneAffinity struct {
	ID        uint   `gorm:"primarykey"`
	BucketID  uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Value     string
}

func (BucketZoneAffinity) TableName() string {
	return "gcp_storage_bucket_zone_affinities"
}

type BucketLabel struct {
	ID        uint   `gorm:"primarykey"`
	BucketID  uint   `neo:"ignore"`
	ProjectID string `gorm:"-"`
	Key       string
	Value     string
}

func (BucketLabel) TableName() string {
	return "gcp_storage_bucket_labels"
}

func (c *Client) transformPolicyBindingsMembers(values []string) []*BucketPolicyBindingsMember {
	var tValues []*BucketPolicyBindingsMember
	for _, v := range values {
		tValues = append(tValues, &BucketPolicyBindingsMember{
			ProjectID: c.projectID,
			Name:      v,
		})
	}
	return tValues
}

func (c *Client) transformPolicyBinding(value *storage.PolicyBindings) *BucketPolicyBinding {
	res := BucketPolicyBinding{
		ProjectID: c.projectID,
		Members:   c.transformPolicyBindingsMembers(value.Members),
		Role:      value.Role,
	}

	if value.Condition != nil {
		res.ConditionDescription = value.Condition.Description
		res.ConditionExpression = value.Condition.Expression
		res.ConditionLocation = value.Condition.Location
		res.ConditionTitle = value.Condition.Title
	}

	return &res
}

func (c *Client) transformPolicyBindings(values []*storage.PolicyBindings) []*BucketPolicyBinding {
	var tValues []*BucketPolicyBinding
	for _, v := range values {
		tValues = append(tValues, c.transformPolicyBinding(v))
	}
	return tValues
}

func (c *Client) transformBucketLabels(values map[string]string) []*BucketLabel {
	var tValues []*BucketLabel
	for k, v := range values {
		tValues = append(tValues, &BucketLabel{
			ProjectID: c.projectID,
			Key:       k,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformBucketAccessControl(value *storage.BucketAccessControl) *BucketACLs {
	res := BucketACLs{
		ProjectID:  c.projectID,
		Bucket:     value.Bucket,
		Domain:     value.Domain,
		Email:      value.Email,
		Entity:     value.Entity,
		EntityId:   value.EntityId,
		Etag:       value.Etag,
		ResourceID: value.Id,
		Kind:       value.Kind,
		Role:       value.Role,
		SelfLink:   value.SelfLink,
	}

	if value.ProjectTeam != nil {
		res.ProjectTeamProjectNumber = value.ProjectTeam.ProjectNumber
		res.ProjectTeamTeam = value.ProjectTeam.Team
	}

	return &res
}

func (c *Client) transformBucketAccessControls(values []*storage.BucketAccessControl) []*BucketACLs {
	var tValues []*BucketACLs
	for _, v := range values {
		tValues = append(tValues, c.transformBucketAccessControl(v))
	}
	return tValues
}

func (c *Client) transformBucketCorsMethods(values []string) []*BucketCorsMethod {
	var tValues []*BucketCorsMethod
	for _, v := range values {
		tValues = append(tValues, &BucketCorsMethod{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformBucketCorsOrigins(values []string) []*BucketCorsOrigin {
	var tValues []*BucketCorsOrigin
	for _, v := range values {
		tValues = append(tValues, &BucketCorsOrigin{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformBucketCorsResponseHeaders(values []string) []*BucketCorsResponseHeader {
	var tValues []*BucketCorsResponseHeader
	for _, v := range values {
		tValues = append(tValues, &BucketCorsResponseHeader{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformBucketCors(value *storage.BucketCors) *BucketCors {
	return &BucketCors{
		ProjectID:      c.projectID,
		MaxAgeSeconds:  value.MaxAgeSeconds,
		Method:         c.transformBucketCorsMethods(value.Method),
		Origin:         c.transformBucketCorsOrigins(value.Origin),
		ResponseHeader: c.transformBucketCorsResponseHeaders(value.ResponseHeader),
	}
}

func (c *Client) transformBucketCorss(values []*storage.BucketCors) []*BucketCors {
	var tValues []*BucketCors
	for _, v := range values {
		tValues = append(tValues, c.transformBucketCors(v))
	}
	return tValues
}

func (c *Client) transformBucketObjectAccessControl(value *storage.ObjectAccessControl) *BucketObjectACLs {
	res := BucketObjectACLs{
		ProjectID:  c.projectID,
		Bucket:     value.Bucket,
		Domain:     value.Domain,
		Email:      value.Email,
		Entity:     value.Entity,
		EntityId:   value.EntityId,
		Etag:       value.Etag,
		Generation: value.Generation,
		ResourceID: value.Id,
		Kind:       value.Kind,
		Object:     value.Object,
		Role:       value.Role,
		SelfLink:   value.SelfLink,
	}
	if value.ProjectTeam != nil {
		res.ProjectTeamProjectNumber = value.ProjectTeam.ProjectNumber
		res.ProjectTeamTeam = value.ProjectTeam.Team
	}
	return &res
}

func (c *Client) transformBucketObjectAccessControls(values []*storage.ObjectAccessControl) []*BucketObjectACLs {
	var tValues []*BucketObjectACLs
	for _, v := range values {
		tValues = append(tValues, c.transformBucketObjectAccessControl(v))
	}
	return tValues
}

func (c *Client) transformBucketLifecycleRuleConditionMatchesStorageClasss(values []string) []*BucketLifecycleRuleConditionMatchesStorageClass {
	var tValues []*BucketLifecycleRuleConditionMatchesStorageClass
	for _, v := range values {
		tValues = append(tValues, &BucketLifecycleRuleConditionMatchesStorageClass{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformBucketLifecycleRule(value *storage.BucketLifecycleRule) *BucketLifecycleRule {
	res := BucketLifecycleRule{
		ProjectID: c.projectID,
	}

	if value.Action != nil {
		res.ActionStorageClass = value.Action.StorageClass
		res.ActionType = value.Action.Type
	}

	if value.Condition != nil {
		res.Age = value.Condition.Age
		res.CreatedBefore = value.Condition.CreatedBefore
		res.CustomTimeBefore = value.Condition.CustomTimeBefore
		res.DaysSinceCustomTime = value.Condition.DaysSinceCustomTime
		res.DaysSinceNoncurrentTime = value.Condition.DaysSinceNoncurrentTime
		res.IsLive = value.Condition.IsLive
		res.MatchesPattern = value.Condition.MatchesPattern
		res.MatchesStorageClass = c.transformBucketLifecycleRuleConditionMatchesStorageClasss(value.Condition.MatchesStorageClass)
		res.NoncurrentTimeBefore = value.Condition.NoncurrentTimeBefore
		res.NumNewerVersions = value.Condition.NumNewerVersions
	}

	return &res
}

func (c *Client) transformBucketLifecycleRules(values []*storage.BucketLifecycleRule) []*BucketLifecycleRule {
	var tValues []*BucketLifecycleRule
	for _, v := range values {
		tValues = append(tValues, c.transformBucketLifecycleRule(v))
	}
	return tValues
}

func (c *Client) transformBucketZoneAffinities(values []string) []*BucketZoneAffinity {
	var tValues []*BucketZoneAffinity
	for _, v := range values {
		tValues = append(tValues, &BucketZoneAffinity{
			ProjectID: c.projectID,
			Value:     v,
		})
	}
	return tValues
}

func (c *Client) transformBucket(value *storage.Bucket) (*Bucket, error) {
	res := Bucket{
		ProjectID:             c.projectID,
		Acl:                   c.transformBucketAccessControls(value.Acl),
		Cors:                  c.transformBucketCorss(value.Cors),
		Labels:                c.transformBucketLabels(value.Labels),
		DefaultEventBasedHold: value.DefaultEventBasedHold,
		DefaultObjectAcl:      c.transformBucketObjectAccessControls(value.DefaultObjectAcl),
		Etag:                  value.Etag,
		ResourceID:            value.Id,
		Kind:                  value.Kind,
		Location:              value.Location,
		LocationType:          value.LocationType,
		Metageneration:        value.Metageneration,
		Name:                  value.Name,
		ProjectNumber:         value.ProjectNumber,
		SelfLink:              value.SelfLink,
		StorageClass:          value.StorageClass,
		TimeCreated:           value.TimeCreated,
		Updated:               value.Updated,
		ZoneAffinity:          c.transformBucketZoneAffinities(value.ZoneAffinity),
	}

	if value.IamConfiguration != nil {
		res.BucketPolicyOnlyEnabled = value.IamConfiguration.BucketPolicyOnly.Enabled
		res.BucketPolicyOnlyLockedTime = value.IamConfiguration.BucketPolicyOnly.LockedTime
		res.UniformBucketLevelAccessEnabled = value.IamConfiguration.UniformBucketLevelAccess.Enabled
		res.UniformBucketLevelAccessLockedTime = value.IamConfiguration.UniformBucketLevelAccess.LockedTime
	}

	if value.Billing != nil {
		res.BillingRequesterPays = value.Billing.RequesterPays
	}

	if value.Encryption != nil {
		res.EncryptionDefaultKmsKeyName = value.Encryption.DefaultKmsKeyName
	}

	if value.Logging != nil {
		res.LoggingLogBucket = value.Logging.LogBucket
		res.LoggingLogObjectPrefix = value.Logging.LogObjectPrefix
	}

	if value.Owner != nil {
		res.OwnerEntity = value.Owner.Entity
		res.OwnerEntityId = value.Owner.EntityId
	}

	if value.RetentionPolicy != nil {
		res.RetentionPolicyEffectiveTime = value.RetentionPolicy.EffectiveTime
		res.RetentionPolicyIsLocked = value.RetentionPolicy.IsLocked
		res.RetentionPolicyRetentionPeriod = value.RetentionPolicy.RetentionPeriod
	}

	if value.Versioning != nil {
		res.VersioningEnabled = value.Versioning.Enabled
	}

	if value.Website != nil {
		res.WebsiteMainPageSuffix = value.Website.MainPageSuffix
		res.WebsiteNotFoundPage = value.Website.NotFoundPage
	}

	if value.Lifecycle != nil {
		res.LifecycleRules = c.transformBucketLifecycleRules(value.Lifecycle.Rule)
	}

	call := c.svc.Buckets.GetIamPolicy(value.Name)
	output, err := call.Do()
	if err != nil {
		if e, ok := err.(*googleapi.Error); ok {
			if e.Code == 403 && len(e.Errors) > 0 && e.Errors[0].Reason == "forbidden"{
				c.log.Info("access denied. skipping.",
					zap.String("project_id", c.projectID), zap.String("resource", "storage.buckets"))
				return &res, nil
			}
		}
		return nil, err
	}
	res.PolicyBindings = c.transformPolicyBindings(output.Bindings)

	return &res, nil
}

func (c *Client) transformBuckets(values []*storage.Bucket) ([]*Bucket, error) {
	var tValues []*Bucket
	for _, v := range values {
		tValue, err := c.transformBucket(v)
		if err != nil {
			return nil, err
		}
		tValues = append(tValues, tValue)
	}
	return tValues, nil
}

type BucketConfig struct {
	Prefix string
}

var BucketTables = []interface{}{
	&Bucket{},
	&BucketACLs{},
	&BucketCors{},
	&BucketObjectACLs{},
	&BucketLifecycleRule{},
	&BucketCorsMethod{},
	&BucketCorsOrigin{},
	&BucketCorsResponseHeader{},
	&BucketZoneAffinity{},
	&BucketLabel{},
	&BucketPolicyBinding{},
	&BucketPolicyBindingsMember{},
}

func (c *Client) buckets(gConfig interface{}) error {
	var config BucketConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}

	nextPageToken := ""
	for {
		call := c.svc.Buckets.List(c.projectID)
		call.Prefix(config.Prefix)
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.Where("project_id", c.projectID).Delete(BucketTables...)
		tValues, err := c.transformBuckets(output.Items)
		if err != nil {
			return err
		}
		c.db.ChunkedCreate(tValues)
		c.log.Info("Fetched resources", zap.String("resource", "storage.buckets"), zap.Int("count", len(output.Items)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
