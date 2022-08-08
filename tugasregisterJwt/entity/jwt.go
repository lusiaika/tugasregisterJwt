package entity

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	jwt.StandardClaims
	Uid string `json:"uid"`
	//Pwd string `json:"pwd"`
}
