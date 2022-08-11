---
title: AWS SSO Tutorial with Google Workspace (Gsuite) as an IdP Step-by-Step
tag: security
date: 2021/10/26
description: >-
  A tutorial walking you through setting up AWS Single Sign-On withGoogle
  Workspace as an IdP
author: mikeelsmore
---

import { BlogHeader } from "../../components/BlogHeader"

<BlogHeader/>


AWS SSO and AWS Organization were released around 2017 and are probably the best way to manage AWS access at scale.

> "AWS Single Sign-On (SSO) is a cloud SSO service that makes it easy to centrally manage SSO access to multiple AWS accounts and business applications. It enables users to sign in to an AWS IAM user with their existing corporate credentials and access all of their assigned accounts and applications from one place."
> _Quote From AWS SSO page_

This is a huge security and operational win, some highlights:

- No need to rotate another new password in AWS IAM
- 2FA is already managed at your IdP (Google Workspace (G Suite)/Okta/AzureAD) level
- When a user is leaving an organization he is automatically removed access from the organization
- Easily automate the provisioning of AWS access when a user joins an organisation or department

In this article we, will go through a step-by-step guide to set-up AWS SSO with Google Workspace (previously G Suite) as an IdP. If you are using Google Workspace and use it as your central directory, this is the guide for you.


## Prerequisite

