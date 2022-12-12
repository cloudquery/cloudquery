package model

import "github.com/cloudquery/cloudquery/plugins/source/vercel/resources/model"

const (
	DomainsURL       = "/v5/domains"
	DomainRecordsURL = "/v4/domains/%s/records"
)

type Domain struct {
	BoughtAt            *model.MilliTime `json:"boughtAt"`
	CdnEnabled          bool             `json:"cdnEnabled"`       // undocumented
	ConfigVerifiedAt    *model.MilliTime `json:"configVerifiedAt"` // undocumented
	CreatedAt           model.MilliTime  `json:"createdAt"`
	ExpiresAt           *model.MilliTime `json:"expiresAt"`
	ID                  string           `json:"id"`
	IntendedNameservers []string         `json:"intendedNameservers"`
	CustomNameservers   []string         `json:"customNameservers"`
	Name                string           `json:"name"`
	Nameservers         []string         `json:"nameservers"`
	NsVerifiedAt        *model.MilliTime `json:"nsVerifiedAt"` // undocumented
	OrderedAt           *model.MilliTime `json:"orderedAt"`
	Renew               bool             `json:"renew"`
	ServiceType         string           `json:"serviceType"`
	TransferStartedAt   *model.MilliTime `json:"transferStartedAt"`
	TransferredAt       *model.MilliTime `json:"transferredAt"`
	TxtVerifiedAt       *model.MilliTime `json:"txtVerifiedAt"`      // undocumented
	VerificationRecord  string           `json:"verificationRecord"` // undocumented
	Verified            bool             `json:"verified"`
	Zone                bool             `json:"zone"` // undocumented
	Creator             struct {
		CustomerID       *string `json:"customerId"`
		Email            string  `json:"email"`
		ID               string  `json:"id"`
		IsDomainReseller *bool   `json:"isDomainReseller"`
		Username         string  `json:"username"`
	} `json:"creator"`
}

type DomainRecord struct {
	ID         string           `json:"id"`
	Slug       string           `json:"slug"`
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	Value      string           `json:"value"`
	MxPriority *int64           `json:"mxPriority"`
	Priority   *int64           `json:"priority"`
	Creator    string           `json:"creator"`
	Created    *model.MilliTime `json:"created"`
	Updated    *model.MilliTime `json:"updated"`
	CreatedAt  *model.MilliTime `json:"createdAt"`
	UpdatedAt  *model.MilliTime `json:"updatedAt"`
	TTL        int64            `json:"ttl"` // undocumented
}
