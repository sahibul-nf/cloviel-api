package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	ValidateToken(encodedToken string) (*jwt.Token, error)
	GenerateToken(userID int) (string, error)
}

type authService struct {
}

func NewService() *authService {
	return &authService{}
}

var (
	secretKey = "JWT_SECRET_KEY"
)

func (s *authService) GenerateToken(userID int) (string, error) {
	// create payload data
	claim := jwt.MapClaims{
		"user_id": userID,
	}

	// hashing payload data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// make signature
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *authService) ValidateToken(encodedToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte([]byte(secretKey)), nil
	})

	if err != nil {
		return token, err
	}

	return token, err
}
