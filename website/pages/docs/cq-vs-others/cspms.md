# CloudQuery vs CSPMs (Cloud Security Posture Management)

CSPMs are set of closed-source enterprise products to monitor, detect and response to threats and mis-configurations in the cloud.

**Key Differences:**

- **Open Source Pluggable Architecture**: CloudQuery has an open-source [pluggable architecture](https://hub.cloudquery.io) which means you can contribute missing resources to existing plugins or you can create your own plugins in order to grab data from proprietary APIs or other SaaS applications.
- **Database Agnostic and Raw Access to data**: CloudQuery supports multiple databases such as PostgreSQL, BigQuery and others. This makes it play nicely with the whole SQL eco-system and gives you the ability to re-use other tools like Grafana/BI. AWS Config is using a proprietary subset of SQL and database and thus doesn't give you the ability to re-use other tools easily.
- **Policy Language**: CloudQuery uses standard SQL as the query engine to define rules. CSPMs use custom query languages and closed source databases.
- **Pricing**: CloudQuery is open-source and thus you will pay only for the hosting of your PostgreSQL (you can use RDS, or any other managed version on whatever cloud provider you would like) and the compute for running [CQ binary](../deployment/overview). Nearly all CSPMs require long term contracts and have variable pricing which makes large accounts and organizations expensive to monitor.
