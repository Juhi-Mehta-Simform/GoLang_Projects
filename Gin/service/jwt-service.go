package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTservice interface {
	GenerateToken(name string, admin bool) string
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTservice {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "simform.com",
	}
}

func getSecretKey() string {
	secret := "secret"
	return secret
}

func (jwtSrv *jwtService) GenerateToken(username string, admin bool) string {
	claims := &jwtCustomClaims{
		username,
		admin,
		jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				time.Now().Add(time.Hour * 72),
			},
			IssuedAt: &jwt.NumericDate{
				time.Now(),
			},
			Issuer: jwtSrv.issuer,
		},
	}
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := tokens.SignedString([]byte(jwtSrv.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}
func (jwtSrv *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected Signing Method: %v", token.Header["alg"])
		}
		return []byte(jwtSrv.secretKey), nil
	})
}
