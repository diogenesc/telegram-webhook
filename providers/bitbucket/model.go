package bitbucket

type Repository struct {
	Name string `json:"name"`
}

type CommitStatus struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	State string `json:"state"`
	Type  string `json:"type"`
}

type Body struct {
	Repository   Repository   `json:"repository"`
	CommitStatus CommitStatus `json:"commit_status"`
}

type BuildStatusMessage struct {
	Title string
	State string
	URL   string
}
