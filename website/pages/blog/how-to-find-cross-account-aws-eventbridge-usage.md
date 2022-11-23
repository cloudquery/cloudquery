---
title: Finding Cross Account AWS EventBridge Usage
tag: security
date: 2022/11/25
description: >-
  How a CloudQuery customer wrote custom queries to find cross account AWS EventBridge Usage to help their teams migrate to new AWS EventBridge security features to increase security and compliance.
author: jsonkao
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

## Overview

Recently, AWS sent out customer notification emails regarding upcoming changes for EventBridge cross account event bus targets.  In short, Amazon EventBridge today does not require a IAM role when sending events to a cross account event bus target.  

## EventBridge Change

Beginning February 16, 2023, Amazon EventBridge will start requiring IAM roles for all new cross account event bus targets.


## Customer Query

We would like to thank `jbarney` for sharing and writing the below query.  We're especially happy when our users bring innovation and layer on advanced queries on top of CloudQuery data to provide value to their organizations.

```sql
SELECT
  account_id,
  arn,
  policy,
  (regexp_match(policy, '[0-9]{12})),
  account_id != (regexp_match(policy, '[0-9]{12}))[1] as allows_cross_account
FROM aws_eventbridge_event_buses
WHERE policy ~ '[0-9]{12}'
and account_id != (regexp_match(policy, '[0-9]{12}))[1];
```

## References and Useful Links

[AWS: Service Control Policies (SCPs)](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps.html)

[AWS: Sending and receiving Amazon EventBridge events between AWS Account](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-cross-account.html)

[CloudQuery: AWS Plugin](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws)