package basepath

import "testing"

func TestFromEnvDefault(t *testing.T) {
	t.Setenv("BASE_PATH", "")
	cfg := FromEnv()
	if cfg.BasePath != DefaultBasePath {
		t.Fatalf("expected default %q, got %q", DefaultBasePath, cfg.BasePath)
	}
}

func TestNormalizeBasePathRoot(t *testing.T) {
	if got := NormalizeBasePath("/"); got != "" {
		t.Fatalf("expected empty for root, got %q", got)
	}
	if got := NormalizeBasePath(""); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestProviderJoin(t *testing.T) {
	p := NewProvider(Config{BasePath: "/mmw"})
	got := p.Join("api", "proxy-provider", "1")
	want := "/mmw/api/proxy-provider/1"
	if got != want {
		t.Fatalf("got %q want %q", got, want)
	}
}
