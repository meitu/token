package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strconv"
)

const (
	//tokenSignLen token default len
	tokenSignLen = 11
)

//Token generation based on TokenInfo
type TokenInfo struct {
	Version   int8   `json:"version"`
	CreateAt  int64  `json:"create_at"`
	Namespace []byte `json:"namespace"`
}

//MarshalBinary Namespace SHOULD NOT contains a colon
func (ti *TokenInfo) MarshalBinary() (data []byte, err error) {
	data = append(data, ti.Namespace...)
	data = append(data, '-')
	data = append(data, []byte(strconv.FormatInt(ti.CreateAt, 10))...)
	data = append(data, '-')
	data = append(data, []byte(strconv.FormatInt(int64(ti.Version), 10))...)
	return data, nil
}

//UnmarshalBinary token base unmarshl
func (ti *TokenInfo) UnmarshalBinary(data []byte) error {
	fields := bytes.Split(data, []byte{'-'})
	l := len(fields)
	if l != 3 {
		return errors.New("invalid token")
	}

	version, err := strconv.ParseInt(string(fields[l-1]), 10, 64)
	if err != nil {
		return err
	}
	ti.Version = int8(version)

	createAt, err := strconv.ParseInt(string(fields[l-2]), 10, 64)
	if err != nil {
		return err
	}
	ti.CreateAt = createAt

	ti.Namespace = bytes.Join(fields[:l-2], []byte(""))

	return nil
}

//Verify token auth
func Verify(token, key []byte) ([]byte, error) {
	encodedSignLen := hex.EncodedLen(tokenSignLen)
	if len(token) < encodedSignLen || len(key) == 0 {
		return nil, errors.New("token or key is parameter illegal")
	}

	sign := make([]byte, tokenSignLen)
	hex.Decode(sign, token[len(token)-encodedSignLen:])

	meta := token[:len(token)-encodedSignLen-1] //counting in the ":"
	mac := hmac.New(sha256.New, key)
	mac.Write(meta)

	if !hmac.Equal(mac.Sum(nil)[:tokenSignLen], sign) {
		return nil, errors.New("token mismatch")
	}

	var t TokenInfo
	if err := t.UnmarshalBinary(meta); err != nil {
		return nil, err
	}
	return t.Namespace, nil
}

//Token token create through key server namespace create time
func Token(key, namespace []byte, createAt int64) ([]byte, error) {
	t := &TokenInfo{Namespace: namespace, CreateAt: createAt, Version: 1}
	data, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}

	mac := hmac.New(sha256.New, key)
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
