---
title: Creating a Cross Project (or Account) Service Account in GCP Step-by-Step
tag: security
date: 2021/11/09
description: >-
  A walkthrough to add Service Accounts in Google Cloud Platform and make them
  cross-project
author: mikeelsmore
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


In Google Cloud Platform (GCP) it is common to have multiple projects for different environments (like `dev`, `staging`, `prod`, `prod-team1`, etc.). It is also a common use-case to have one set of credentials (service account) to access multiple accounts, For example:

- **Auditing:** one service account with read-only access to all projects
- **Multi-project access/communication:** one service in one project might need access/communicate to other services in different projects.

In this tutorial we will show you how to create one service account in GCP that can access multiple projects either under the same organization/account or even completely different accounts (for AWS users this is the GCP's assume role equivalent).



## How do you set up a Service Account in GCP?

Assuming you’ve got your project setup (we are going to use Project A & Project B to test all this), you’ll want to navigate to Project A and then do the following steps:

1. Within the `IAM & Admin` menu select `Service Accounts`

   ![Service Accounts is under the IAM & Admin menu](/images/blog/gcp-cross-project-service-account/image3.png)

2. Select `+ CREATE SERVICE ACCOUNT`

   ![Select + CREATE SERVICE ACCOUNT](/images/blog/gcp-cross-project-service-account/image1.png)

3. Fill in the Service Accounts details, as it’s going to be used cross-projects make sure it’s clearly defined as such (you will be using the `Service account ID` later). Then select `CREATE AND CONTINUE`

   ![Fill in all the descriptive details for the account](/images/blog/gcp-cross-project-service-account/image5.png)

4. Now apply the permissions you want this Service Account to have, I’m using the `Viewer` permission, you can also add any [conditions](https://cloud.google.com/iam/docs/conditions-overview) to the permissions

   ![Select the level of permissions for the service account](/images/blog/gcp-cross-project-service-account/image11.png)

5. Once you have applied all your desired permissions to the Service Account select `CONTINUE`

   ![Click CONTINUE to move to the next step](/images/blog/gcp-cross-project-service-account/image9.png)

6. If you’d like to grant specific users access to this Service Account (for modification or to see what it’s doing) you can add those users here

   ![Grant specific users access to the service account, this is optional](/images/blog/gcp-cross-project-service-account/image10.png)

7. After adding any users you wish to grant access, select `DONE` and you should be sent to a screen with the Service Account and it’s status etc

   ![List of your service accounts](/images/blog/gcp-cross-project-service-account/image14.png)

## How does do we grant it access to other projects?

Once we have a working Service Account, we now have to go through a slightly different process to add it to other projects.

1. Firstly, using the project navigation in the top menu select your second project. In my case this is Project B

   ![Moving to the second project](/images/blog/gcp-cross-project-service-account/image12.png)

2. Like before we need to select `IAM & Admin` from the menu, be this time we select `IAM`

   ![Go to the IAM under the IAM & Admin menu](/images/blog/gcp-cross-project-service-account/image4.png)

3. Once here simply select `+ADD`

   ![Select +ADD](/images/blog/gcp-cross-project-service-account/image2.png)

4. From this new menu, you will need to use the `Service account ID` from the previous flow of creating the Service Account

   ![Use the Service account ID from creating the service account](/images/blog/gcp-cross-project-service-account/image8.png)

5. And add the role you want to have assigned to the Service Account within this Project, I’m going with `Viewer` again

   ![Select the role with permissions for the account inside this project](/images/blog/gcp-cross-project-service-account/image7.png)

6. After applying all the roles and permissions the Service Account needs, click `SAVE`

   ![Select SAVE to complete the process](/images/blog/gcp-cross-project-service-account/image6.png)

7. After the policy has updated, you’ll be able to see your user in the IAM list

   ![Once saved the Service account is listed](/images/blog/gcp-cross-project-service-account/image13.png)

## Summary

And that’s it, your Service Account created in Project A now has access to both Project A and Project B, enjoy.
