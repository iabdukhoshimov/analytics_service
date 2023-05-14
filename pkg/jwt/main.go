package jwt

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
	PassCode     TokenType = "passcode"
)

type Cred struct {
	Secret string
	Expire time.Duration
}

type Jwt struct {
	Creds map[TokenType]Cred
}

type Authenticator interface {
	GenerateJWT(mapToken map[string]interface{}, tokenType TokenType) (string, error)
	CheckClaims(params CheckClaimsParams) error
	GetAuthInfo(c *gin.Context) (*AuthResp, error)
}

type CheckClaimsParams struct {
	Token     string
	Code      string
	Phone     string
	TokenType TokenType
	Dev       bool
}

type AuthResp struct {
	ID   string
	Role int32
}

func New(params map[TokenType]Cred) Authenticator {
	return &Jwt{
		Creds: params,
	}
}

// GenerateToken generates a new JWT token string - tokenType - [access, passcode, refresh]
func (j *Jwt) GenerateJWT(mapToken map[string]interface{}, tokenType TokenType) (string, error) {
	var (
		token       = jwt.New(jwt.SigningMethodHS256)
		tokenString string
		err         error
	)

	creds, ok := j.Creds[tokenType]
	if !ok {
		return "", errors.New("invalid token type")
	}

	claims := token.Claims.(jwt.MapClaims)

	for key, value := range mapToken {
		claims[key] = value
	}

	claims["iss"] = creds.Secret
	claims["aud"] = creds.Secret
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(creds.Expire).Unix()

	tokenString, err = token.SignedString([]byte(creds.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *Jwt) extractClaims(tokenString string, tokenType TokenType) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	creds, ok := j.Creds[tokenType]
	if !ok {
		return nil, errors.New("invalid token type")
	}

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(creds.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (j *Jwt) CheckClaims(params CheckClaimsParams) error {
	claims, err := j.extractClaims(params.Token, params.TokenType)
	if err != nil {
		return err
	}

	passCodeToken := claims["hashed_code"].(string)
	err = bcrypt.CompareHashAndPassword(
		[]byte(passCodeToken),
		[]byte(fmt.Sprintf("%s-%s", params.Code, params.Phone)),
	)

	return err
}

func (j *Jwt) GetAuthInfo(ctx *gin.Context) (*AuthResp, error) {
	var (
		auth AuthResp
	)

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		return &auth, nil
	}

	claims, err := j.extractClaims(tokenString, AccessToken)
	if err != nil {
		return nil, err
	}

	id, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("invalid id")
	}

	role, ok := claims["role"].(float64)
	if !ok {
		return nil, errors.New("invalid role")
	}

	auth.ID = id
	auth.Role = int32(role)

	return &auth, nil
}
