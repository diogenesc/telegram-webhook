package github

type Repository struct {
	FullName string `json:"full_name"`
}

type WorkflowRun struct {
	Name       string `json:"name"`
	URL        string `json:"html_url"`
	Status     string `json:"status"`
	Conclusion string `json:"conclusion"`
}

type Sender struct {
	Login string `json:"login"`
}

type Body struct {
	Repository  Repository  `json:"repository"`
	WorkflowRun WorkflowRun `json:"workflow_run"`
	Sender      Sender      `json:"sender"`
}

type BuildStatusMessage struct {
	RepositoryFullName string
	Title              string
	Status             string
	Conclusion         string
	Sender             string
	URL                string
}
