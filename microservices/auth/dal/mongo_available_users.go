package dal

import (
	"auth/config"
	"fmt"
	"log"
	"time"

	"github.com/headzoo/surf/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewDbAvailableUsers(settings *config.Settings) *DbAvailableUserRepository {
	return &DbAvailableUserRepository{
		settings: settings,
	}
}

type DbAvailableUserRepository struct {
	settings *config.Settings
}

type userAttempt struct {
	Id         bson.ObjectId `bson:"_id"`
	UserName   string        `bson:"UserName"`
	Date       time.Time     `bson:"Date"`
	ErrMessage string        `bson:"ErrMessage"`
}

func (r *DbAvailableUserRepository) UserIsAvailable(userName string) bool {
	err := r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("AvailableMobileUsers")
		if n, err := collection.Find(bson.M{"UserName": userName}).Count(); err != nil || n < 1 {
			log.Println(n, "записей", err)
			return errors.New("not found")
		}
		return nil
	})
	return err == nil
}

func (r *DbAvailableUserRepository) SaveUser(userName string, errMessage string) error {
	err := r.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(r.settings.MongoDbNameAccount).C("LoginAttempt")

		attempt := userAttempt{
			Id:         bson.NewObjectId(),
			UserName:   userName,
			Date:       time.Now(),
			ErrMessage: errMessage,
		}

		if err := collection.Insert(attempt); err != nil {
			return errors.New("cant write attempt")
		}
		return nil
	})
	return err
}

func (r *DbAvailableUserRepository) mongoSession(action func(*mgo.Session) error) error {
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
