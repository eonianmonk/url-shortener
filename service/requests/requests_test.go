package requests

import (
	"net/http"
	"testing"
)

func TestRequestsParsing(t *testing.T) {
	t.Run("test-get-url", func(t *testing.T) {
		r, _ := http.NewRequest("PUT", "/link", nil)
		r.Host = "site.com"
		surl := NewGetRequest(r)
		if surl != "link" {
			t.Fatalf("failed to get interesting part, got \"%s\" instead", surl)
		}
	})
}
