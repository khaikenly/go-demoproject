package repositories

import "go-demoproject/pkg/models"

type IUserRepository interface {
	Select() (users []models.User, err error)
	Insert(u *models.User) (err error)
}
