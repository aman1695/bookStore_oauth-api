package access_token

import (
	"github.com/aman1695/bookStore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const(
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId 		int64  `json:"user_id"`
	ClientId 	int64  `json:"client_id"`
	Expired 	int64  `json:"expired"`
}
func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid Access Token!!!!")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("Invalid User ID!!!")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid Client ID!!!")
	}
	if at.Expired <= 0 {
		return errors.NewBadRequestError("Invalid Expire Time!!!")
	}
	return nil
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expired:     time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func (at AccessToken) isExpired() bool {
	return time.Unix(at.Expired,0).Before(time.Now().UTC())
}
