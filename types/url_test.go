package types

import "testing"

func TestURL(t *testing.T) {
	t.Run("test-validation", func(t *testing.T) {
		uri := Url("x.com")

		uri.Verify()

		if uri != "http://x.com" {
			t.Fatalf("failed basic verification")
		}
	})
}
