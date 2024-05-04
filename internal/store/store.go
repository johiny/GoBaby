package store

import (
	"GoBaby/internal/models"
)

var MaxSize = 10

var UserStore = &models.UserStore{
	MaxSize: MaxSize,
	Users:   make([]map[string]string, 0, MaxSize),
}

// AddUser agrega un nuevo usuario a la cola de usuarios activamente logueados
func AddUser(email string, jwt_token string, user string) {
	if len(UserStore.Users) >= UserStore.MaxSize {
		// Si la cola está llena, eliminar el elemento más antiguo
		UserStore.Users = UserStore.Users[1:]
	}
	newUser := map[string]string{
		"email":     email,
		"jwt_token": jwt_token,
		"user":      user,
	}
	// Agregar el nuevo elemento al final de la cola
	UserStore.Users = append(UserStore.Users, newUser)
}

func GetUser(email string) map[string]string {
	for _, user := range UserStore.Users {
		if user["email"] == email {
			return user
		}
	}
	return nil
}
