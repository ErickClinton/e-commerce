package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenManager struct {
	secretKey   []byte
	tokenExpiry time.Duration
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func NewTokenManager(secretKey string, tokenExpiry time.Duration) *TokenManager {
	return &TokenManager{
		secretKey:   []byte(secretKey),
		tokenExpiry: tokenExpiry,
	}
}

func (tm *TokenManager) GenerateToken(userID string, role string) (string, error) {
	expirationTime := time.Now().Add(tm.tokenExpiry)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tm.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %v", err)
	}

	return tokenString, nil
}

func (tm *TokenManager) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tm.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (tm *TokenManager) RefreshToken(tokenString string) (string, error) {
	claims, err := tm.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Minute {
		return tokenString, nil
	}

	return tm.GenerateToken(claims.UserID, claims.Role)
}
