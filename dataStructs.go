package main

type offer struct {
	OfferCode             string `json:"offerCode"`
	VersionIndexURL       string `json:"versionIndexUrl"`
	CurrentVersionURL     string `json:"currentVersionUrl"`
	CurrentRegionIndexURL string `json:"currentRegionIndexUrl"`
}

type offerIndex struct {
	FormatVersion   string           `json:"formatVersion"`
	Disclaimer      string           `json:"disclaimer"`
	PublicationDate string           `json:"publicationDate"`
	Offers          map[string]offer `json:"offers"`
}

type region struct {
	RegionCode        string `json:"regionCode"`
	CurrentVersionURL string `json:"currentVersionUrl"`
}

type regionIndex struct {
	FormatVersion   string            `json:"formatVersion"`
	Disclaimer      string            `json:"disclaimer"`
	PublicationDate string            `json:"publicationDate"`
	Regions         map[string]region `json:"regions"`
}

type productAttr struct {
	Servicecode  string `json: servicecode`
	Location     string `json: location`
	LocationType string `json: locationType`
	Group        string `json: group`
	Usagetype    string `json: usagetype`
	Operation    string `json: operation`
	Servicename  string `json: servicename`
}

type product struct {
	Sku           string                 `json:sku`
	ProductFamily string                 `json:productFamily`
	Attributes    map[string]productAttr `json:attributes`
}

// type termAttr struct{}
// type appliesTo struct{}
// type pricePerUnit struct {
// 	USD string
// }

// type priceDimension struct {
// 	RateCode     string                  `json: "rateCode"`
// 	Description  string                  `json: "description"`
// 	BeginRange   string                  `json: "beginRange"`
// 	EndRange     string                  `json: "endRange"`
// 	Unit         string                  `json: "unit"`
// 	PricePerUnit map[string]pricePerUnit `json: "pricePerUnit"`
// 	AppliesTo    []appliesTo             `json: "appliesTo"`
// }

// type productSkuPerTerm struct {
// 	OfferTermCode   string                    `json: "offerTermCode"`
// 	Sku             string                    `json: "sku"`
// 	EffectiveDate   string                    `json: "effectiveDate"`
// 	PriceDimensions map[string]priceDimension `json: "priceDimensions"`
// }

type productIndex struct {
	FormatVersion   string             `json:"formatVersion"`
	Disclaimer      string             `json:"disclaimer"`
	OfferCode       string             `json:offerCode`
	PublicationDate string             `json:"publicationDate"`
	Products        map[string]product `json:"products"`
	Terms           interface{}        `json:"terms"`
}
