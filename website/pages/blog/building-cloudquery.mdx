---
title: "Building CloudQuery: High Performance Data Integration Framework in Go"
tag: engineering
date: 2022/11/28
description: >-
  Deep dive on how we built CloudQuery and the design decisions we made along the way.
author: yevgenypats
---


import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>

Data Integration is a 30+ year-old problem of extracting data from APIs, normalizing it and loading it into various different destinations such as databases, data lakes and files. In the last 20 years the problem has become acute, especially in the marketing and business analytics space, but recently also in the infrastructure world, with the explosion of cloud vendors and infrastructure services.

Even though data integration is an old problem and will never be [completely](https://twitter.com/mattrickard/status/1542193426979909634) solved (at least while SaaS services don’t get better at having consistent machine readable APIs), we believe there is a big opportunity to address this problem at scale and achieve a new local optimum, and it can be done right now.

Addressing this problem comes with significant challenges. In this post, I’ll go over these challenges, why they’re hard, and the solutions we landed on in CloudQuery’s design.

# Engineering Requirements and Challenges

Before we go over the technical and architecture decisions we made in CloudQuery, let’s cover the main challenges and requirements of an ELT framework:

- **API Coverage:** Scaling and supporting the number of APIs covered in an efficient manner (i.e less developers, more APIs supported)--leveraging the community is key.
- **Scaling Destinations:** Unlike sources, destinations are finite but the right architecture should ensure sources and destinations are decoupled correctly so new destinations will work out-of-the-box and won't require any changes in sources (otherwise development work will grow by `n*m` where `n` is number of source plugins and `m` is number of destination plugins).
- **Performance:** Running ELT workloads is compute-intensive and can get expensive, growing with the number of APIs being extracted from. Having a high-performance and concurrent architecture is key to both driving costs down and transferring information from source to destination as fast as possible.

# Technical Deep Dive

This section will discuss the most important decisions and solutions we landed on, organized by the challenges listed above.

## API Coverage

### Open Source

There is no vendor that can support all integrations in-house, thus the best and only way in our opinion is to open-source the software and give users the ability to contribute to our official plugins and develop community plugins without being blocked by a single vendor. Good recent examples that dealt with this “infinite API integrations” problem and took the open source route successfully are Terraform and Pulumi.

### Pluggable Sources

Pluggable sources are key to scale API coverage both from development and usability perspective and solve the following challenges:

- **Binary Size:** Users download only the plugins they want to use to ensure good performance and user experience without a binary that grows infinitely and will eventually cause not only binary size issues but library clashes from different services.
- **Versioning:** Given plugins are developed as standalone gRPC binaries users can use different versions of different plugins depending on their needs.
- **Independent development:** Developers can develop their own plugins in their own repositories without being blocked by a vendor (us in this case) to review PRs.

Our [pluggable system](https://www.cloudquery.io/docs/developers/architecture) is based on gRPC to ensure our plugins can be cross-platform, independent and performant.

### SDK for Source Plugins

Writing ELT code involves a lot of boilerplate code, testing code, tricky performance issues and parallelization code. We’ve built a [Go SDK](https://github.com/cloudquery/plugin-sdk/) that enables both us and other developers to focus only on the E(Extract) code while CloudQuery SDK takes care of parallelization, transformation, loading and testing.

### Code Generation

A big challenge is scaling and continuously supporting a large number of APIs. Before we dive into how we leveraged code-generation to generate CloudQuery plugins, let’s take a quick look at how client library generation for multiple languages often work.

Usually there is some intermediate language such as gRPC/GraphQL/Smithy that generates server stubs and clients automatically for multiple languages instead of manually maintaining clients for a number of languages and keeping up with the server APIs.

There were some significant advances around client library generation in the recent years that we also took advantage of.

We could take advantage of the schema definition and generate the source plugin from that, but we decided to do that one step further in the pipeline: from the Go Code.

Go is incredibly suited for Code Analysis so we created a [code generation library](https://github.com/cloudquery/plugin-sdk/tree/main/codegen) that generates CloudQuery schema from Go structs. This helps achieve a number of things:

1. Keep CloudQuery schema up-to-date with APIs automatically
2. Enable developers to use the Go Client which is built for developers and include important capabilities when talking to the service API.
3. Our current set of plugins are >80% auto-generated!

## Scaling Destinations

### Pluggable Destinations

Another key requirement is supporting multiple destinations. The number of destinations is actually limited, but if destinations won't be decoupled from sources, source plugins will need to be updated every time there is a new destination, which will require a growing `n*m` (`n` number of source plugins and `m` number of destinations plugins) number of code-changes and make development unsustainable.

Destination plugins are implemented in a similar way to source plugins as gRPC servers.

### Type System

The job of a pluggable ELT platform is to do two things:

1. Transform and normalize the data from an API
2. For each transformed field, add type data.

CloudQuery transforms every single field to its own rich type system, which contains more than [17 types](https://github.com/cloudquery/plugin-sdk/tree/main/schema) (including things like IP Addresses, MAC Addresses). This ensures all the validation is happening in the transformation phase. Destination plugins then only need transform this, depending on what types each destination supports. This is a big shift from how Singer or Airbyte works. These systems use what is available in JSON and [JSON Schema](https://json-schema.org/), as they couple the encoding together with the type system.

Because of these two design decisions, CloudQuery already supports [5 destinations](/docs/plugins/destinations/overview) only two months after our initial [V1 release](https://www.cloudquery.io/blog/cloudquery-v1-release).

## Performance

Cloud Infrastructure hyperscale created new challenges for the ELT world. For example, some companies have more than 10,000 GCP and AWS accounts, with more than 100 million resources in total. How do you keep your inventory of these up-to-date, and fetch them on a daily (or twice daily) basis? To address this we made two architecture decisions:

### Concurrency and Scheduling Model

As part of our SDK, which provides the concurrency model for all the source plugins, we took advantage of the excellent concurrency support in Go. Go's goroutines enable us to scale to tens of thousands of API calls with a low number of OS threads and low memory overhead.

The second thing we did here is to create a good scheduler. Extracting APIs in a concurrent manner involves restrictions on memory. This is similar to the classic computer science problem of crawling, and how to efficiently crawl the web when you don’t know the number of links in each “depth”. CQ SDK is built in such a way that can support multiple scheduling algorithms and mechanisms with the default one being a concurrent DFS algorithm which makes sure memory usage is being kept at a given limit while taking advantage of all compute resources available for concurrency.

### Horizontal Scaling

The previously described concurrency and scheduling algorithm gives a good utilization of compute and memory for one CQ process. This makes it easy to scale vertically, but what about horizontal scaling? To be able to scale both vertically and horizontally, we designed CQ as stateless - i.e it doesn’t have any backend and it runs with just the right credentials and configuration file. This gives the ability for users to split configuration (for example configuration per account) and run CQ on as many nodes as needed without any need for orchestration.

# Future

This gives a short walkthrough on the main design decisions we took in CloudQuery. If anything is of particular interest, feel free to reach out to us on [GitHub](https://github.com/cloudquery/cloudquery) or [Discord](https://cloudquery.io/discord) and I’ll be happy to write a follow-up blog-post.
