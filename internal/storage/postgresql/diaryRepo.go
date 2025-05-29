package postgresql

import (
	"fmt"

	"app-diary/internal/models"
)

func (s Storage) SaveDiary(login, date, project, task, description string, timeHours float64) (models.Diary, error) {
	diary := models.Diary{
		Login:       login,
		Date:        date,
		Project:     project,
		Task:        task,
		Description: description,
		TimeHours:   timeHours,
	}
	fmt.Printf("SaveDiary\n %+v", diary)
	fmt.Printf("login ")

	err := s.DB.Create(&diary).Error

	if err != nil {
		return diary, fmt.Errorf("Failed to create diary: %w", err)
	}

	return diary, nil

}

func (s Storage) GetDiary(login string, date string) (models.Diary, error) {
	var diary models.Diary

	fmt.Printf("Get diary: login: %s, date: %s\n", login, date)

	err := s.DB.Where("login = ? and date = ?", login, date).First(&diary).Error
	if err != nil {
		return diary, fmt.Errorf("Failed to get diary: %w", err)
	}
	return diary, nil
}

func (s Storage) DeleteDiary(login string, date string) error {
	var diary models.Diary
	fmt.Printf("Delted diary: login: %s, date: %s\n", login, date)
	err := s.DB.Where("login = ? and date = ?", login, date).First(&diary).Error

	if err != nil {
		return fmt.Errorf("Failed to delete diary: %w", err)
	}

	err = s.DB.Delete(&diary).Error

	if err != nil {
		return fmt.Errorf("Failed to delete diary: %w", err)
	}

	return nil

}
