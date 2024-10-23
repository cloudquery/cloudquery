---
title: Instrumenting a Paid Integration
description: Learn how to modify your integration code to publish a paid (premium or open-core) integration
---

# Instrumenting a Paid Integration

This page is aimed at integration developers. CloudQuery integrations can be published as free, open-core or premium. In order for rows to be counted as paid in an open-core or premium integration, you will need to add some additional instrumentation code. Instrumenting a paid integration to check quotas and count the number of rows synced is relatively simple and can be done using the [`github.com/cloudquery/plugin-sdk/v4/premium`](http://github.com/cloudquery/plugin-sdk/v4/premium) package.

## Steps

1. Ensure that the integration’s team, name and kind are passed in. For example:

   ```go
   var (
       Name    = "your-integration-name" // TODO: replace with your integration name
       Kind    = "source"           // TODO: replace with your integration kind (source / destination)
       Team    = "your-team-name"   // TODO: replace with your team name
       Version = "development"
   )
   
   func Plugin() *plugin.Plugin {
       return plugin.NewPlugin(
           Name,
           Version,
           Configure,
           plugin.WithKind(Kind),
           plugin.WithTeam(Team),
       )
   }
   ```

2. Inside `resources/plugin/client.go`, add `usage premium.UsageClient` to the `Client` struct.
3. Instantiate the `premium.UsageClient` inside `Configure`:

   ```go
   uc, err := premium.NewUsageClient(
         opts.PluginMeta,
         premium.WithLogger(logger),
     )
     if err != nil {
         return nil, fmt.Errorf("failed to initialize usage client: %w", err)
     }
 
   return &Client{
     // ...
         usage:  uc,
     // ...
   }
   ```

4. Add the following methods to the `Client`:

   ```go
   // OnBeforeSend increases the usage count for every message. If some messages should not be counted,
   // they can be ignored here.
   func (c *Client) OnBeforeSend(_ context.Context, msg message.SyncMessage) (message.SyncMessage, error) {
       if c.usage == nil {
           return msg, nil
       }
       if si, ok := msg.(*message.SyncInsert); ok {
           if err := c.usage.Increase(uint32(si.Record.NumRows())); err != nil {
               return msg, fmt.Errorf("failed to increase usage: %w", err)
           }
       }
       return msg, nil
   }
   
   // OnSyncFinish is used to ensure the final usage count gets reported
   func (c *Client) OnSyncFinish(_ context.Context) error {
      if c.usage != nil {
         return c.usage.Close()
      }
      return nil
   }
   ```

5. Inside the `Client` `Sync` method, create a new context using `premium.WithCancelOnQuotaExceeded`. This will do two things: 1. stop the sync from happening if the user has no remaining quota, and 2. periodically check that the user still has remaining quota, canceling the context if not.

   ```go
   newCtx, err := premium.WithCancelOnQuotaExceeded(ctx, c.usage)
   if err != nil {
       return fmt.Errorf("failed to configure quota monitor: %w", err)
   }
   return c.scheduler.Sync(newCtx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID))
   ```

   If there is a `stateClient` the above block should read:

   ```go
   newCtx, err := premium.WithCancelOnQuotaExceeded(ctx, c.usage)
   if err != nil {
       return fmt.Errorf("failed to configure quota monitor: %w", err)
   }
   if err := c.scheduler.Sync(newCtx, schedulerClient, tt, res, scheduler.WithSyncDeterministicCQID(options.DeterministicCQID)); err != nil {
       return fmt.Errorf("failed to sync: %w", err)
   }
   return stateClient.Flush(ctx)
   ```

6. If all tables are paid: `return premium.MakeAllTablesPaid(tables)` in `getTables`.
If only some tables are paid: add `isPaid: true` to the relevant Table definitions.
