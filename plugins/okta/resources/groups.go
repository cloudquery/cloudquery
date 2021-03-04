package resources

import (
	"github.com/okta/okta-sdk-golang/v2/okta"
	"time"
)

type Group struct {
	Id                    string        `json:"id,omitempty" gorm:"primaryKey"`
	Created               *time.Time    `json:"created,omitempty"`
	LastMembershipUpdated *time.Time    `json:"lastMembershipUpdated,omitempty"`
	LastUpdated           *time.Time    `json:"lastUpdated,omitempty"`
	Profile               *GroupProfile `json:"profile,omitempty" gorm:"embedded;embeddedPrefix:profile_"`
	Type                  string        `json:"type,omitempty"`
	// User To Groups Many2Many this will make gorm auto create these tables
	Users []*User `json:"users,omitempty" gorm:"many2many:okta_user_groups"`
}

func (g Group) TableName() string {
	return "okta_groups"
}

type GroupProfile struct {
	Name        string
	Description string
}

func TransformGroup(group *okta.Group) *Group {
	if group == nil {
		return nil
	}

	return &Group{
		Id:                    group.Id,
		Created:               group.Created,
		LastMembershipUpdated: group.LastMembershipUpdated,
		LastUpdated:           group.LastUpdated,
		Profile:               TransformProfile(group.Profile),
		Type:                  group.Type,
	}

}

func TransformGroups(gg []*okta.Group) []*Group {
	tg := make([]*Group, len(gg))
	for i, g := range gg {
		tg[i] = TransformGroup(g)
	}
	return tg
}


func TransformProfile(profile *okta.GroupProfile) *GroupProfile {
	if profile == nil {
		return nil
	}
	return &GroupProfile{
		Name:        profile.Name,
		Description: profile.Description,
	}
}
