package gitlab

const (
	// GitLab hook types.
	TagEvents                 Event  = "Tag Push Hook"
	JobEvents                 Event  = "Job Hook"
	PushEvents                Event  = "Push Hook"
	BuildEvents               Event  = "Build Hook"
	IssuesEvents              Event  = "Issue Hook"
	CommentEvents             Event  = "Note Hook"
	ReleaseEvents             Event  = "Release Hook"
	WikiPageEvents            Event  = "Wiki Page Hook"
	PipelineEvents            Event  = "Pipeline Hook"
	DeploymentEvents          Event  = "Deployment Hook"
	SystemHookEvents          Event  = "System Hook"
	MergeRequestEvents        Event  = "Merge Request Hook"
	ConfidentialIssuesEvents  Event  = "Confidential Issue Hook"
	ConfidentialCommentEvents Event  = "Confidential Note Hook"
	objectTag                 string = "tag_push"
	objectPush                string = "push"
	objectBuild               string = "build"
	objectMergeRequest        string = "merge_request"
	eventKeyCreate            string = "key_create"
	eventUserCreate           string = "user_create"
	eventUserRename           string = "user_rename"
	eventKeyDestroy           string = "key_destroy"
	eventUserDestroy          string = "user_destroy"
	eventGroupRename          string = "group_rename"
	eventGroupCreate          string = "group_create"
	eventGroupDestroy         string = "group_destroy"
	eventProjectCreate        string = "project_create"
	eventProjectRename        string = "project_rename"
	eventProjectUpdate        string = "project_update"
	eventUserAddToTeam        string = "user_add_to_team"
	eventProjectDestroy       string = "project_destroy"
	eventUserAddToGroup       string = "user_add_to_group"
	eventUserFailedLogin      string = "user_failed_login"
	eventProjectTransfer      string = "project_transfer"
	eventUserUpdateForTeam    string = "user_update_for_team"
	eventUserRemoveFromTeam   string = "user_remove_from_team"
	eventUserUpdateForGroup   string = "user_update_for_group"
	eventUserRemoveFromGroup  string = "user_remove_from_group"
)

// Event defines a GitLab hook event type by the X-Gitlab-Event Header.
type Event string

// Webhook instance contains all methods needed to process events.
type Webhook struct {
	secretHash []byte
}
