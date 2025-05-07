package utils

import (
    "errors"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "webtechproject/config"
)

func GenerateToken(userID uint, role string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(config.JWT_SECRET)
}

func ValidateToken(tokenString string) (uint, string, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return config.JWT_SECRET, nil
    })
    if err != nil || !token.Valid {
        return 0, "", errors.New("invalid token")
    }

    claims := token.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64))
    role := claims["role"].(string)
    return userID, role, nil
}
