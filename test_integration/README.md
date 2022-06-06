# Snapshots tests

Tests in this directory use a testing methodology called snapshot testing to verify the output of the CloudQuery CLI.
On first run of the test, a snapshot of the CLI output is created and on sequential runs the CLI output is compared to the snapshot to verify the output is the same.

If the tests fail, and the output change is expected, update the snapshots by running `UPDATE_SNAPSHOTS=true go test ./...`.

> Snapshot testing is useful for specific cases where the output is not expected to change very often
