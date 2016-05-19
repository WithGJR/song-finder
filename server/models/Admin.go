package models

type Admin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (this *Admin) IsValid() bool {
	return this.Username == "admin" && this.Password == "admin"
}
