package userapi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"../../config"
	"../../dao"
	a "../../dao/abstractdao"
	"../../entities"
	"gopkg.in/mgo.v2/bson"
)

// @route   GET api/users/test
// @desc    Tests users route
// @access  Public
func Test(resp http.ResponseWriter, req *http.Request) {
	respondWithJson(resp, http.StatusOK, map[string]string{"msg": "Users Work"})
}

// @route   GET api/users/current
// @desc    Return current user
// @access  Private
func Current(resp http.ResponseWriter, req *http.Request) {
	stringToken := req.Header.Get("Authorization")
	var user entities.User
	claims, err := user.ObterLogado(stringToken)
	if err != nil {
		respondWithError(resp, http.StatusUnauthorized, err.Error())
		return
	}
	db, err := config.Connect()
	if err != nil {
		respondWithError(resp, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
		userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
		user, err2 := userDAO.FindByEmail(claims.Email)
		if err2 != nil {
			respondWithError(resp, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(resp, http.StatusOK, user)
		}
	}
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
		userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
		user, err2 := userDAO.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, user)
		}
	}
}

func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
		userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		user, err2 := userDAO.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, user)
		}
	}
}

// Add new user
func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
		userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
		var user entities.User
		user.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&user)
		if validErrs := user.IsValid(); len(validErrs) > 0 {
			respondWithValidationErrors(response, validErrs)
		} else {
			if err2 != nil {
				respondWithError(response, http.StatusBadRequest, err2.Error())
			} else {
				err3 := userDAO.Create(&user)
				if err3 != nil {
					respondWithError(response, http.StatusBadRequest, err3.Error())
				} else {
					respondWithJson(response, http.StatusOK, user)
				}
			}
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
		userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := userDAO.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, entities.User{})
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
		userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		var user entities.User
		err2 := json.NewDecoder(request.Body).Decode(&user)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := userDAO.Update(id, &user)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, user)
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
