package models

type UserStore struct {
	Users       []map[string]string
	MaxSize     int
	CurrentUser map[string]string
}
