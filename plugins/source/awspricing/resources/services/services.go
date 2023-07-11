package services

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/awspricing/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

const ROOT_URL = "https://pricing.us-east-1.amazonaws.com"

func Services() *schema.Table {
	return &schema.Table{
		Name:                "awspricing_services",
		Title:               "Services from the AWS Price List API",
		Description:         "https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/price-changes.html",
		Resolver:            fetchServicesTable,
		PreResourceResolver: getPricingFile,
		Transform: transformers.TransformWithStruct(&PricingFile{},
			transformers.WithSkipFields("Products", "Terms"),
			transformers.WithPrimaryKeys("OfferCode", "Version", "PublicationDate"),
		),
		Relations: []*schema.Table{
			products(),
			terms(),
		},
	}
}

func fetchServicesTable(_ context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	links, err := getPricingFileLinks(c)
	if err != nil {
		return err
	}
	res <- links
	return nil
}

func validateOffersToSync(c *client.Client, offerCode string) bool {
	// Can't use glob.Glob because currently it is in plugin-sdk/internal
	for _, includeOfferCode := range c.OfferCodes {
		// if glob.Glob(includeOfferCode, offerCode) {
		if includeOfferCode == "*" || includeOfferCode == offerCode {
			return true
		}
	}
	return false
}

func getPricingFileLinks(c *client.Client) ([]string, error) {
	resp, err := http.Get(ROOT_URL + "/offers/v1.0/aws/index.json")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	offersFile := make(map[string]any, 0)

	if err := json.NewDecoder(resp.Body).Decode(&offersFile); err != nil {
		return nil, err
	}
	links := make([]string, 0)
	for _, offer := range offersFile["offers"].(map[string]any) {
		offer := offer.(map[string]any)
		if !validateOffersToSync(c, offer["offerCode"].(string)) {
			continue
		}
		full_url := ROOT_URL + offer["currentRegionIndexUrl"].(string)

		regionalLinks, err := getRegionalPricingFileLinks(c, full_url)
		if err != nil {
			return nil, err
		}
		links = append(links, regionalLinks...)
	}
	return links, nil
}

func validateRegionsToSync(c *client.Client, regionCode string) bool {
	// Can't use glob.Glob because currently it is in plugin-sdk/internal
	for _, includeRegionCode := range c.RegionCodes {
		// if glob.Glob(includeOfferCode, offerCode) {
		if includeRegionCode == "*" || includeRegionCode == regionCode {
			return true
		}
	}
	return false
}

func getRegionalPricingFileLinks(c *client.Client, link string) ([]string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	offerFiles := make(map[string]any, 0)

	if err := json.NewDecoder(resp.Body).Decode(&offerFiles); err != nil {
		return nil, err
	}
	links := make([]string, 0)
	regions := offerFiles["regions"].(map[string]any)
	for _, region := range regions {
		region := region.(map[string]any)
		if !validateRegionsToSync(c, region["regionCode"].(string)) {
			continue
		}
		links = append(links, ROOT_URL+region["currentVersionUrl"].(string))
	}
	return links, nil
}

func getPricingFile(_ context.Context, _ schema.ClientMeta, resource *schema.Resource) error {
	link := resource.Item.(string)
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rawPricingFile := make(map[string]any, 0)

	if err := json.NewDecoder(resp.Body).Decode(&rawPricingFile); err != nil {
		return err
	}
	pricingFile, _ := transformRawPricingFile(rawPricingFile)
	resource.Item = pricingFile
	return nil
}

func transformRawPricingFile(rawStruct map[string]any) (PricingFile, error) {
	var pricingFile PricingFile
	pricingFile.FormatVersion = rawStruct["formatVersion"].(string)
	pricingFile.Disclaimer = rawStruct["disclaimer"].(string)
	pricingFile.OfferCode = rawStruct["offerCode"].(string)
	pricingFile.Version = rawStruct["version"].(string)

	pubDate, err := time.Parse(time.RFC3339, rawStruct["publicationDate"].(string))
	if err != nil {
		return PricingFile{}, err
	}
	pricingFile.PublicationDate = pubDate

	pricingFile.Products = extractProducts(rawStruct["products"].(map[string]any))

	rawTerms := rawStruct["terms"].(map[string]any)
	terms, err := extractAllTerms(rawTerms)
	if err != nil {
		return PricingFile{}, err
	}
	pricingFile.Terms = terms

	return pricingFile, nil
}

func extractProducts(rawProducts map[string]any) []Product {
	products := make([]Product, len(rawProducts))
	counter := 0
	for _, val := range rawProducts {
		val := val.(map[string]any)
		product := Product{
			Sku:        val["sku"].(string),
			Attributes: extractAttributes(val["attributes"].(map[string]any)),
		}
		if productFamily, ok := val["productFamily"]; ok {
			product.ProductFamily = productFamily.(string)
		}
		products[counter] = product
		counter++
	}
	return products
}

