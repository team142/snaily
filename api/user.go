package api

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
	"github.com/team142/snaily/utils"
	"io/ioutil"
	"net/http"
	"time"
)

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var err error
	if b, err = ioutil.ReadAll(r.Body); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := model.User{}
	if err = json.Unmarshal(b, &user); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	exists := controller.UserExists(conn, user.Email)
	if exists {
		if err := utils.WriteXToWriter(w, model.MessageRegisterResponseV1{OK: false}); err != nil {
			logrus.Errorln(err)
			return
		}
	} else {
		user.ID = uuid.NewV4().String()
		user.SaltAndSetPassword()
		if err = controller.InsertUser(conn, &user); err != nil {
			logrus.Errorln(err)
			http.Error(w, "Database write error", http.StatusInternalServerError)
			return
		}
		if err = utils.WriteXToWriter(w, model.MessageRegisterResponseV1{OK: true}); err != nil {
			logrus.Errorln(err)
		}
	}
	return
}

func handleLoginUser(w http.ResponseWriter, r *http.Request) {
	var b []byte
	var err error

	if b, err = ioutil.ReadAll(r.Body); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := model.User{}
	if err = json.Unmarshal(b, &user); err != nil {
		logrus.Errorln(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Persist or store
	conn, err := db.Connect(db.DefaultConfig)
	if err != nil {
		logrus.Errorln(err)
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	var dbUser *model.User
	dbUser, err = controller.GetUserByEmail(conn, user.Email)
	if dbUser != nil && dbUser.ID != "" && dbUser.CheckPassword(user.Password) {

		db.GlobalSessionCache.SetSession(dbUser.ID, dbUser.ID, 24*time.Hour)
		if err = utils.WriteXToWriter(w, model.MessageLoginResponseV1{
			OK:  true,
			Key: dbUser.ID,
		}); err != nil {
			logrus.Errorln(err)
		}

	} else {
		if err = utils.WriteXToWriter(w, model.MessageLoginResponseV1{OK: false}); err != nil {
			logrus.Errorln(err)
		}
	}
	return
}
