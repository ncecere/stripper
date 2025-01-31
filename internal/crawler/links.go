package crawler

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

// extractLinks parses HTML content and returns all unique links
func extractLinks(content string, baseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return nil, err
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	links := make(map[string]struct{})
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href := a.Val
					if href == "" || strings.HasPrefix(href, "#") {
						continue
					}

					u, err := url.Parse(href)
					if err != nil {
						continue
					}

					// Resolve relative URLs
					resolved := base.ResolveReference(u)
					// Clean the URL (remove fragments, etc.)
					resolved.Fragment = ""
					resolved.RawQuery = ""

					// Only keep http(s) URLs
					if resolved.Scheme != "http" && resolved.Scheme != "https" {
						continue
					}

					links[resolved.String()] = struct{}{}
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	result := make([]string, 0, len(links))
	for link := range links {
		result = append(result, link)
	}
	return result, nil
}

// extractTitle extracts the title from HTML content
func extractTitle(content string) string {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		return ""
	}

	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				title = n.FirstChild.Data
				return
			}
		}
		for c := n.FirstChild; c != nil && title == ""; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return strings.TrimSpace(title)
}

// isSameDomain checks if two URLs belong to the same domain
func isSameDomain(baseURL, targetURL string) bool {
	base, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	target, err := url.Parse(targetURL)
	if err != nil {
		return false
	}

	return base.Hostname() == target.Hostname()
}
