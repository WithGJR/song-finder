package models

type Song struct {
	BaseModel
	Name     string
	Country  string
	Lyricist string
	Composer string
	Length   string

	SingerId uint
	AlbumId  uint
}
