package models

type Song struct {
	Model
	Name     string
	Country  string
	Lyricist string
	Composer string
	Length   string

	SingerId uint
	AlbumId  uint
}

func FindSongs(songs *[]Song) error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	db.Find(songs)
	db.Close()
	return nil
}

func (this *Song) Create() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	db.Create(this)
	db.Close()
	return nil
}

func (this *Song) Update() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	db.Save(this)
	db.Close()
	return nil
}

func (this *Song) Delete() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	if err := db.First(this).Error; err != nil {
		return err
	}
	db.Delete(this)
	db.Close()
	return nil
}
