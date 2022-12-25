# Creating Connected Guide

Full documentation is available [here](https://help.salesforce.com/articleView?id=connected_app_create.htm&type=5). This guide will walk you through the steps to create a connected app.

## Pre-requisites

- Salesforce developer account to be able to access the REST APIs
- A Salesforce user with the permissions to create connected a

## Step 1: Go To Connected App

Go to Setup Screen and search for "App Manager" and click on "New Connected App".

![step1](/images/docs/sf/step1.png)

## Step 2: Create Connected App

Fill in the fields as the following and click `save`:

![step2](/images/docs/sf/step2.png)

## Step 3: Copy Consumer Key and Consumer Secret

Now that this is configured click `Manage Consumer Details` to view your Consumer Key and Consumer Secret (those will be used as `client_id` and `client_secret` to configure your CloudQuery Salesforce source plugin).

![step3](/images/docs/sf/step3.png)

And copy the Consumer Key and Consumer Secret somewhere safe so you can provide them as environment variables to the Salesforce plugin.

![step4](/images/docs/sf/step4.png)

## Step 4: Final Configuration

Depending on where you run CloudQuery your might need to change `IP Relxation` to `Relax IP restrictions` or configure the IP address of the machine you run CloudQuery on.

![step5](/images/docs/sf/step5.png)


