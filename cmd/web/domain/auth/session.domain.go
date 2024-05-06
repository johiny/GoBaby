package authDomain

import (
	"GoBaby/internal/models"
	"GoBaby/internal/store"
	"GoBaby/internal/utils"
	authUtils "GoBaby/internal/utils/auth"
	"GoBaby/ui"
	"errors"
	"fmt"
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

func CheckSession(r *http.Request) error {
	cookie, err := r.Cookie("jwt_token")
	fmt.Println("jwt token", cookie)
	if err != nil {
		fmt.Println("error obteniendo la cookie")
		return err
	}
	email, err := authUtils.ValidateToken(cookie.Value)
	fmt.Println("email", email)
	if err != nil {
		return err
	}
	if store.UserStore.CurrentUser["email"] != email {
		return errors.New("no es el mismo usuario")
	}
	return nil
}

func SetCurrentUser(w http.ResponseWriter, email string) error {
	user := store.GetUser(email)
	if user == nil {
		return errors.New("usuario no encontrado")
	}
	store.UserStore.CurrentUser = user
	authUtils.SetTokenOnCookie(w, user["email"])
	return nil
}

func ChangeCurrentUser(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	fmt.Println("correo que viene del query params", email)
	err := SetCurrentUser(w, email)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func UserMenuView(w http.ResponseWriter, r *http.Request) {
	utils.CheckIfPath(w, r, models.RoutesInstance.USERMENU)
	utils.ParseTemplateFiles(w, "usersMenu", store.UserStore.Users, utils.EmptyFuncMap, ui.Content, "html/pages/auth/usersMenu.tmpl.html")
}
