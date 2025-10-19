package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Email string
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data JWTData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": data.Email,
	})

	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	return s, nil
}

func (j *JWT) Parse(tokenString string) (bool, *JWTData) {
	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})

	if err != nil {
		return false, nil
	}
	email, _ := t.Claims.(jwt.MapClaims)["email"].(string)

	return t.Valid, &JWTData{Email: email}
}
