# wh [![CI](https://github.com/pchchv/wh/workflows/CI/badge.svg)](https://github.com/pchchv/wh/actions?query=workflow%3ACI+event%3Apush) [![Godoc Reference](https://pkg.go.dev/badge/github.com/pchchv/wh)](https://pkg.go.dev/github.com/pchchv/wh) [![Go Report Card](https://goreportcard.com/badge/github.com/pchchv/wh)](https://goreportcard.com/report/github.com/pchchv/wh)

The `wh` package allows for easy receiving and parsing of GitHub, Bitbucket, GitLab, Docker Hub, Gogs and Azure DevOps Webhook Events.

## Features:

* Parses the entire payload, not just a few fields.
* Fields + Schema directly lines up with webhook posted json

## Examples:

```go
package main

import (
	"fmt"

	"net/http"

	"github.com/pchchv/wh/github"
)

const path = "/webhooks"

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect...?"))
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
			}
		}

		switch payload.(type) {
		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// do whatever you want from here...
			fmt.Printf("%+v", release)
		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}
```
