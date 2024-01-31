package repositories

import "github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// TODO: create user method
func (ur *UserRepository) Create(user models.User) {
	generateInsertQuery(user)
}
