package test

import (
	"testing"
	"blockchain/internal/block"
	"time"
)

func TestBlock(t *testing.T) {
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	tm, _ := time.Parse(layout, "Jan 1, 2022 at 8:07pm (PST)")
   
	b := block.New(tm, "lastHash", "hash", "test_data")
	found := b.ToString()

	const expected = `{"Timestamp":"2022-01-01T20:07:00-08:00","LastHash":"lastHash","Hash":"hash","Data":"test_data"}`
	if expected != b.ToString() {
		t.Fatal("Expected block (%s), found (%s)", expected, found)
	}
}
