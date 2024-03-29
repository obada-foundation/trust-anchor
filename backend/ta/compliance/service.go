package compliance

import (
	"github.com/obada-foundation/trust-anchor/ta"
)

type Service struct {
	taTokenSvc *ta.TokenService
}

// NewService initialize compliance service
func New(tokenSvc *ta.TokenService) *Service {
	return &Service{
		taTokenSvc: tokenSvc,
	}
}

// CheckByToken checks if owner of token is compliant
func (s Service) CheckByToken(token string) (bool, error) {
	_, err := s.taTokenSvc.VerifyToken(token)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s Service) CheckByUser() (bool, error) {
	return false, nil
}
