package api

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/bus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
	"github.com/team142/snaily/utils"
	"io/ioutil"
	"net/http"
	"sync"
)

func handleCreateItem(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item := &model.Item{}
	err = json.Unmarshal(b, item)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created, errorMsg := bus.CreateItem(item)
	if !created {
		http.Error(w, errorMsg, http.StatusInternalServerError)
		logrus.Errorln(errorMsg)
		return
	}

	if err := utils.WriteXToWriter(w, model.MessageNewItemResponseV1{OK: true, ID: item.ID}); err != nil {
		logrus.Errorln(err)
	}

}

func handleGetItem(w http.ResponseWriter, r *http.Request) {

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

func handleGetMyItems(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var err error
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req := model.MessageMyItemsRequestV1{}
	if err = json.Unmarshal(b, &req); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var conn *pgx.Conn
	if conn, err = db.Connect(db.DefaultConfig); err != nil {
		logrus.Errorln(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	//TODO: FOR NOW THE KEY IS THE USER ID. THIS MUST CHANGE

	user, err := controller.GetUser(conn, req.Key)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Authentication error", http.StatusBadRequest)
		return
	}

	result := model.MessageMyItemsResponseV1{
		CreatedByMe:  make([]*model.Item, 0),
		WaitingForMe: make([]*model.Item, 0),
		Users:        make(model.MessageUsersV1, 0),
	}

	var wgItems sync.WaitGroup

	in := make(chan string, 10)

	go func(in chan string) {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		for ID := range in {
			if !result.Users.Contains(ID) {
				u, err := controller.GetUser(conn, ID)
				if err != nil {
					logrus.Errorln(err)
					wgItems.Done()
					continue
				}
				if u == nil {
					logrus.Errorln("Could not find ID: ", ID)
					wgItems.Done()
					continue
				}
				result.Users = append(result.Users, u.GetUserMessage())
			}
			wgItems.Done()
		}

	}(in)

	wgItems.Add(1)
	go func() {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		var out chan *model.Item
		if out, err = controller.GetItemsByCreatedBy(conn, user.ID); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		for item := range out {
			wgItems.Add(1)
			in <- item.WaitingFor
			wgItems.Add(1)
			in <- item.CreatedBy
			result.CreatedByMe = append(result.CreatedByMe, item)
		}
		wgItems.Done()
	}()

	wgItems.Add(1)
	go func() {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		var out chan *model.Item
		if out, err = controller.GetItemsByWaitingFor(conn, user.ID); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database query error", http.StatusInternalServerError)
			return
		}
		for item := range out {
			wgItems.Add(1)
			in <- item.WaitingFor
			wgItems.Add(1)
			in <- item.CreatedBy
			result.WaitingForMe = append(result.WaitingForMe, item)
		}
		wgItems.Done()
	}()

	wgItems.Wait()
	close(in)

	if err = utils.WriteXToWriter(w, result); err != nil {
		logrus.Errorln(err)
	}

}
