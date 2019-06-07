package sync

import (
	"encoding/json"
	"github.com/just1689/entity-sync/entitysync"
	"github.com/just1689/entity-sync/entitysync/shared"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/bus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
)

func SetupSync(es entitysync.EntitySync) {

	/*
		`home` entity is for the users home page
	*/
	var entityHome shared.EntityType = "home"
	es.RegisterEntityAndDBHandler(entityHome, func(entityKey shared.EntityKey, secret string, handler shared.ByteHandler) {
		ok, ID := db.GlobalSessionCache.SessionValid(secret)
		if !ok {
			logrus.Errorln("Access denied based on secret", secret)
			return
		}
		result, err := bus.GetMyItems(ID)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		b, err := json.Marshal(result)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		handler(b)
	})
	//Give the controller a way to notify all of change
	controller.NotifyChangeHome = func(userID string) {
		es.Bridge.NotifyAllOfChange(shared.EntityKey{
			Entity: entityHome,
			ID:     userID,
		})
	}

	/*
		`items` entity is for the the items table in the db and its related entities
	*/
	var entityItems shared.EntityType = "items"
	es.RegisterEntityAndDBHandler(entityItems, func(entityKey shared.EntityKey, secret string, handler shared.ByteHandler) {
		ok, _ := db.GlobalSessionCache.SessionValid(secret)
		if !ok {
			logrus.Errorln("Access denied based on secret", secret)
			return
		}

		result, err := bus.GetItem(entityKey.ID)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		b, err := json.Marshal(result)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		handler(b)
	})
	//Give the controller a way to notify all of change
	controller.NotifyChangeItems = func(ID string) {
		es.Bridge.NotifyAllOfChange(shared.EntityKey{
			Entity: entityItems,
			ID:     ID,
		})
	}

}
