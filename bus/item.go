package bus

import (
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
)

func CreateItem(item *model.Item) (created bool, errorMessage string) {
	item.GenerateID()

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		return false, "Database connection problem"
	}
	defer conn.Close()

	u, err := controller.GetUserByEmail(conn, item.WaitingFor)
	if err != nil {
		logrus.Errorln(err)
		return false, "Database read problem"
	}
	if u == nil {
		u = model.NewUserFromEmail(item.WaitingFor)
		if err = controller.InsertUser(conn, u); err != nil {
			logrus.Errorln(err)
			return false, "Database write problem"
		}
	}
	item.WaitingFor = u.ID

	if err := controller.InsertItem(conn, &item); err != nil {
		logrus.Errorln(err)
		return false, "Database write problem"
	}

	created = true
	return

}
