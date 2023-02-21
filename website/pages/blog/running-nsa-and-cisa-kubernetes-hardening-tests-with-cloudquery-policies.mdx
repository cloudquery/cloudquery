---
title: >-
  Running and Customizing NSA, CISA Kubernetes hardening guidance with
  CloudQuery Policies
tag: security
date: 2022/01/11
description: >-
  Automate, customize, codify and run NSA, CISA Kubernetes hardening CloudQuery
  Policies.
author: danielspangenberg
---

import { BlogHeader } from "../../components/BlogHeader"
import { Callout } from 'nextra-theme-docs'

<BlogHeader/>

<Callout type="warning">
HCL policies were deprecated - see up-to-date policy documentation [here](https://www.cloudquery.io/docs/core-concepts/policies).
</Callout>

On August 29th, 2022, USA's National Security Agency (NSA) and the Cybersecurity and Infrastructure Security Agency (CISA) released, [“Kubernetes Hardening Guidance”](https://media.defense.gov/2022/Aug/29/2003066362/-1/-1/0/CTR_KUBERNETES_HARDENING_GUIDANCE_1.2_20220829.PDF). The guide describes in great detail the challenges in the security k8s environment, base threat model and guidance on how to provide secure configuration to minimize risk.

As with any security guidelines, what is missing, or up to the user/security team, is how to validate, automate, customize, and implement those guidelines. Kubernetes environments vary widely, depending on usage, version, managed version (like GKE, EKS), requirements and capacity of the security team. All those factors will impact how you would want to implement those guidelines.

CloudQuery policies gives you a powerful way to automate, customize, codify, and run your cloud security & compliance continuously with HCL and SQL.

## Overview

CloudQuery NSA and CISA Kubernetes Hardening Guidance Policy contains more than 60 checks and is available on [GitHub](https://github.com/cloudquery-policies/k8s/tree/main/nsa_cisa_v1). The documentation for all the checks and queries is available on [GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/k8s/policies_v1)

Here is a snippet from the NSA & CISA Kubernetes policy:

```hcl
// policy.hcl
policy "nsa_cisa_v1" {
  title ="NSA and CISA Kubernetes Hardening Guidance v1.0"
  doc = file("nsa_cisa_v1/README.md")
  configuration {
    provider "k8s" {
      version = ">= v0.3.0"
    }
  }

  policy "pod_security" {
    source = file("nsa_cisa_v1/pod_security.hcl")
  }

  policy "network_hardening" {
    source = file("nsa_cisa_v1/network_hardening.hcl")
  }
}
```

And here is an example of how we check if a container has privileged access.

```sql
-- queries/pod_security/daemonset_container_privilege_disabled.sql
SELECT uid,
       name AS pod_name,
       namespace,
       context
FROM k8s_apps_daemon_sets,
     JSONB_ARRAY_ELEMENTS(template -> 'spec' -> 'containers') AS c
WHERE c -> 'securityContext' ->> 'privileged' = 'true';
```

The policy is split into sections (services) as sub-policies, so you can run either the whole policy, sub-policy or even a one specific check. The query itself is defined in a separate file, so we can re-use it in other policies in the future.

```bash
# execute specific policy pack
cloudquery policy run k8s//nsa_cisa_v1

# execute specific section in NSA and CISA policy pack
cloudquery policy run k8s//nsa_cisa_v1/pod_security
```

You are also free to fork this repository and create your own policy to adopt the guidelines to your needs.

## Running

Running this is as simple as ensuring your database has the latest cloud asset configuration with the fetch command and then executing each of pre-made queries with the policy run command.

Following is a quick start to run the policy. Otherwise, checkout full details on our [docs](/docs/core-concepts/policies).

### Prerequisite

Please follow the [quickstart guide](/docs/quickstart) documentation on how to `install`, `init`, and `fetch` the [K8S Provider](/docs/plugins/sources/overview).

### Running

```bash
# describe all available policies and sub-policies defined the K8s NSA and CISA Kubernetes Hardening Guidance
cloudquery policy describe k8s

# execute the whole policy pack
cloudquery policy run k8s

# execute specific policy pack
cloudquery policy run k8s//nsa_cisa_v1

# execute specific section in NSA and CISA policy pack
cloudquery policy run k8s//nsa_cisa_v1/pod_security
```

You can also output the results into a JSON and pass them to downstream processing for automated monitoring and alerting.

```bash
cloudquery policy run k8s//nsa_cisa_v1 --output-dir=results
```

## Build your own and share!

Do you have a policy that you want to codify, or you’ve been running it with python or bash scripts? You are welcome to try out codifying it with CloudQuery Policies. Feel free to drop on [discord](https://www.cloudquery.io/discord) or [GitHub](https://github.com/cloudquery/cloudquery/issues) to get any help, and we will share your policy on [CloudQuery Hub](https://www.cloudquery.io/).
