package jwt

import (
	"dot_conf/configs"
	"dot_conf/constants"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

const (
	expirationTime = time.Hour * 6
)

var (
	secretKey = configs.GetJwtSecretKey()
)

func Generate(username, role string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"iss": constants.AppName,
		"aud": role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(expirationTime).Unix(),
	})

	token, err := claims.SignedString(secretKey)
	if err != nil {
		log.Error("Error creating JWT token", err)
		return "", err
	}

	return token, nil
}

func Verify(role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				log.Info("JWT Token not found", reqToken)
				http.Error(w, "Invalid/Missing Token", http.StatusUnauthorized)
				return
			}
			reqToken = splitToken[1]
			claims := &jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				log.Error("Invalid token", err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			audience, err := claims.GetAudience()
			if err != nil {
				log.Error("Error getting the claims", err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if audience[0] != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func GetUsername(r *http.Request) (string, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		log.Info("JWT Token not found")
		return "", errors.New("missing token")
	}
	reqToken = splitToken[1]
	claims := &jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(reqToken, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		log.Error("Invalid token", err)
		return "", errors.New("invalid token")
	}

	return claims.GetSubject()
}
