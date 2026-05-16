/** Path prefix from Vite base (fork: /mmw/; upstream: ./ or /). */
function normalizedBase(): string {
  const raw = (import.meta.env.BASE_URL ?? '/').trim()
  if (raw === './' || raw === '/' || raw === '') {
    return ''
  }
  return raw.replace(/\/$/, '')
}

export const basePath = normalizedBase()

/** Prepend BASE_PATH to an app route, e.g. /login → /mmw/login */
export function withBase(path: string): string {
  const p = path.startsWith('/') ? path : `/${path}`
  return basePath ? `${basePath}${p}` : p
}

/** Origin + base path for copyable links and proxy-provider URLs */
export function appOrigin(): string {
  if (typeof window === 'undefined') {
    return ''
  }
  return window.location.origin + basePath
}

/** Full absolute URL for copy/share (origin + BASE_PATH + path) */
export function absoluteURL(path: string): string {
  if (typeof window === 'undefined') {
    return withBase(path)
  }
  const p = path.startsWith('/') ? path : `/${path}`
  if (basePath && (p === basePath || p.startsWith(`${basePath}/`))) {
    return window.location.origin + p
  }
  return window.location.origin + withBase(p)
}

/** Axios baseURL: origin + base path (API paths still use /api/...) */
export function apiBaseURL(): string {
  if (typeof window === 'undefined') {
    return ''
  }
  const { protocol, host, hostname } = window.location
  const suffix = basePath
  if (
    hostname === 'localhost' ||
    hostname === '127.0.0.1' ||
    hostname === '::1'
  ) {
    return `${protocol}//${hostname}:8080${suffix}`
  }
  return `${protocol}//${host}${suffix}`
}
