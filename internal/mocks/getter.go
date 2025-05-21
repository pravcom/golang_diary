package mocks

import (
	"fmt"

	"app-diary/internal/models"
)

type MockGetter struct {
	Diary []models.Diary
}

func (m *MockGetter) GetDiary(login string, date string) (models.Diary, error) {
	for _, d := range m.Diary {
		if d.Login == login && d.Date == date {
			return d, nil
		}
	}
	var d models.Diary

	return d, fmt.Errorf("Dont find diary")
}

func (m *MockGetter) SetDiary(data models.Diary) {
	m.Diary = append(m.Diary, data)
}
