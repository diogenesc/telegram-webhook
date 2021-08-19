package bitbucket

type Repository struct {
	Name string `json:"name"`
}

type CommitStatus struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	State string `json:"state"`
}

type Body struct {
	Repository   Repository   `json:"repository"`
	CommitStatus CommitStatus `json:"commit_status"`
}
