package api

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/team142/snaily/controller"
	"github.com/team142/snaily/db"
	"github.com/team142/snaily/model"
	"github.com/team142/snaily/utils"
	"io/ioutil"
	"net/http"
)

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	fmt.Println(string(b))
	user := model.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	//Persist or store
	conn, err := db.Connect()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer conn.Close()

	exists := controller.UserExists(conn, user.Email)
	if exists {
		err := utils.WriteXToWriter(w, model.MessageRegisterResponseV1{OK: false})
		if err != nil {
			logrus.Errorln(err)
		}

	} else {
		user.ID = uuid.NewV4().String()
		err = controller.InsertUser(conn, &user)
		if err != nil {
			logrus.Errorln(err)
			return
		}
		err := utils.WriteXToWriter(w, model.MessageRegisterResponseV1{OK: true})
		if err != nil {
			logrus.Errorln(err)
		}

	}
	return
}

func handleLoginUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	user := model.User{}
	err = json.Unmarshal(b, &user)
	if err != nil {
		logrus.Errorln(err)
		return
	}

	//Persist or store
	conn, err := db.Connect()
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer conn.Close()

	var dbUser *model.User
	dbUser, err = controller.GetUserByEmail(conn, user.Email)
	if dbUser != nil && dbUser.ID != "" {
		err := utils.WriteXToWriter(w, model.MessageLoginResponseV1{
			OK:  false,
			Key: dbUser.ID,
		})
		if err != nil {
			logrus.Errorln(err)
		}

	} else {
		err := utils.WriteXToWriter(w, model.MessageLoginResponseV1{OK: false})
		if err != nil {
			logrus.Errorln(err)
		}

	}
	return
}
