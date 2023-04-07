package legacy

import (
	"context"
	"encoding/json"
	"net/url"
)

// ListReportingIssuesResponse represents the response from the Snyk API when
// listing issues.
type ListReportingIssuesResponse struct {
	Results []ListReportingIssueResult `json:"results"`
	Total   int                        `json:"total"`
}

// ListReportingIssueResult represents a single issue in the response from the
// Snyk API.
type ListReportingIssueResult struct {
	Issue          ListReportingIssue          `json:"issue"`
	Projects       []ListReportingIssueProject `json:"projects"` // When groupBy is used
	Project        ListReportingIssueProject   `json:"project"`  // When groupBy is not used
	IsFixed        bool                        `json:"isFixed"`
	IntroducedDate string                      `json:"introducedDate"`
	PatchedDate    string                      `json:"patchedDate"`
	FixedDate      string                      `json:"fixedDate"`
}

// ListReportingIssue represents an issue in the response from the Snyk API.
type ListReportingIssue struct {
	URL                  string   `json:"url"`
	ID                   string   `json:"id"`
	Title                string   `json:"title"`
	Type                 string   `json:"type"`
	Package              string   `json:"package"`
	Version              string   `json:"version"`
	Severity             string   `json:"severity"`
	OriginalSeverity     string   `json:"originalSeverity"`
	UniqueSeveritiesList []string `json:"uniqueSeveritiesList"`
	ExploitMaturity      string   `json:"exploitMaturity"`
	IsUpgradable         bool     `json:"isUpgradable"`
	IsPatchable          bool     `json:"isPatchable"`
	IsPinnable           bool     `json:"isPinnable"`
	JiraIssueURL         string   `json:"jiraIssueUrl"`
	PublicationTime      string   `json:"publicationTime"`
	DisclosureTime       string   `json:"disclosureTime"`
	Language             string   `json:"language"`
	PackageManager       string   `json:"packageManager"`
	Identifiers          struct {
		CVE   []string `json:"CVE"`
		CWE   []string `json:"CWE"`
		OSVDB []string `json:"OSVDB"`
	}
	Credit        []string `json:"credit"`
	CVSSv3        string   `json:"CVSSv3"`
	PriorityScore int      `json:"priorityScore"`
	CVSSScore     float64  `json:"CVSSScore"`
	Patches       []struct {
		ID               string   `json:"id"`
		ModificationTime string   `json:"modificationTime"`
		Urls             []string `json:"urls"`
		Comments         string   `json:"comments"`
		Version          string   `json:"version"`
	}
	IsIgnored bool `json:"isIgnored"`
	IsPatched bool `json:"isPatched"`
	Semver    struct {
		Vulnerable []string `json:"vulnerable"`
		Unaffected string   `json:"unaffected"`
	}
	Ignored []struct {
		Reason  string `json:"reason"`
		Expires string `json:"expires"`
		Source  string `json:"source"`
	} `json:"ignored"`
}

// ListReportingIssueProject represents a project in the response from the Snyk
// API.
type ListReportingIssueProject struct {
	URL            string `json:"url"`
	ID             string `json:"id"`
	Name           string `json:"name"`
	Source         string `json:"source"`
	PackageManager string `json:"packageManager"`
	TargetFile     string `json:"targetFile"`
}

// ListReportingIssuesRequest represents the request to the Snyk API when
// listing issues.
type ListReportingIssuesRequest struct {
	Page    int
	PerPage int    // max 1000
	SortBy  string // Possible values: severity, issueTitle, projectName, isFixed, isPatched, isIgnored, introducedDate, isUpgradable, isPatchable, priorityScore
	Order   string // example: "asc"
	GroupBy string // only allowed value is "issue"
}

// ListLatestReportingIssues lists the latest issues for the organization.
func (c *Client) ListLatestReportingIssues(ctx context.Context, req ListReportingIssuesRequest) (*ListReportingIssuesResponse, error) {
	q := url.Values{
		"page":    {string(req.Page)},
		"perPage": {string(req.PerPage)},
		"sortBy":  {req.SortBy},
		"order":   {req.Order},
		"groupBy": {req.GroupBy},
	}
	b, err := c.post(ctx, "/v1/reporting/issues/latest", q)
	if err != nil {
		return nil, err
	}
	r := &ListReportingIssuesResponse{}
	if err := json.Unmarshal(b, r); err != nil {
		return nil, err
	}
	return r, nil
}
