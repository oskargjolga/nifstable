package nifs

import (
	"net/http"
	"testing"
)

func TestFetchMatches(t *testing.T) {
	c := NewNifsClient("https://api.nifs.no")

	matches := c.FetchMatches("5", "682936")
	if len(matches) != 240 {
		// 16 teams -> 16*16-16 matches
		t.Error("Expected 240 matches, got", len(matches))
	}

	if matches[0].HomeTeam.Name != "Aalesund" {
		t.Error("Expected Aalesund, got", matches[0].HomeTeam.Name)
	}
}

func TestFetchMatch(t *testing.T) {
	c := NewNifsClient("https://api.nifs.no")

	match := c.FetchMatch("1800721")
	if match.Name != "Aalesund - Molde" {
		t.Error("Expected Aalesund - Molde, got", match.Name)
	}
}

func TestPerformRequest(t *testing.T) {
	tests := []struct {
		url string
		err error
	}{
		{"https://api.nifs.no/api/v1/matches/1800721", nil},
	}

	for _, test := range tests {
		c := NewNifsClient("https://api.nifs.no")
		req, err := http.NewRequest(http.MethodGet, test.url, nil)
		if err != nil {
			t.Error(err)
		}
		_, err = c.PerformRequest(req)
		if err != test.err {
			t.Error("Expected", test.err, "got", err)
		}
	}

}
