package model

type CatImage struct {
	Id     string `json:"id"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
