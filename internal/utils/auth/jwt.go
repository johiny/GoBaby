package authUtils

import (
	"GoBaby/internal/store"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// esto deberia ir en un archivo .env - temporal
var secretKey = []byte("GranBigote")

func GenerateToken(email string) (string, error) {
	// Crear un nuevo token JWT
	token := jwt.New(jwt.SigningMethodHS256)

	// Configurar claims (datos) del token
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Expira en 24 horas

	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (string, error) {
	// Parsear el token con la clave secreta
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	// Validar si el token es válido
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["email"].(string)
		return user, nil
	}

	return "", fmt.Errorf("token inválido")
}

func SetTokenOnCookie(w http.ResponseWriter, r *http.Request) {
	// obtener email del usuario de la request
	userEmail := r.FormValue("user_email")
	user := store.GetUser(userEmail)
	if user == nil {
		fmt.Println("usuario no encontrado")
		return
	}
	// generar token
	cookie := &http.Cookie{
		Name:     "jwt_token",
		Value:    user["jwt_token"],
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
}
