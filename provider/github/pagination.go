package github

import "strings"

type pageLinks map[string]string

func parseLinkHeader(header string) pageLinks {
	links := make(pageLinks)
	for _, part := range strings.Split(header, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		segments := strings.Split(part, ";")
		if len(segments) < 2 {
			continue
		}
		urlPart := strings.TrimSpace(segments[0])
		urlPart = strings.TrimPrefix(urlPart, "<")
		urlPart = strings.TrimSuffix(urlPart, ">")

		relPart := strings.TrimSpace(segments[1])
		relPart = strings.TrimPrefix(relPart, "rel=")
		relPart = strings.Trim(relPart, "\"")

		links[relPart] = urlPart
	}
	return links
}

func (l pageLinks) Next() string {
	return l["next"]
}

func (l pageLinks) Prev() string {
	return l["prev"]
}

func (l pageLinks) Last() string {
	return l["last"]
}

func (l pageLinks) First() string {
	return l["first"]
}
