package models

type User struct {
	UserName string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Logs     []Log  `bson:"logs"`
	Id       int    `bson:"_id, omitempty"`
}
