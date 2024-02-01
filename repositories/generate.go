package repositories

import (
	"github.com/nguyenduclam1711/go-htmx-postgre-tryout/models"
)

var MapRepositories = struct {
	User CoreRepositoryer[models.User]
}{}

func GenerateMapRepositories() {
	MapRepositories.User = NewUserRepository()
}
