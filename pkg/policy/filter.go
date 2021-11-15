package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config"
)

func FilterPolicies(args []string, configPolicies []*config.Policy, policyName, subPath string) ([]*config.Policy, error) {
	var policies []*config.Policy

	if len(args) > 0 {
		remotePolicy, err := ParsePolicyFromArgs(args)
		if err != nil {
			return nil, err
		}
		policyConfig, err := remotePolicy.ToPolicyConfig()
		policyConfig.SubPath = subPath
		if err != nil {
			return nil, err
		}
		policies = append(policies, policyConfig)
	} else {
		policies = configPolicies
	}

	if len(policies) == 0 {
		return nil, fmt.Errorf(`
Could not find policies to run.
Please add policy to block to your config file.
`)
	}
	policiesToRun := make([]*config.Policy, 0)

	// select policies to run
	for _, p := range policies {
		if policyName != "" {
			// request to run only specific policy
			if policyName == p.Name {
				// override subPath if specified
				if subPath != "" {
					p.SubPath = subPath
				}
				policiesToRun = append(policiesToRun, p)
				break
			}
		}
		policiesToRun = append(policiesToRun, p)
	}

	return policiesToRun, nil
}
