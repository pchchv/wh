// azure devops does not send an event header, this BasicEvent is provided to get the EventType.
package azure

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
