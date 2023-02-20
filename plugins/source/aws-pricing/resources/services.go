package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/aws-pricing/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

var ROOT_PATH = "https://pricing.us-east-1.amazonaws.com"

func Services() *schema.Table {
	return &schema.Table{
		Name:                "awspricing_services",
		Resolver:            fetchSampleTable,
		PreResourceResolver: getPricingFile,
		Transform:           transformers.TransformWithStruct(&PricingFile{}, transformers.WithSkipFields("Products", "Terms")),
		Relations: []*schema.Table{
			products(),
			terms(),
		},
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	links, err := getPricingFileLinks(c.Endpoint)
	if err != nil {
		return err
	}
	res <- links
	return nil
}

func getPricingFileLinks(endpoint string) ([]string, error) {
	resp, err := http.Get(endpoint + "/offers/v1.0/aws/index.json")
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
		full_url := endpoint + offer["currentRegionIndexUrl"].(string)
		regionalLinks, err := getRegionalPricingFileLinks(full_url)
		if err != nil {
			return nil, err
		}
		links = append(links, regionalLinks...)
	}
	return links, nil

}

func getRegionalPricingFileLinks(link string) ([]string, error) {
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
		links = append(links, region["currentVersionUrl"].(string))
	}
	return links, nil
}

func getPricingFile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
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

	products, err := extractProducts(rawStruct["products"].(map[string]any))
	if err != nil {
		return PricingFile{}, err
	}
	pricingFile.Products = products

	rawTerms := rawStruct["terms"].(map[string]any)
	terms, err := extractAllTerms(rawTerms)
	if err != nil {
		return PricingFile{}, err
	}
	pricingFile.Terms = terms

	return pricingFile, nil
}

func extractProducts(rawProducts map[string]any) ([]Product, error) {
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
	return products, nil
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
			counter = counter + len(typeTermsRaw.(map[string]any))
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
		priceDimension = extractPriceDimensions(val.(map[string]any))
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

func extractPriceDimensions(priceDimensions map[string]any) []PriceDimension {
	priceDimensionsList := make([]PriceDimension, len(priceDimensions))
	counter := 0
	for _, val := range priceDimensions {
		val := val.(map[string]any)
		var pricePerUnit PricePerUnit
		byteArray, err := json.Marshal(val["pricePerUnit"])
		if err != nil {
			return nil
		}
		json.Unmarshal(byteArray, &pricePerUnit)

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

	return priceDimensionsList
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
	RateCode     string        `json:"rateCode"`
	Description  string        `json:"description"`
	BeginRange   string        `json:"beginRange"`
	EndRange     string        `json:"endRange"`
	Unit         string        `json:"unit"`
	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
	AppliesTo    []interface{} `json:"appliesTo"`
}
type Term struct {
	Type            string            `json:"type"`
	OfferTermCode   string            `json:"offerTermCode"`
	Sku             string            `json:"sku"`
	EffectiveDate   time.Time         `json:"effectiveDate"`
	PriceDimensions []PriceDimension  `json:"priceDimensions"`
	TermAttributes  map[string]string `json:"termAttributes"`
}
