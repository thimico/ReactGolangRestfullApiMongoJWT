package dao

import (
	"../entities"
	. "./abstractdao"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type UserDAO struct {
	AbstractDAO
}

func (userDAO UserDAO) Create(user *entities.User) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	return userDAO.DB.C(userDAO.COLLECTION).Insert(&user)
}

func (userDAO UserDAO) CheckUsernameAndPassword(username, password string) bool {
	var user entities.User
	err := userDAO.DB.C(userDAO.COLLECTION).Find(bson.M{
		"username": username,
	}).One(&user)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
	}
}

func (userDAO UserDAO) CheckEmailAndPassword(email, password string) bool {
	var user entities.User
	err := userDAO.DB.C(userDAO.COLLECTION).Find(bson.M{
		"email": email,
	}).One(&user)
	if err != nil {
		return false
	} else {
		return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
	}
}

func (userDAO UserDAO) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := userDAO.DB.C(userDAO.COLLECTION).Find(bson.M{
		"email": email,
	}).One(&user)
	return user, err
}

func (userDAO UserDAO) FindByID(id string) (entities.User, error) {
	var user entities.User
	err := userDAO.DB.C(userDAO.COLLECTION).FindId(bson.M{
		"_id": id,
	}).One(&user)
	return user, err
}

func (dao UserDAO) FindOne(id string) (entities.User, error) {
	var entity entities.User
	err := dao.DB.C(dao.COLLECTION).FindId(bson.ObjectIdHex(id)).One(&entity)
	return entity, err
}