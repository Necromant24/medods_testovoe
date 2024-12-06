package services

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"medods/auth-service/interfaces"
	"medods/auth-service/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type TokensService struct {
	interfaces.ITokensRepository
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
	jwt.RegisteredClaims
}

func (service *TokensService) RefreshTokens(accessToken string, refreshToken string) (string, string, error) {
	_, err := Hash(refreshToken, cost)

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return SecretKey, nil
	})

	if err != nil {
		fmt.Println("Ошибка разбора токена: %v", err)
	}

	getIp := ""
	getUserId := ""

	// Проверяем валидность токена и извлекаем данные
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token valid!")
		for key, value := range claims {
			if key == "ip" {
				getIp = value.(string)
			}
			if key == "user_id" {
				getUserId = value.(string)
			}
		}
	} else {
		fmt.Println("Invalid token")
	}

	baseToken, err := service.ITokensRepository.GetRefreshToken(getUserId)

	if getIp != baseToken.UserIp {
		// mock send email warning

		err = fmt.Errorf("IP changed")

		return "", "", err
	}

	var tokenCorrect = false

	hash := sha256.Sum256([]byte(refreshToken))
	preHashedPassword := hex.EncodeToString(hash[:])
	err = bcrypt.CompareHashAndPassword([]byte(baseToken.Token), []byte(preHashedPassword))
	if err != nil {
		fmt.Println(err)
	} else {
		tokenCorrect = true
	}

	if tokenCorrect {
		// do refresh

		access, refresh, err := GenerateTokenPair(getUserId, getIp)

		bhash, err := Hash(refresh, cost)

		err = service.ITokensRepository.SaveToken(models.RefreshToken{Token: bhash, UserId: getUserId, UserIp: getIp})

		return access, refresh, err
	} else {
		err = fmt.Errorf("smthng wrong with tokens")
		return "", "", err
	}

}

func (service *TokensService) GetTokensPair(userId string, userIp string) (string, string, error) {
	access, refresh, err := GenerateTokenPair(userId, userIp)

	bhash, err := Hash(refresh, cost)

	err = service.ITokensRepository.SaveToken(models.RefreshToken{Token: bhash, UserId: userId, UserIp: userIp})

	return access, refresh, err
}

// TODO: init from config
var SecretKey = []byte("your-secret-key")

// TODO: init from config
// from 4 to 31 range
var cost int = 5

func GenerateTokenPair(userID, clientIP string) (string, string, error) {
	// Генерация Access токена
	accessClaims := CustomClaims{
		UserID: userID,
		IP:     clientIP,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), // Access токен действует 30 минут
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims)
	accessTokenString, err := accessToken.SignedString(SecretKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign access token: %w", err)
	}

	// Генерация Refresh токена
	refreshPayload := fmt.Sprintf("%s|%s|%d", userID, clientIP, time.Now().UnixNano())
	hash := sha512.Sum512([]byte(refreshPayload))
	refreshToken := base64.StdEncoding.EncodeToString(hash[:])

	return accessTokenString, refreshToken, nil
}

func Hash(password string, cost int) (string, error) {
	hash := sha256.Sum256([]byte(password))
	preHashedPassword := hex.EncodeToString(hash[:])
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(preHashedPassword), cost)
	if err != nil {
		return "", fmt.Errorf("failed to generate bcrypt hash: %w", err)
	}

	return string(bcryptHash), nil
}
