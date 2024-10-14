package database

import "users/server/users/models"

type Database interface {
	GetUserById(id int) (models.User, error)
	PostUser(user models.User) (int, error)
	PutUser(user models.User) error
	DeleteUser(id int) error
}
