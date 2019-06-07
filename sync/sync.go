package sync

import (
	"encoding/json"
	"github.com/just1689/entity-sync/entitysync"
	"github.com/just1689/entity-sync/entitysync/shared"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/bus"
	"github.com/team142/snaily/db"
)

func SetupSync(es entitysync.EntitySync) {

	es.RegisterEntityAndDBHandler("home", func(entityKey shared.EntityKey, secret string, handler shared.ByteHandler) {
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

}
