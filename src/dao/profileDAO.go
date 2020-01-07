package dao

import (
	. "ReactGolangRestfullApiMongoJWT/src/dao/abstractdao"
	"gopkg.in/mgo.v2/bson"
	"ReactGolangRestfullApiMongoJWT/src/entities"
)

type ProfileDAO struct {
	AbstractDAO
}

func (profileDAO ProfileDAO) Handle(handle string) (entities.Profile, error) {
	var profile entities.Profile
	err := profileDAO.DB.C(profileDAO.COLLECTION).Find(bson.M{
		"handle": handle,
	}).One(&profile)
	return profile, err
}

func (profileDAO ProfileDAO) FindByUser(user entities.User) (entities.Profile, error) {
	var profile entities.Profile
	err := profileDAO.DB.C(profileDAO.COLLECTION).Find(bson.M{
		"user": user.Id,
	}).One(&profile)
	return profile, err
}