package jwtauth

import (
	"encoding/json"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"../../entities"
	"../../config"
	"../../dao"
	a "../../dao/abstractdao"
)

var secretKey = "MySecretKey"

type JWTToken struct {
	Token string `json:"token"`
}

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var user entities.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		respondWithError(response, http.StatusUnauthorized, err.Error())
	} else {
		db, err2 := config.Connect()
		if err2 != nil {
			respondWithError(response, http.StatusUnauthorized, err2.Error())
		} else {
			abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
			userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
			if validErrs := user.IsLoginValid(); len(validErrs) > 0 {
				respondWithValidationErrors(response, validErrs)
			} else {
				valid := userDAO.CheckEmailAndPassword(user.Email, user.Password)
				if valid {
					token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
						"id":    user.Id,
						"name":    user.Name,
						"email":    user.Email,
						"password": user.Password,
						"exp":      time.Now().Add(time.Hour * 72).Unix(),
					})
					tokenString, err2 := token.SignedString([]byte(secretKey))
					if err2 != nil {
						respondWithError(response, http.StatusUnauthorized, err.Error())
					} else {
						respondWithJson(response, http.StatusOK, JWTToken{Token: tokenString})
					}
				} else {
					respondWithError(response, http.StatusUnauthorized, "User Invalid")
				}
			}
		}

	}
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithValidationErrors(w http.ResponseWriter, validErrs interface{}) {
	err := map[string]interface{}{"errors": validErrs}
	w.Header().Set("Content-type", "applciation/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(err)
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
