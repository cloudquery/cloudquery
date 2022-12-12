package model

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/resources/model"
)

const (
	TeamsURL       = "/v2/teams"
	TeamMembersURL = "/v2/teams/%s/members"
)

type Team struct {
	ID        string          `json:"id"`
	Slug      string          `json:"slug"`
	Name      *string         `json:"name"`
	Avatar    *string         `json:"avatar"`
	CreatedAt model.MilliTime `json:"createdAt"`
	// not included (dupe): Created                 time.Time        `json:"created"`

	Membership              interface{}      `json:"membership"`
	EnablePreviewFeedback   interface{}      `json:"enablePreviewFeedback"`
	CreatorID               string           `json:"creatorId"`
	UpdatedAt               *model.MilliTime `json:"updatedAt"`
	PlatformVersion         interface{}      `json:"platformVersion"`
	Billing                 interface{}      `json:"billing"`
	Description             interface{}      `json:"description"`
	Profiles                []interface{}    `json:"profiles"`
	StagingPrefix           string           `json:"stagingPrefix"`
	ResourceConfig          interface{}      `json:"resourceConfig"`
	PreviewDeploymentSuffix interface{}      `json:"previewDeploymentSuffix"`
	SoftBlock               interface{}      `json:"softBlock"`
	RemoteCaching           interface{}      `json:"remoteCaching"`
	EnabledInvoiceItems     interface{}      `json:"enabledInvoiceItems"`
}

type TeamMember struct {
	Avatar    *string           `json:"avatar"`
	Confirmed bool              `json:"confirmed"`
	Email     string            `json:"email"`
	Github    *MemberSubAccount `json:"github"`
	Gitlab    *MemberSubAccount `json:"gitlab"`
	Bitbucket *MemberSubAccount `json:"bitbucket"`

	Role     string  `json:"role"`
	UID      string  `json:"uid"`
	Username string  `json:"username"`
	Name     *string `json:"name"`

	CreatedAt         model.MilliTime  `json:"createdAt"`
	AccessRequestedAt *model.MilliTime `json:"accessRequestedAt"`
	JoinedFrom        interface{}      `json:"joinedFrom"`
}

type MemberSubAccount struct {
	AccountID *string `json:"accountId,omitempty"`
	Email     *string `json:"email,omitempty"`
	Login     *string `json:"login,omitempty"`
	UserID    *int64  `json:"userId,omitempty"`
}
