package client

import (
	"github.com/go-gandi/go-gandi/certificate"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/email"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/go-gandi/go-gandi/simplehosting"
)

//go:generate mockgen -package=mocks -destination=./mocks/mock_domain_client.go . DomainClient
type DomainClient interface {
	//domain.Domain
	ListDomains() ([]domain.ListResponse, error)
	GetDomain(domain string) (domain.Details, error)
	GetNameServers(domain string) ([]string, error)
	GetContacts(domain string) (domain.Contacts, error)
	ListDNSSECKeys(domain string) ([]domain.DNSSECKey, error)
	ListGlueRecords(domain string) ([]domain.GlueRecord, error)
	ListWebRedirections(domain string) ([]domain.WebRedirection, error)
	GetLiveDNS(domain string) (domain.LiveDNS, error)
	GetTags(domain string) ([]string, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_email_client.go . EmailClient
type EmailClient interface {
	ListMailboxes(domain string) ([]email.ListMailboxResponse, error)
	GetMailbox(domain, mailbox_id string) (email.MailboxResponse, error)
	GetForwards(domain string) ([]email.GetForwardRequest, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_livedns_client.go . LiveDNSClient
type LiveDNSClient interface {
	ListDomains() ([]livedns.Domain, error)
	GetDomainNS(fqdn string) ([]string, error)
	ListSnapshots(fqdn string) ([]livedns.Snapshot, error)
	// not implemented: axfr, keys, domain records
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_simplehosting_client.go . SimpleHostingClient
type SimpleHostingClient interface {
	ListInstances() ([]simplehosting.Instance, error)
	//GetInstance(instanceId string) (simplehosting.Instance, error)
	ListVhosts(instanceId string) ([]simplehosting.Vhost, error)
	//GetVhost(instanceId string, fqdn string) (simplehosting.Vhost, error)
}

//go:generate mockgen -package=mocks -destination=./mocks/mock_certificate_client.go . CertificateClient
type CertificateClient interface {
	ListCertificates() ([]certificate.CertificateType, error)
	GetCertificate(certificateId string) (certificate.CertificateType, error)
	ListPackages() ([]certificate.Package, error)
}
