package policy

import (
	"errors"
	"fmt"
)

// TODO: move PolicyName to be also as part of args i.e first arg is policy name second arg is subpath, the rest is taken from config,
// TODO: if first arg isn't found in config, execute remote hub download only policy

func FilterPolicies(args []string, policies Policies, policyName string) (Policies, error) {

	//if len(args) > 0 {
	//	remotePolicy, err := ParsePolicyFromArgs(args)
	//	if err != nil {
	//		return nil, err
	//	}
	//	policyConfig, err := remotePolicy.ToPolicyConfig()
	//	if len(args) == 2 {
	//		policyConfig.SubPath = args[1]
	//	}
	//	if err != nil {
	//		return nil, err
	//	}
	//	policies = append(policies, policyConfig)
	//} else {
	//	policies = configPolicies
	//}

	if len(policies) == 0 {
		return nil, errors.New("could not find policies to run. Please add policy to block to your config file")
	}

	// run them all
	if policyName == "" {
		return policies, nil
	}
	// select policies to run
	for _, p := range policies {
		// request to run only specific policy
		if policyName == p.Name {
			return Policies{p}, nil
		}
	}
	return nil, fmt.Errorf("no policy with name %s found. Available: %s", policyName, policies.All())
}
