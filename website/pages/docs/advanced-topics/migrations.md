# Overview

This section describes how CloudQuery is dealing with schema changes in plugins.
The overall idea is to not have breaking changes, but rather always add columns, because it is common for users to build views on top which we don't want to break. Those migration tactics are usually implemented in the destination plugins as source plugins are database agnostic and just send back JSON objects.

## Addition

Column addition is the easy case where a source plugin adds additional column the destination plugin will not drop old one and will add a new one.

## Removal

If a source plugin removes a column the destination plugin will not drop it and it will be up to the user to drop it if necessary or to do any other logic.

## Rename

Rename is basically Removal + Addition, meaning the destination plugin will just add additional column and new data will be saved there. It will be up to the user to do any migrations if needed.

CloudQuery is an open-source cloud asset inventory powered by SQL, and as such, when plugins change their schema (change/remove columns) some migrations from the previous run are required. CloudQuery automatically drops and recreates those tables automatically as needed.

## Type Change

This is the most complex situation most probably and as a source plugin developer you want to avoid it if possible. On the destination plugin side, the column will be dropped an re-added.

