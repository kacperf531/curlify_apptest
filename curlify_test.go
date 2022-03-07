package curlify_apptest

import "testing"

func TestParsing(t *testing.T) {

	input := `{{'source': 'Configuration API Public'} irrelevant text REQUEST:{"x": "test"}{'source': 'Configuration API Public'} irrelevant text REQUEST DETAILS:{"method": "POST"}`

	t.Run("Parsing to struct with 2 strings: Details & Payload", func(t *testing.T) {
		got := Parse(input)
		want := ParsedInput{`{"x": "test"}`, `{"method": "POST"}`}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Parsing request details", func(t *testing.T) {
		request_details := `{
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
		got := ParseDetails(request_details)
		want := ParsedDetails{"POST", "https://api.livechatinc.com/v3.5/configuration/action/list_agents", Headers{"apptest", "application/json", "Basic XYZ="}}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
