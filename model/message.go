package model

type MessageNewItemResponseV1 struct {
	OK bool   `json:"ok"`
	ID string `json:"id"`
}

type MessageRegisterResponseV1 struct {
	OK bool `json:"ok"`
}

type MessageLoginResponseV1 struct {
	OK  bool   `json:"ok"`
	Key string `json:"key"`
}

type MessageMyItemsRequestV1 struct {
	Key string `json:"key"`
}

type MessageMyItemsResponseV1 struct {
	CreatedByMe  []*Item        `json:"createdByMe"`
	WaitingForMe []*Item        `json:"WaitingForMe"`
	Users        MessageUsersV1 `json:"users"`
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
	ID        string `json:"ID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
