---
title: 'AWS, Log4j and Finding Unrestricted Outbound Access'
tag: security
date: 2021/12/15
description: >-
  The Log4jshell (log4j) vulnerability (CVE-2021-44228) emphasized more than
  ever the importance of setting network controls & policies not only on inbound
  traffic but also on outbound traffic.
author: benjamin
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

The Log4shell (log4j) vulnerability (CVE-2021-44228) emphasized more than ever the importance of setting network controls & policies not only on inbound traffic but also on outbound traffic.

In this blog we will go through:

- What are the requirements for log4j exploitability?
- What are the possible ways to expose different AWS resources to the internet via outbound access?.
- How to find resources unrestricted outbound with CloudQuery open-source cloud asset inventory. This will help both to prioritize updates in the current situation as well as help apply network best practices in general.

## Log4j Vulnerability and Exploitability

The vulnerability in the popular log4j library is a critical remote code execution described in a simple and understandable way in the [Swiss Government Emergency and Response site](https://www.govcert.ch/blog/zero-day-exploit-targeting-popular-java-library-log4j/). The main takeaway here is that for an application to be exploitable it has to answer 3 requirements:
Java application with the vulnerable log4j version (Any version of log4j between versions 2.0 and 2.14.1 are affected).
Application is logging user controllable strings. This one is really hard to detect or to know what log lines are or are not controllable by a user, So better to assume it is always user controllable.
The application running the vulnerable version has unrestricted outbound access (Internet) to download the malicious payload from the attacker server or exfiltrate data.

The third point is a “critical” requirement, as without it your application cannot/less likely to be exploited. This is not to say you shouldn’t patch all of your vulnerable java applications ASAP.

## Finding EC2 Instances with Unrestricted Outbound Access

There are numerous ways to allow/disallow outbound access from an EC2. Following is a diagram:

![](/images/blog/outbound-architecture.png)

### Security Groups

The most important and usually the most commonly used is a security group. Security groups act as a virtual firewall and are attached directly to an instance (EC2 network interface).

Following is a query to identify all security groups with unrestricted outbound access.

**Prerequisite**:

- Run `cloudquery sync` (See our [quickstart guide](https://www.cloudquery.io/docs/quickstart))
- Create this view

```sql
-- Create Temporary View
create temporary view view_aws_security_group_egress_rules as
select
    account_id,
    region,
    group_name,
    arn,
    group_id as id,
    vpc_id,
    (i->>'FromPort')::integer AS from_port,
        (i->>'ToPort')::integer AS to_port,
        i->>'IpProtocol' AS ip_protocol,
    ip_ranges->>'CidrIp' AS ip,
    ip6_ranges->>'CidrIpv6' AS ip6
from aws_ec2_security_groups, JSONB_ARRAY_ELEMENTS(aws_ec2_security_groups.ip_permissions_egress) as i
    LEFT JOIN JSONB_ARRAY_ELEMENTS(i->'IpRanges') as ip_ranges ON true
    LEFT JOIN JSONB_ARRAY_ELEMENTS(i->'Ipv6Ranges') as ip6_ranges ON true;
```

- Run the following query

```sql
-- Find all AWS instances that have a security group that allows unrestricted egress
select aws_ec2_instances.account_id, 
       aws_ec2_instances.region, 
       aws_ec2_instances.instance_id, 
       sg->>'GroupId' AS security_group_id 
from aws_ec2_instances, jsonb_array_elements(security_groups) sg
    --  Find all instances that have egress rule that allows access to all ip addresses
    inner join view_aws_security_group_egress_rules on id = sg->>'GroupId'
where (ip = '0.0.0.0/0' or ip6 = '::/0');
```

You can also run the `public_egress` query straight from the policy pack located on our [GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies)

```bash
cloudquery policy run aws//public_egress
```

Depending on the size/number of accounts this is probably going to return a bunch of results. A good best practice (though it will be a project) would be to go through each one of the returned results and understand if this unrestricted outbound access is needed or can we tighten it up? If this is needed the best way to address it is to create the following tags, for example: “egress: true”, “egress-reason: this is our A microservice which needs to access A,B,Z….”

Even though your security groups should have least privilege network access configured, having unrestricted access to a security group doesn’t necessarily mean it really has outbound access to the internet. This means we can filter further to help find and prioritize resources that not only have wide open security groups but really have outbound internet access.

### NACLs

We won’t talk about NACLs here as they are less widely used and usually you need to configure the tightest security groups in any case.

### Internet Gateways & Egress Only Internet Gateways

The second requirement for outbound access is an internet gateway. Here is a query that checks if an instance with an unrestricted outbound access security group resides in a VPC with an internet gateway.

**Prerequisite**:

- Run `cloudquery sync` (See our [quickstart guide](https://www.cloudquery.io/docs/quickstart))
- Create this view

```sql
-- Create Temporary View
create temporary view view_aws_security_group_egress_rules as
select
    account_id,
    region,
    group_name,
    arn,
    group_id as id,
    vpc_id,
    (i->>'FromPort')::integer AS from_port,
        (i->>'ToPort')::integer AS to_port,
        i->>'IpProtocol' AS ip_protocol,
    ip_ranges->>'CidrIp' AS ip,
    ip6_ranges->>'CidrIpv6' AS ip6
from aws_ec2_security_groups, JSONB_ARRAY_ELEMENTS(aws_ec2_security_groups.ip_permissions_egress) as i
    LEFT JOIN JSONB_ARRAY_ELEMENTS(i->'IpRanges') as ip_ranges ON true
    LEFT JOIN JSONB_ARRAY_ELEMENTS(i->'Ipv6Ranges') as ip6_ranges ON true;
```

- Run the following query

```sql
-- Find all AWS instances that have a security group that allows unrestricted egress, and are in a VPC with an internet gateway
select aws_ec2_instances.account_id, 
       aws_ec2_instances.region, 
       aws_ec2_instances.instance_id, 
       sg->>'GroupId' AS security_group_id
from aws_ec2_instances, jsonb_array_elements(security_groups) sg
    --  Find all instances that have egress rule that allows access to all ip addresses
    inner join view_aws_security_group_egress_rules on id = sg->>'GroupId'
where (ip = '0.0.0.0/0' or ip6 = '::/0')
and ((aws_ec2_instances.vpc_id in (
    select value->>'VpcId' FROM aws_ec2_internet_gateways, jsonb_array_elements(aws_ec2_internet_gateways.attachments) AS value))
  or (aws_ec2_instances.vpc_id in (
    select value->>'VpcId' FROM aws_ec2_egress_only_internet_gateways, jsonb_array_elements(aws_ec2_egress_only_internet_gateways.attachments) AS value)));
```

### Other Methods

This doesn’t cover all the methods but does cover pretty much the most popular connectivity structure. For example, there are ways to peer VPCs which enables VPCs to share an Internet Gateway, this will require different queries. If we missed some other connectivity scenarios please open an issue in our [GitHub](https://github.com/cloudquery/cloudquery)

### Other Resources

EC2 instances are just one type of compute resource that can run a vulnerable application. Applications can also be run on other services including ECS, Lambda, AppRunner, Lightsail, and more.

If you found this tutorial/policy useful and you would like to see more of these feel free to either open an [issue](https://github.com/cloudquery/cloudquery), hop on [discord](https://www.cloudquery.io/discord) or [tweet](https://twitter.com/cloudqueryio) us.
