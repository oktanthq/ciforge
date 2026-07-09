package github

import "strings"

type Links map[string]string

func ParseLinkHeader(header string) Links {
	links := make(Links)
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

func (l Links) Next() string {
	return l["next"]
}

func (l Links) Prev() string {
	return l["prev"]
}

func (l Links) Last() string {
	return l["last"]
}

func (l Links) First() string {
	return l["first"]
}
