package utils

import (
    "github.com/golang-jwt/jwt"
    "github.com/stretchr/testify/assert"
    "testing"
    "time"
)

func Test_Valid_success(t *testing.T) {
    now := time.Now()
    token := JwtToken{
        UserId: []byte{1},
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: now.Add(100 * time.Second).Unix(),
            NotBefore: now.Unix(),
            IssuedAt:  now.Unix(),
        },
    }
    err := token.Valid()

    assert.Nil(t, err)
}

func Test_Valid_expiresAt_is_0(t *testing.T) {
    now := time.Now()
    token := JwtToken{
        UserId: []byte{1},
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: 0,
            NotBefore: now.Unix(),
            IssuedAt:  now.Unix(),
        },
    }
    err := token.Valid()

    assert.Error(t, err)
}

func Test_Valid_NotBefore_is_0(t *testing.T) {
    now := time.Now()
    token := JwtToken{
        UserId: []byte{1},
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: now.Add(100 * time.Second).Unix(),
            NotBefore: 0,
            IssuedAt:  now.Unix(),
        },
    }
    err := token.Valid()

    assert.Error(t, err)
}

func Test_Valid_IssuedAt_is_0(t *testing.T) {
    now := time.Now()
    token := JwtToken{
        UserId: []byte{1},
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: now.Add(100 * time.Second).Unix(),
            NotBefore: now.Unix(),
            IssuedAt:  0,
        },
    }
    err := token.Valid()

    assert.Error(t, err)
}
