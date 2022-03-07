package curlify_apptest

import "strings"

func Parse(input string) ParsedInput {
	payload := strings.Split(strings.SplitAfter(input, "REQUEST:")[1], "{'source")[0]
	details := strings.Trim(strings.SplitAfter(input, "REQUEST DETAILS:")[1], "\n\t")
	return ParsedInput{payload, details}
}

type ParsedInput struct {
	Payload string
	Details string
}
