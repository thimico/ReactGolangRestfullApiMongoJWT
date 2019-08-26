package dao

import (
	. "./abstractdao"
	"gopkg.in/mgo.v2/bson"
	"../entities"
)

type ProfileDAO struct {
	AbstractDAO
}

func (profileDAO ProfileDAO) handle(handle string) (entities.Profile, error) {
	var profile entities.Profile
	err := profileDAO.DB.C(profileDAO.COLLECTION).FindId(bson.M{
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