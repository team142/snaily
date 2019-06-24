package model

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

type MessageNewItemResponseV1 struct {
	OK bool   `json:"ok"`
	ID string `json:"id"`
}

type MessageOKResponseV1 struct {
	OK bool `json:"ok"`
}

type MessageRegisterResponseV1 MessageOKResponseV1

type MessageLoginResponseV1 struct {
	OK  bool   `json:"ok"`
	Key string `json:"key"`
	ID  string `json:"id"`
}

type MessageMyItemsResponseV1 struct {
	CreatedByMe  []*Item        `json:"createdByMe"`
	WaitingForMe []*Item        `json:"waitingForMe"`
	Users        MessageUsersV1 `json:"users"`
}

type MessageGetItemResponseV1 struct {
	Item  *Item          `json:"item"`
	Users MessageUsersV1 `json:"users"`
}

type MessageUsersV1 []*MessageUserV1

func (m *MessageUsersV1) Contains(ID string) bool {
	for _, r := range *m {
		if r.ID == ID {
			return true
		}
	}
	return false
}

type MessageUserV1 struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type MessageCloseItemResponseV1 struct {
	ID string `json:"id"`
}

func ReadCloserToMessageCloseItemResponseV1(body io.ReadCloser) (result *MessageCloseItemResponseV1, err error) {
	var b []byte
	b, err = ioutil.ReadAll(body)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	result = &MessageCloseItemResponseV1{}
	err = json.Unmarshal(b, result)
	if err != nil {
		logrus.Errorln(err)
	}
	return
}
