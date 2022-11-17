---
title: Announcing CIS Policies for Kubernetes
tag: announcement
date: 2022/11/20
description: Announcing CIS Policy for Kubernetes
author: sckelemen
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

We are excited to announce the release of support for the CIS Kubernetes policy!

CloudQuery policies now help you ensure security and compliance in your Kubernetes environments.

CloudQuery brings the power of SQL to your cloud infrastructure, providing easy monitoring, governance and security.


## Run the policy

The Kubernetes CIS Policy can be run

```bash
psql -U {$PSQL_USER} -h localhost -f /plugins/source/k8s/policies/cis_v1.6.0/policy.sql
```

Or all Kubernetes policies can be run:

```bash
psql -U {$PSQL_USER} -h localhost -f /plugins/source/k8s/policies/policy.sql
```

## What’s Next?

We are always eager to hear feedback so feel free to file feature-requests/bugs/issues at [github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery/issues).

Also, we have more exciting features coming up to enhance the new policy feature so subscribe to our [twitter](https://twitter.com/cloudqueryio) and/or mailing list.
