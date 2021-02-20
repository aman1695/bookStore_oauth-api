package access_token

import "github.com/aman1695/bookStore_oauth-api/src/utils/errors"

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
	Create(AccessToken) *errors.RestErr
	UpdateExpirationTime(AccessToken) *errors.RestErr
}
type service struct {
	repository Repository
}
func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
func (s* service) GetById(accessTokenId string) (*AccessToken, *errors.RestErr) {
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken,nil
}

func (s* service) Create(accessToken AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}
	return s.repository.Create(accessToken)
}

func (s* service) UpdateExpirationTime(accessToken AccessToken) *errors.RestErr {
	if err := accessToken.Validate(); err != nil {
		return err
	}
	return s.repository.UpdateExpirationTime(accessToken)
}