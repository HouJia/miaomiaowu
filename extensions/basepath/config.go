package basepath

import (
	"os"
	"strings"
)

const DefaultBasePath = "/mmw"

// Config holds sub-path deployment settings.
type Config struct {
	BasePath      string // e.g. "/mmw"; "/" or "" disables strip
	PublicOrigin  string // optional scheme+host only, no path
}

// FromEnv reads BASE_PATH (default /mmw) and optional PUBLIC_ORIGIN.
func FromEnv() Config {
	bp := os.Getenv("BASE_PATH")
	if bp == "" {
		bp = DefaultBasePath
	}
	return Config{
		BasePath:     NormalizeBasePath(bp),
		PublicOrigin: strings.TrimRight(strings.TrimSpace(os.Getenv("PUBLIC_ORIGIN")), "/"),
	}
}

// NormalizeBasePath returns "" for root mode, otherwise a path like "/mmw".
func NormalizeBasePath(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" || raw == "/" {
		return ""
	}
	if !strings.HasPrefix(raw, "/") {
		raw = "/" + raw
	}
	return strings.TrimRight(raw, "/")
}

// Enabled reports whether sub-path middleware should strip a prefix.
func (c Config) Enabled() bool {
	return c.BasePath != ""
}
