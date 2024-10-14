package database

import (
	"errors"
	"users/server/users/models"
)

type MockDatabase struct {
}

func InitMockDatabase() *MockDatabase {
	return &MockDatabase{}
}

func (pg MockDatabase) GetUserById(id int) (models.User, error) {
	user := models.User{
		ID:   1,
		Name: "John",
		Age:  33,
	}
	return user, nil
}

func (pg MockDatabase) PostUser(user models.User) (int, error) {
	return 1, nil
}

func (pg MockDatabase) PutUser(user models.User) error {
	if user.ID == 0 {
		return errors.New("wrong id")
	}
	return nil
}

func (pg MockDatabase) DeleteUser(id int) error {
	return nil
}
