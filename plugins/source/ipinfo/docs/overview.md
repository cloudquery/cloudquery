---
name: IPinfo
stage: GA
title: IPinfo Source Plugin
description: CloudQuery IPinfo Source Plugin documentation
---
# IPinfo Source Plugin

:badge

The CloudQuery IPinfo plugin extracts information from supported services by ipinfo.io and loads into any supported CloudQuery destination (e.g, Sqlite, and [more](/docs/plugins/destinations/overview)). It is based on the github.com/ipinfo/go library.

## Authentication

:authentication

## Configuration

The following example sets up the HubSpot plugin, and connects it to a postgresql destination:

:configuration

### IPinfo Spec

This is the specs that can be used by the IPinfo source Plugin.

- `ip` (`string`, optional):
  To lookup information of a particular IP address, for example, 8.8.8.8

- `token` (`string`, optional)

  Rate limit per second for requests done HubSpot API, this will depend on your IPinfo plan (https://ipinfo.io/pricing)