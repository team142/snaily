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
