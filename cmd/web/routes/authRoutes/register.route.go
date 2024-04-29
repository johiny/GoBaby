package authRoutes

import (
	authDomain "GoBaby/cmd/web/domain/auth"
	"GoBaby/cmd/web/routes"
	"GoBaby/internal/models"
)

func RegisterRender() {
	routes.GetMuxInstance().HandleFunc("GET "+models.RoutesInstance.REGISTER, authDomain.RegisterView)
}
