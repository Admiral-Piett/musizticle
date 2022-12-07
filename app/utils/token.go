package utils

import (
    "errors"
    "github.com/golang-jwt/jwt"
)

type JwtToken struct {
    UserId         []byte             `json:"userId"`
    StandardClaims jwt.StandardClaims `json:"claims"`
}

func (t JwtToken) Valid() error {
    if t.StandardClaims.ExpiresAt == 0 || t.StandardClaims.NotBefore == 0 || t.StandardClaims.IssuedAt == 0 {
        return errors.New("TokenInvalid: Missing field")
    }
    // This validates the above 3 fields, but it will pass them if they're the default, so we'll check them above.
    err := t.StandardClaims.Valid()
    if err != nil {
        return err
    }

    return nil
}
