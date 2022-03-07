package curlify_apptest

import (
	"encoding/json"
	"strings"
)

func Parse(input string) ParsedInput {
	payload := strings.Split(strings.SplitAfter(input, "REQUEST:")[1], "{'source")[0]
	details := strings.Trim(strings.SplitAfter(input, "REQUEST DETAILS:")[1], "\n\t")
	return ParsedInput{payload, details}
}

func ParseDetails(details string) ParsedDetails {
	var result ParsedDetails
	json.Unmarshal([]byte(details), &result)
	return result
}

type ParsedInput struct {
	Payload string
	Details string
}

type Headers struct {
	UserAgent     string `json:"User-Agent"`
	ContentType   string `json:"Content-Type"`
	Authorization string `json:"Authorization"`
}

type ParsedDetails struct {
	Method  string  `json:"method"`
	URL     string  `json:"url"`
	Headers Headers `json:"headers"`
}
