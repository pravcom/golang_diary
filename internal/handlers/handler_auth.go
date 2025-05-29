package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"app-diary/internal/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type SaverUser interface {
	SaveUser(login, password string) (models.Users, error)
}

type GetterUser interface {
	FindUserByLogin(login string) (models.Users, error)
}

var JWTSecret = []byte("secret_key")

func SaveUser(saver SaverUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.Users

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

		if err != nil {
			http.Error(w, "could not hash password", http.StatusBadRequest)
			return
		}

		user.Password = string(hashPass)

		userCreated, err := saver.SaveUser(user.Login, user.Password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		claims := jwt.MapClaims{
			"login": userCreated.Login,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		}

		tokenString, err := generateToken(claims)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
			"user":  userCreated.Login,
		})

	}
}

func Login(getter GetterUser) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var inputUser models.Users

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&inputUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		user, err := getter.FindUserByLogin(inputUser.Login)
		if err != nil {
			http.Error(w, "Dont find user", http.StatusBadRequest)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputUser.Password)); err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		claims := jwt.MapClaims{
			"login": inputUser.Login,
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		}

		tokenString, err := generateToken(claims)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// 3. Устанавливаем заголовок Authorization
		w.Header().Set("Authorization", "Bearer "+tokenString)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})

	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Удаляем токен из заголовка Authorization (для клиентских приложений)
		w.Header().Set("Authorization", "")

		//Возвращаем успешный статус
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Successfully logged out",
		})
	}
}

func generateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return tokenString, fmt.Errorf("Could not generate token")
	}

	return tokenString, nil
}
