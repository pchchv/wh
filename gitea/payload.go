package gitea

// PusherType define the type to push.
type PusherType string

// StateType issue state type.
type StateType string

// HookRepoAction an action that happens to a repo.
type HookRepoAction string

// HookIssueAction FIXME.
type HookIssueAction string

// HookReleaseAction defines hook release action type.
type HookReleaseAction string

// HookIssueCommentAction defines hook issue comment action.
type HookIssueCommentAction string

// Label a label to an issue or a pr.
type Label struct {
	ID          int64  `json:"id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Description string `json:"description"`
}

// Permission represents a set of permissions.
type Permission struct {
	Pull  bool `json:"pull"`
	Push  bool `json:"push"`
	Admin bool `json:"admin"`
}

// InternalTracker represents settings for internal tracker.
type InternalTracker struct {
	EnableTimeTracker                bool `json:"enable_time_tracker"`
	EnableIssueDependencies          bool `json:"enable_issue_dependencies"`
	AllowOnlyContributorsToTrackTime bool `json:"allow_only_contributors_to_track_time"`
}

// ExternalTracker represents settings for external tracker.
type ExternalTracker struct {
	ExternalTrackerURL    string `json:"external_tracker_url"`
	ExternalTrackerStyle  string `json:"external_tracker_style"`
	ExternalTrackerFormat string `json:"external_tracker_format"`
}
