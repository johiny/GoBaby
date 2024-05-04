package models

type UserStore struct {
	Users   []map[string]string
	MaxSize int
}
