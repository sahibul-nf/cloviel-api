package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	// ValidateToken(token string) (*jwt.Token, error)
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

func GenerateToken(userID int) (string, error) {
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

// func ValidateToken(token string)  {

// }
