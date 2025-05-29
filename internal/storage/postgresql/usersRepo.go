package postgresql

import (
	"fmt"

	"app-diary/internal/models"
)

func (s *Storage) SaveUser(login, password string) (models.Users, error) {
	var users models.Users

	users.Login = login
	users.Password = password

	err := s.DB.Create(&users).Error
	if err != nil {
		return users, fmt.Errorf("Failed to create user: %w", err)
	}

	return users, nil
}

func (s *Storage) FindUserByLogin(login string) (models.Users, error) {
	var users models.Users

	err := s.DB.Where("login = ?", login).First(&users).Error
	if err != nil {
		return users, fmt.Errorf("Failed to find user by login: %w", err)
	}

	return users, nil
}
