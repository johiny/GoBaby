package authDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/store"
	authUtils "GoBaby/internal/utils/auth"
	"net/http"
)

func SetSession(user models.User, w http.ResponseWriter) error {
	token, err := authUtils.GenerateToken(user.Email)
	if err != nil {
		return err
	}
	store.AddUser(user, token)
	authUtils.SetTokenOnCookie(w, user.Email)
	return nil
}
