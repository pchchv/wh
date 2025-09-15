// azure devops does not send an event header, this BasicEvent is provided to get the EventType.
package azure

import "net/http"

const (
	// Azure DevOps Server hook types.
	GitPushEventType               Event = "git.push"
	BuildCompleteEventType         Event = "build.complete"
	GitPullRequestMergedEventType  Event = "git.pullrequest.merged"
	GitPullRequestCreatedEventType Event = "git.pullrequest.created"
	GitPullRequestUpdatedEventType Event = "git.pullrequest.updated"
)

// Event defines an Azure DevOps server hook event type.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	username string
	password string
}

func (hook Webhook) verifyBasicAuth(r *http.Request) bool {
	if hook.username == "" && hook.password == "" {
		// skip validation if username or password was not provided
		return true
	}

	username, password, ok := r.BasicAuth()
	return ok && username == hook.username && password == hook.password
}
