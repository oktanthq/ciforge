package circleci_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktanthq/ciforge/provider/circleci"
)

func TestGetMyInfo(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/me")
		assertMethod(t, r, http.MethodGet)
		assertAuth(t, r)

		json.NewEncoder(w).Encode(map[string]any{
			"id": "u1", "login": "ada", "name": "Ada Lovelace", "avatar_url": "https://example.com/a.png",
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	user, err := client.GetMyInfo(context.Background())
	assertNoError(t, err)
	assertEqual(t, "u1", user.ID)
	assertEqual(t, "ada", user.Login)
}

func TestGetCollaborations(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/me/collaborations")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode([]map[string]any{
			{"id": "org-1", "name": "CircleCI-Public", "vcs-type": "github", "slug": "gh/CircleCI-Public"},
		})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	orgs, err := client.GetCollaborations(context.Background())
	assertNoError(t, err)
	assertEqual(t, 1, len(orgs))
	assertEqual(t, "CircleCI-Public", orgs[0].Name)
	assertEqual(t, "github", orgs[0].VCSType)
}

func TestGetUserInfo(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "/user/u1")
		assertMethod(t, r, http.MethodGet)

		json.NewEncoder(w).Encode(map[string]any{"id": "u1", "login": "ada", "name": "Ada Lovelace"})
	}))
	defer srv.Close()

	client := circleci.New("token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	user, err := client.GetUserInfo(context.Background(), "u1")
	assertNoError(t, err)
	assertEqual(t, "ada", user.Login)
}
