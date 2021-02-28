package crm

import (
	"google.golang.org/api/cloudresourcemanager/v1"
)


type Project struct {
	ID uint `gorm:"primarykey"`
	CreateTime string
	LifecycleState string
	Name string
	Labels []*ProjectLabel `gorm:"constraint:OnDelete:CASCADE;"`
	ParentId string
	ParentType string

	ProjectId string
	ProjectNumber int64
}

func (Project) TableName() string {
	return "gcp_crm_projects"
}

type ProjectLabel struct {
	ID uint `gorm:"primarykey"`
	ProjectID uint
	Key string
	Value string
}

func (ProjectLabel) TableName() string {
	return "gcp_crm_project_labels"
}

func (c *Client) transformProjects(value *cloudresourcemanager.Project) *Project {
	tValue := Project {
		CreateTime: value.CreateTime,
		LifecycleState: value.LifecycleState,
		Name: value.Name,
		ProjectId: value.ProjectId,
		ProjectNumber: value.ProjectNumber,
	}
	if value.Parent != nil {
		tValue.ParentId = value.Parent.Id
		tValue.ParentType = value.Parent.Type
	}

	for key, value := range value.Labels {
		tValue.Labels = append(tValue.Labels, &ProjectLabel{
			Key:       key,
			Value:     value,
		})
	}

	return &tValue
}


var ProjectTables = []interface{} {
	&Project{},
	&ProjectLabel{},
}

func (c *Client)projects(_ interface{}) error {

	c.db.Where("project_id", c.projectID).Delete(ProjectTables...)
	call := c.svc.Projects.Get(c.projectID)
	output, err := call.Do()
	if err != nil {
		return err
	}

	c.db.InsertOne(c.transformProjects(output))
	c.log.Info("populating Projects", "count", 1)
	return nil
}

