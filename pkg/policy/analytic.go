package policy

import "github.com/cloudquery/cloudquery/internal/getter"

type Analytic struct {
	// Whether policy will persist in database
	Persistence bool
	// Name of the policy
	Name string
	// Type of the policy i.e S3/Hub/Git
	Type string
	// The selector used for the policy
	Selector string
	// Whether policy is private
	Private bool
}

func (p Policy) Analytic(dbPersistence bool) Analytic {
	pa := Analytic{
		Persistence: dbPersistence,
		Name:        p.Name,
		Type:        p.SourceType(),
		Selector:    p.SubPolicy(),
		Private:     p.SourceType() != "hub",
	}
	if !p.HasMeta() {
		policyName, subPath := getter.ParseSourceSubPolicy(p.Source)
		dp, _, _ := DetectPolicy(policyName, subPath)
		pa.Type = dp.SourceType()
		pa.Selector = subPath
		pa.Name = policyName
		pa.Private = p.SourceType() != "hub"
	}
	if pa.Private {
		pa.Selector = "private"
		pa.Name = p.Sha256Hash()
	}
	return pa
}

func (a Analytic) Properties() map[string]interface{} {
	return map[string]interface{}{
		"policy_persistence": a.Persistence,
		"policy_name":        a.Name,
		"policy_type":        a.Type,
		"policy_is_private":  a.Private,
		"policy_selector":    a.Selector,
	}
}
