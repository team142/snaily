package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/bus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
	"github.com/team142/snaily/utils"
	"io/ioutil"
	"net/http"
)

func handleCreateItem(w http.ResponseWriter, r *http.Request, ID string) {
	item, err := model.ReadCloserToItem(r.Body)
	if err != nil {
		http.Error(w, "Invalid request or body", http.StatusBadRequest)
		logrus.Errorln(err)
		return
	}

	var created bool

	if created, err = bus.CreateItem(item, ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Errorln(err)
		return
	}

	if err := utils.WriteXToWriter(w, model.MessageNewItemResponseV1{OK: created, ID: item.ID}); err != nil {
		logrus.Errorln(err)
	}

}

func handleGetItem(w http.ResponseWriter, r *http.Request, ID string) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	itemReq := model.Item{}
	if err = json.Unmarshal(b, &itemReq); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Database connection problem", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	result := model.MessageGetItemResponseV1{}

	var item *model.Item
	if item, err = controller.GetItem(conn, itemReq.ID); err != nil {
		logrus.Errorln(err)
		http.Error(w, "Database read problem", http.StatusInternalServerError)
		return
	}

	if item == nil {
		logrus.Errorln("Item not found")
		http.Error(w, "Not found", http.StatusNotFound)
		return

	}

	result.Item = item

	uCreated, err := controller.GetUser(conn, item.CreatedBy)
	if err != nil {
		logrus.Errorln(err)
	} else {
		result.Users = append(result.Users, uCreated.GetUserMessage())
	}

	uFor, err := controller.GetUser(conn, item.WaitingFor)
	if err != nil {
		logrus.Errorln(err)
	} else {
		result.Users = append(result.Users, uFor.GetUserMessage())
	}

	if err = utils.WriteXToWriter(w, result); err != nil {
		logrus.Errorln(err)
	}

}

func handleGetMyItems(w http.ResponseWriter, r *http.Request, ID string) {
	result, err := bus.GetMyItems(ID)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	if err = utils.WriteXToWriter(w, result); err != nil {
		logrus.Errorln(err)
	}

}
