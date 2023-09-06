package entities

type Config struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Name string `json:"name"`
	Key  string `json:"key"`
	Host string `json:"host"`
}
