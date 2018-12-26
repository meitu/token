package token

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

const (
	//CurrentVesion is the default version
	CurrentVesion = 1
	tokenSignLen  = 11
)

// Token is generated based on version and key
type Token struct {
	key []byte
}

// New return a token object
func New(key []byte) *Token {
	return &Token{key: key}
}

// Sign used to generate signatures
func (t *Token) Sign(payload []byte) ([]byte, error) {
	m := &message{version: CurrentVesion, createAt: int64(time.Now().Unix()), payload: payload}
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

// Verify a token
func (t *Token) Verify(sign []byte) error {
	encodedSignLen := hex.EncodedLen(tokenSignLen)
	if len(sign) < encodedSignLen {
		return errors.New("invalid size")
	}

	s := make([]byte, tokenSignLen)
	hex.Decode(s, sign[len(sign)-encodedSignLen:])

	meta := sign[:len(sign)-encodedSignLen-1] //counting in the "-"
	mac := hmac.New(sha256.New, t.key)
	mac.Write(meta)

	if !hmac.Equal(mac.Sum(nil)[:tokenSignLen], s) {
		return errors.New("token mismatch")
	}

	return nil
}

// Auth a token and renturn payload
func (t *Token) Auth(sign []byte) ([]byte, error) {
	encodedSignLen := hex.EncodedLen(tokenSignLen)
	if len(sign) < encodedSignLen {
		return nil, errors.New("auth invalid size")
	}

	s := make([]byte, tokenSignLen)
	hex.Decode(s, sign[len(sign)-encodedSignLen:])

	meta := sign[:len(sign)-encodedSignLen-1] //counting in the ":"
	mac := hmac.New(sha256.New, t.key)
	mac.Write(meta)

	if !hmac.Equal(mac.Sum(nil)[:tokenSignLen], s) {
		return nil, errors.New("token mismatch")
	}

	var m message

	if err := m.UnmarshalBinary(meta); err != nil {
		return nil, err
	}
	return m.payload, nil
}

// message contains the necessary constituent fields for a signature
type message struct {
	version  int64
	createAt int64
	payload  []byte
}

// MarshalBinary is used to binary code data
func (m *message) MarshalBinary() (data []byte, err error) {
	data = append(data, m.payload...)
	data = append(data, '-')
	data = append(data, []byte(strconv.FormatInt(m.createAt, 10))...)
	data = append(data, '-')
	data = append(data, []byte(strconv.FormatInt(m.version, 10))...)
	return data, nil
}

func (m *message) UnmarshalBinary(data []byte) error {
	fields := bytes.Split(data, []byte{'-'})
	l := len(fields)
	if l < 3 {
		return errors.New("invalid token")
	}

	version, err := strconv.ParseInt(string(fields[l-1]), 10, 64)
	if err != nil {
		return err
	}
	m.version = version

	createAt, err := strconv.ParseInt(string(fields[l-2]), 10, 64)
	if err != nil {
		return err
	}
	m.createAt = createAt

	m.payload = bytes.Join(fields[:l-2], []byte(""))

	return nil
}
