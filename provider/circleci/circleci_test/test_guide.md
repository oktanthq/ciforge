# CircleCI Test Guide

Tests live in `circleci_test/` as an external package (`package circleci_test`) importing `provider/circleci`. Run: `go test ./circleci_test/`.

## Success template

```go
func Test<Method>_Success(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertPath(t, r, "<path>")
		assertMethod(t, r, http.Method<VERB>)
		assertAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`<response body>`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	<result>, err := client.<Method>(<args>)

	assertNoError(t, err)
	assertEqual(t, <want>, <result>.<Field>)
	// slices: check len with t.Fatalf, then assert elements
}
```

## Error template

```go
func Test<Method>_<Error>(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.Status<CODE>)
		w.Write([]byte(`{"message": "<msg>"}`))
	}))
	defer srv.Close()

	client := circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))
	_, err := client.<Method>(<args>)

	assertError(t, err)
	if !errors.Is(err, circleci.<ErrXXX>) {
		t.Errorf("err = %v, want %v", err, circleci.<ErrXXX>)
	}
}
```

## Rules

- Naming: `Test<Method>_<Scenario>` — at least one success + one error test per method
- Request assertions go in the handler before writing the response
- Client construction: `circleci.New("test-token", circleci.WithBaseURL(srv.URL), circleci.WithHTTPClient(srv.Client()))`
- Errors: use `circleci.ErrNotFound`, `circleci.ErrUnauthorized`, `circleci.ErrRateLimited`
- Use realistic API JSON payloads with `Content-Type: application/json`
- Helpers in `helper_test.go`: `assertPath`, `assertMethod`, `assertAuth`, `assertNoError`, `assertError`, `assertEqual` (all with `t.Helper()`)
