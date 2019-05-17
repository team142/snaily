package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
	"github.com/team142/snaily/utils"
	"io/ioutil"
	"net/http"
)

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	item := model.Item{}
	err = json.Unmarshal(b, &item)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	item.GenerateID()

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer conn.Close()

	if err := controller.InsertItem(conn, &item); err != nil {
		logrus.Errorln(err)
		if err := utils.WriteXToWriter(w, model.MessageNewItemResponseV1{OK: false}); err != nil {
			logrus.Errorln(err)
		}
		return
	}

	if err := utils.WriteXToWriter(w, model.MessageNewItemResponseV1{OK: true, ID: item.ID}); err != nil {
		logrus.Errorln(err)
	}

}

func handleGetItemAll(w http.ResponseWriter, r *http.Request) {

}
