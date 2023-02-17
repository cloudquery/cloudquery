package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

var ROOT_PATH = "https://pricing.us-east-1.amazonaws.com"

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:                "awspricing_services",
		Resolver:            fetchSampleTable,
		PreResourceResolver: getPricingFile,
		Transform:           transformers.TransformWithStruct(&PricingFile{}, transformers.WithSkipFields("Products", "Terms", "AttributesList")),
		Relations: []*schema.Table{
			products(),
			terms(),
		},
	}
}

func products() *schema.Table {
	return &schema.Table{
		Name:      "awspricing_service_products",
		Resolver:  fetchProducts,
		Transform: transformers.TransformWithStruct(&Product{}),
		Relations: []*schema.Table{},
	}
}

func terms() *schema.Table {
	return &schema.Table{
		Name:      "awspricing_service_terms",
		Resolver:  fetchTerms,
		Transform: transformers.TransformWithStruct(&Term{}),
		Relations: []*schema.Table{},
	}
}

func fetchTerms(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	pricingFile := parent.Item.(PricingFile)
	res <- pricingFile.Terms
	return nil
}

func fetchProducts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	pricingFile := parent.Item.(PricingFile)
	res <- pricingFile.Products
	return nil
}

func getRegionalPricingFileLinks(link string) ([]string, error) {
	resp, err := http.Get(ROOT_PATH + link)
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

func getPricingFileLinks() ([]string, error) {
	resp, err := http.Get(ROOT_PATH + "/offers/v1.0/aws/index.json")
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
		regionalLinks, err := getRegionalPricingFileLinks(offer["currentRegionIndexUrl"].(string))
		if err != nil {
			return nil, err
		}
		links = append(links, regionalLinks...)
	}
	return links, nil

}

func getPricingFile(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	link := resource.Item.(string)
	fmt.Printf("Fetching %s%s \n", ROOT_PATH, link)

	resp, err := http.Get(ROOT_PATH + link)
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
func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	links, err := getPricingFileLinks()

	if err != nil {
		return err
	}
	res <- links

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
	terms, err := extractTerms(rawTerms)
	if err != nil {
		return PricingFile{}, err
	}
	pricingFile.Terms = terms

	return pricingFile, nil
}

func extractProducts(rec map[string]any) ([]Product, error) {
	products := make([]Product, len(rec))
	counter := 0
	for _, val := range rec {
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

func countItems(termsRaw map[string]any) int {
	counter := 0
	if onDemand, ok := termsRaw["OnDemand"]; ok {
		onDemand := onDemand.(map[string]any)
		for _, term := range onDemand {
			term := term.(map[string]any)
			counter = counter + len(term)
		}
	}

	if reserved, ok := termsRaw["Reserved"]; ok {
		reserved := reserved.(map[string]any)
		for _, term := range reserved {
			term := term.(map[string]any)
			counter = counter + len(term)
		}
	}

	return counter
}

func extractTerms(termsRaw map[string]any) ([]Term, error) {

	onDemand := termsRaw["OnDemand"].(map[string]any)

	terms := make([]Term, countItems(termsRaw)+1)
	counter := 0
	for _, term := range onDemand {
		term := term.(map[string]any)
		for _, val := range term {
			val := val.(map[string]any)
			effectiveDate, err := time.Parse(time.RFC3339, val["effectiveDate"].(string))
			if err != nil {
				return nil, err
			}
			terms[counter] = Term{
				Type:            "OnDemand",
				OfferTermCode:   val["offerTermCode"].(string),
				Sku:             val["sku"].(string),
				EffectiveDate:   effectiveDate,
				TermAttributes:  extractAttributes(val["termAttributes"].(map[string]any)),
				PriceDimensions: extractPriceDimensions(val["priceDimensions"].(map[string]any)),
			}
			counter++
		}
	}
	if _, ok := termsRaw["Reserved"]; !ok {
		return terms, nil
	}
	reserved := termsRaw["Reserved"].(map[string]any)
	for _, term := range reserved {
		term := term.(map[string]any)
		for _, val := range term {
			val := val.(map[string]any)
			effectiveDate, err := time.Parse(time.RFC3339, val["effectiveDate"].(string))
			if err != nil {
				return nil, err
			}

			var priceDimension []PriceDimension
			if val, ok := val["priceDimensions"]; ok {
				priceDimension = extractPriceDimensions(val.(map[string]any))
			}

			terms[counter] = Term{
				Type:            "Reserved",
				OfferTermCode:   val["offerTermCode"].(string),
				Sku:             val["sku"].(string),
				EffectiveDate:   effectiveDate,
				TermAttributes:  extractAttributes(val["termAttributes"].(map[string]any)),
				PriceDimensions: priceDimension,
			}
			counter++
		}
	}

	return terms, nil
}

func extractPriceDimensions(priceDimensions map[string]any) []PriceDimension {
	priceDimensionsList := make([]PriceDimension, len(priceDimensions))

	for _, val := range priceDimensions {
		val := val.(map[string]any)
		var pricePerUnit PricePerUnit
		byteArray, err := json.Marshal(val["pricePerUnit"])
		if err != nil {
			return nil
		}
		json.Unmarshal(byteArray, &pricePerUnit)

		priceDimension := PriceDimension{
			RateCode:     val["rateCode"].(string),
			Description:  val["description"].(string),
			Unit:         val["unit"].(string),
			PricePerUnit: pricePerUnit,
			AppliesTo:    val["appliesTo"].([]interface{}),
		}
		if val, ok := val["beginRange"]; ok {
			priceDimension.BeginRange = val.(string)
		}
		if val, ok := val["endRange"]; ok {
			priceDimension.EndRange = val.(string)
		}

		priceDimensionsList = append(priceDimensionsList, priceDimension)
	}
	return priceDimensionsList
}

type PricingFile struct {
	FormatVersion   string         `json:"formatVersion"`
	Disclaimer      string         `json:"disclaimer"`
	OfferCode       string         `json:"offerCode"`
	Version         string         `json:"version"`
	PublicationDate time.Time      `json:"publicationDate"`
	Products        []Product      `json:"products"`
	Terms           []Term         `json:"terms"`
	AttributesList  AttributesList `json:"attributesList"`
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
type TermAttributes struct {
}
type Term struct {
	Type            string            `json:"type"`
	OfferTermCode   string            `json:"offerTermCode"`
	Sku             string            `json:"sku"`
	EffectiveDate   time.Time         `json:"effectiveDate"`
	PriceDimensions []PriceDimension  `json:"priceDimensions"`
	TermAttributes  map[string]string `json:"termAttributes"`
}

type AttributesList struct {
}
