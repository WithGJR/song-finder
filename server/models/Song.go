package models

type Song struct {
	BaseModel
	Name     string
	Lyricist string
	Composer string
	Length   string

	SingerId uint
	Singer   Singer

	AlbumId uint
	Album   Album
}
