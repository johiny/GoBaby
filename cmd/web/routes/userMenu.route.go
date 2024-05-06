package routes

import (
	authDomain "GoBaby/cmd/web/domain/auth"
	"GoBaby/internal/models"
)

func UserMenuRender() {
	mux.HandleFunc("GET "+models.RoutesInstance.USERMENU, authDomain.UserMenuView)
	mux.HandleFunc("GET "+models.RoutesInstance.CHANGEUSER, authDomain.ChangeCurrentUser)
}
