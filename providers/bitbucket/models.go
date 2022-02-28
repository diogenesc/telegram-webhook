package bitbucket

// Request parse
type Body struct {
	Repository   Repository   `json:"repository"`
	CommitStatus CommitStatus `json:"commit_status"`
	PullRequest  PullRequest  `json:"pullrequest"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

type CommitStatus struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	State  string `json:"state"`
	Type   string `json:"type"`
	Commit Commit `json:"commit"`
}

type PullRequest struct {
	Title       string      `json:"title"`
	Links       Links       `json:"links"`
	State       string      `json:"state"`
	Destination Destination `json:"destination"`
	Source      Source      `json:"source"`
}

type Commit struct {
	Author Author `json:"author"`
}

type Author struct {
	User User `json:"user"`
}

type User struct {
	DisplayName string `json:"display_name"`
}

type Source struct {
	Branch Branch `json:"branch"`
}

type Destination struct {
	Branch Branch `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}

type Links struct {
	Html Html `json:"html"`
}

type Html struct {
	Href string `json:"href"`
}

// Messages fields
type BuildPipelineMessage struct {
	RepositoryFullName string
	Title              string
	State              string
	Author             string
	URL                string
}

type BuildPullRequestMessage struct {
	RepositoryFullName string
	Title              string
	Source             string
	Destination        string
	State              string
	URL                string
}
