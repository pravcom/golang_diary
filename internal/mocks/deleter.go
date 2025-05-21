package mocks

import (
	"fmt"

	"app-diary/internal/models"
)

type MockDeleter struct {
	Diary []models.Diary
}

func (d *MockDeleter) DeleteDiary(login string, date string) error {
	var diaryCopy []models.Diary
	var isFound bool

	for _, d := range d.Diary {
		if d.Login == login && d.Date == date {
			isFound = true
			continue
		}
		diaryCopy = append(diaryCopy, d)
	}

	if isFound == false {
		return fmt.Errorf("Not found")
	}

	d.Diary = diaryCopy
	return nil

}

func (m *MockDeleter) SetDiary(data models.Diary) {
	m.Diary = append(m.Diary, data)
}
