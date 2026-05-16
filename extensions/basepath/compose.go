package basepath

import (
	"net/http"
	"strings"
)

// Wrap applies sub-path middleware: root redirect, StripPrefix on ingress, Location rewrite on egress.
func Wrap(next http.Handler, cfg Config) http.Handler {
	if next == nil || !cfg.Enabled() {
		return next
	}
	inner := stripMiddleware(cfg.BasePath, next)
	base := strings.TrimRight(cfg.BasePath, "/")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "" || path == "/" {
			http.Redirect(w, r, base+"/", http.StatusFound)
			return
		}
		if path == base {
			http.Redirect(w, r, base+"/", http.StatusFound)
			return
		}
		rw := &responseWriter{ResponseWriter: w, basePath: cfg.BasePath}
		inner.ServeHTTP(rw, r)
	})
}
