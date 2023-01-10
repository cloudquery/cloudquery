---
title: How to use CloudQuery for Attack Surface Management and Graph Visualization
tag: tutorial
description: >-
  How to use CloudQuery for Attack Surface Management (ASM) and Graph Visualization with Neo4j
author: jsonkao
---

import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

In this guide, we will demonstrate how to set up CloudQuery for customizable Attack Surface Management (ASM) and how to get started with utilizing graph visualization to execute security queries.

## Walkthrough

### Prerequisites

In this guide, we will use Neo4j as a destination and AWS as a source.  For more information on how to set those up, see our documentation on [Neo4j](https://www.cloudquery.io/docs/plugins/destinations/neo4j/overview) and [AWS](https://www.cloudquery.io/docs/plugins/sources/aws/overview).

Refer to Neo4j's installation documentation (https://neo4j.com/docs/operations-manual/current/installation/) for help setting up Neo4j. For this walkthrough, make sure a local instance of Neo4j is up and running.  Also make sure to install [Awesome Procedures on Cypher (APOC)](https://neo4j.com/labs/apoc/) for Neo4j as we'll be using useful functionality in APOC to assist with our attack surface management use cases.  

### Step 1: Install or Deploy CloudQuery

To get started with CloudQuery, check out our [quickstart guide](/docs/quickstart) and [AWS source plugin](/docs/plugins/sources/aws/overview) for how to configure CloudQuery and run it locally on your machine.

### Step 2: Sync Data from CloudQuery to Neo4j

The following command will sync data from AWS as a source to Neo4J as a destination:

`cloudquery sync aws-neo.yml neo4j.yml`

For more information on configuration files, see [AWS source configuration](https://www.cloudquery.io/docs/plugins/sources/aws/configuration) and [Neo4j destination configuration](https://www.cloudquery.io/docs/plugins/destinations/neo4j/overview)

You should see a `sync completed successfully` message.  CloudQuery has now synced your AWS data to Neo4j.

### Step 3: Test Out Neo4j

If running Neo4j locally, navigate to `http://localhost:7474/browser/` for the Neo4j browser. 

Let's start with a simple query to find all our IAM Roles.  

`MATCH (n:aws_iam_roles) RETURN n`

![IAM Roles](/images/how-to-guides/attack-surface-management-with-graph/iam-roles.png)

### Step 4: Run Custom ASM Queries and Create Relationships in Neo4j

* Example 1: IAM User Access Keys and their linked permissions.

Let's start with IAM User Access Keys.  With this query, we'll look for the 4 distinct ways with identity policies an IAM User can be granted permissions and link those to the IAM Users and to the IAM User Access Keys.

An IAM User can be granted permissions from:
* Direct Inline Policies of the IAM User
* Directly Attached Managed Policies to the IAM User
* Inline Policies via IAM Group Membership
* Attached Managed Policies via IAM Group Membership

By running the following command, we create a relationship between IAM User nodes and IAM User Access Key nodes.

```cypher 
MATCH
  (a:aws_iam_user_access_keys),
  (b:aws_iam_users)
WHERE a.`_cq_parent_id` =  b.`_cq_id`
CREATE (b)-[r:has_access_keys]->(a)
RETURN type(r)
```

Next, let's create a relationship between IAM Users and IAM Groups.  In AWS, IAM Users can be members of IAM Groups and inherit their IAM policies.

```cypher
MATCH
  (iamusers:aws_iam_users),
  (usergroups:aws_iam_user_groups),
  (iamgroups:aws_iam_groups)
WHERE iamusers.arn = usergroups.user_arn and iamgroups.arn = usergroups.arn
CREATE (iamusers)-[r:is_a_member_of]->(iamgroups)
RETURN type(r)
```

Now, we'll create relationships between IAM Users and all IAM User inline policies.

```cypher
MATCH
(iamusers:aws_iam_users),
(inlinep:aws_iam_user_policies)
WHERE iamusers.arn = inlinep.user_arn
CREATE (iamusers)-[r:has_inline_policy]->(inlinep)
RETURN type(r)
```

Now, we'll create relationships between IAM Users and directly attached managed policies. 

```cypher
MATCH
(iamusers:aws_iam_users),
(attachp:aws_iam_user_attached_policies)
WHERE iamusers.arn = inlinep.user_arn
CREATE (iamusers)-[r:has_attached_policy]->(attachp)
RETURN type(r)
```

Next, we'll create relationships between IAM Groups and their inline policies.

```cypher
MATCH
(iamgroups:aws_iam_groups),
(groupinline:aws_iam_group_policies)
WHERE iamgroups.arn = groupinline.group_arn
CREATE (iamgroups)-[r:has_inline_policy]->(groupinline)
RETURN type(r)
```

Lastly, we'll create relationships between IAM Groups and their attached managed policies.

```cypher
MATCH (n:aws_iam_groups), (policies:aws_iam_policies) 
UNWIND (keys(apoc.convert.fromJsonMap(n.policies))) as y 
WITH y, policies, n
WHERE y = policies.arn
CREATE (n)-[r:has_attached_policy]->(policies)
RETURN type(r)
```

After all these relationships have been created, we can run a `MATCH (n:aws_iam_users) return n` to return all IAM Users.  

In the UI, feel free to play around with node labels, colors, and expansion of nodes and relationships. 

In our sample environment, we have 3 IAM Users.  The following image shows the following and their relationships:
* IAM User Access Keys in Green
* IAM Managed Policies and Inline Policies in Red 
* IAM Users are in Gray
* IAM Groups are in Blue

![Sample Graph of IAM Users](/images/how-to-guides/attack-surface-management-with-graph/graph-users.png)

* Example 2: Data in RDS

Create RDS Instances Relationship with KMS Keys
```cypher
MATCH (rdsinstances:aws_rds_instances), (kmskeys:aws_kms_keys)
WHERE rdsinstances.kms_key_id = kmskeys.arn
CREATE (rdsinstances)-[r:is_encrypted_by_key]->(kmskeys)
RETURN type(r)
```

Connect KMS Keys with all their Key Grants
```cypher
MATCH (keygrants:aws_kms_key_grants), (kmskeys:aws_kms_keys)
WHERE keygrants.key_arn = kmskeys.arn
CREATE (kmskeys)-[r:has_kms_key_grant]->(keygrants)
RETURN type(r)
```

Connect RDS Instances with their Security Groups and Networking

```cypher
MATCH (rds_is:aws_rds_instances), (sgs:aws_ec2_security_groups) WHERE apoc.convert.fromJsonList(rds_is.vpc_security_groups)[0]['VpcSecurityGroupId'] = sgs.group_id CREATE (rds_is)-[r:uses_security_group]->(sgs) 
RETURN type(r)
```

Connect KMS Keys with KMS Key Policies
```cypher 
MATCH (keypolicies:aws_kms_key_policies), (keys:aws_kms_keys) 
WHERE keypolicies.key_arn = keys.arn
CREATE (keys)-[r:has_key_policy]->(keypolicies)
RETURN type(r)
```


## Summary

We have demonstrated how to get started with Attack Surface Management (ASM) and graph visualization with Neo4j along with some starter queries.  Now you should be able to customize and create more queries, relationships, and utilize CloudQuery to help improve the security posture of your organization!

If you have use cases, custom queries, and examples from using CloudQuery, we would love to hear from you! Reach out to us on [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://cloudquery.io/discord)!