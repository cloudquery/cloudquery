---
title: Introducing CloudQuery SDK
tag: product
date: 2021/04/07
description: Introducing simple and extensible SDK to query your cloud
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Today we are pleased to announce the release of [CloudQuery SDK](https://github.com/cloudquery/plugin-sdk)!

We released [CloudQuery](https://github.com/cloudquery/cloudquery) at the end of last year to give developers,
SREs and security engineers a better and open-source alternative to gain deep visibility into their cloud infrastructure.
We made a few decisions like using SQL as a query and policy engine so developers won’t need to learn yet another query or policy engine!

**Bonus**: Reuse and take advantage of the whole huge SQL ecosystem.

CloudQuery grew over 1.3K stars in under 5 months!
This led us to develop a better and simpler way to extend CloudQuery with new resources and customs providers.

## Enter CloudQuery SDK

So far adding support for new cloud plugins and resources to CloudQuery required developers to implement **ET** (In **ETL** - Extract, Transform, Load).

Now, CloudQuery SDK means you as a developer will only have to implement the **E** (in ETL), and the SDK will take care of the rest.
Also, you will benefit from easy testable code, new features like policy packs and others that your plugins will get out of the box as the SDK develops.

Full Documentation is available [here](https://www.cloudquery.io/docs/developers/architecture).

For a quick snippet continue reading!

### Architecture Overview

CloudQuery has a pluggable architecture and is using gRPC to load, run and communicate between plugins.

![](/images/blog/cloudquery-sdk-architecture.png)

### Example

Here is a snippet of how an AWS resource implementation looks like with the new SDK

```go
func FlowLogs() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_flow_logs",
		Resolver:    fetchEc2FlowLogs,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveFlowLogArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "deliver_cross_account_role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DeliverCrossAccountRole"),
			},
			...
}

func fetchEc2FlowLogs(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeFlowLogsInput
	c := meta.(*client.Client)
	svc := c.Services().Ec2
	for {
		output, err := svc.DescribeFlowLogs(ctx, &config)
		if err != nil {
			return err
		}
		res <- output.FlowLogs
		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
```

Essentially you have to implement two things:

1. Define how the SQL table will look like and CloudQuery will automatically get the data from the 3rd party SDK struct via default naming convention (which you can override)
2. Implement the extract part of fetching the data via 3rd party SDK (In this case AWS SDK).

More Documentation available at:

- [Docs](/docs/developers/creating-new-plugin)
- [AWS provider](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws)

### What's next?

More plugins, integrations and features coming up!
Subscribe to our mailing list for updates and hit that [star](https://github.com/cloudquery/cloudquery) button on our GitHub!
