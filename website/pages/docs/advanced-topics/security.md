---
title: Security
description: This section will list key security points regarding CloudQuery. Make sure you follow best practices if you decide to "host it yourself."
---

# Security

This section will list key security points regarding CloudQuery. Make sure you follow best practices if you decide to "host it yourself."

## Integration Authentication Credentials

- Integration Authentication Credentials should always be read-only.
- The machine where CloudQuery is running should be secured with the correct permissions, as it contains the credentials to your cloud infrastructure.

## CloudQuery Database

Even though in most cases the CloudQuery database contains only configuration and meta-data, you should protect it and keep it secure with correct access and permissions.
