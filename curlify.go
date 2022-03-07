package curlify_apptest

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Parse(input string) ParsedInput {
	payload := strings.Split(strings.SplitAfter(input, "REQUEST:")[1], "{'source")[0]
	details := ParseDetails(strings.SplitAfter(input, "REQUEST DETAILS:")[1])
	return ParsedInput{payload, details}
}

func ParseDetails(details string) ParsedDetails {
	var result ParsedDetails
	json.Unmarshal([]byte(details), &result)
	return result
}

type ParsedInput struct {
	Payload string
	Details ParsedDetails
}

type ParsedDetails struct {
	Method  string  `json:"method"`
	URL     string  `json:"url"`
	Headers Headers `json:"headers"`
}

type Headers struct {
	UserAgent     string `json:"User-Agent"`
	ContentType   string `json:"Content-Type"`
	Authorization string `json:"Authorization"`
}

func Curlify(pi ParsedInput) string {
	return fmt.Sprintf(`curl --location --request %s '%s' \
		--header 'Authorization: %s' \
		--header 'Content-Type: %s' \
		--header 'User-Agent: %s' \
		--data '%s'`, pi.Details.Method, pi.Details.URL, pi.Details.Headers.Authorization, pi.Details.Headers.ContentType, pi.Details.Headers.UserAgent, pi.Payload)
}
