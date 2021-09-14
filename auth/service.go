package auth

import "github.com/dgrijalva/jwt-go"

//JWT -> gnerate token -> validasi token

type Service interface {
	GenerateToken(userID int) (string, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("secret")

func NewService() *jwtService{
	return &jwtService{}
}
func (s *jwtService) GenerateToken(userID int) (string, error) {
	//payload == claims
	payload := jwt.MapClaims{}
	payload["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
