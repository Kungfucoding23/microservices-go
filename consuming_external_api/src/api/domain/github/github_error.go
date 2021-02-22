package github

//ErrorResponse is the error struct for a bad request
type ErrorResponse struct {
	StatusCode       int            `json:"status_code"`
	Message          string         `json:"message"`
	Errors           []ErrorsGithub `json:"errors"`
	DocumentationURL string         `json:"documentation_url"`
}

func (r ErrorResponse) Error() string {
	return r.Message
}

//ErrorsGithub is the error struct
type ErrorsGithub struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
