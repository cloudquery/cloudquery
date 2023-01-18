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
	log zerolog.Logger

	baseURL    string
	hc         HTTPDoer
	maxRetries int64

	apiKey, apiSecret, accessToken string

	lim *rate.Limiter
}

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

func New(log zerolog.Logger, hc HTTPDoer, apiKey, apiSecret, accessToken, shopURL string, maxRetries int64) (*Client, error) {
	if accessToken == "" && (apiKey == "" || apiSecret == "") {
		return nil, fmt.Errorf("missing shopify access token, api key or secret")
	}
	if shopURL == "" {
		return nil, fmt.Errorf("missing shop url")
	}

	return &Client{
		log:        log,
		baseURL:    strings.TrimRight(shopURL, "/") + "/",
		hc:         hc,
		maxRetries: maxRetries,

		apiKey:      apiKey,
		apiSecret:   apiSecret,
		accessToken: accessToken,

		lim: rate.NewLimiter(rate.Limit(80), 120),
	}, nil
}

func (s *Client) request(ctx context.Context, edge string) (retResp *http.Response, retErr error) {
	tries := int64(0)

	defer func() {
		if retErr != nil {
			s.log.Error().Err(retErr).Str("edge", edge).Msg("request failed")
		} else if tries > 5 {
			s.log.Debug().Str("edge", edge).Int64("num_tries", tries).Msg("success after tries")
		}
	}()

	for tries < s.maxRetries {
		if !s.lim.Allow() {
			s.log.Debug().Msg("waiting for rate limiter...")
			if err := s.lim.Wait(ctx); err != nil {
				return nil, err
			}
			s.log.Debug().Msg("wait complete")
		}

		r, wait, err := s.retryableRequest(ctx, edge)
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
			break
		}

		tries++
		if wait == nil { // no retry-after returned, linear backoff
			w := time.Duration(tries) * 1 * time.Second
			wait = &w
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(*wait):
		}
	}

	return nil, errors.New("exceeded max retries")
}

func (s *Client) retryableRequest(ctx context.Context, edge string) (*http.Response, *time.Duration, error) {
	u := fmt.Sprintf("%v%v", s.baseURL, edge)
	log := s.log.With().Str("edge", edge).Logger()

	var (
		body []byte
		err  error
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}

	if s.accessToken != "" {
		req.Header.Add("X-Shopify-Access-Token", s.accessToken)
	} else {
		req.SetBasicAuth(s.apiKey, s.apiSecret)
	}
	req.Header.Add("Content-type", "application/json")

	resp, err := s.hc.Do(req)
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

func (s *Client) GetProducts(ctx context.Context, pageUrl string) (*GetProductsResponse, string, error) {
	var ret GetProductsResponse

	const pageSize = 250

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/products.json?limit=%v", APIVersion, pageSize)
	}

	resp, err := s.request(ctx, pageUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = pageSize

	return &ret, nextPage, nil
}

func (s *Client) GetOrders(ctx context.Context, pageUrl string) (*GetOrdersResponse, string, error) {
	var ret GetOrdersResponse

	const pageSize = 250

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/orders.json?limit=%v", APIVersion, pageSize)
	}

	resp, err := s.request(ctx, pageUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = pageSize

	return &ret, nextPage, nil
}

func (s *Client) GetCustomers(ctx context.Context, pageUrl string) (*GetCustomersResponse, string, error) {
	var ret GetCustomersResponse

	const pageSize = 250

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/customers.json?limit=%v", APIVersion, pageSize)
	}

	resp, err := s.request(ctx, pageUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = pageSize

	return &ret, nextPage, nil
}

func (s *Client) GetAbandonedCheckouts(ctx context.Context, pageUrl string) (*GetCheckoutsResponse, string, error) {
	var ret GetCheckoutsResponse

	const pageSize = 20

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/checkouts.json?limit=%v", APIVersion, pageSize)
	}

	resp, err := s.request(ctx, pageUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = pageSize

	return &ret, nextPage, nil
}

func (s *Client) GetPriceRules(ctx context.Context, pageUrl string) (*GetPriceRulesResponse, string, error) {
	var ret GetPriceRulesResponse

	const pageSize = 250

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/price_rules.json?limit=%v", APIVersion, pageSize)
	}

	resp, err := s.request(ctx, pageUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = pageSize

	return &ret, nextPage, nil
}

func (s *Client) GetDiscountCodes(ctx context.Context, priceRuleID int64, pageUrl string) (*GetDiscountCodesResponse, string, error) {
	var ret GetDiscountCodesResponse

	const pageSize = 250

	if pageUrl == "" {
		pageUrl = fmt.Sprintf("admin/api/%s/price_rules/%d/discount_codes.json?limit=%v", APIVersion, priceRuleID, pageSize)
	}

	resp, err := s.request(ctx, pageUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	nextPage := getNextPage(resp.Header)

	err = json.NewDecoder(resp.Body).Decode(&ret)
	if err != nil {
		return nil, "", err
	}

	ret.PageSize = pageSize

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
