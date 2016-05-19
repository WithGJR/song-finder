package models

type Admin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (this *Admin) IsValid() bool {
	return this.Username == "admin" && this.Password == "admin"
}