You should have the [AWS Organization](https://aws.amazon.com/organizations/) (If you are not using it, This service combined with AWS SSO is a real game changer) set-up.

You need to sign up from the main account (also called **"management account"** ) and with enough permissions (usually Administrator permissions).

You will also need to make sure that you have access to the Google Workspace Admin and the relevant permissions to manage it.

### Setting up AWS

Now that you have all the relevant permissions, everything is ready to configure for AWS SSO. Here is the step by step to set it all up:

1. From within the [AWS Management Console](http://console.aws.amazon.com/) search so `single sign on` and go to the [AWS Single Sign-on](https://console.aws.amazon.com/singlesignon).

   ![Find AWS Single Sign-on](/images/blog/aws-sso-gsuite/image26.png)

2. Once on the service page, click the `Enable AWS SSO` button to start the service. This will take a few moments to complete.

   ![Enable AWS Single Sign-on](/images/blog/aws-sso-gsuite/image29.png)

3. Now that SSO is enabled, we need to change from the AWS directory to using an external provider. Select `Choose your identity source`.

   ![Choose your identity source](/images/blog/aws-sso-gsuite/image1.png)

4. Within the Settings page, select `Change` under the `Identity source` section.

   ![Identity source](/images/blog/aws-sso-gsuite/image31.png)

5. Now we can change from the AWS SSO directory to an Active Directory (not what we need), or an `External identity provider` which is what we need to configure Google Workspace as the provider.

   ![External identity provider](/images/blog/aws-sso-gsuite/image28.png)

6. After you have selected `External identity provider`, scroll down to `Service provider metadata` and click `Show individual metadata values`.

   ![Show individual metadata values](/images/blog/aws-sso-gsuite/image15.png)

7. You should now be presented with three fields that you can use to configure the next step on Google Workspace in the Google Admin console.

   ![Configuration for Google Workspace](/images/blog/aws-sso-gsuite/image16.png)

8. Don't close this screen, you will need it shortly after you have done the next section.

### Google Workspace SAML setup

With the SSO URLs for our AWS organization, we can go to our Google Workspace Admin console and configure it.

1. When inside the [Google Workspace Admin console](https://admin.google.com/), go to the `Web and mobile apps` settings. You can find this in the left-hand navigation menu under `Apps`.

   ![Web and mobile apps](/images/blog/aws-sso-gsuite/image12.png)

2. Then select `Add App` from the top navigation, then `Add custom SAML app`.

   ![Add custom SAML app](/images/blog/aws-sso-gsuite/image2.png)

3. Add an `App name` for the integration, I'm using `AWS SSO` to make it easier to find later.

   ![Set the App name](/images/blog/aws-sso-gsuite/image18.png)

4. We suggest you download the Google IdP metadata ready to put it back into AWS, this is under `Option 1: Download IdP metadata`.

   ![Option 1: Download IdP metadata](/images/blog/aws-sso-gsuite/image5.png)

5. Now to add the AWS SSO URLs from earlier to configure Google Workspace to point to the correct location. The mapping of data is:

   ![Add the AWS SSO URLs](/images/blog/aws-sso-gsuite/image22.png)

   - For ACS URL, enter the AWS SSO ACS URL.
   - For Entity ID, enter the AWS SSO Issue URL.
   - For Start URL, leave the field blank.
   - For Name ID format, choose EMAIL.
   - For Name ID, choose Basic Information > Primary email.

6. We don't need to apply anything to the `Attribute mapping` settings, so you can just click `FINISH` to move forward.

   ![Click FINISH](/images/blog/aws-sso-gsuite/image24.png)

7. Once that's saved, it is time to enable it for everyone. In the `User access` section, open the settings by selecting the karat in the top right corner.

   ![User access settings](/images/blog/aws-sso-gsuite/image30.png)

8. Now that you're in the Service status screen, select `ON for everyone` and `SAVE`. This will enable the service and allow you to manage who can have access to AWS, but to configure what they can access you need to do that in AWS SSO as Google Workspace is unaware of all the possible options.

   ![ON for everyone](/images/blog/aws-sso-gsuite/image11.png)

### Adding Google Workspace configuration to AWS SSO

Now that the AWS SSO service is enabled, and the Google Workspace SAML app exists, it's time to make them talk to each other.

1. Go back to the `Change identity source` screen in AWS SSO. Scroll to the bottom and add the `GoogleIDPMetadata.xml` file you downloaded a few moments ago, then click `Next: Review`.

   ![Add GoogleIDPMetadata.xml](/images/blog/aws-sso-gsuite/image17.png)

2. To confirm this new identity source, you will need to type `ACCEPT` into the field under the warnings and then select `Change identity source`.

   ![Confirm new identity source](/images/blog/aws-sso-gsuite/image20.png)

3. And now you are done with configuring the SSO and SAML connection between AWS SSO and Google Workspace. However, you aren't quite done as you need to configure the user provisioning at this point.

   ![Completed AWS SSO configuration](/images/blog/aws-sso-gsuite/image23.png)

### Setting up Users and Permissions

As of writing this, you can't automatically sync users between AWS and Google (this is being worked on over at OpenID) so we are limited to two options; manually creating the user (which we will go through) and using [https://github.com/awslabs/ssosync](https://github.com/awslabs/ssosync) to automate the process.

To manually add users, you will want to follow these instructions.

1. Go to `Users` in the sidebar of the `AWS SSO` service. Then select `Add user`.

   ![Add user](/images/blog/aws-sso-gsuite/image4.png)

2. Use the primary Google Workplace email address as the `Username` as well as the `Email address`, and fill the other fields accordingly. Then hit `Next: Groups` to save.

   ![Use Primary Google Workspace email](/images/blog/aws-sso-gsuite/image13.png)

3. As part of this process, we aren't going to be adding groups so we can skip these phases by selecting `Add user` in the bottom right.

   ![Skip groups and add user](/images/blog/aws-sso-gsuite/image25.png)

4. The user now needs to be associated with an AWS account. So select `AWS accounts` from the left navigation, select the checkbox next to the user, and click `Assign users` to attach them to the account.

   ![Assign users](/images/blog/aws-sso-gsuite/image8.png)

5. On the next screen select the user again so that we can move on to permissions by clicking `Next: Permissions sets`.

   ![Select user to assign permissions](/images/blog/aws-sso-gsuite/image10.png)

6. As we haven’t configured and permission sets before we will have to do that now by clicking `Create new permission set`.

   ![Create new permission set](/images/blog/aws-sso-gsuite/image21.png)

7. We will be using `Use an existing job function policy`, these are like [AWS managed policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html) that you will be aware of if you have configured permissions inside AWS previously. Now `Next: Details` to select the policy.

   ![Use an existing job function policy](/images/blog/aws-sso-gsuite/image27.png)

8. From here you can select the policy you want to assign, I'll be using `AdministratorAccess` as this is for myself. But you could use `PowerUserAccess` as this would allow the user to build whatever they want, but not mess with other users and groups. Then click `Next: Tags` to apply this to the user.

   ![Select policy to use](/images/blog/aws-sso-gsuite/image7.png)

9. Tags are optional, but they are advised for auditing and search at a later point. But I don't need them so we select `Next: Review` to move forward.

   ![Tags are optional](/images/blog/aws-sso-gsuite/image6.png)

10. A quick once over to make sure everything is set correctly, then we can click `Create` to do so.

    ![Complete new permissions set](/images/blog/aws-sso-gsuite/image3.png)

11. Back to our `Assign Users` screen, we can click the refresh icon to view our permission sets. From here we select the checkbox for the permissions set we want for the user, then select `Finish` to apply it.

    ![Assign newly created permissions set](/images/blog/aws-sso-gsuite/image9.png)

12. It'll take a moment to provision, but you should get a Complete screen saying you are done.

    ![Provisioning user complete](/images/blog/aws-sso-gsuite/image14.png)

And we are done, now the user can authenticate and log in from Google Workplace using the handy link in the `Google apps` selector.

![Google Apps selector](/images/blog/aws-sso-gsuite/image19.png)

## Summary

By now you should have AWS SSO configured with Google Workspace as an IdP and you can manage access & permissions to your AWS in the AWS SSO service.
