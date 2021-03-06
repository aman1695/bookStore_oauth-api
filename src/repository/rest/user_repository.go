package rest

import (
	"encoding/json"
	"github.com/aman1695/bookStore_oauth-api/src/domain/users"
	"github.com/aman1695/bookStore_oauth-api/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)
var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Microsecond,
	}
)
type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}
type usersRepository struct {}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}
func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", request)
	if response == nil || response.Response == nil {
		return nil,errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestErr
		err := json.Unmarshal(response.Bytes(),restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface ehrn trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(),&user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}