func extractAttributes(attributes map[string]any) map[string]string {
	returnVal := make(map[string]string)
	for key, val := range attributes {
		returnVal[key] = val.(string)
	}
	return returnVal
}

var termTypes = []string{"OnDemand", "Reserved"}

func countTermItems(termsRaw map[string]any) int {
	counter := 0
	for _, termType := range termTypes {
		if _, ok := termsRaw[termType]; !ok {
			continue
		}
		for _, typeTermsRaw := range termsRaw[termType].(map[string]any) {
			counter += len(typeTermsRaw.(map[string]any))
		}
	}
	return counter
}

func extractAllTerms(termsRaw map[string]any) ([]Term, error) {
	terms := make([]Term, countTermItems(termsRaw))

	counter := 0
	for _, termType := range termTypes {
		if _, ok := termsRaw[termType]; !ok {
			continue
		}
		for _, typeTermsRaw := range termsRaw[termType].(map[string]any) {
			for _, val := range typeTermsRaw.(map[string]any) {
				term, err := extractTerm(val.(map[string]any), termType)
				if err != nil {
					return nil, err
				}
				terms[counter] = term
				counter++
			}
		}
	}

	return terms, nil
}

func extractTerm(rawTerm map[string]any, termType string) (Term, error) {
	effectiveDate, err := time.Parse(time.RFC3339, rawTerm["effectiveDate"].(string))
	if err != nil {
		return Term{}, err
	}

	var priceDimension []PriceDimension
	if val, ok := rawTerm["priceDimensions"]; ok {
		priceDimension, err = extractPriceDimensions(val.(map[string]any))
		if err != nil {
			return Term{}, err
		}
	}

	return Term{
		Type:            termType,
		OfferTermCode:   rawTerm["offerTermCode"].(string),
		Sku:             rawTerm["sku"].(string),
		EffectiveDate:   effectiveDate,
		TermAttributes:  extractAttributes(rawTerm["termAttributes"].(map[string]any)),
		PriceDimensions: priceDimension,
	}, nil
}

func extractPriceDimensions(priceDimensions map[string]any) ([]PriceDimension, error) {
	priceDimensionsList := make([]PriceDimension, len(priceDimensions))
	counter := 0
	for _, val := range priceDimensions {
		val := val.(map[string]any)
		var pricePerUnit PricePerUnit
		byteArray, err := json.Marshal(val["pricePerUnit"])
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(byteArray, &pricePerUnit)
		if err != nil {
			return nil, err
		}

		// This makes sure there are no nil values
		appliesTo := make([]any, 0)
		if val, ok := val["appliesTo"]; ok {
			appliesTo = val.([]any)
		}

		priceDimension := PriceDimension{
			RateCode:     val["rateCode"].(string),
			Description:  val["description"].(string),
			Unit:         val["unit"].(string),
			PricePerUnit: pricePerUnit,
			AppliesTo:    appliesTo,
		}
		if val, ok := val["beginRange"]; ok {
			priceDimension.BeginRange = val.(string)
		}
		if val, ok := val["endRange"]; ok {
			priceDimension.EndRange = val.(string)
		}

		priceDimensionsList[counter] = priceDimension
		counter++
	}

	return priceDimensionsList, nil
}

type PricingFile struct {
	FormatVersion   string    `json:"formatVersion"`
	Disclaimer      string    `json:"disclaimer"`
	OfferCode       string    `json:"offerCode"`
	Version         string    `json:"version"`
	PublicationDate time.Time `json:"publicationDate"`
	Products        []Product `json:"products"`
	Terms           []Term    `json:"terms"`
}
type Product struct {
	Sku           string            `json:"sku"`
	ProductFamily string            `json:"productFamily"`
	Attributes    map[string]string `json:"attributes"`
}
type PricePerUnit struct {
	Usd string `json:"USD"`
}
type PriceDimension struct {
	RateCode     string       `json:"rateCode"`
	Description  string       `json:"description"`
	BeginRange   string       `json:"beginRange"`
	EndRange     string       `json:"endRange"`
	Unit         string       `json:"unit"`
	PricePerUnit PricePerUnit `json:"pricePerUnit"`
	AppliesTo    []any        `json:"appliesTo"`
}
type Term struct {
	Type            string            `json:"type"`
	OfferTermCode   string            `json:"offerTermCode"`
	Sku             string            `json:"sku"`
	EffectiveDate   time.Time         `json:"effectiveDate"`
	PriceDimensions []PriceDimension  `json:"priceDimensions"`
	TermAttributes  map[string]string `json:"termAttributes"`
}
