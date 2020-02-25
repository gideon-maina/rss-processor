package serverss

import (
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

var SigningKey = []byte("thisisabigsecretkey====***okay?")

type Token struct {
	JWTTokenValue string
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header.Get("Token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					log.Println("Cannot verify token")
					return nil, fmt.Errorf("Error in token verification")
				}
				return SigningKey, nil
			})
			if err != nil {
				log.Println("Error in IsAuthorized check", err)
				fmt.Fprintf(w, "Your Token is InValid")
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Token is missing from your request")
		}
	})
}

// Genearate a JSON Web Token to securely reach our protected endpoints.
// This will be called by a public endpoint and will provide them with a valid token to
// then query the endpoints (N.B This is for demo purposes of auth and will need to change for prod app)
func GetToken(w http.ResponseWriter, r *http.Request) {
	tokenToGive, err := generateJWT()
	if err != nil {
		log.Println("Failed to get JWT")
		fmt.Fprintf(w, "Failed to generate web token. Try again later.")
		return
	}
	tokenResponse := Token{
		JWTTokenValue: tokenToGive}
	jsonToken, err := json.Marshal(tokenResponse)
	if err != nil {
		log.Println("Failed to marshal token")
		fmt.Fprintf(w, "Failed to generate web token. Try again later.")
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonToken)

}

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "Demo Auth User"
	claims["exp"] = time.Now().Add(time.Minute * 100).Unix()

	tokenString, err := token.SignedString(SigningKey)

	if err != nil {
		log.Println("Something went wrong in getting token string", err)
		return "", err
	}

	return tokenString, nil
}
