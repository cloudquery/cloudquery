package firebase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
)

const (
	CloudQueryRegistryURL              = "https://firestore.googleapis.com/v1/projects/hub-cloudquery/databases/(default)/documents/orgs/"
	CloudQueryRegistryURLWithProviders = CloudQueryRegistryURL + "%s/providers/%s"
	providersVersionsPath              = "%s/providers/%s/versions"
	policiesVersionPath                = "%s/policies/%s/versions"
)

type FirebaseClient struct {
	url string
}

func NewFirebaseClient(registryUrl string) *FirebaseClient {
	f := &FirebaseClient{
		url: registryUrl,
	}

	return f
}

func (f *FirebaseClient) GetLatestProviderRelease(ctx context.Context, organization, providerName string) (string, error) {
	versions, err := url.Parse(fmt.Sprintf(f.url+providersVersionsPath, organization, providerName))
	if err != nil {
		return "", err
	}
	qv := versions.Query()
	qv.Set("pageSize", "1")
	qv.Set("orderBy", "v_major desc, v_minor desc, v_patch desc, published_at desc")
	qv.Set("mask.fieldPaths", "tag")
	versions.RawQuery = qv.Encode()

	hc := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, versions.String(), nil)
	res, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	var doc struct {
		Documents []struct {
			Name   string `json:"name"`
			Fields struct {
				Tag struct {
					Val string `json:"stringValue"`
				} `json:"tag"`
			} `json:"fields"`
		} `json:"documents"`
	}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return "", err
	}

	if len(doc.Documents) == 0 || doc.Documents[0].Fields.Tag.Val == "" {
		return "", fmt.Errorf("failed to find provider %s latest version", providerName)
	}
	return doc.Documents[0].Fields.Tag.Val, nil
}

func (f *FirebaseClient) GetLatestPolicyRelease(ctx context.Context, organization, policyName string) (string, error) {
	versions, err := url.Parse(fmt.Sprintf(f.url+policiesVersionPath, organization, policyName))
	if err != nil {
		return "", err
	}
	qv := versions.Query()
	qv.Set("mask.fieldPaths", "policy.Name")
	versions.RawQuery = qv.Encode()

	hc := &http.Client{Timeout: 15 * time.Second}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, versions.String(), nil)
	res, err := hc.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", res.StatusCode)
	}

	var doc struct {
		Documents []struct {
			Name string `json:"name"`
		} `json:"documents"`
	}

	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return "", err
	}

	valid := make([]*version.Version, 0)
	for _, d := range doc.Documents {
		nameParts := strings.Split(d.Name, "/")
		if v, err := version.NewSemver(nameParts[len(nameParts)-1]); err == nil {
			valid = append(valid, v)
		}
	}

	sort.SliceStable(valid, func(i, j int) bool {
		return valid[i].GreaterThan((valid[j]))
	})

	if len(valid) == 0 || valid[0].Original() == "" {
		return "", fmt.Errorf("failed to find policy %s latest version", policyName)
	}
	return valid[0].Original(), nil
}
