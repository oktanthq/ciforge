package github_test

import (
	"testing"

	"github.com/abdulahshoaib/ciforge/provider/github"
)

func TestParseLinkHeader(t *testing.T) {
	t.Parallel()

	header := `<https://api.github.com/repos/owner/repo/actions/workflows?page=2>; rel="next", <https://api.github.com/repos/owner/repo/actions/workflows?page=5>; rel="last"`
	links := github.ParseLinkHeader(header)

	assertEqual(t, "https://api.github.com/repos/owner/repo/actions/workflows?page=2", links.Next())
	assertEqual(t, "https://api.github.com/repos/owner/repo/actions/workflows?page=5", links.Last())
}

func TestParseLinkHeaderEmpty(t *testing.T) {
	t.Parallel()

	links := github.ParseLinkHeader("")
	assertEqual(t, "", links.Next())
}
