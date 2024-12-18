package blue

import (
	"encoding/json"
	"testing"

	"github.com/prebid/prebid-server/v3/openrtb_ext"
)

func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderBlue, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected blue params: %s", validParam)
		}
	}
}

func TestInvalidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderBlue, json.RawMessage(invalidParam)); err == nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

var validParams = []string{
	`{"publisherId":"1234", "placementId":"12345","bidFloor":1,"currency":"USD"}`,
	`{"publisherId":"", "placementId":"","bidFloor":0,"currency":""}`,
}

var invalidParams = []string{
	``,
	`null`,
	`true`,
	`5`,
	`4.2`,
	`[]`,
	`{}`,
	`{"placementId":"12345","bidFloor":1,"currency":"USD"}`,
	`{"publisherId":"1234", "bidFloor":1,"currency":"USD"}`,
	`{"publisherId":"1234", "placementId":"12345","currency":"USD"}`,
	`{"publisherId":"1234", "placementId":"12345","bidFloor":1}`,
	`{"publisherId":"Y9Evrh40ejsrCR4EtidUt1cSxhJsz8X1", "placementId":"12345","bidFloor":1,"currency":"USD"}`,
	`{"publisherId":"1234", "placementId":"alNYtemWggraDVbhJrsOs9pXc3Eld32E","bidFloor":1,"currency":"USD"}`,
	`{"publisherId":"1234", "placementId":"12345","bidFloor":1,"currency":"Bq7PfZuLxKcMnT5RwV2Gy9AoHiJD4S6F"}`,
}
