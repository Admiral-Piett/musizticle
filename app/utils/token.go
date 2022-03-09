package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JwtToken struct {
	UserId []byte `json:"userId"`
	StandardClaims jwt.StandardClaims `json:"claims"`
}

//TODO - HERE - I need to verify the token signature in here, or
// at worst (and only for the moment) I have to do a DB look up on the decrypted user id.
func (t JwtToken) Valid() error {
	if t.StandardClaims.ExpiresAt == 0 || t.StandardClaims.NotBefore == 0 || t.StandardClaims.IssuedAt == 0 {
		return errors.New("TokenInvalid: Missing field")
	}
	// This validates the above 3 fields, but it will pass them if they're falsey, so we'll check them now.
	err := t.StandardClaims.Valid()
	if err != nil {
		return err
	}

	// If we can't decrypt the token then we know it's invalid, so stop.
	_, err = Decrypt(t.UserId)
	if err != nil {
		return errors.New(fmt.Sprintf("TokenInvalid: %s", err))
	}
	return nil
}

