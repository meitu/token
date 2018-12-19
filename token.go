package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

const (
	CurrentVesion = 1 //CurrentVesion is the default initial version
	tokenSignLen  = 11
)

//Token is generated based on version and key
type Token struct {
	key []byte
}

// New return a token object
func New(key []byte) *Token {
	return &Token{key: key}
}

// Sign used to generate signatures
func (t *Token) Sign(data []byte) ([]byte, error) {
	m := &message{version: CurrentVesion, createAt: int64(time.Now().Unix()), payload: data}
	data, err := m.MarshalBinary()
	if err != nil {
		return nil, err
	}
	mac := hmac.New(sha256.New, t.key)
	mac.Write(data)
	sign := mac.Sum(nil)

	//truncate to 32 byte: https://tools.ietf.org/html/rfc2104#section-5
	// we have 11 byte rigth of hmac,so the rest of data is token message
	sign = sign[:tokenSignLen]

	encodedSign := make([]byte, hex.EncodedLen(len(sign)))
	hex.Encode(encodedSign, sign)
	var token []byte
	token = append(token, data...)
	token = append(token, '-')
	token = append(token, encodedSign...)
	return token, nil
}

// Verify used to token auth
func (t *Token) Verify(sign []byte) error {
	encodedSignLen := hex.EncodedLen(tokenSignLen)
	if len(sign) < encodedSignLen || len(t.key) == 0 {
		return errors.New("token or key is parameter illegal")
	}

	s := make([]byte, tokenSignLen)
	hex.Decode(s, sign[len(sign)-encodedSignLen:])

	meta := sign[:len(sign)-encodedSignLen-1] //counting in the ":"
	mac := hmac.New(sha256.New, t.key)
	mac.Write(meta)

	if !hmac.Equal(mac.Sum(nil)[:tokenSignLen], s) {
		return errors.New("token mismatch")
	}

	return nil
}

// message contains the necessary constituent fields for a signature
type message struct {
	version  int64
	createAt int64
	payload  []byte
}

func (m *message) MarshalBinary() (data []byte, err error) {
	data = append(data, m.payload...)
	data = append(data, '-')
	data = append(data, []byte(strconv.FormatInt(m.createAt, 10))...)
	data = append(data, '-')
	data = append(data, []byte(strconv.FormatInt(m.version, 10))...)
	return data, nil
}
