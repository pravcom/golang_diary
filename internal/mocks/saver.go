package mocks

import (
	"errors"

	"app-diary/internal/models"
)

type MockStorage struct {
	ShouldFail bool         // Флаг для эмуляции ошибки
	SavedDiary models.Diary // Сохранённая запись (для проверки в тестах)
	CallCount  int          // Количество вызовов
}

func (m *MockStorage) SaveDiary(login, date, project, task, description string, timeHours float64) (models.Diary, error) {
	m.CallCount++

	diary := models.Diary{
		Login:       login,
		Date:        date,
		Project:     project,
		Task:        task,
		Description: description,
		TimeHours:   timeHours,
	}

	m.SavedDiary = diary // Сохраняем для последующей проверки

	if m.ShouldFail {
		return diary, errors.New("mock save error")
	}

	return diary, nil
}
