# What is ASM?

As cyber threats continue to increase in complexity and frequency, it's more important than ever for businesses to adopt effective cybersecurity measures. One critical component of cybersecurity is Attack Surface Management (ASM). In this post we will talk shortly about what Attack Surface Management (ASM) is, and how to build an open source ASM on top of your infrastructure data lake.

Attack Surface Management is the process of identifying and managing all the potential points of attack on a company's IT infrastructure. This includes everything from networks and servers to applications and endpoints. By identifying these vulnerabilities, businesses can take steps to reduce their risk of cyber attacks.

## Why is Attack Surface Management important?

There are several reasons why Attack Surface Management is a crucial component of a robust cybersecurity strategy, including:

1. Risk Reduction - ASM helps businesses identify and mitigate potential vulnerabilities before they can be exploited by cybercriminals, reducing the risk of successful attacks.
2. Compliance - Many regulatory requirements and industry standards require businesses to conduct regular assessments of their attack surface to ensure that they are in compliance.
3. Cost Savings - By identifying and mitigating potential vulnerabilities early on, businesses can save money on potential damages and loss of revenue associated with cyber attacks.

## How does Attack Surface Management work?

ASM typically involves several key steps, including:

1. Discovery - The first step in ASM is to discover all of the potential points of attack on a company's IT infrastructure. This may include conducting vulnerability scans, network mapping, and other assessments.
2. Analysis - Once the attack surface has been identified, businesses can analyze each potential vulnerability to determine its severity and potential impact on the organization.
3. Remediation - After analyzing the potential vulnerabilities, businesses can take steps to remediate them. This may involve implementing new security measures, updating software, or patching vulnerabilities.
4. Monitoring - Attack Surface Management is an ongoing process, and businesses must monitor their attack surface regularly to identify new vulnerabilities and mitigate them as quickly as possible.

## Open Source ASM

In [What is an Infrastructure Data Lake](/docs/glossary/what-is-infrastructure-data-lake) we covered what an Infrastructure Data Lake is and how to build one. With an infrastructure data lake in place, you can easily build your own customizable ASM with just standard SQL or Cypher (Neo4j) queries and views that you can monitor and visualize with your go-to BI tools, avoiding yet-another-dashboard fatigue and the need for learning new proprietary query languages. Check out our full guide on how to use CloudQuery, pre-built queries, views and neo4j dashboards to build an [open-source attack surface management](/how-to-guides/attack-surface-management-with-graph).