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
	CreatedByMe  []Item          `json:"createdByMe"`
	WaitingForMe []Item          `json:"WaitingForMe"`
	Users        []MessageUserV1 `json:"users"`
}

type MessageUserV1 struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
