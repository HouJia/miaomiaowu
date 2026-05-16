package basepath

import (
	"net/http"
	"strings"
)

// responseWriter rewrites Location headers to include the base path prefix.
type responseWriter struct {
	http.ResponseWriter
	basePath string
}

func (w *responseWriter) WriteHeader(code int) {
	if loc := w.Header().Get("Location"); loc != "" && w.basePath != "" {
		w.Header().Set("Location", rewriteLocation(loc, w.basePath))
	}
	w.ResponseWriter.WriteHeader(code)
}

func rewriteLocation(loc, basePath string) string {
	if loc == "" || basePath == "" {
		return loc
	}
	// Skip absolute URLs and protocol-relative URLs
	if strings.Contains(loc, "://") || strings.HasPrefix(loc, "//") {
		return loc
	}
	if !strings.HasPrefix(loc, "/") {
		return loc
	}
	if strings.HasPrefix(loc, basePath+"/") || loc == basePath {
		return loc
	}
	return basePath + loc
}

// stripMiddleware removes BASE_PATH from incoming requests.
func stripMiddleware(basePath string, next http.Handler) http.Handler {
	if basePath == "" {
		return next
	}
	return http.StripPrefix(basePath, next)
}
