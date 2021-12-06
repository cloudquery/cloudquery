package policy

import (
	"fmt"

	"github.com/cloudquery/cloudquery/pkg/config"
)

// TODO: move PolicyName to be also as part of args i.e first arg is policy name second arg is subpath, the rest is taken from config,
// TODO: if first arg isn't found in config, execute remote hub download only policy

func FilterPolicies(args []string, configPolicies []*config.Policy, policyName string) ([]*config.Policy, error) {
	var policies []*config.Policy

	if len(args) > 0 {
		remotePolicy, err := ParsePolicyFromArgs(args)
		if err != nil {
			return nil, err
		}
		policyConfig, err := remotePolicy.ToPolicyConfig()
		if len(args) == 2 {
			policyConfig.SubPath = args[1]
		}
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
Please add policy to block to your config file`)
	}
	policiesToRun := make([]*config.Policy, 0)

	// select policies to run
	for _, p := range policies {
		if policyName != "" {
			// request to run only specific policy
			if policyName == p.Name {
				policiesToRun = append(policiesToRun, p)
				break
			}
		}
		policiesToRun = append(policiesToRun, p)
	}

	return policiesToRun, nil
}
