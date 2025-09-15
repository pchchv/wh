package docker

// Event defines a Docker hook event type.
type Event string

// Docker hook types (only one for now).
const BuildEvent Event = "build"
