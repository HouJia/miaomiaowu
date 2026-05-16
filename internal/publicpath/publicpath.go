package publicpath

import (
	"net/http"
	"path"
	"strings"
)

// Provider supplies path prefix and absolute URL helpers for sub-path deployment.
type Provider interface {
	Prefix() string
	Join(elem ...string) string
	AbsURL(pathAndQuery string) string
	PublicOrigin() string
}

type noopProvider struct{}

func (noopProvider) Prefix() string { return "" }

func (noopProvider) Join(elem ...string) string { return joinSegments(elem...) }

func (noopProvider) AbsURL(pathAndQuery string) string { return joinSegments(strings.Split(strings.TrimPrefix(pathAndQuery, "/"), "/")...) }

func (noopProvider) PublicOrigin() string { return "" }

var current Provider = noopProvider{}

// SetProvider installs the active provider (e.g. extensions/basepath). Nil resets to noop.
func SetProvider(p Provider) {
	if p == nil {
		current = noopProvider{}
		return
	}
	current = p
}

// Prefix returns the configured path prefix (e.g. "/mmw") or "" for root deployment.
func Prefix() string { return current.Prefix() }

// Join builds a path with the active prefix: Join("api", "x") → "/api/x" or "/mmw/api/x".
func Join(elem ...string) string { return current.Join(elem...) }

// AbsURL returns an absolute URL when PUBLIC_ORIGIN or request context is available.
func AbsURL(pathAndQuery string) string { return current.AbsURL(pathAndQuery) }

// PublicOrigin returns optional fixed origin (scheme+host only), or "".
func PublicOrigin() string { return current.PublicOrigin() }

// AbsURLFromRequest builds absolute URL using request Host when no PUBLIC_ORIGIN is set.
func AbsURLFromRequest(r *http.Request, pathAndQuery string) string {
	if origin := PublicOrigin(); origin != "" {
		return strings.TrimRight(origin, "/") + ensureLeadingSlash(pathAndQuery)
	}
	if r == nil {
		return Join(strings.Split(strings.Trim(strings.TrimPrefix(pathAndQuery, "/"), "/"), "/")...)
	}
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	if proto := r.Header.Get("X-Forwarded-Proto"); proto != "" {
		scheme = proto
	}
	host := r.Host
	if host == "" {
		host = r.URL.Host
	}
	p := pathAndQuery
	if !strings.HasPrefix(p, "/") {
		p = "/" + p
	}
	return scheme + "://" + host + p
}

func joinSegments(elem ...string) string {
	var parts []string
	for _, e := range elem {
		e = strings.Trim(e, "/")
		if e != "" {
			parts = append(parts, e)
		}
	}
	if len(parts) == 0 {
		return "/"
	}
	return "/" + path.Join(parts...)
}

func ensureLeadingSlash(p string) string {
	if p == "" || strings.HasPrefix(p, "/") {
		return p
	}
	return "/" + p
}
