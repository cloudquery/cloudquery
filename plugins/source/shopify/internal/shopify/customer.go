package shopify

import "time"

type Customer struct {
	ID                        int64             `json:"id"`
	Email                     string            `json:"email"`
	AcceptsMarketing          bool              `json:"accepts_marketing"`
	CreatedAt                 time.Time         `json:"created_at"`
	UpdatedAt                 time.Time         `json:"updated_at"`
	FirstName                 string            `json:"first_name"`
	LastName                  string            `json:"last_name"`
	OrdersCount               int               `json:"orders_count"`
	State                     string            `json:"state"`
	TotalSpent                string            `json:"total_spent"`
	LastOrderID               int64             `json:"last_order_id"`
	Note                      any               `json:"note"`
	VerifiedEmail             bool              `json:"verified_email"`
	MultipassIdentifier       any               `json:"multipass_identifier"`
	TaxExempt                 bool              `json:"tax_exempt"`
	Tags                      string            `json:"tags"`
	LastOrderName             string            `json:"last_order_name"`
	Currency                  string            `json:"currency"`
	Phone                     any               `json:"phone"`
	Addresses                 []CustomerAddress `json:"addresses"`
	AcceptsMarketingUpdatedAt *time.Time        `json:"accepts_marketing_updated_at"`
	MarketingOptInLevel       string            `json:"marketing_opt_in_level"`
	TaxExemptions             []any             `json:"tax_exemptions"`
	EmailMarketingConsent     struct {
		State            string     `json:"state"`
		OptInLevel       string     `json:"opt_in_level"`
		ConsentUpdatedAt *time.Time `json:"consent_updated_at"`
	} `json:"email_marketing_consent"`
	SmsMarketingConsent any             `json:"sms_marketing_consent"`
	AdminGraphqlAPIID   string          `json:"admin_graphql_api_id"`
	DefaultAddress      CustomerAddress `json:"default_address"`
}

type CustomerAddress struct {
	ID           int64   `json:"id"`
	CustomerID   int64   `json:"customer_id"`
	Name         string  `json:"name"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Company      *string `json:"company"`
	Address1     *string `json:"address1"`
	Address2     *string `json:"address2"`
	City         *string `json:"city"`
	Province     *string `json:"province"`
	Country      *string `json:"country"`
	Zip          *string `json:"zip"`
	Phone        *string `json:"phone"`
	ProvinceCode *string `json:"province_code"`
	CountryCode  *string `json:"country_code"`
	CountryName  *string `json:"country_name"`
	Default      bool    `json:"default"`
}

type GetCustomersResponse struct {
	Customers []Customer `json:"customers"`
	PageSize  int        `json:"page_size"`
}
