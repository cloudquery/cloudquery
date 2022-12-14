package vercel

import (
	"context"
	"fmt"
)

const (
	domainsURL       = "/v5/domains"
	domainRecordsURL = "/v4/domains/%s/records"
)

type Domain struct {
	BoughtAt            *MilliTime `json:"boughtAt"`
	CdnEnabled          bool       `json:"cdnEnabled"`       // undocumented
	ConfigVerifiedAt    *MilliTime `json:"configVerifiedAt"` // undocumented
	CreatedAt           MilliTime  `json:"createdAt"`
	ExpiresAt           *MilliTime `json:"expiresAt"`
	ID                  string     `json:"id"`
	IntendedNameservers []string   `json:"intendedNameservers"`
	CustomNameservers   []string   `json:"customNameservers"`
	Name                string     `json:"name"`
	Nameservers         []string   `json:"nameservers"`
	NsVerifiedAt        *MilliTime `json:"nsVerifiedAt"` // undocumented
	OrderedAt           *MilliTime `json:"orderedAt"`
	Renew               bool       `json:"renew"`
	ServiceType         string     `json:"serviceType"`
	TransferStartedAt   *MilliTime `json:"transferStartedAt"`
	TransferredAt       *MilliTime `json:"transferredAt"`
	TxtVerifiedAt       *MilliTime `json:"txtVerifiedAt"`      // undocumented
	VerificationRecord  string     `json:"verificationRecord"` // undocumented
	Verified            bool       `json:"verified"`
	Zone                bool       `json:"zone"` // undocumented
	Creator             struct {
		CustomerID       *string `json:"customerId"`
		Email            string  `json:"email"`
		ID               string  `json:"id"`
		IsDomainReseller *bool   `json:"isDomainReseller"`
		Username         string  `json:"username"`
	} `json:"creator"`
}

type DomainRecord struct {
	ID         string `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Value      string `json:"value"`
	MxPriority *int64 `json:"mxPriority"`
	Priority   *int64 `json:"priority"`
	Creator    string `json:"creator"`
	// duplicate: Created    *MilliTime `json:"created"`
	// duplicate: Updated    *MilliTime `json:"updated"`
	CreatedAt *MilliTime `json:"createdAt"`
	UpdatedAt *MilliTime `json:"updatedAt"`
	TTL       int64      `json:"ttl"` // undocumented
}

func (v *Client) ListDomains(ctx context.Context, pag *Paginator) ([]Domain, *Paginator, error) {
	var list struct {
		Domains    []Domain  `json:"domains"`
		Pagination Paginator `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, domainsURL, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Domains, &list.Pagination, nil
}

func (v *Client) ListDomainRecords(ctx context.Context, domainName string, pag *Paginator) ([]DomainRecord, *Paginator, error) {
	u := fmt.Sprintf(domainRecordsURL, domainName)

	var list struct {
		Records    []DomainRecord `json:"records"`
		Pagination Paginator      `json:"pagination"`
	}

	var until *int64
	if pag != nil {
		until = pag.Next
	}

	err := v.Request(ctx, u, until, &list)
	if err != nil {
		return nil, nil, err
	}
	return list.Records, &list.Pagination, nil
}
