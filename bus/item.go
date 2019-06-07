package bus

import (
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
)

func CreateItem(item *model.Item, createdBy string) (created bool, err error) {
	item.GenerateID()
	item.CreatedBy = createdBy

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		return false, errors.New("Database connection problem")
	}
	defer conn.Close()

	u, err := controller.GetUserByEmail(conn, item.WaitingFor)
	if err != nil {
		logrus.Errorln(err)
		return false, errors.New("Database read problem")
	}
	if u == nil {
		u = model.NewUserFromEmail(item.WaitingFor)
		if err = controller.InsertUser(conn, u); err != nil {
			logrus.Errorln(err)
			return false, errors.New("Database write problem")
		}
	}
	item.WaitingFor = u.ID

	if err := controller.InsertItem(conn, item); err != nil {
		logrus.Errorln(err)
		return false, errors.New("Database write problem")
	}

	created = true
	return

}
