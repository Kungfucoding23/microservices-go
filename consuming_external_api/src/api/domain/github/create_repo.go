package github

/*
	{
		"name": "Golang microservices Tutorial",
		"description": "This is a repo created using github api",
		"homepage": "https://github.com",
		"private": false,
		"has_issues": true,
		"has_projects": true,
		"has_wiki": true
	}
*/

//CreateRepoRequest is the request struct
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

//CreateRepoResponse is the response struct
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

//RepoOwner is the owner struct
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

//RepoPermissions is the permissions struct
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPush bool `json:"push"`
	HasPull bool `json:"pull"`
}

//CreateRepo ...
// func CreateRepo() {
// 	//this is not recomended
// 	// request := map[string]interface{}{
// 	// 	"name":    "Hello-World",
// 	// 	"private": false,
// 	// }
// 	// //beacuse to acces a key this is needed
// 	// private := request["private"].(bool) // this is not going to scale
// 	//so we create a struct "CreateRepoRequest" instade of this

// }
