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

- Run `cloudquery fetch`
- Create this view

```sql
-- Create Temporary View
CREATE TEMPORARY VIEW aws_security_group_egress_rules AS
(
    WITH sg_rules_ports AS (
        SELECT sg.account_id,
               sg.region,
               sg.group_name,
               sg.arn,
               sg.id,
               p.from_port,
               p.to_port,
               p.ip_protocol,
               p.cq_id AS permission_id
        FROM aws_ec2_security_groups sg
                 LEFT JOIN aws_ec2_security_group_ip_permissions p
                           ON sg.cq_id = p.security_group_cq_id
    )
    SELECT sgs.*, r.cidr AS ip
    FROM sg_rules_ports sgs
             LEFT JOIN aws_ec2_security_group_ip_permission_ip_ranges r
                       ON sgs.permission_id = r.security_group_ip_permission_cq_id
);
```

- Run the following query

```sql
-- Find all AWS instances that have a security group that allows unrestricted egress
SELECT id,
	region,
	account_id,
	vpc_id
FROM aws_ec2_instances
WHERE  cq_id in
	-- 	Find all instances that have egress rule that allows access to all ip addresses
	(SELECT instance_cq_id
		FROM aws_ec2_instance_security_groups
		JOIN aws_security_group_egress_rules ON group_id = id
		WHERE (ip = '0.0.0.0/0' OR ip = '::/0'));
```

You can also run the query straight from the policy pack located on our [GitHub](https://github.com/cloudquery-policies/aws/blob/main/public_egress/policy.hcl)

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

- Run `cloudquery fetch`
- Create this view

```sql
-- Create Temporary View
CREATE TEMPORARY VIEW aws_security_group_egress_rules AS
(
    WITH sg_rules_ports AS (
        SELECT sg.account_id,
               sg.region,
               sg.group_name,
               sg.arn,
               sg.id,
               p.from_port,
               p.to_port,
               p.ip_protocol,
               p.cq_id AS permission_id
        FROM aws_ec2_security_groups sg
                 LEFT JOIN aws_ec2_security_group_ip_permissions p
                           ON sg.cq_id = p.security_group_cq_id
    )
    SELECT sgs.*, r.cidr AS ip
    FROM sg_rules_ports sgs
             LEFT JOIN aws_ec2_security_group_ip_permission_ip_ranges r
                       ON sgs.permission_id = r.security_group_ip_permission_cq_id
);
```

- Run the following query

```sql
-- Find all AWS instances that are in a subnet that includes a catchall route
SELECT id,
       region,
       account_id,
       vpc_id
FROM aws_ec2_instances
WHERE subnet_id in
    --  Find all subnets that include a route table that inclues a catchall route
      (SELECT subnet_id
       FROM aws_ec2_route_tables
                JOIN aws_ec2_route_table_associations ON aws_ec2_route_table_associations.route_table_cq_id = aws_ec2_route_tables.cq_id
       WHERE aws_ec2_route_tables.cq_id in
                 --  Find all routes in any route table that contains a route to 0.0.0.0/0 or ::/0
             (SELECT route_table_cq_id
              FROM aws_ec2_route_table_routes
              WHERE destination_cidr_block = '0.0.0.0/0'
                 OR destination_ipv6_cidr_block = '::/0'))
  AND cq_id in
    -- 	Find all instances that have egress rule that allows access to all ip addresses
      (SELECT instance_cq_id
       FROM aws_ec2_instance_security_groups
                JOIN aws_security_group_egress_rules ON group_id = id
       WHERE (ip = '0.0.0.0/0' OR ip = '::/0'));
```

You can also run the query straight from the policy pack located on our [GitHub](https://github.com/cloudquery-policies/aws/blob/main/public_egress/policy.hcl)

### Other Methods

This doesn’t cover all the methods but does cover pretty much the most popular connectivity structure. For example, there are ways to peer VPCs which enables VPCs to share an Internet Gateway, this will require different queries. If we missed some other connectivity scenarios please open an issue in our [GitHub](https://github.com/cloudquery-policies/aws)

### Other Resources

EC2 instances are just one type of compute resource that can run a vulnerable application. Applications can also be run on other services including ECS, Lambda, AppRunner, Lightsail, and more.

If you found this tutorial/policy useful and you would like to see more of these feel free to either open an [issue](https://github.com/cloudquery-policies/aws), hop on [discord](https://www.cloudquery.io/discord) or [tweet](https://twitter.com/cloudqueryio) us.
