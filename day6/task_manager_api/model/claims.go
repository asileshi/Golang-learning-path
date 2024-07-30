package model

import "github.com/golang-jwt/jwt"

type Claim struct{
	Email      string `json:"email"`
	Role          string `json:"role"`
	jwt.StandardClaims
}
