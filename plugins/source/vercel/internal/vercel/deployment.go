package vercel

import (
	"context"
	"fmt"
)

const (
	deploymentsURL      = "/v6/deployments"
	deploymentChecksURL = "/v1/deployments/%s/checks"
)

type Deployment struct {
	UID                 string     `json:"uid"`
	Name                string     `json:"name"`
	URL                 *string    `json:"url,omitempty"`
	Source              *string    `json:"source,omitempty"`
	State               *string    `json:"state,omitempty"`
	Type                string     `json:"type"`
	Creator             any        `json:"creator"`
	InspectorURL        string     `json:"inspectorUrl"`
	Meta                any        `json:"meta"`
	Target              any        `json:"target"`
	AliasError          any        `json:"aliasError"`
	AliasAssigned       any        `json:"aliasAssigned"`
	IsRollbackCandidate *bool      `json:"isRollbackCandidate,omitempty"`
	Ready               *MilliTime `json:"ready,omitempty"`
	ChecksState         *string    `json:"checksState,omitempty"`
	ChecksConclusion    *string    `json:"checksConclusion,omitempty"`

	// duplicate: Created             MilliTime       `json:"created"`
	CreatedAt  MilliTime `json:"createdAt"`
	BuildingAt MilliTime `json:"buildingAt"`
}

type DeploymentCheck struct {
	ID            string     `json:"id"`
	CreatedAt     MilliTime  `json:"createdAt"`
	CompletedAt   *MilliTime `json:"completedAt,omitempty"`
	Conclusion    *string    `json:"conclusion,omitempty"`
	DetailsURL    *string    `json:"detailsUrl,omitempty"`
	IntegrationID *string    `json:"integrationId,omitempty"`
	Name          *string    `json:"name,omitempty"`
	Output        any        `json:"output"`
	Path          *string    `json:"path,omitempty"`
	Rerequestable bool       `json:"rererequestable"`
	StartedAt     *MilliTime `json:"startedAt,omitempty"`
	UpdatedAt     *MilliTime `json:"updatedAt,omitempty"`
	Status        string     `json:"status"`
}

func (v *Client) ListDeployments(ctx context.Context, pag *Paginator) ([]Deployment, *Paginator, error) {
	var list struct {
		Deployments []Deployment `json:"deployments"`
		Pagination  Paginator    `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, deploymentsURL, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Deployments, &list.Pagination, nil
}

func (v *Client) ListDeploymentChecks(ctx context.Context, deploymentId string, pag *Paginator) ([]DeploymentCheck, *Paginator, error) {
	u := fmt.Sprintf(deploymentChecksURL, deploymentId)

	var list struct {
		Checks     []DeploymentCheck `json:"checks"`
		Pagination Paginator         `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, u, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Checks, &list.Pagination, nil
}
