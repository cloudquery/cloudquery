#!/bin/bash
set -e

# Those are packages that were deprecated but still show up in the proxy

packages_re="github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/[a-z-]+/[a-z]+"

# Those are deprecated packages
bad_packages_re="github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/web/armweb|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtestservice/armloadtestservice|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsans/armelasticsans|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"

# Those are packages that are not used at least as far as I can tell.
# Also there is not much value for packages that talk to "partner" apis which can just grow exponentially
# I've no idea why there are in the API in the first place.
not_relevant_packages_re="github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood|\
github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"

# Get the content of the specified URL using curl
content=$(curl -s https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk)

# Use grep to search for a string that matches a regular expression
# and exclude a list of predefined strings
grep -oE "$packages_re" <<< "$content" | grep -v -E "$bad_packages_re" | xargs go get -u

# go mod tidy