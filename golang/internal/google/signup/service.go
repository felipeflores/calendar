package signup

import (
	"context"
	"iot/pkg/google"
	googleModel "iot/pkg/google/model"
)

type SignupService struct {
	clientGoogle *google.ClientGoogle
}

func NewSignupService(clientGoogle *google.ClientGoogle) *SignupService {
	return &SignupService{
		clientGoogle: clientGoogle,
	}
}

func (s *SignupService) SetCredentials(ctx context.Context, credentials googleModel.Credentials) (string, error) {
	url, err := s.clientGoogle.Setup(ctx, credentials)
	if err != nil {
		return "", err
	}
	return url, nil
}
func (s *SignupService) SetCode(ctx context.Context, code Code) error {
	err := s.clientGoogle.GenerateToken(ctx, code.Code)
	if err != nil {
		return err
	}
	return nil
}
