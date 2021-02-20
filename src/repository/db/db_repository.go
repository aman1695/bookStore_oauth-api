package db

import (
	"github.com/aman1695/bookStore_oauth-api/src/client/cassandra"
	"github.com/aman1695/bookStore_oauth-api/src/domain/access_token"
	"github.com/aman1695/bookStore_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?"
	queryCreateToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
	queryUpdateExpires = "UPDATE access_tokens SET expires=? WHERE access_Token=?"
	)
func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {

}

func (r *dbRepository)GetById(id string) (*access_token.AccessToken,*errors.RestErr) {
	var accessToken access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken,id).Scan(&accessToken.AccessToken,&accessToken.UserId,
		&accessToken.ClientId,&accessToken.Expired); err != nil {
		if err == gocql.ErrNotFound{
			return nil, errors.NewNotFoundError("no access token exists with the id: "+id)
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &accessToken,nil
}

func (r *dbRepository) Create(accessToken access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateToken,accessToken.AccessToken,accessToken.UserId,
		accessToken.ClientId,accessToken.Expired).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(accessToken access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,accessToken.Expired,
		accessToken.AccessToken).Exec(); err != nil {
		if err == gocql.ErrNotFound{
			return errors.NewNotFoundError("no access token exists with the id: "+accessToken.AccessToken)
		}
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}