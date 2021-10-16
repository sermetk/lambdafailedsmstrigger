package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/aws/aws-lambda-go/lambda"
)

type CloudWatchSMSFailure struct {
	Notification Notification `json:"notification"`
	Delivery     Delivery     `json:"delivery"`
	Status       string       `json:"status"`
}

type Delivery struct {
	PhoneCarrier              string  `json:"phoneCarrier"`
	Mnc                       int64   `json:"mnc"`
	NumberOfMessageParts      int64   `json:"numberOfMessageParts"`
	Destination               string  `json:"destination"`
	PriceInUSD                float64 `json:"priceInUSD"`
	SMSType                   string  `json:"smsType"`
	Mcc                       int64   `json:"mcc"`
	ProviderResponse          string  `json:"providerResponse"`
	DwellTimeMS               int64   `json:"dwellTimeMs"`
	DwellTimeMSUntilDeviceACK int64   `json:"dwellTimeMsUntilDeviceAck"`
}

type Notification struct {
	MessageID string `json:"messageId"`
	Timestamp string `json:"timestamp"`
}

func handler(ctx context.Context, b json.RawMessage) error {

	cloudWatchSMSFailure, err := UnmarshalCloudWatchSMSFailure(b)

	if err != nil {
		log.Fatal("Could not unmarshal scheduled event: ", err)
	} else if cloudWatchSMSFailure.Status == "FAILURE" {
		destinationNumber := trimFirstChar(cloudWatchSMSFailure.Delivery.Destination)
		//retry logic
	}
	return nil
}

func UnmarshalCloudWatchSMSFailure(data []byte) (CloudWatchSMSFailure, error) {
	var r CloudWatchSMSFailure
	err := json.Unmarshal(data, &r)
	return r, err
}

func trimFirstChar(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
func main() {
	lambda.Start(handler)
}
