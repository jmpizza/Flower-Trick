package models

type Pagination struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
}
type NamedResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type URLResource struct {
	URL string `json:"url"`
}
type NamedResponse struct {
	Pagination
	Results []NamedResource `json:"results"`
}

type UnnamedResponse struct {
	Pagination
	Results []URLResource `json:"results"`
}
