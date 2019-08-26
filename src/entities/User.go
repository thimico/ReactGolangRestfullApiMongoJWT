package entities

import (
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"regexp"
	_ "net/url"
	"github.com/dgrijalva/jwt-go"
)

var regexpEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name 	 string        	`json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email" validate:"required,email"`
	Password string        `json:"password" bson:"password"`
	Password2 string        `json:"password2" bson:"password2"`
	Avatar   string          `json:"avatar" bson:"avatar"`
	Date     string        `json:"date" bson:"date"`
}

type Claims struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}
var jwtKey = []byte("MySecretKey")

func (e *User) New() IEntity {
	return e
}

func (e *User) ObterLogado(stringToken string) (*Claims, error) {
	claims := &Claims{}
	tkn, err :=  jwt.ParseWithClaims(stringToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return claims, err
		}
		return claims, err
	}
	if !tkn.Valid {
		return claims, err
	}

	return claims, nil
}

func (a User) IsValid() (errs url.Values) {
	// check if the name empty
	errs = make(url.Values)
	if a.Name == "" {
		errs.Add("name", "The name is required!")
	}
	// check the name field is between 3 to 120 chars
	if len(a.Name) < 2 || len(a.Name) > 60 {
		errs.Add("name", "The name field must be between 2-60 chars!")
	}
	if a.Email == "" {
		errs.Add("email", "The email field is required!")
	}

	if !regexpEmail.MatchString(a.Email) {
		errs.Add("email", "Email is invalid")
	}

	if a.Password == "" {
		errs.Add("password", "Password must be at least 6 characters")
	}

	if a.Password2 == "" {
		errs.Add("password2", "Confirm Password field is required")
	}

	return errs
}

func (a User) IsLoginValid() (errs url.Values) {
	// check if the name empty
	errs = make(url.Values)

	if a.Email == "" {
		errs.Add("email", "The email field is required!")
	}

	if !regexpEmail.MatchString(a.Email) {
		errs.Add("email", "Email is invalid")
	}

	if a.Password == "" {
		errs.Add("password", "Password field is required")
	}

	return errs
}