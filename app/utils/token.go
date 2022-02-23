package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JwtToken struct {
	UserId []byte `json:"userId"`
	StandardClaims jwt.StandardClaims `json:"claims"`
}

//TODO - HERE - I need to verify the token signature in here, or
// at worst (and only for the moment) I have to do a DB look up on the decrypted user id.
func (t JwtToken) Valid() error {
	now := time.Now().Unix()

	// If we have passed the expiration time then it's expired, so we don't need to decrypt.
	if now > t.StandardClaims.ExpiresAt {
		return errors.New("token expired")
	}

	// If we can't decrypt the token then we know it's invalid, so stop.
	_, err := Decrypt(t.UserId)
	if err != nil {
		return errors.New("token invalid")
	}


	return nil
}

