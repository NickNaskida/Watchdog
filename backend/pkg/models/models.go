package models

type Alert struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Message  string `json:"message"`
}
