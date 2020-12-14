package idp

import "encoding/json"

// IDP struct.
type IDP struct {
	Sub           string `json:"sub"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Username      string `json:"preferred_username"`
	Email         string `json:"email"`
}

// Extract json string to IDP.
func Extract(jsonString string) (IDP, error) {
	var idp IDP
	err := json.Unmarshal([]byte(jsonString), &idp)
	if err != nil {
		return idp, err
	}
	return idp, nil
}
