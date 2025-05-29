package handlers

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка `Authorization`
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		//Извлекаем токен (удаляем "Bearer ")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		//tokenCookie, err := r.Cookie("jwt")
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusUnauthorized)
		//	return
		//}
		//
		//tokenString := tokenCookie.String()

		// Парсим токен
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Если токен валиден, передаем управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}
