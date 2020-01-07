package entities

import (
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"time"
)

type Profile struct {
	Id       			bson.ObjectId 				`json:"id" bson:"_id"`
	User      			bson.ObjectId 				`json:"user" bson:"user"`
	Handle 	 			string        				`json:"handle" bson:"handle"`
	Company    			string        				`json:"company" bson:"company"`
	Website 			string        				`json:"website" bson:"website"`
	Location 			string        				`json:"location" bson:"location"`
	Status   			string          			`json:"status" bson:"status"`
	Skills    			string        				`json:"skills" bson:"skills"`
	Bio     			string        			 	`json:"bio" bson:"bio"`
	Githubusername     	string        				`json:"githubusername" bson:"githubusername"`
	Experience    		*[]Experience        		`json:"experience" bson:"experience"`
	Education     		*[]Education        		`json:"education" bson:"education"`
	Social     			*Social        				`json:"social" bson:"social"`
	Date     			time.Time         			`json:"date" bson:"date"`
}

type Experience struct {
	Id     				bson.ObjectId 		`bson:"_id"`
	Title  				string        		`bson:"title"`
	Company  			string       		`bson:"company"`
	Location  			string       		`bson:"location"`
	From  				time.Time       	`bson:"from"`
	To  				time.Time        	`bson:"to"`
	Current 			bool          		`bson:"current"`
	Description 		string      		`bson:"description"`
}

type Education struct {
	Id     				bson.ObjectId 		`bson:"_id"`
	School  			string        		`bson:"school"`
	Degree  			string       		`bson:"degree"`
	Fieldofstudy  		string    			`bson:"fieldofstudy"`
	From  				time.Time       	`bson:"from"`
	To  				time.Time        	`bson:"to"`
	Current 			bool          		`bson:"current"`
	Description 		string      		`bson:"description"`
}

type Social struct {
	Youtube  string        	`bson:"youtube"`
	Twitter  string       	`bson:"twitter"`
	Facebook  string       	`bson:"facebook"`
	Linkedin string      	`bson:"linkedin"`
	Instagram string      	`bson:"instagram"`
}


func (e *Profile) New() IEntity {
	return e
}


func (a Profile) IsValid() (errs url.Values) {
	// check if the name empty
	errs = make(url.Values)
	if a.Handle == "" {
		errs.Add("handle", "The handle is required!")
	}

	return errs
}
