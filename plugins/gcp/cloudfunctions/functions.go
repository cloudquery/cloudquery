package cloudfunctions
import (
	"go.uber.org/zap"
	"google.golang.org/api/cloudfunctions/v1"
)


type CloudFunction struct {
	ID uint `gorm:"primarykey"`
	ProjectID string
	AvailableMemoryMb int64
	BuildId string
	BuildWorkerPool string
	Description string
	EntryPoint string

	EventTriggerEventType string

	EventTriggerResource string
	EventTriggerService string

	HttpsTriggerSecurityLevel string
	HttpsTriggerUrl string

	IngressSettings string
	MaxInstances int64
	Name string
	Network string
	Runtime string
	ServiceAccountEmail string
	SourceArchiveUrl string

	SourceRepositoryDeployedUrl string
	SourceRepositoryUrl string

	SourceToken string
	SourceUploadUrl string
	Status string
	Timeout string
	UpdateTime string
	VersionId int64
	VpcConnector string
	VpcConnectorEgressSettings string
}

func (CloudFunction) TableName() string {
	return "gcp_cloudfunctions_functions"
}

func (c *Client) transformCloudFunctions(values []*cloudfunctions.CloudFunction) []*CloudFunction {
	var tValues []*CloudFunction
	for _, value := range values {
		tValue := CloudFunction {
			ProjectID: c.projectID,
			AvailableMemoryMb: value.AvailableMemoryMb,
			BuildId: value.BuildId,
			BuildWorkerPool: value.BuildWorkerPool,
			Description: value.Description,
			EntryPoint: value.EntryPoint,
			IngressSettings: value.IngressSettings,
			MaxInstances: value.MaxInstances,
			Name: value.Name,
			Network: value.Network,
			Runtime: value.Runtime,
			ServiceAccountEmail: value.ServiceAccountEmail,
			SourceArchiveUrl: value.SourceArchiveUrl,
			SourceToken: value.SourceToken,
			SourceUploadUrl: value.SourceUploadUrl,
			Status: value.Status,
			Timeout: value.Timeout,
			UpdateTime: value.UpdateTime,
			VersionId: value.VersionId,
			VpcConnector: value.VpcConnector,
			VpcConnectorEgressSettings: value.VpcConnectorEgressSettings,
		}
		if value.EventTrigger != nil {

			tValue.EventTriggerEventType = value.EventTrigger.EventType
			tValue.EventTriggerResource = value.EventTrigger.Resource
			tValue.EventTriggerService = value.EventTrigger.Service

		}
		if value.HttpsTrigger != nil {

			tValue.HttpsTriggerSecurityLevel = value.HttpsTrigger.SecurityLevel
			tValue.HttpsTriggerUrl = value.HttpsTrigger.Url

		}
		if value.SourceRepository != nil {

			tValue.SourceRepositoryDeployedUrl = value.SourceRepository.DeployedUrl
			tValue.SourceRepositoryUrl = value.SourceRepository.Url

		}
		tValues = append(tValues, &tValue)
	}
	return tValues
}


var CloudFunctionTables = []interface{} {
	&CloudFunction{},
}

func (c *Client)functions(_ interface{}) error {

	nextPageToken := ""
	c.db.Where("project_id", c.projectID).Delete(CloudFunctionTables...)
	for {
		call := c.svc.Projects.Locations.Functions.List("projects/" + c.projectID + "/locations/-")
		call.PageToken(nextPageToken)
		output, err := call.Do()
		if err != nil {
			return err
		}

		c.db.ChunkedCreate(c.transformCloudFunctions(output.Functions))
		c.log.Info("populating CloudFunctions", zap.Int("count", len(output.Functions)))
		if output.NextPageToken == "" {
			break
		}
		nextPageToken = output.NextPageToken
	}
	return nil
}

