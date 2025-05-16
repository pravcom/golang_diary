package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app-diary/internal/models"
)

type SaverDiary interface {
	SaveDiary(login, date, project, task, description string, timeHours float64) (models.Diary, error)
}

type GetterDiary interface {
	GetDiary(login string, date string) (models.Diary, error)
}

type DeleterDiary interface {
	DeleteDiary(login string, date string) error
}

func SaveDiary(creater SaverDiary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Метод %s не поддерживается", r.Method)
			return
		}

		var diary models.Diary

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&diary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(diary)

		diary, err = creater.SaveDiary(diary.Login, diary.Date, diary.Project, diary.Task, diary.Description, diary.TimeHours)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(diary)
	}
}

func GetDiary(getter GetterDiary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Метод %s не поддерживается", r.Method)
			return
		}

		var diary models.Diary

		err := json.NewDecoder(r.Body).Decode(&diary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		diary, err = getter.GetDiary(diary.Login, diary.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(diary)
	}
}

func DeleteDiary(deleter DeleterDiary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Метод %s не поддерживается", r.Method)
			return
		}

		var diary models.Diary

		err := json.NewDecoder(r.Body).Decode(&diary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = deleter.DeleteDiary(diary.Login, diary.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusOK)

	}
}
