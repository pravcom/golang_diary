package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"app-diary/internal/mocks"
	"app-diary/internal/models"
	"github.com/stretchr/testify/require"
)

func TestSaveDiary(t *testing.T) {
	cases := []struct {
		name           string
		method         string
		bodyRequest    string
		expectedStatus int
		responseError  string
	}{
		{
			name:   "Success",
			method: "POST",
			bodyRequest: `{
				"date": "2025-05-12",
				"login": "Denis",
				"project": "Lukoil",
				"task": "КШ_REP_CA_193.19",
				"description": "test",
				"time_hours": 8.0
}`,
			expectedStatus: http.StatusCreated,
			responseError:  "",
		},

		{
			name:   "Login failed",
			method: "POST",
			bodyRequest: `{
				"date": "2025-05-12",
				"login": "",
				"project": "Lukoil",
				"task": "КШ_REP_CA_193.19",
				"description": "test",
				"time_hours": 8.0
}`,
			expectedStatus: http.StatusBadRequest,
			responseError:  ErrLoginIsEmpty.Error(),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			saverMock := &mocks.MockStorage{ShouldFail: false}

			r := httptest.NewRequest(c.method, "/diary/create", strings.NewReader(c.bodyRequest))

			rr := httptest.NewRecorder()

			handler := SaveDiary(saverMock)
			handler(rr, r)

			require.Equal(t, rr.Code, c.expectedStatus)

			if !strings.Contains(rr.Body.String(), c.responseError) {
				t.Errorf("handler return unexpected body: got %v want %v", rr.Body.String(), c.responseError)
			}

			body := rr.Body.String()

			var diary models.Diary
			var diaryExpected models.Diary

			if c.responseError == "" {
				require.NoError(t, json.Unmarshal([]byte(body), &diary))
				require.NoError(t, json.Unmarshal([]byte(c.bodyRequest), &diaryExpected))

				require.Equal(t, diaryExpected, diary)
			}
		})
	}
}

func TestGetDiary(t *testing.T) {
	cases := []struct {
		name           string
		method         string
		bodyRequest    string
		expectedStatus int
		responseError  string
	}{
		{
			name:   "Success",
			method: "GET",
			bodyRequest: `{
				"date": "2025-05-12",
                "login": "Denis"
			}`,
			expectedStatus: http.StatusOK,
			responseError:  "",
		},

		{
			name:   "Fail method",
			method: "POST",
			bodyRequest: `{
				"date": "2025-05-12",
                "login": "Denis"
			}`,
			expectedStatus: http.StatusMethodNotAllowed,
			responseError:  fmt.Sprintf(ErrBadMethod.Error(), http.MethodPost),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var diary models.Diary

			diary.Date = "2025-05-12"
			diary.Login = "Denis"
			diary.Project = "X5"
			diary.Task = "test"
			diary.Description = "test"
			diary.TimeHours = 8.0

			getterMock := mocks.MockGetter{}

			getterMock.SetDiary(diary)

			r := httptest.NewRequest(c.method, "/diary/get", strings.NewReader(c.bodyRequest))

			rr := httptest.NewRecorder()

			handler := GetDiary(&getterMock)
			handler(rr, r)

			require.Equal(t, rr.Code, c.expectedStatus)

			if !strings.Contains(rr.Body.String(), c.responseError) {
				t.Errorf("handler return unexpected body: got %v want %v", rr.Body.String(), c.responseError)
			}

			if c.responseError == "" {
				body := rr.Body.String()

				var diaryGot models.Diary

				require.NoError(t, json.Unmarshal([]byte(body), &diaryGot))

				require.Equal(t, diaryGot, diary)
			}

		})
	}
}

func TestDeleteDiary(t *testing.T) {
	cases := []struct {
		name           string
		method         string
		bodyRequest    string
		expectedStatus int
		responseError  string
	}{
		{
			name:   "Success",
			method: "DELETE",
			bodyRequest: `{
				"date": "2025-05-12",
                "login": "Denis"
			}`,
			expectedStatus: http.StatusOK,
			responseError:  "",
		},

		{
			name:   "wrong method",
			method: http.MethodPost,
			bodyRequest: `{
				"date": "2025-05-12",
                "login": "Denis"
			}`,
			expectedStatus: http.StatusMethodNotAllowed,
			responseError:  fmt.Sprintf(ErrBadMethod.Error(), http.MethodPost),
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var diary models.Diary

			diary.Date = "2025-05-12"
			diary.Login = "Denis"
			diary.Project = "X5"
			diary.Task = "test"
			diary.Description = "test"
			diary.TimeHours = 8.0

			mockDeleter := mocks.MockDeleter{}

			mockDeleter.SetDiary(diary)

			r := httptest.NewRequest(c.method, "/diary/delete", strings.NewReader(c.bodyRequest))
			rr := httptest.NewRecorder()

			handler := DeleteDiary(&mockDeleter)
			handler(rr, r)

			require.Equal(t, c.expectedStatus, rr.Code)

			if !strings.Contains(rr.Body.String(), c.responseError) {
				t.Errorf("handler return unexpected body: got %v want %v", rr.Body.String(), c.responseError)
			}

		})
	}
}
