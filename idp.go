package idp

import (
	"encoding/base64"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// IDP struct.
type IDP struct {
	ID            string `json:"id"`
	Sub           string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Username      string `json:"preferred_username"`
	Email         string `json:"email"`
	Active        bool   `json:"active"`
}

// Extract json string to IDP.
func Extract(jsonString []byte) (IDP, error) {
	var idp IDP
	err := json.Unmarshal(jsonString, &idp)
	if err != nil {
		return idp, err
	}
	return idp, nil
}

// SetIDP method.
func SetIDP(c *gin.Context, headerString string) (IDP, error) {
	var (
		encodedUserInfo = c.Request.Header.Get(headerString)
	)

	decodedUserInfo, err := base64.StdEncoding.DecodeString(encodedUserInfo)
	if err != nil {
		return IDP{}, err
	}

	return Extract(decodedUserInfo)
}
