package nifs

import "testing"

func TestGetMatches(t *testing.T) {
	c := NewNifsClient()

	matches := c.GetMatches()
	if len(matches) != 240 {
		t.Error("Expected 240 matches, got", len(matches)) // should be 240 (16*16 -16)
	}
}
