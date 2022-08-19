---
title: Inventory Microsoft Azure with CloudQuery
tag: guest
date: 2021/11/30
description: Inventory Microsoft Azure with CloudQuery
author: giselatorres
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


[Originally posted at [https://www.returngis.net/2021/11/haz-inventario-de-microsoft-azure-con-cloudquery/](https://www.returngis.net/2021/11/haz-inventario-de-microsoft-azure-con-cloudquery/)]

I have recently discovered a tool that in different scenarios, especially in the inventory, can be useful to us. It's called [CloudQuery](https://www.cloudquery.io/) and it allows you to export the data of the resources of your subscriptions, from the different cloud providers, to be able to execute queries on it by **launching SQL statements**, since the result is stored in a Postgres. In this article I tell you how to configure it for Microsoft Azure.


## Install CloudQuery

The first thing you need to do is install the CloudQuery tool on your machine. In my case I am using Mac, so I have run the following command using [Homebrew](https://brew.sh/):

```bash
brew install cloudquery/tap/cloudquery
```

If you use another operating system you can see the different [releases here](https://github.com/cloudquery/cloudquery/releases).

## Start CloudQuery

Now that you have the CloudQuery tool installed on your machine, create a directory, I have called it **cloudquery**, and initialize the configuration within it, with the following command:

```bash
cloudquery init azure
```

By doing this, it generates a file called **config.hcl** which we can customize, indicating which subscriptions we want to bring the data from, what types of resources, and so on.

![cloudquery init generates the config.hcl file](/images/blog/inventory-microsoft-azure-with-cloudquery/cloudquery-config.hcl-file-1536x922.png 'cloudquery init generates the config.hcl file')

If we do not modify anything, all the resources of all the subscriptions to which they have access will be brought. Now we have our project ready to retrieve the information of our subscriptions.

## Create a Postgres database in Docker

As I already mentioned at the beginning of this article, the information retrieved is exported to a **Postgres**-type database , so we will need one. The simple way is using a [Docker](https://www.returngis.net/2019/02/hoy-empiezo-con-docker/) container:

```bash
#Create a database in Docker
docker run -d --name postgresdb \
-p 5432:5432 \
-e POSTGRES_PASSWORD=pass \
postgres
```

## Create a Service Principal

Now that we know that we want to retrieve the information from Azure, in order to do so we need to create a main service that has access to the subscriptions we want to export:

```bash
SUBSCRIPTION_ID=<YOUR_SUBSCRIPTION_ID>
az account set --subscription $SUBSCRIPTION_ID

#Need to register Microsoft.Security
az provider register --namespace 'Microsoft.Security'

#Create a service principal
az ad sp create-for-rbac --name cloudquery --scopes /subscriptions/$SUBSCRIPTION_ID  > auth.json
```

Once the response is created, and stored in the auth.json file , I use the `jq` tool to store the main service information in these environment variables:

```bash
#Set variables
export AZURE_TENANT_ID=$(jq -r '.tenant' auth.json)
export AZURE_CLIENT_ID=$(jq -r '.appId' auth.json)
export AZURE_CLIENT_SECRET=$(jq -r '.password' auth.json)
export AZURE_SUBSCRIPTION_ID=$(az account show --query id -o tsv)
```

## Assign the primary service to the Azure AD "Application Administrator" role

To finish the configuration, we need the main service that we just created to be associated with an Azure Active Directory role called **Application administrator**, which you can find in the Azure AD resource, in the **Roles and administrators section**:

![Azure AD - Roles and administrators - Application Administrator](/images/blog/inventory-microsoft-azure-with-cloudquery/Azure-AD-Application-administrator-role.png 'Azure AD - Roles and administrators - Application Administrator')

Once in it, look for the main service that we have called **cloudquery**, through the Add assignments button, and associate it permanently.

![Add the main service to the Application Administrator role](/images/blog/inventory-microsoft-azure-with-cloudquery/Azure-AD-Application-administrator-cloudquery-assigment-2048x542.png 'Add the main service to the Application Administrator role')

## Dump the information in Postgres with CloudQuery

Now the only thing left is to retrieve the information with CloudQuery. To do this, you only have to execute a single command:

```bash
#Fetch the information into the database
cloudquery fetch --dsn "postgres://postgres:pass@localhost:5432/postgres?sslmode=disable"
```

This uses the environment variables that we have previously configured, with the information of our main service, and the postgres that we have generated in Docker. Once the process finishes, you will see that you have a bunch of generated tables (for this example I have used [DataGrip](https://www.jetbrains.com/datagrip/) as GUI):

![92 CloudQuery generated tables in Postgres](/images/blog/inventory-microsoft-azure-with-cloudquery/92-tablas-generadas-por-CloudQuery-en-postgres-1095x1536.png '92 CloudQuery generated tables in Postgres')

You can make queries like these, simply to retrieve resources of a specific type:

```sql
SELECT * from azure_compute_virtual_machines;

SELECT * from azure_web_apps;
```

Or go further and consult about these aspects that may be important to you and you need to validate or generate a report on them. For example: "Tell me which storage accounts have public access enabled." It would be something as simple as this:

```sql
SELECT * from azure_storage_accounts where allow_blob_public_access is null
```

Greetings!
