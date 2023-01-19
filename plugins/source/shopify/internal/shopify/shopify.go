package shopify

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/shopify/internal/httperror"
	"github.com/rs/zerolog"
	"golang.org/x/time/rate"
)

const (
	APIVersion = "2022-10"
)

type Client struct {
	opts *ClientOptions

	baseURL string
	lim     *rate.Limiter
}

type ClientOptions struct {
	Log zerolog.Logger

	HC         HTTPDoer
	MaxRetries int64
	PageSize   int

	ApiKey, ApiSecret, AccessToken string
	ShopURL                        string
}

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(opts ClientOptions) (*Client, error) {
	if opts.AccessToken == "" && (opts.ApiKey == "" || opts.ApiSecret == "") {
		return nil, fmt.Errorf("missing shopify access token, api key or secret")
	}
	if opts.ShopURL == "" {
		return nil, fmt.Errorf("missing shop url")
	}

	return &Client{
		opts:    &opts,
		baseURL: strings.TrimRight(opts.ShopURL, "/") + "/",
		lim:     rate.NewLimiter(rate.Limit(80), 120),
	}, nil
}

func (s *Client) request(ctx context.Context, edge string, params url.Values) (retResp *http.Response, retErr error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("limit", strconv.FormatInt(int64(s.opts.PageSize), 10))

	tries := int64(0)

	log := s.opts.Log.With().Str("edge", edge).Interface("query_params", params).Logger()

	defer func() {
		if retErr != nil {
			log.Error().Err(retErr).Msg("request failed")
		} else if tries > 0 {
			log.Debug().Int64("num_tries", tries).Msg("success after tries")
		}
	}()

	for {
		if !s.lim.Allow() {
			log.Debug().Msg("waiting for rate limiter...")
			if err := s.lim.Wait(ctx); err != nil {
				return nil, err
			}
			log.Debug().Msg("wait complete")
		}

		r, wait, err := s.retryableRequest(ctx, edge, params)
		if err == nil {
			return r, nil
		}

		temporary := false
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			temporary = true
		} else if he, ok := err.(httperror.Error); ok {
			temporary = he.Temporary()
		}
		if !temporary {
			return nil, fmt.Errorf("request failed with error: %w", err)
		}

		tries++
		if tries >= s.opts.MaxRetries {
			return nil, fmt.Errorf("exceeded max retries (%d): %w", s.opts.MaxRetries, err)
		}

		if wait == nil { // no retry-after returned, linear backoff
			w := time.Duration(tries) * 1 * time.Second
			wait = &w
		}

		log.Warn().Err(err).Float64("backoff_seconds", wait.Seconds()).Msg("retryable request failed, will retry")

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(*wait):
		}
	}
}

func (s *Client) retryableRequest(ctx context.Context, edge string, params url.Values) (*http.Response, *time.Duration, error) {
	log := s.opts.Log.With().Str("edge", edge).Interface("query_params", params).Logger()

	u := s.baseURL + edge + "?" + params.Encode()

	var (
		body []byte
		err  error
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}

	if s.opts.AccessToken != "" {
		req.Header.Add("X-Shopify-Access-Token", s.opts.AccessToken)
	} else {
		req.SetBasicAuth(s.opts.ApiKey, s.opts.ApiSecret)
	}
	req.Header.Add("Content-type", "application/json")

	resp, err := s.opts.HC.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("do %s: %w", edge, err)
	}

	if dr := resp.Header.Get("X-Shopify-Api-Deprecated-Reason"); dr != "" {
		log.Warn().Str("deprecated_reason", dr).Msg("Deprecated API call detected")
	}

	var wait *time.Duration
	if ra := resp.Header.Get("Retry-After"); ra != "" {
		rr, err := strconv.ParseFloat(ra, 64)
		if err != nil {
			log.Warn().Str("retry_after", ra).Err(err).Msg("Unknown Retry-After received")
		} else {
			t := time.Duration(rr) * time.Second
			wait = &t
		}
	}

	if resp.StatusCode != http.StatusOK {
		bdy, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		var bodyStr string
		if bdy != nil {
			bodyStr = string(bdy)
		}
		if bodyStr == "" {
			b, _ := json.Marshal(resp.Header)
			bodyStr = "headers: " + string(b)
		}

		return nil, wait, httperror.New(resp.StatusCode, http.MethodGet, edge, resp.Status, bodyStr)
	}

	return resp, wait, nil
}

func (s *Client) GetProducts(ctx context.Context, pageUrl string, params url.Values) (*GetProductsResponse, string, error) {
	var ret GetProductsResponse

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/products.json", APIVersion)
	}

	resp, err := s.request(ctx, pageUrl, params)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = s.opts.PageSize

	return &ret, nextPage, nil
}

func (s *Client) GetOrders(ctx context.Context, pageUrl string, params url.Values) (*GetOrdersResponse, string, error) {
	var ret GetOrdersResponse

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/orders.json", APIVersion)
	}

	resp, err := s.request(ctx, pageUrl, params)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = s.opts.PageSize

	return &ret, nextPage, nil
}

func (s *Client) GetCustomers(ctx context.Context, pageUrl string, params url.Values) (*GetCustomersResponse, string, error) {
	var ret GetCustomersResponse

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/customers.json", APIVersion)
	}

	resp, err := s.request(ctx, pageUrl, params)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = s.opts.PageSize

	return &ret, nextPage, nil
}

func (s *Client) GetAbandonedCheckouts(ctx context.Context, pageUrl string, params url.Values) (*GetCheckoutsResponse, string, error) {
	var ret GetCheckoutsResponse

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/checkouts.json", APIVersion)
	}

	resp, err := s.request(ctx, pageUrl, params)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = s.opts.PageSize

	return &ret, nextPage, nil
}

func (s *Client) GetPriceRules(ctx context.Context, pageUrl string, params url.Values) (*GetPriceRulesResponse, string, error) {
	var ret GetPriceRulesResponse

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/price_rules.json", APIVersion)
	}

	resp, err := s.request(ctx, pageUrl, params)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = s.opts.PageSize

	return &ret, nextPage, nil
}

func (s *Client) GetDiscountCodes(ctx context.Context, priceRuleID int64, pageUrl string) (*GetDiscountCodesResponse, string, error) {
	var ret GetDiscountCodesResponse

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/price_rules/%d/discount_codes.json", APIVersion, priceRuleID)
	}

	resp, err := s.request(ctx, pageUrl, nil)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = s.opts.PageSize

	return &ret, nextPage, nil
}

func getNextPage(hdr http.Header) string {
	for _, link := range textproto.MIMEHeader(hdr)[textproto.CanonicalMIMEHeaderKey("link")] {
		for _, part := range strings.Split(link, ",") {
			if !strings.HasSuffix(part, `; rel="next"`) {
				continue
			}

			if u := strings.Trim(strings.TrimSuffix(part, `; rel="next"`), `<> `); u != "" {
				ur, _ := url.Parse(u)
				return ur.Path + "?" + ur.RawQuery
			}
		}
	}

	return ""
}
