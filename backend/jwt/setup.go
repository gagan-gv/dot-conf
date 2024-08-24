package jwt

import (
	"dot_conf/configs"
	"dot_conf/constants"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
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

func Verify(role string) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			tokenStr := cookie.Value
			claims := &jwt.MapClaims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
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
