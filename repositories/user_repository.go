package repositories

import "github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"

func NewUserRepository() CoreRepositoryer[models.User] {
	return NewRepository[models.User](models.User{})
}
