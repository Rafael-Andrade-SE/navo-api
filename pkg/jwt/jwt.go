package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GerarToken(usuarioID uint, email string, role string) (string, error) {
	segredo := os.Getenv("JWT_SECRET")
	expMin := os.Getenv("JWT_EXPIRATION_MINUTES")

	duracao, err := time.ParseDuration(expMin + "m")
	if err != nil {
		duracao = time.Hour // fallback padrão
	}

	claims := jwt.MapClaims{
		"id":    usuarioID,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(duracao).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(segredo))
}
