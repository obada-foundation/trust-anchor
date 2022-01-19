package ta

import (
	"errors"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type TokenService struct {
	alg      jwa.SignatureAlgorithm
	verifier jwt.ParseOption
	signKey  interface{}
}

func New(signKey, verifyKey interface{}) *TokenService {
	ja := jwa.SignatureAlgorithm("EdDSA")

	return &TokenService{
		alg:      ja,
		verifier: jwt.WithVerify(ja, verifyKey),
		signKey:  signKey,
	}
}

func (ts TokenService) CreateToken(DID string) (t jwt.Token, tokenStr string, err error) {
	t = jwt.New()

	t.Set("did", DID)

	payload, err := jwt.Sign(t, ts.alg, ts.signKey)
	if err != nil {
		return nil, "", err
	}

	tokenStr = string(payload)

	return
}

// VerifyToken returns trust anchor token if token string is valid
func (ts TokenService) VerifyToken(token string) (*jwt.Token, error) {
	if len(token) == 0 {
		return nil, errors.New("missing token")
	}

	t, err := jwt.Parse([]byte(token), ts.verifier)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	return &t, nil
}
