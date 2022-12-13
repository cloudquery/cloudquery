package azparser

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var packagesToSkip = map[string]bool{
	// Manually generated recipes
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault": true,
	// This is a special API and we create those recipes manually
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources": true,
	// this can be written manually and potentially we can also get it from armresources
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription": true,
	// this can be written manually and potentially we can also get it from armresources
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions": true,
	// seems something is not right with that package
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups": true,
	// skipping for now
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices":                       true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder": true,

	// deprecated packages
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsight/armsecurityinsight":                 true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/web/armweb":                                         true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtestservice/armloadtestservice":                 true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsans/armelasticsans":                         true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/videoanalyzer/armvideoanalyzer":                     true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices": true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadmonitor/armworkloadmonitor":                 true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsight":                true,

	// not relevant packages
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain": true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood":     true,
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos":           true,
}

var subpackageRe = regexp.MustCompile(`github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/[a-z-]+/[a-z]+`)

func DiscoverSubpackages() ([]string, error) {
	var subpackages = make([]string, 0)
	resp, err := http.Get("https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	matches := subpackageRe.FindAllString(string(body), -1)
	for _, match := range matches {
		if _, ok := packagesToSkip[match]; !ok {
			subpackages = append(subpackages, match)
		}
	}

	return subpackages, nil
}
