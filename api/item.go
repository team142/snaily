package api

import (
	"encoding/json"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
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
	item := model.Item{}
	err = json.Unmarshal(b, &item)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item.GenerateID()

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Database connection problem", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	u, err := controller.GetUserByEmail(conn, item.WaitingFor)
	if err != nil {
		http.Error(w, "Database read problem", http.StatusInternalServerError)
		logrus.Errorln(err)
		return
	}
	if u == nil {
		u = model.NewUserFromEmail(item.WaitingFor)
		if err = controller.InsertUser(conn, u); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database write problem", http.StatusInternalServerError)
			return
		}
	}
	item.WaitingFor = u.ID

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
			wgItems.Done()
			if !result.Users.Contains(ID) {
				u, err := controller.GetUser(conn, ID)
				if err != nil {
					logrus.Errorln(err)
					continue
				}
				if u == nil {
					logrus.Errorln("Could not find ID: ", ID)
					continue
				}
				result.Users = append(result.Users, &model.MessageUserV1{ID: u.ID, Email: u.Email, FirstName: u.FirstName, LastName: u.LastName})
			}
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
