package access_token

import (
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.isExpired(),"brand new access token should not be expired")
	assert.True(t, at.UserId == 0,"new access token should not have an associated user id")
	assert.True(t, at.AccessToken == "","new access token should not have defined access token id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.False(t, !at.isExpired(),"empty access token should be expired by default")
	at.Expired = time.Now().UTC().Add(3*time.Hour).Unix()
	assert.False(t, at.isExpired(),"access token expiring 3 hours from now should not be expired")
}
