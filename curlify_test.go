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
}
