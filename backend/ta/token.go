package ta

import (
	"errors"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type TokenService struct {
	alg      jwa.SignatureAlgorithm
	verifier jwt.ParseOption
}

func New(signKey, verifyKey interface{}) *TokenService {
	ja := jwa.SignatureAlgorithm("EdDSA")

	return &TokenService{
		alg:      ja,
		verifier: jwt.WithVerify(ja, verifyKey),
	}
}

func (s TokenService) CreateToken() *jwt.Token {
	return nil
}

// VerifyToken returns trust anchor token if token string is valid
func (s TokenService) VerifyToken(token string) (*jwt.Token, error) {
	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	t, err := jwt.Parse([]byte(token), s.verifier)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return &t, nil
}
