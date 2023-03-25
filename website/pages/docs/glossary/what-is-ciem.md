# What is CIEM?

In this blog we will talk shortly about what a CSPM is and how to build an open source CIEM on top of your infrastructure data lake.

# Overview

As more and more businesses move their operations to the cloud, the need for robust security measures has become increasingly apparent. One of the most critical aspects of cloud security is managing access to cloud resources. That's where Cloud Infrastructure Entitlement Management (CIEM) comes in.

CIEM is a set of tools and processes designed to help businesses manage access to their cloud resources. This includes managing user identities, permissions, and entitlements. The goal of CIEM is to ensure that only authorized users have access to cloud resources and that they have the appropriate level of access to perform their jobs.

CIEM solutions typically include several key components, such as:

1. Identity and Access Management (IAM) - IAM is a critical component of CIEM, as it allows businesses to manage user identities and permissions across different cloud environments.
2. Entitlement Management - This component allows businesses to manage access to specific resources, such as data or applications, based on the user's role or job function.
3. Compliance Monitoring - CIEM solutions also typically include compliance monitoring tools that help businesses ensure that they are meeting regulatory requirements and industry standards.

# Open Source CIEM

Previously we covered [what is the modern data stack](https://www.cloudquery.io/blog/what-is-the-modern-data-stack) and how we can build an infrastructure data lake. By having an infrastructure data lake you can easily build your own customizable CIEM with just standard SQL queries and views that you can monitor and visualize with your go-to BI tools and avoid the yet-another-dashboard fatigue and learning new proprietary query languages. Check out our full guide on how to use CloudQuery, pre-built queries, views and Grafana dashboards to build an open-source CIEM (a TODO blog)