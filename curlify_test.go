package main

import "testing"

func TestCurlify(t *testing.T) {
	t.Run("Test input parsing", func(t *testing.T) {
		input := `{{'source': 'Configuration API Public'} irrelevant text REQUEST:{"x": "test"}{'source': 'Configuration API Public'} irrelevant text REQUEST DETAILS:{
			"method": "POST",
			"url": "https://api.livechatinc.com/v3.5/configuration/action/list_agents",
			"headers": {
				"User-Agent": "apptest",
				"Accept-Encoding": "gzip, deflate",
				"Accept": "*/*",
				"Connection": "keep-alive",
				"Content-Type": "application/json",
				"Authorization": "Basic XYZ=",
				"Content-Length": "2"
			},
			"request_send_time": "10:45:22.171454",
			"response_duration": "0.27426 second(s)"
		}`
		got := Parse(input)
		want := ParsedInput{`{"x": "test"}`, ParsedDetails{"POST", "https://api.livechatinc.com/v3.5/configuration/action/list_agents", Headers{"apptest", "application/json", "Basic XYZ="}}}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Test the output", func(t *testing.T) {
		got := Curlify(ParsedInput{`{"suspended": true}`, ParsedDetails{"POST", "https://api.livechatinc.com/v3.5/configuration/action/list_agents", Headers{"apptest", "application/json", "Basic XYZ="}}})
		want := `curl --location --request POST 'https://api.livechatinc.com/v3.5/configuration/action/list_agents' \
		--header 'Authorization: Basic XYZ=' \
		--header 'Content-Type: application/json' \
		--header 'User-Agent: apptest' \
		--data '{"suspended": true}'`
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
