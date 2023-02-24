package homebrew

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Days string

const (
	Days30  Days = "30d"
	Days90  Days = "90d"
	Days365 Days = "365d"
)

func (d Days) String() string {
	return string(d)
}

const layout = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	f := d.Time.Format(layout)
	return []byte(fmt.Sprintf("\"%s\"", f)), nil
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return err
	}
	d.Time, err = time.Parse(layout, s)
	return err
}

type InstallsResponse struct {
	Category   string `json:"category"`
	TotalItems int    `json:"total_items"`
	TotalCount int    `json:"total_count"`
	StartDate  Date   `json:"start_date"`
	EndDate    Date   `json:"end_date"`
	Items      []struct {
		Number  int    `json:"number"`
		Formula string `json:"formula"`
		Count   string `json:"count"`
		Percent string `json:"percent"`
	}
}

type InstallItem struct {
	Number  int     `json:"number"`
	Formula string  `json:"formula"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type Installs struct {
	Category   string `json:"category"`
	TotalItems int    `json:"total_items"`
	TotalCount int    `json:"total_count"`
	StartDate  Date   `json:"start_date"`
	EndDate    Date   `json:"end_date"`
	Items      []InstallItem
}

func (c *Client) GetInstalls(ctx context.Context, days Days) (Installs, error) {
	path := fmt.Sprintf("/api/analytics/install/%s.json", days)
	return c.getInstalls(ctx, path)
}

func (c *Client) GetInstallOnRequestEvents(ctx context.Context, days Days) (Installs, error) {
	path := fmt.Sprintf("/api/analytics/install-on-request/%s.json", days)
	return c.getInstalls(ctx, path)
}

type CaskInstallsResponse struct {
	Category   string `json:"category"`
	TotalItems int    `json:"total_items"`
	TotalCount int    `json:"total_count"`
	StartDate  Date   `json:"start_date"`
	EndDate    Date   `json:"end_date"`
	Items      []struct {
		Number  int    `json:"number"`
		Cask    string `json:"cask"`
		Count   string `json:"count"`
		Percent string `json:"percent"`
	}
}

type CaskInstallItem struct {
	Number  int     `json:"number"`
	Cask    string  `json:"cask"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type CaskInstalls struct {
	Category   string `json:"category"`
	TotalItems int    `json:"total_items"`
	TotalCount int    `json:"total_count"`
	StartDate  Date   `json:"start_date"`
	EndDate    Date   `json:"end_date"`
	Items      []CaskInstallItem
}

func (c *Client) getInstalls(ctx context.Context, path string) (Installs, error) {
	r, err := c.get(ctx, path)
	if err != nil {
		return Installs{}, err
	}
	defer r.Close()
	var resp InstallsResponse
	d := json.NewDecoder(r)
	err = d.Decode(&resp)
	if err != nil {
		return Installs{}, err
	}
	installs := Installs{
		Category:   resp.Category,
		TotalItems: resp.TotalItems,
		TotalCount: resp.TotalCount,
		StartDate:  resp.StartDate,
		EndDate:    resp.EndDate,
		Items:      make([]InstallItem, len(resp.Items)),
	}
	for i, item := range resp.Items {
		count, _ := strconv.Atoi(strings.ReplaceAll(item.Count, ",", ""))
		percent, _ := strconv.ParseFloat(item.Percent, 64)
		installs.Items[i] = InstallItem{
			Number:  item.Number,
			Formula: item.Formula,
			Count:   count,
			Percent: percent,
		}
	}
	return installs, nil
}

func (c *Client) GetCaskInstalls(ctx context.Context, days Days) (CaskInstalls, error) {
	path := fmt.Sprintf("/api/analytics/cask-install/%s.json", days)
	return c.getCaskInstalls(ctx, path)
}

func (c *Client) getCaskInstalls(ctx context.Context, path string) (CaskInstalls, error) {
	r, err := c.get(ctx, path)
	if err != nil {
		return CaskInstalls{}, err
	}
	defer r.Close()
	var resp CaskInstallsResponse
	d := json.NewDecoder(r)
	err = d.Decode(&resp)
	if err != nil {
		return CaskInstalls{}, err
	}
	installs := CaskInstalls{
		Category:   resp.Category,
		TotalItems: resp.TotalItems,
		TotalCount: resp.TotalCount,
		StartDate:  resp.StartDate,
		EndDate:    resp.EndDate,
		Items:      make([]CaskInstallItem, len(resp.Items)),
	}
	for i, item := range resp.Items {
		count, _ := strconv.Atoi(strings.ReplaceAll(item.Count, ",", ""))
		percent, _ := strconv.ParseFloat(item.Percent, 64)
		installs.Items[i] = CaskInstallItem{
			Number:  item.Number,
			Cask:    item.Cask,
			Count:   count,
			Percent: percent,
		}
	}
	return installs, nil
}

type BuildErrorsResponse struct {
	Category   string `json:"category"`
	TotalItems int    `json:"total_items"`
	TotalCount int    `json:"total_count"`
	StartDate  Date   `json:"start_date"`
	EndDate    Date   `json:"end_date"`
	Items      []struct {
		Number  int    `json:"number"`
		Formula string `json:"formula"`
		Count   string `json:"count"`
		Percent string `json:"percent"`
	}
}

type BuildErrorItem struct {
	Number  int     `json:"number"`
	Formula string  `json:"formula"`
	Count   int     `json:"count"`
	Percent float64 `json:"percent"`
}

type BuildErrors struct {
	Category   string `json:"category"`
	TotalItems int    `json:"total_items"`
	TotalCount int    `json:"total_count"`
	StartDate  Date   `json:"start_date"`
	EndDate    Date   `json:"end_date"`
	Items      []BuildErrorItem
}

func (c *Client) GetBuildErrors(ctx context.Context, days Days) (BuildErrors, error) {
	path := fmt.Sprintf("/api/analytics/build-error/%s.json", days)
	return c.getBuildErrors(ctx, path)
}

func (c *Client) getBuildErrors(ctx context.Context, path string) (BuildErrors, error) {
	r, err := c.get(ctx, path)
	if err != nil {
		return BuildErrors{}, err
	}
	defer r.Close()
	var resp BuildErrorsResponse
	d := json.NewDecoder(r)
	err = d.Decode(&resp)
	if err != nil {
		return BuildErrors{}, err
	}
	errors := BuildErrors{
		Category:   resp.Category,
		TotalItems: resp.TotalItems,
		TotalCount: resp.TotalCount,
		StartDate:  resp.StartDate,
		EndDate:    resp.EndDate,
		Items:      make([]BuildErrorItem, len(resp.Items)),
	}
	for i, item := range resp.Items {
		count, _ := strconv.Atoi(strings.ReplaceAll(item.Count, ",", ""))
		percent, _ := strconv.ParseFloat(item.Percent, 64)
		errors.Items[i] = BuildErrorItem{
			Number:  item.Number,
			Formula: item.Formula,
			Count:   count,
			Percent: percent,
		}
	}
	return errors, nil
}
