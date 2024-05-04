package store

import (
	"GoBaby/internal/models"
	"fmt"
)

var MaxSize = 10

var UserStore = &models.UserStore{
	MaxSize: MaxSize,
	Users:   make([]map[string]string, 0, MaxSize),
}

// AddUser agrega un nuevo usuario a la cola de usuarios activamente logueados
func AddUser(user models.User, jwt_token string) {
	if len(UserStore.Users) >= UserStore.MaxSize {
		// Si la cola está llena, eliminar el elemento más antiguo
		UserStore.Users = UserStore.Users[1:]
	}
	newUser := map[string]string{
		"email":     user.Email,
		"jwt_token": jwt_token,
		"user":      user.UserName,
	}
	// Agregar el nuevo elemento al final de la cola
	UserStore.Users = append(UserStore.Users, newUser)
	fmt.Println(UserStore.Users)
}

func GetUser(email string) map[string]string {
	for _, user := range UserStore.Users {
		if user["email"] == email {
			return user
		}
	}
	return nil
}
