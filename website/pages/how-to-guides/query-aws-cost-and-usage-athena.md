---
title: How to Query AWS Cost and Usage Data with Amazon Athena
tag: tutorial
date: 2022/12/15
description: >-
  
author: jsonkao
---
import { BlogHeader } from "../../components/BlogHeader"
import { HowToGuideHeader } from "../../components/HowToGuideHeader"

<HowToGuideHeader/>

## Introduction

In a [previous guide](www.cloudquery.io/how-to-guides/how-to-load-infrastructure-data-into-athena), we talked about how to load infrastructure data into Amazon Athena.  In this tutorial, we will show you how to load AWS's Cost and Usage Reports into Amazon Athena and enrich infrastructure data with the data provided by Cost and Usage Reports.  This will enable you to query your data and execute different cost optimization and FinOps use cases to understand and manage the financials of your cloud infrastructure.

## Steps

### 1. Infrastructure Data Prerequisites

We've already followed the previous guide to load infrastructure data into Amazon Athena via CloudQuery.  If you haven't done so already, follow the [CloudQuery guide](www.cloudquery.io/how-to-guides/how-to-load-infrastructure-data-into-athena).  

### 2. Start Gathering Cost and Usage Reports

To set up Cost and Usage Reports, we'll follow AWS documentation for how to set up the infrastructure necessary for Cost and Usage Reports.

By the end of this step, we'll have created the following:
* A S3 Bucket to hold our Cost and Usage Reports.
* A new Cost and Usage Report.

Let's start by [creating a new S3 bucket](https://docs.aws.amazon.com/cur/latest/userguide/cur-s3.html) for our Cost and Usage Reports.  AWS's `billingreports.amazonaws.com` will deliver these reports to our new S3 Bucket.  Keep in mind that it may take up to 24 hours for AWS to start delivering reports to the Amazon S3 bucket.  If you are creating a fresh S3 bucket, we suggest creating the bucket and coming back the next day to complete the rest of this guide!  Storing the billing reports data in Amazon S3 will incur standard [S3 costs](https://docs.aws.amazon.com/cur/latest/userguide/billing-cur-limits.html).  



https://docs.aws.amazon.com/cur/latest/userguide/cur-query-athena.html

### 3. Sync Cost and Usage Reports to Athena


### 4. Optional: Set up multi account Cost and Usage Report informatino


### 5. Run Athena Queries on Infrastructure and Cost Data

Let's start by looking at our biggest expenses:

** Show biggest expenses and correlate with resources.

Now let's look at one straightforward method to reduce costs by using more cost-effective EBS volume types.

** Use gp3 volumes instead of gp2 volumes for Amazon EBS.
https://aws.amazon.com/blogs/storage/migrate-your-amazon-ebs-volumes-from-gp2-to-gp3-and-save-up-to-20-on-costs/

** Correlate Groups of Resources based off Tags

** Search by a CloudFormation Stack (Maybe)

## Summary

If you have comments or questions about using CloudQuery for Cloud Cost Management and Optimization, we would love to hear from you! Reach out to us on [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://cloudquery.io/discord)!

## References

[CloudQuery: AWS Source Plugin](https://www.cloudquery.io/docs/plugins/sources/aws/overview)

[AWS: Querying Cost and Usage Reports using Amazon Athena](https://docs.aws.amazon.com/cur/latest/userguide/cur-query-athena.html)

[AWS: What are AWS Cost and Usage Reports](https://docs.aws.amazon.com/cur/latest/userguide/what-is-cur.html)

[AWS Blogs: Starting your Cloud Financial Management journey: Cost visibility](https://aws.amazon.com/blogs/aws-cloud-financial-management/op-starting-your-cloud-financial-management-journey-cost-visibility/)


