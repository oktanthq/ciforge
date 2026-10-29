# ciforge

[![CI](https://github.com/oktanthq/ciforge/actions/workflows/ci.yml/badge.svg)](https://github.com/oktanthq/ciforge/actions/workflows/ci.yml)

ciforge is a Go library for talking to CI provider APIs. Each provider is its own package; they share an internal HTTP client for authentication, request execution, and common error mapping.

Today that means:

- **GitHub Actions** — `provider/github`
- **GitLab CI** — `provider/gitlab`
- **CircleCI** — `provider/circleci`

**Jenkins** (`provider/jenkins`) is planned; the package is an empty placeholder.

There is no unified SDK interface yet. You import a provider and call its methods directly. A shared interface is tracked in [#10](https://github.com/oktanthq/ciforge/issues/10).

## Providers

| Provider | Package | Status | Docs |
| --- | --- | --- | --- |
| GitHub Actions | [`provider/github`](provider/github) | Supported | [`provider/github/README.md`](provider/github/README.md) |
| GitLab CI | [`provider/gitlab`](provider/gitlab) | Supported | [`provider/gitlab/README.md`](provider/gitlab/README.md) |
| CircleCI | [`provider/circleci`](provider/circleci) | Supported | — |
| Jenkins | [`provider/jenkins`](provider/jenkins) | Planned | — |

Worked examples live under [`examples/`](examples/) (`examples/github` and `examples/gitlab`). CircleCI does not have an example or provider README yet ([#8](https://github.com/oktanthq/ciforge/issues/8)).

The shared HTTP layer lives in [`provider/internal/httpclient`](provider/internal/httpclient). Provider clients embed it, then set their own auth headers (`Authorization: Bearer` for GitHub, `PRIVATE-TOKEN` for GitLab, `Circle-Token` for CircleCI). Common failures map to `ErrUnauthorized`, `ErrNotFound`, and `ErrRateLimited`.

## Installation

```bash
go get github.com/oktanthq/ciforge
```

Import a provider:

```go
import (
    gh "github.com/oktanthq/ciforge/provider/github"
    gl "github.com/oktanthq/ciforge/provider/gitlab"
    cci "github.com/oktanthq/ciforge/provider/circleci"
)
```

Requires Go 1.26.3 or later (see `go.mod`).

## Quick start

Every provider is constructed with `New(token)` plus optional `WithHTTPClient` and `WithBaseURL`. A client is safe to reuse.

### GitHub Actions

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    gh "github.com/oktanthq/ciforge/provider/github"
)

func main() {
    client := gh.New(os.Getenv("GITHUB_TOKEN"))
    // Optional: gh.New(token, gh.WithHTTPClient(httpClient), gh.WithBaseURL("https://github.example.com/api/v3"))

    ctx := context.Background()
    owner, repo := "owner", "repo"

    workflows, err := client.ListWorkflows(ctx, repo, owner)
    if err != nil {
        log.Fatal(err)
    }
    for _, w := range workflows {
        fmt.Printf("%s (id=%d, state=%s)\n", w.Name, w.ID, w.State)
    }

    runs, err := client.ListRunsForWorkflow(ctx, repo, owner, workflows[0].ID, gh.ListRunsForWorkflowOptions{})
    if err != nil {
        log.Fatal(err)
    }
    for _, r := range runs {
        fmt.Printf("run #%d status=%s conclusion=%s\n", r.RunNumber, r.Status, r.Conclusion)
    }
}
```

See [`examples/github/main.go`](examples/github/main.go) and [`provider/github/README.md`](provider/github/README.md) for workflows, runs, jobs, artifacts, logs, and concurrency groups.

### GitLab CI

The GitLab client's default base URL is `https://gitlab.example.com/api/v4`. Point it at gitlab.com (or your instance) with `WithBaseURL`.

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    gl "github.com/oktanthq/ciforge/provider/gitlab"
)

func main() {
    client := gl.New(os.Getenv("GITLAB_TOKEN"),
        gl.WithBaseURL("https://gitlab.com/api/v4"),
    )

    ctx := context.Background()
    projectID := 12345

    pipelines, err := client.ListPipelines(ctx, projectID, &gl.ListPipelinesOptions{Ref: "main"})
    if err != nil {
        log.Fatal(err)
    }
    for _, p := range pipelines {
        fmt.Printf("pipeline #%d status=%s ref=%s\n", p.ID, p.Status, p.Ref)
    }

    created, err := client.CreatePipeline(ctx, projectID, "main")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("created pipeline #%d\n", created.ID)
}
```

See [`examples/gitlab/main.go`](examples/gitlab/main.go) and [`provider/gitlab/README.md`](provider/gitlab/README.md) for pipelines, jobs, artifacts, schedules, trigger tokens, project variables, and runners.

### CircleCI

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    cci "github.com/oktanthq/ciforge/provider/circleci"
)

func main() {
    client := cci.New(os.Getenv("CIRCLECI_TOKEN"))
    // Optional: cci.New(token, cci.WithHTTPClient(httpClient), cci.WithBaseURL("https://circleci.com/api/v2"))

    ctx := context.Background()
    slug := "gh/org/project"

    project, err := client.GetProject(ctx, slug)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("project %s (%s)\n", project.Name, project.Slug)

    pipelines, err := client.GetAllPipelines(ctx, slug, cci.GetPipelineOptions{Branch: "main"})
    if err != nil {
        log.Fatal(err)
    }
    for _, p := range pipelines {
        fmt.Printf("pipeline #%d state=%s id=%s\n", p.Number, p.State, p.ID)
    }

    if err := client.TriggerPipeline(ctx, slug, cci.TriggerPipelineBody{Branch: "main"}); err != nil {
        log.Fatal(err)
    }
}
```

The CircleCI package also covers workflows, jobs, contexts, projects, webhooks, insights, usage, deploy, and rollback. A dedicated example and README are tracked in [#8](https://github.com/oktanthq/ciforge/issues/8).

## Tests

httptest coverage lives next to each provider:

- GitHub Actions: `provider/github/github_test`
- GitLab CI: `provider/gitlab/gitlab_test`
- CircleCI: `provider/circleci/circleci_test`

```bash
go test ./...
```

Race-detector run:

```bash
go test -race ./...
```

## Contributing

1. Open an issue or pick one from the [roadmap](#roadmap).
2. Fork the repo and work on a feature branch.
3. Match the surrounding style: one provider package, methods on `*Client`, options structs with `query` tags, sentinel errors via `errors.Is`.
4. Add httptest coverage next to the GitHub tests where you change behavior.
5. Open a pull request.

Please keep changes scoped. Provider APIs should stay on the shared `provider/internal/httpclient` client rather than introducing a second HTTP stack.

## Roadmap

### Docs & packaging

- [#2](https://github.com/oktanthq/ciforge/issues/2) — Add a LICENSE

### Test coverage & CI

- [#4](https://github.com/oktanthq/ciforge/issues/4) — Add tests for the GitLab provider
- [#5](https://github.com/oktanthq/ciforge/issues/5) — Add tests for the CircleCI provider
- [#6](https://github.com/oktanthq/ciforge/issues/6) — Add a CI workflow to build, test and lint

### Provider completion

- [#7](https://github.com/oktanthq/ciforge/issues/7) — Implement the Jenkins provider
- [#8](https://github.com/oktanthq/ciforge/issues/8) — Add a CircleCI example and README
- [#9](https://github.com/oktanthq/ciforge/issues/9) — Rename `provider/circleci/tigger.go` to `trigger.go`

### Unified SDK interface

- [#10](https://github.com/oktanthq/ciforge/issues/10) — Define a shared provider interface
