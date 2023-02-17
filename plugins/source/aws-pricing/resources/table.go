package resources

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:      "awspricing_services",
		Resolver:  fetchSampleTable,
		Transform: transformers.TransformWithStruct(&PricingFile{}),
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {

	resp, err := http.Get("https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/comprehend/20230210221619/us-east-1/index.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rawPricingFile := make(map[string]any, 0)

	if err := json.NewDecoder(resp.Body).Decode(&rawPricingFile); err != nil {
		return err
	}
	pricingFile, _ := transformRawPricingFile(rawPricingFile)
	res <- pricingFile
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

	terms := rawStruct["terms"].(map[string]any)
	onDemandTerms, err := extractOnDemandTerms(terms["OnDemand"].(map[string]any))
	if err != nil {
		return PricingFile{}, err
	}
	pricingFile.Terms.OnDemand = onDemandTerms

	return pricingFile, nil
}

func extractProducts(rec map[string]any) ([]Product, error) {
	var products []Product
	for _, val := range rec {
		val := val.(map[string]any)
		var attributes Attributes
		byteArray, err := json.Marshal(val["attributes"])
		if err != nil {
			return nil, err
		}
		json.Unmarshal(byteArray, &attributes)
		if err != nil {
			return nil, err
		}
		products = append(products, Product{
			Sku:           val["sku"].(string),
			ProductFamily: val["productFamily"].(string),
			Attributes:    attributes,
		})
	}
	return products, nil
}
func extractOnDemandTerms(onDemand map[string]any) ([]OnDemand, error) {
	var onDemandTerms []OnDemand
	for _, terms := range onDemand {
		terms := terms.(map[string]any)
		for _, val := range terms {
			val := val.(map[string]any)
			effectiveDate, err := time.Parse(time.RFC3339, val["effectiveDate"].(string))
			if err != nil {
				return nil, err
			}
			onDemandTerms = append(onDemandTerms, OnDemand{
				OfferTermCode:   val["offerTermCode"].(string),
				Sku:             val["sku"].(string),
				EffectiveDate:   effectiveDate,
				TermAttributes:  val["termAttributes"].(map[string]any),
				PriceDimensions: extractPriceDimensions(val["priceDimensions"].(map[string]any)),
			})
		}
	}
	return onDemandTerms, nil
}

func extractPriceDimensions(priceDimensions map[string]any) []PriceDimension {
	var priceDimensionsList []PriceDimension
	for _, val := range priceDimensions {
		val := val.(map[string]any)
		var pricePerUnit PricePerUnit
		byteArray, err := json.Marshal(val["pricePerUnit"])
		if err != nil {
			return nil
		}
		json.Unmarshal(byteArray, &pricePerUnit)
		priceDimensionsList = append(priceDimensionsList, PriceDimension{
			RateCode:     val["rateCode"].(string),
			Description:  val["description"].(string),
			BeginRange:   val["beginRange"].(string),
			EndRange:     val["endRange"].(string),
			Unit:         val["unit"].(string),
			PricePerUnit: pricePerUnit,
			AppliesTo:    val["appliesTo"].([]interface{}),
		})
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
	Terms           Terms          `json:"terms"`
	AttributesList  AttributesList `json:"attributesList"`
}
type Attributes struct {
	Servicecode      string `json:"servicecode"`
	Location         string `json:"location"`
	LocationType     string `json:"locationType"`
	Group            string `json:"group"`
	GroupDescription string `json:"groupDescription"`
	Usagetype        string `json:"usagetype"`
	Operation        string `json:"operation"`
	Platofeaturetype string `json:"platofeaturetype"`
	RegionCode       string `json:"regionCode"`
	Servicename      string `json:"servicename"`
}
type Product struct {
	Sku           string     `json:"sku"`
	ProductFamily string     `json:"productFamily"`
	Attributes    Attributes `json:"attributes"`
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
type OnDemand struct {
	OfferTermCode   string           `json:"offerTermCode"`
	Sku             string           `json:"sku"`
	EffectiveDate   time.Time        `json:"effectiveDate"`
	PriceDimensions []PriceDimension `json:"priceDimensions"`
	TermAttributes  map[string]any   `json:"termAttributes"`
}
type Terms struct {
	OnDemand []OnDemand `json:"OnDemand"`
}
type AttributesList struct {
}
