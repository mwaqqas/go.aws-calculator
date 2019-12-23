package main

import (
	// "fmt"
	"os"
	"strings"
)

// Docs:
// https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/using-ppslong.html

const (
	offerIndexURL       = "https://pricing.us-east-1.amazonaws.com/offers/v1.0/aws/index.json"
	rawDataDirPath      = "./raw_data"
	offerIndexLocalPath = rawDataDirPath + "/offer_index.json"
	baseURL             = "https://pricing.us-east-1.amazonaws.com"
)

func main() {
	dirCreation(rawDataDirPath)
	downloadFile(offerIndexURL, offerIndexLocalPath)
	parsedOfferIndex, _ := parseOfferIndex(offerIndexLocalPath)
	for _, v := range parsedOfferIndex.Offers {
		offerDir := strings.Join([]string{rawDataDirPath, v.OfferCode}, "/")
		os.MkdirAll(offerDir, 0775)
		regionIndexFile := offerDir + "/region_index.json"
		downloadFile(
			baseURL+v.CurrentRegionIndexURL,
			regionIndexFile,
		)
		parsedRegionIndex, _ := parseRegionIndex(regionIndexFile)
		for _, v := range parsedRegionIndex.Regions {
			offerFile := offerDir + "/" + v.RegionCode + ".json"
			downloadFile(
				baseURL+v.CurrentVersionURL,
				offerFile,
			)
		}
	}
}
