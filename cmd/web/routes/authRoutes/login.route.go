package authRoutes

import (
	authDomain "GoBaby/cmd/web/domain/auth"
	"GoBaby/cmd/web/routes"
	"GoBaby/internal/models"
)

func LoginRender() {
	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesInstance.LOGIN, authDomain.LoginView)
}
