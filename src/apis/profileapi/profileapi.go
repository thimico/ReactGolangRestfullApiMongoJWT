package profileapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"../../config"
	"../../dao"
	a "../../dao/abstractdao"
	"../../entities"
	"gopkg.in/mgo.v2/bson"
)

// @route   GET api/profile/test
// @desc    Tests profiles route
// @access  Public
func Test(resp http.ResponseWriter, req *http.Request) {
	respondWithJson(resp, http.StatusOK, map[string]string{"msg": "Profile Work"})
}

// @route   GET api/profile/current
// @desc    Current profile route
// @access  Public
func Current(resp http.ResponseWriter, req *http.Request) {
	stringToken := req.Header.Get("Authorization")
	var user entities.User
	claims, err := user.ObterLogado(stringToken)
	log.Print(claims.Email)
	if err != nil {
		respondWithError(resp, http.StatusUnauthorized, err.Error())
		return
	}
	db, err := config.Connect()
	abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
	userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
	userSearch, err2 := userDAO.FindByEmail(claims.Email)
	log.Println(userSearch)
	log.Println(err2)
	if err2 != nil {
		respondWithError(resp, http.StatusBadRequest, err2.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		log.Println(userSearch)
		profile, err3 := profileDAO.FindByUser(userSearch)
		if err3 != nil {
			respondWithError(resp, http.StatusBadRequest, err3.Error())
		} else {
			data := struct {
				User 		entities.User `json:"user"`
				Handle 		string
				Company		string


			}{ userSearch, profile.Handle, profile.Company}

			if err = json.NewEncoder(resp).Encode(data); err != nil {
				log.Fatalln("Error on encode json of address: %v", err)
				respondWithError(resp, http.StatusBadRequest, "Falha ao recuperar a unidade de entrega (cod=2.1)")
				return
			}
		}
	}
}


func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		profile, err2 := profileDAO.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, profile)
		}
	}
}

func Find(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		profile, err2 := profileDAO.Find(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, profile)
		}
	}
}

// Add new profile
func Create(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {

		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		var profile entities.Profile
		profile.Id = bson.NewObjectId()
		err2 := json.NewDecoder(request.Body).Decode(&profile)
		if validErrs := profile.IsValid(); len(validErrs) > 0 {
			respondWithValidationErrors(response, validErrs)
		} else {
			if err2 != nil {
				respondWithError(response, http.StatusBadRequest, err2.Error())
			} else {
				err3 := profileDAO.Create(&profile)
				if err3 != nil {
					respondWithError(response, http.StatusBadRequest, err3.Error())
				} else {
					respondWithJson(response, http.StatusOK, profile)
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
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		err2 := profileDAO.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			respondWithJson(response, http.StatusOK, entities.Profile{})
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		id := vars["id"]
		var profile entities.Profile
		err2 := json.NewDecoder(request.Body).Decode(&profile)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			err3 := profileDAO.Update(id, &profile)
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				respondWithJson(response, http.StatusOK, profile)
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