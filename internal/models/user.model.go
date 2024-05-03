package models

type User struct {
	UserName string `bson:"username"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
	Logs     []Log  `bson:"logs"`
	Id       string `bson:"_id"`
}
