package basepath

import (
	"strings"
)

// Provider implements publicpath.Provider for sub-path deployment.
type Provider struct {
	cfg Config
}

// NewProvider returns a publicpath provider for the given config.
func NewProvider(cfg Config) *Provider {
	return &Provider{cfg: cfg}
}

func (p *Provider) Prefix() string {
	return p.cfg.BasePath
}

func (p *Provider) Join(elem ...string) string {
	if !p.cfg.Enabled() {
		return joinRoot(elem...)
	}
	parts := append([]string{strings.Trim(p.cfg.BasePath, "/")}, trimElems(elem)...)
	return joinRoot(parts...)
}

func (p *Provider) AbsURL(pathAndQuery string) string {
	path := pathAndQuery
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	if !p.cfg.Enabled() {
		if p.cfg.PublicOrigin != "" {
			return p.cfg.PublicOrigin + path
		}
		return path
	}
	if !strings.HasPrefix(path, p.cfg.BasePath) {
		path = p.cfg.BasePath + path
	}
	if p.cfg.PublicOrigin != "" {
		return p.cfg.PublicOrigin + path
	}
	return path
}

func (p *Provider) PublicOrigin() string {
	return p.cfg.PublicOrigin
}


func joinRoot(elem ...string) string {
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
	return "/" + strings.Join(parts, "/")
}

func trimElems(elem []string) []string {
	var out []string
	for _, e := range elem {
		e = strings.Trim(e, "/")
		if e != "" {
			out = append(out, e)
		}
	}
	return out
}
