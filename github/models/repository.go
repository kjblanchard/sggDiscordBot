package github

type Repository struct {
	Id int `json:"id"`
	Name string `json:"name"`
	// Fork bool `json:"fork"`
	// UnknownFields map[string]interface{} `json:"-"`
}
