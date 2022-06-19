#!/bin/bash
set -e

echo "Running sanity checks"

echo "Running fetch"
go run ./main.go fetch --config=internal/test/test_config.hcl --enable-console-log

echo "Fetch Multiple Provider"
go run ./main.go fetch --config=internal/test/test_double_provider_config.hcl --enable-console-log

echo "Init"
go run ./main.go init test --config=test_init_config.hcl

echo "Policy Describe"
go run ./main.go policy describe k8s//nsa_cisa_v1/pod_security --config=internal/test/test_config.hcl

echo "Policy Download"
go run ./main.go policy download github.com/roneli/aws//cis_v1.2.0 --config=internal/test/test_config.hcl

echo "Policy Run bad subpath"
go run ./main.go policy run aws//path/not/exist --config=internal/test/test_aws.hcl --disable-fetch-check && echo "test: 'Policy Run bad subpath' failed" && exit 1

echo "Sync (upgrade)"
go run ./main.go provider upgrade test --config=internal/test/test_config.hcl

echo "Purge"
go run ./main.go provider purge test --config=internal/test/test_config.hcl

echo "Sync (downgrade)"
go run ./main.go provider downgrade test --config=internal/test/test_config.hcl

echo "Drop"
go run ./main.go provider drop test --config=internal/test/test_config.hcl --force