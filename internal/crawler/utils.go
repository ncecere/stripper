package crawler

import "strings"

// shouldIgnoreURL checks if a URL should be ignored based on its extension.
// It handles URLs with query parameters and fragments, and ensures consistent
// extension comparison by normalizing both the URL and extension format.
func shouldIgnoreURL(urlStr string, ignoreExts []string) bool {
	// Remove query parameters and fragments
	if idx := strings.IndexAny(urlStr, "?#"); idx != -1 {
		urlStr = urlStr[:idx]
	}

	// Convert to lowercase for case-insensitive comparison
	lower := strings.ToLower(urlStr)

	// Check each extension
	for _, ext := range ignoreExts {
		// Ensure extension has a leading dot
		ext = "." + strings.TrimPrefix(ext, ".")
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}
	return false
}
