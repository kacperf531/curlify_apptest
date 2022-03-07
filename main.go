package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Parse(input string) ParsedInput {
	payload := strings.ReplaceAll(strings.Split(strings.SplitAfter(input, "REQUEST:")[1], "{'source")[0], "\n", "")
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

func main() {
	scn := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Hello, this is a tool which converts output from apptest logs to cURL command which can be imported in e.g. postman.")
		fmt.Println("Insert request log from apptest report, followed by ctrl+], insert ctrl+D to exit.")
		var lines []string
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				// Group Separator (GS ^]): ctrl-]
				if line[0] == '\x1D' {
					break
				}
			}
			lines = append(lines, line)
		}

		if len(lines) > 0 {
			fmt.Println()
			fmt.Println("Result:")
			input := strings.Join(lines, "\n")
			parsed_input := Parse(input)
			fmt.Println(Curlify(parsed_input))
			fmt.Println()
		}

		if err := scn.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		if len(lines) == 0 {
			break
		}
	}

}
