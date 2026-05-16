package basepath

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWrapRootRedirect(t *testing.T) {
	cfg := Config{BasePath: "/mmw"}
	var innerPath string
	inner := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		innerPath = r.URL.Path
		w.WriteHeader(http.StatusOK)
	})
	h := Wrap(inner, cfg)

	for _, path := range []string{"/", "/mmw"} {
		req := httptest.NewRequest(http.MethodGet, path, nil)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		if rec.Code != http.StatusFound {
			t.Fatalf("%s: want 302, got %d", path, rec.Code)
		}
		loc := rec.Header().Get("Location")
		if loc != "/mmw/" {
			t.Fatalf("%s: Location=%q want /mmw/", path, loc)
		}
		if innerPath != "" {
			t.Fatalf("%s: inner handler should not run", path)
		}
	}

	req := httptest.NewRequest(http.MethodGet, "/mmw/api/setup/status", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if innerPath != "/api/setup/status" {
		t.Fatalf("stripped path=%q want /api/setup/status", innerPath)
	}
}
