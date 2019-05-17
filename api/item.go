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

	result := model.MessageMyItemsResponseV1{}

	var wg sync.WaitGroup

	in := make(chan string, 10)
	stop := make(chan bool)

	go func(in chan string, stop chan bool) {
		var conn *pgx.Conn
		if conn, err = db.Connect(db.DefaultConfig); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer conn.Close()
		for {
			select {
			case ID := <-in:
				if !result.Users.Contains(ID) {
					u, err := controller.GetUser(conn, ID)
					if err != nil {
						logrus.Errorln(err)
						continue
					}
					result.Users = append(result.Users, &model.MessageUserV1{ID: u.ID, Email: u.Email, FirstName: u.FirstName, LastName: u.LastName})
				}
			case <-stop:
				return
			}
		}

	}(in, stop)

	wg.Add(1)
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
			result.CreatedByMe = append(result.CreatedByMe, item)
		}
		wg.Done()
	}()

	wg.Add(1)
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
			result.WaitingForMe = append(result.WaitingForMe, item)
		}
		wg.Done()
	}()

	wg.Wait()
	stop <- true
	close(in)
	close(stop)

	if err = utils.WriteXToWriter(w, result); err != nil {
		logrus.Errorln(err)
	}

}
