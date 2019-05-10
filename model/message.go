package model

type MessageRegisterResponseV1 struct {
	OK bool `json:"ok"`
}

type MessageLoginResponseV1 struct {
	OK  bool   `json:"ok"`
	Key string `json:"key"`
}
