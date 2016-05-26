package models

type Album struct {
	BaseModel
	Name          string
	Company       string
	PublishedDate string
	Language      string
	Reward        string
	Photo         string
}
