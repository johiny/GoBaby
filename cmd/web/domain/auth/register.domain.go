package authDomain

import (
	errorDomain "GoBaby/cmd/web/domain/error"
	repository_adapters "GoBaby/cmd/web/domain/repository/adapters"
	"GoBaby/internal/models"
	"GoBaby/internal/utils"
	AuthUtils "GoBaby/internal/utils/auth"
	formUtils "GoBaby/internal/utils/form"
	"GoBaby/ui"
	"fmt"
	"net/http"
)

func RegisterView(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"html/base.html",
		"html/pages/auth/register.tmpl.html",
	}
	utils.CheckIfPath(w, r, models.RoutesInstance.REGISTER)
	utils.ParseTemplateFiles(w, "base", utils.EmptyStruct, utils.EmptyFuncMap, ui.Content, files...)
}

func RegisterAction(w http.ResponseWriter, r *http.Request) {
	formValues, err := formUtils.GetFormValues(r, []string{"Email", "Password", "Username"})
	if err != nil {
		errorDomain.ErrorTemplate(w, r, err)
		return
	}

	encryptedPass, err := AuthUtils.EncryptPassword(formValues["Password"])
	if err != nil {
		errorDomain.ErrorTemplate(w, r, err)
		return
	}
	user := models.User{
		Email:    formValues["Email"],
		Password: encryptedPass,
		UserName: formValues["Username"],
	}
	err = repository_adapters.SetUser(&user)
	if err != nil {
		fmt.Println("Error:", err.Err.Error())
		errorDomain.ErrorTemplate(w, r, err)
		return
	}
	err1 := SetSession(user, w)
	if err1 != nil {
		fmt.Println("Error:", err1.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
