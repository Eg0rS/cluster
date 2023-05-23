package dal

import (
	"auth/config"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewDbRefreshTokenRepository(settings *config.Settings) RefreshTokenRepository {
	return &DbRefreshTokenRepository{settings}
}

type DbRefreshTokenRepository struct {
	settings *config.Settings
}

func (r *DbRefreshTokenRepository) Save(token *RefreshToken) error {
	var resErr error
	resErr = r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("RefreshTokens")
		if err := collection.Insert(token); err != nil {
			return fmt.Errorf("save refresh token error: %s", err.Error())
		}
		return nil
	})
	return resErr
}

func (r *DbRefreshTokenRepository) Get(token string, userId string) (result *RefreshToken, err error) {
	err = r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("RefreshTokens")
		return collection.Find(bson.M{"token": token, "user_id": userId}).One(&result)
	})
	return
}

func (r *DbRefreshTokenRepository) TokenExists(token string) (b2 bool) {
	count := 0
	err := r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("RefreshTokens")
		var err error
		count, err = collection.Find(bson.M{"token": token}).Count()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return count == 0
}

func (r *DbRefreshTokenRepository) AccessTokenExists(token string) (b2 bool) {
	count := 0
	err := r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("RefreshTokens")
		var err error
		count, err = collection.Find(bson.M{"access_token": token}).Count()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false
	}
	return count == 0
}

func (r *DbRefreshTokenRepository) Delete(token string, userId string) (err error) {
	err = r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("RefreshTokens")
		if err := collection.Remove(bson.M{"token": token, "user_id": userId}); err != nil {
			return fmt.Errorf("delete refresh token error: %s", err.Error())
		}
		return nil
	})
	return err
}

func (r *DbRefreshTokenRepository) DeleteByUserId(userId string) (err error) {
	err = r.mongoSession(func(session *mgo.Session) error {
		if _, err := session.DB(r.settings.MongoDbNameAccount).C("RefreshTokens").RemoveAll(
			bson.M{"user_id": userId}); err != nil {
			return fmt.Errorf("delete refresh token error: %s", err.Error())
		}
		return nil
	})
	return err
}

func (r *DbRefreshTokenRepository) mongoSession(action func(*mgo.Session) error) error {
	session, err := mgo.Dial(r.settings.MongoDbConnectionString)
	if err != nil {
		return fmt.Errorf("MongoDB session creation error: %s", err.Error())
	}
	defer session.Close()

	if err := action(session); err != nil {
		return err
	}
	return nil
}
