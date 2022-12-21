package vercel

import (
	"context"
	"fmt"
)

const (
	teamsURL       = "/v2/teams"
	teamMembersURL = "/v2/teams/%s/members"
)

type Team struct {
	ID        string    `json:"id"`
	Slug      string    `json:"slug"`
	Name      *string   `json:"name"`
	Avatar    *string   `json:"avatar"`
	CreatedAt MilliTime `json:"createdAt"`
	// duplicate: Created                 time.Time        `json:"created"`

	Membership              any        `json:"membership"`
	EnablePreviewFeedback   any        `json:"enablePreviewFeedback"`
	CreatorID               string     `json:"creatorId"`
	UpdatedAt               *MilliTime `json:"updatedAt"`
	PlatformVersion         any        `json:"platformVersion"`
	Billing                 any        `json:"billing"`
	Description             any        `json:"description"`
	Profiles                []any      `json:"profiles"`
	StagingPrefix           string     `json:"stagingPrefix"`
	ResourceConfig          any        `json:"resourceConfig"`
	PreviewDeploymentSuffix any        `json:"previewDeploymentSuffix"`
	SoftBlock               any        `json:"softBlock"`
	RemoteCaching           any        `json:"remoteCaching"`
	EnabledInvoiceItems     any        `json:"enabledInvoiceItems"`
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

	CreatedAt         MilliTime  `json:"createdAt"`
	AccessRequestedAt *MilliTime `json:"accessRequestedAt"`
	JoinedFrom        any        `json:"joinedFrom"`
}

type MemberSubAccount struct {
	AccountID *string `json:"accountId,omitempty"`
	Email     *string `json:"email,omitempty"`
	Login     *string `json:"login,omitempty"`
	UserID    *int64  `json:"userId,omitempty"`
}

func (v *Client) ListTeams(ctx context.Context, pag *Paginator) ([]Team, *Paginator, error) {
	var list struct {
		Teams      []Team    `json:"teams"`
		Pagination Paginator `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, teamsURL, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Teams, &list.Pagination, nil
}

func (v *Client) ListTeamMembers(ctx context.Context, teamID string, pag *Paginator) ([]TeamMember, *Paginator, error) {
	u := fmt.Sprintf(teamMembersURL, teamID)

	var list struct {
		TeamMembers []TeamMember `json:"members"`
		Pagination  Paginator    `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, u, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.TeamMembers, &list.Pagination, nil
}
