package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)


func GetPublic() (*rsa.PublicKey, error) {
	keyData, err := ioutil.ReadFile("jwt/private.pem")

	if err != nil {
		return nil, err
	}

	private, err := jwt.ParseRSAPrivateKeyFromPEMWithPassword(keyData, "5642254")
	if err != nil {
		return nil, err
	}
	return &private.PublicKey, nil
}