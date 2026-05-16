package publicpath

import "testing"

func TestNoopJoin(t *testing.T) {
	SetProvider(nil)
	got := Join("api", "proxy-provider", "42")
	want := "/api/proxy-provider/42"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
