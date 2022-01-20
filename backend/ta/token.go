package ta

import (
	"errors"
	"fmt"
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
)

type TokenService struct {
	issuer    string
	verifyUrl string
	alg       jwa.SignatureAlgorithm
	verifier  jwt.ParseOption
	signKey   interface{}
}

func New(issuer, taUrl string, signKey, verifyKey interface{}) *TokenService {
	ja := jwa.SignatureAlgorithm("EdDSA")

	return &TokenService{
		alg:       ja,
		verifier:  jwt.WithVerify(ja, verifyKey),
		signKey:   signKey,
		issuer:    issuer,
		verifyUrl: fmt.Sprintf("%s/api/v1/verify", taUrl),
	}
}

func (ts TokenService) CreateToken(DID, userID string) (t jwt.Token, tokenStr string, err error) {
	t = jwt.New()

	t.Set("nft", DID)
	t.Set(jwt.IssuerKey, ts.issuer)
	t.Set(jwt.IssuedAtKey, time.Now())
	t.Set(jwt.SubjectKey, userID)
	t.Set("url", ts.verifyUrl)

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
