package crawler

import "strings"

// shouldIgnoreURL checks if a URL should be ignored based on its extension
func shouldIgnoreURL(urlStr string, ignoreExts []string) bool {
	lower := strings.ToLower(urlStr)
	for _, ext := range ignoreExts {
		if strings.HasSuffix(lower, "."+strings.TrimPrefix(ext, ".")) {
			return true
		}
	}
	return false
}
