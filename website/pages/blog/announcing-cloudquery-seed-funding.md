---
title: Announcing CloudQuery Seed Funding
tag: announcement
date: 2021/11/11
description: CloudQuery raises $3.5 million in seed funding within our first year
author: yevgenypats
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


Today, we announced that CloudQuery has raised $3.5 million in seed funding led by [Boldstart Ventures](https://boldstart.vc/), with participation from [work-bench](https://work-bench.com/), [Mango Capital](https://www.linkedin.com/company/mangocapital/) and [Haystack](https://haystack.vc/). We’d love to share a bit about CloudQuery, our journey, and what the future holds for our open-source project and the community.

[CloudQuery](https://github.com/cloudquery/cloudquery) is the open-source cloud asset inventory powered by SQL, enabling developers to assess, audit, and monitor the configuration of your cloud assets.


## CloudQuery Story

CloudQuery came into the world from our personal pain. As developers in various roles we interacted with cloud infrastructure across security, SRE, DevOps, development. Often for both basic and complex tasks we couldn’t find the right abstraction layer and always had to fallback to writing the same code over and over again. Just a few examples to keep it short but concrete:

- **Visibility:** We had to look up in which project the ec2 instance we were debugging is located, as well as some of its meta-data. This usually involved us talking to the production engineer, who in turn ran their secret python scripts against the production environment and returned us the answer.
- **Security:** Every company has its own set of security & compliance rules, and every company’s security team ends up maintaining a bulk of custom python/bash scripts.
- **Monitoring:** In every company we worked at, we had products and workflows in place for visualization, monitoring, alerting (such as Grafana). As developers, we wanted to use the same products and workflows for our cloud assets, rather than learn and maintain yet-another-dashboard and alerting systems.

The common pattern that you are probably starting to see here:

1. You talk to someone with the right access.
2. You (or someone with enough access) run a bunch of scripts that also have a common pattern:
   1. Extract cloud asset configuration.
   2. Transform and normalize them to some workable data structure.
   3. Run your business logic on the normalized, transformed data.

It became clear that for every task we do 80% of the time is spent on similar busywork and only 20% of the time is spent on the actual business objectives. This is why we created CloudQuery - to eliminate all this developer toil, and create an open-source cloud asset inventory powered by standard SQL.

## What we achieved in the last 11 months

- Our [GitHub repository](https://github.com/cloudquery/cloudquery) grew to over 1.8k stars and over 40 contributors across all our repositories.
- [CloudQuery SDK](https://www.cloudquery.io/blog/introducing-cloudquery-sdk) - Helps developers write their own providers (plugins) to CloudQuery.
- [CloudQuery Policies](https://www.cloudquery.io/blog/announcing-cloudquery-policies) - Turns security & compliance rules into a data problem instead of a code problem.
- [CloudQuery Hub](https://hub.cloudquery.io/) - Hosts both official and community CloudQuery plugins and policies, thus growing the eco-system and enabling knowledge sharing.
- [Providers](https://hub.cloudquery.io) - We grew our official cloud provider coverage across all 3 big cloud providers (AWS, GCP, Azure), DigitalOcean, Kubernetes, and Okta. We also welcomed our first community provider - Thanks, Yandex Cloud!

## What the future holds

We are going to focus on expanding CloudQuery platform both in terms of in-depth features and breadth integrations, and, most importantly, building a vibrant community which we believe is a key to the success of the project. In terms of features, We won't spoil some of our news in the upcoming weeks so stay tuned as we have exciting stuff coming up!

## Word of love to community

As already noted CloudQuery couldn’t and can’t exist without the community which means that everyone who’s ever fixed a bug, updated docs, written a feature, opened an issue or a feature request deserves a big thank you! And of-course to our team who is working on making CloudQuery what it is every single day.

## Growing our team

For CloudQuery to become the open-source standard with the big eco-system and roadmap that we envision, we are looking for smart people to join us. So if you are into working on open-source projects, Go, Security, and Cloud, please drop us a line at jobs [at] our domain.

## Getting started with CloudQuery

- Start using CloudQuery today: [github.com/cloudquery/cloudquery](https://github.com/cloudquery/cloudquery)
- Follow us on [Twitter](https://twitter.com/cloudqueryio)
- Subscribe to our newsletter 👇
