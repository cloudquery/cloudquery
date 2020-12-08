package storage

import (
	"github.com/cloudquery/cloudquery/providers/common"
	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap"
	"google.golang.org/api/storage/v1"
	"log"
)

type Bucket struct {
	ID                          uint `gorm:"primarykey"`
	ProjectID                   string
	Acl                         []*BucketAccessControl `gorm:"constraint:OnDelete:CASCADE;"`
	BillingRequesterPays        bool
	Cors                        []*BucketCors `gorm:"constraint:OnDelete:CASCADE;"`
	DefaultEventBasedHold       bool
	DefaultObjectAcl            []*BucketObjectAccessControl `gorm:"constraint:OnDelete:CASCADE;"`
	EncryptionDefaultKmsKeyName string
	Etag                        string

	BucketPolicyOnlyEnabled            bool
	BucketPolicyOnlyLockedTime         string
	UniformBucketLevelAccessEnabled    bool
	UniformBucketLevelAccessLockedTime string

	ResourceID                     string
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

type BucketPolicyBinding struct {
	ID                   uint `gorm:"primarykey"`
	BucketID             uint
	ConditionDescription string
	ConditionExpression  string
	ConditionLocation    string
	ConditionTitle       string
	Members              []*BucketPolicyBindingsMember `gorm:"constraint:OnDelete:CASCADE;"`
	Role                 string
}

type BucketPolicyBindingsMember struct {
	ID                    uint `gorm:"primarykey"`
	BucketPolicyBindingID uint
	Name                  string
}

type BucketAccessControl struct {
	ID                       uint `gorm:"primarykey"`
	BucketID                 uint
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

type BucketCors struct {
	ID             uint `gorm:"primarykey"`
	BucketID       uint
	MaxAgeSeconds  int64
	Method         []*BucketCorsMethod         `gorm:"constraint:OnDelete:CASCADE;"`
	Origin         []*BucketCorsOrigin         `gorm:"constraint:OnDelete:CASCADE;"`
	ResponseHeader []*BucketCorsResponseHeader `gorm:"constraint:OnDelete:CASCADE;"`
}

type BucketCorsMethod struct {
	ID           uint `gorm:"primarykey"`
	BucketCorsID uint
	Value        string
}

type BucketCorsOrigin struct {
	ID           uint `gorm:"primarykey"`
	BucketCorsID uint
	Value        string
}

type BucketCorsResponseHeader struct {
	ID           uint `gorm:"primarykey"`
	BucketCorsID uint
	Value        string
}

type BucketObjectAccessControl struct {
	ID                       uint `gorm:"primarykey"`
	BucketID                 uint
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

type BucketLifecycleRule struct {
	ID                      uint `gorm:"primarykey"`
	BucketID                uint
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

type BucketLifecycleRuleConditionMatchesStorageClass struct {
	ID                    uint `gorm:"primarykey"`
	BucketLifecycleRuleID uint
	Value                 string
}

type BucketZoneAffinity struct {
	ID       uint `gorm:"primarykey"`
	BucketID uint
	Value    string
}

type BucketLabel struct {
	ID       uint `gorm:"primarykey"`
	BucketID uint
	Key      string
	Value    string
}

func (c *Client) transformPolicyBindingsMembers(values []string) []*BucketPolicyBindingsMember {
	var tValues []*BucketPolicyBindingsMember
	for _, v := range values {
		tValues = append(tValues, &BucketPolicyBindingsMember{
			Name: v,
		})
	}
	return tValues
}

func (c *Client) transformPolicyBinding(value *storage.PolicyBindings) *BucketPolicyBinding {
	res := BucketPolicyBinding{
		Members: c.transformPolicyBindingsMembers(value.Members),
		Role:    value.Role,
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
			Key:   k,
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformBucketAccessControl(value *storage.BucketAccessControl) *BucketAccessControl {
	res := BucketAccessControl{
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

func (c *Client) transformBucketAccessControls(values []*storage.BucketAccessControl) []*BucketAccessControl {
	var tValues []*BucketAccessControl
	for _, v := range values {
		tValues = append(tValues, c.transformBucketAccessControl(v))
	}
	return tValues
}

func (c *Client) transformBucketCorsMethods(values []string) []*BucketCorsMethod {
	var tValues []*BucketCorsMethod
	for _, v := range values {
		tValues = append(tValues, &BucketCorsMethod{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformBucketCorsOrigins(values []string) []*BucketCorsOrigin {
	var tValues []*BucketCorsOrigin
	for _, v := range values {
		tValues = append(tValues, &BucketCorsOrigin{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformBucketCorsResponseHeaders(values []string) []*BucketCorsResponseHeader {
	var tValues []*BucketCorsResponseHeader
	for _, v := range values {
		tValues = append(tValues, &BucketCorsResponseHeader{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformBucketCors(value *storage.BucketCors) *BucketCors {
	return &BucketCors{
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

func (c *Client) transformBucketObjectAccessControl(value *storage.ObjectAccessControl) *BucketObjectAccessControl {
	res := BucketObjectAccessControl{
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

func (c *Client) transformBucketObjectAccessControls(values []*storage.ObjectAccessControl) []*BucketObjectAccessControl {
	var tValues []*BucketObjectAccessControl
	for _, v := range values {
		tValues = append(tValues, c.transformBucketObjectAccessControl(v))
	}
	return tValues
}

func (c *Client) transformBucketLifecycleRuleConditionMatchesStorageClasss(values []string) []*BucketLifecycleRuleConditionMatchesStorageClass {
	var tValues []*BucketLifecycleRuleConditionMatchesStorageClass
	for _, v := range values {
		tValues = append(tValues, &BucketLifecycleRuleConditionMatchesStorageClass{
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformBucketLifecycleRule(value *storage.BucketLifecycleRule) *BucketLifecycleRule {
	res := BucketLifecycleRule{}

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
			Value: v,
		})
	}
	return tValues
}

func (c *Client) transformBucket(value *storage.Bucket) *Bucket {
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
		// we should return this err instead of calling log.Fatal
		log.Fatal(err)
	}
	res.PolicyBindings = c.transformPolicyBindings(output.Bindings)

	return &res
}

func (c *Client) transformBuckets(values []*storage.Bucket) []*Bucket {
	var tValues []*Bucket
	for _, v := range values {
		tValues = append(tValues, c.transformBucket(v))
	}
	return tValues
}

type BucketConfig struct {
	Prefix string
}

func (c *Client) buckets(gConfig interface{}) error {
	var config BucketConfig
	err := mapstructure.Decode(gConfig, &config)
	if err != nil {
		return err
	}
	if !c.resourceMigrated["storageBucket"] {
		err := c.db.AutoMigrate(
			&Bucket{},
			&BucketAccessControl{},
			&BucketCors{},
			&BucketObjectAccessControl{},
			&BucketLifecycleRule{},
			&BucketCorsMethod{},
			&BucketCorsOrigin{},
			&BucketCorsResponseHeader{},
			&BucketZoneAffinity{},
			&BucketLabel{},
			&BucketPolicyBinding{},
			&BucketPolicyBindingsMember{},
		)
		if err != nil {
			return err
		}
		c.resourceMigrated["storageBucket"] = true
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

		c.log.Debug("deleting previous Buckets", zap.String("project_id", c.projectID))
		c.db.Where("project_id = ?", c.projectID).Delete(&Bucket{})
		common.ChunkedCreate(c.db, c.transformBuckets(output.Items))
		c.log.Info("populating Buckets", zap.Int("count", len(output.Items)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}
