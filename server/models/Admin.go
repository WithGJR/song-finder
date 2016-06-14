package models

import (
	"song-finder/server/constants"
)

type Admin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (this *Admin) IsValid() bool {
	return this.Username == constants.AdminUsername && this.Password == constants.AdminPassword
}
