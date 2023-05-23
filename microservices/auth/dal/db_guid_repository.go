package dal

import (
	"auth/config"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewDbUUIDRepository(settings *config.Settings) UUIDRepository {
	return DbUUIDRepository{
		settings: settings,
	}
}

type DbUUIDRepository struct {
	settings *config.Settings
}

func (d DbUUIDRepository) Save(uuid UUID) error {
	var resErr error
	resErr = d.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(d.settings.MongoDbNameAccount).C("uuids")
		if err := collection.Insert(uuid); err != nil {
			return fmt.Errorf("save uuid token error: %s", err.Error())
		}
		return nil
	})
	return resErr
}

func (d DbUUIDRepository) Get(uuid string) (*UUID, error) {
	var (
		result UUID
		resErr error
	)
	resErr = d.mongoSession(func(session *mgo.Session) error {
		collection := session.DB(d.settings.MongoDbNameAccount).C("uuids")
		if err := collection.Find(bson.M{"uuid": uuid}).One(&result); err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			return fmt.Errorf("get uuid error: %s", err.Error())
		}
		return nil
	})
	if resErr != nil {
		return nil, resErr
	}
	return &result, nil
}

func (d *DbUUIDRepository) mongoSession(action func(*mgo.Session) error) error {
	session, err := mgo.Dial(d.settings.MongoDbConnectionString)
	if err != nil {
		return fmt.Errorf("MongoDB session creation error: %s", err.Error())
	}
	defer session.Close()

	if err := action(session); err != nil {
		return err
	}
	return nil
}
