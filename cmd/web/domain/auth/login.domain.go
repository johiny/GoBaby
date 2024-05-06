package authDomain

import (
	repository_adapters "GoBaby/cmd/web/domain/repository/adapters"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	authUtils "GoBaby/internal/utils/auth"
	"GoBaby/ui"
	"net/http"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"html/base.html",
		"html/pages/auth/login.tmpl.html",
	}
	utils.CheckIfPath(w, r, models.RoutesInstance.LOGIN)
	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
}

func LoginAction(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear el formulario", http.StatusInternalServerError)
		return
	}
	Email := r.Form.Get("Email")
	Password := r.Form.Get("Password")
	user, dbErr := repository_adapters.GetUserByEmail(Email)
	if dbErr != nil {
		http.Error(w, "Error al obtener el usuario", http.StatusInternalServerError)
		return
	}
	passErr := authUtils.CheckPassword(Password, user.Password)
	if passErr != nil {
		http.Error(w, "Error al verificar la contraseña", http.StatusInternalServerError)
		return
	}
	err = SetSession(user, w)
	if err != nil {
		http.Error(w, "Error al crear la sesión", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
