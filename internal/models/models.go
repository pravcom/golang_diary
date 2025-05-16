package models

type Diary struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Date        string  `json:"date" time_format:"2006-01-02"  gorm:"type:date"`
	Login       string  `json:"login" gorm:"index"`
	Project     string  `json:"project"`
	Task        string  `json:"task"`
	Description string  `json:"description"`
	TimeHours   float64 `json:"time_hours"`
}
