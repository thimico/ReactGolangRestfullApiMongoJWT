package profileapi

import (
	"ReactGolangRestfullApiMongoJWT/src/config"
	"ReactGolangRestfullApiMongoJWT/src/dao"
	a "ReactGolangRestfullApiMongoJWT/src/dao/abstractdao"
	"ReactGolangRestfullApiMongoJWT/src/entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
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
				Id       			bson.ObjectId 				`json:"id" bson:"_id"`
				User 				entities.User 				`json:"user"`
				Handle 				string						`json:"handle"`
				Company				string						`json:"company"`
				Website 			string        				`json:"website" bson:"website"`
				Location 			string        				`json:"location" bson:"location"`
				Status   			string          			`json:"status" bson:"status"`
				Skills    			string        				`json:"skills" bson:"skills"`
				Bio     			string        			 	`json:"bio" bson:"bio"`
				Githubusername     	string        				`json:"githubusername" bson:"githubusername"`


			}{ profile.Id, userSearch, profile.Handle, profile.Company, profile.Website, profile.Location,
				profile.Status, profile.Skills, profile.Bio, profile.Githubusername}

			if err = json.NewEncoder(resp).Encode(data); err != nil {
				log.Fatalln("Error on encode json of address: ", err)
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

func Handle(response http.ResponseWriter, request *http.Request) {
	db, err := config.Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	} else {
		abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "profile"}
		profileDAO := dao.ProfileDAO{AbstractDAO: abstractDAO}
		vars := mux.Vars(request)
		handle := vars["handle"]
		profile, err2 := profileDAO.Handle(handle)
		log.Println(profile)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
			return
		} else {
			abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
			userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
			log.Println(profile.User)
			userSearch, err3 := userDAO.FindOne("5d62e72e26b8cc61b1e96951")
			if err3 != nil {
				respondWithError(response, http.StatusBadRequest, err3.Error())
			} else {
				data := struct {
					User 		entities.User `json:"user"`
					Handle 		string
					Company		string


				}{ userSearch, profile.Handle, profile.Company}
				if err = json.NewEncoder(response).Encode(data); err != nil {
					log.Fatalln("Error on encode json of address: ", err)
					respondWithError(response, http.StatusBadRequest, "Falha ao recuperar a unidade de entrega (cod=2.1)")
					return
				}
			}
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
				stringToken := request.Header.Get("Authorization")
				var user entities.User
				claims, err := user.ObterLogado(stringToken)
				log.Print(claims.Email)
				if err != nil {
					respondWithError(response, http.StatusUnauthorized, err.Error())
					return
				}
				db, err := config.Connect()
				abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
				userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
				userSearch, _ := userDAO.FindByEmail(claims.Email)
				profile.User = userSearch.Id
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
			stringToken := request.Header.Get("Authorization")
			var user entities.User
			claims, err := user.ObterLogado(stringToken)
			log.Print(claims.Email)
			if err != nil {
				respondWithError(response, http.StatusUnauthorized, err.Error())
				return
			}
			db, err := config.Connect()
			abstractDAO := a.AbstractDAO{DB: db, COLLECTION: "user"}
			userDAO := dao.UserDAO{AbstractDAO: abstractDAO}
			userSearch, _ := userDAO.FindByEmail(claims.Email)
			profile.User = userSearch.Id
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