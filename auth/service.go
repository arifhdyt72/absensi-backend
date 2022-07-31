package auth

import (
	"absensi-backend/user"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(data user.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRETKEY = []byte("ABSENSI_s3cr3t_k3Y")

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(data user.User) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = data.ID
	claim["role"] = data.Role

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRETKEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRETKEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
