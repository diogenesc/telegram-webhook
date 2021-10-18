package bitbucket

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

type Commit struct {
	Author Author `json:"author"`
}

type Author struct {
	User User `json:"user"`
}

type User struct {
	DisplayName string `json:"display_name"`
}

type Body struct {
	Repository   Repository   `json:"repository"`
	CommitStatus CommitStatus `json:"commit_status"`
}

type BuildStatusMessage struct {
	RepositoryFullName string
	Title              string
	State              string
	Author             string
	URL                string
}
