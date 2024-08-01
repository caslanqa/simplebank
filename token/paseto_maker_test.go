package token

import (
	"github.com/CharlieAlphaQA/simplebank/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err1 := maker.CreateToken(username, duration)
	require.NoError(t, err1)
	require.NotEmpty(t, token)

	payload, err2 := maker.VerifyToken(token)
	require.NoError(t, err2)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, duration)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, duration)

}

func TestExpiredPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err1 := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err1)
	require.NotEmpty(t, token)

	payload, err2 := maker.VerifyToken(token)
	require.Error(t, err2)
	require.EqualError(t, err2, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidPasetoMaker(t *testing.T) {
	payload, err1 := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err1)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err2 := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err2)

	maker, err3 := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err3)

	payload, err1 = maker.VerifyToken(token)
	require.Error(t, err1)
	require.EqualError(t, err1, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
