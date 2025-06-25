package utils

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`

	jwt.RegisteredClaims
}

func getPrivateKey() (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile("keys/private.pem")
	if err != nil {
		return nil, fmt.Errorf("gagal membaca private key: %w", err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("gagal mem-parse private key: %w", err)
	}

	return privateKey, nil
}
func getPublicKey() (*rsa.PublicKey, error) {
	keyBytes, err := os.ReadFile("keys/public.pem")
	if err != nil {
		return nil, fmt.Errorf("gagal membaca public key: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("gagal mem-parse public key: %w", err)
	}

	return publicKey, nil
}

func GenerateJWT(userID, role string) (string, error) {

	privateKey, err := getPrivateKey()
	if err != nil {
		return "", err
	}
	claims := MyCustomClaims{
		userID,
		role,
		jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-auth-server",
			Subject:   "user-login-token",
		},
	}

	// 3. Buat token baru dengan signing method RS256 dan claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// 4. Tandatangani token dengan private key untuk mendapatkan string token
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", fmt.Errorf("gagal menandatangani token: %w", err)
	}

	return signedToken, nil
}

// 1. Baca dan parse public key dari file

func ValidateJWT(tokenString string) (*MyCustomClaims, error) {
	// Dapatkan public key
	publicKey, err := getPublicKey()
	if err != nil {
		return nil, err
	}

	// 2. Parse token dengan claims dan key function
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 3. Validasi signing method! Ini sangat penting untuk keamanan.
		// Pastikan algoritma yang digunakan adalah RSA.
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Jika OK, kembalikan public key untuk verifikasi
		return publicKey, nil
	})

	if err != nil {
		// Cek error spesifik dari JWT, misal: token kedaluwarsa
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		return nil, fmt.Errorf("gagal mem-parse token: %w", err)
	}

	// 4. Cek apakah token valid dan ambil claims-nya
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token tidak valid")
}
