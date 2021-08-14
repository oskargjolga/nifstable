package nifs

import "testing"

func TestGetMatches(t *testing.T) {
	c := NewNifsClient("https://api.nifs.no")

	matches := c.GetMatches("5", "682936")
	if len(matches) != 240 {
		// 16 teams -> 16*16-16 matches
		t.Error("Expected 240 matches, got", len(matches))
	}
}