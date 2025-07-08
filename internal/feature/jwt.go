package feature

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims tetap sama.
type CustomClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

// Service interface tetap sama.
type Service interface {
	GenerateToken(uuid string) (string, error)
	ValidateToken(tokenString string) (*CustomClaims, error)
}

// JwtServiceConfig sekarang menggunakan kunci RSA.
type JwtServiceConfig struct {
	PrivateKey   *rsa.PrivateKey
	PublicKey    *rsa.PublicKey
	TokenTTLHour time.Duration // Token Time To Live dalam jam
}

// JwtService sekarang memegang konfigurasi dengan kunci RSA.
type JwtService struct {
	config JwtServiceConfig
}

// NewJwtService sekarang membaca kunci dari file PEM.
func NewJwtService() (Service, error) {
	// Ganti path ini dengan path absolut atau gunakan variabel environment
	privateKeyPath := "./configs/private.pem"
	publicKeyPath := "./configs/public.pem"

	// Baca private key
	privateKeyBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not read private key file: %w", err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key: %w", err)
	}

	// Baca public key
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not read public key file: %w", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse public key: %w", err)
	}

	// Konfigurasi TTL (Time-To-Live) tetap sama
	ttlStr := "72" // Sebaiknya diambil dari env var
	if ttlStr == "" {
		ttlStr = "72" // Default TTL
	}
	ttlHour, err := strconv.Atoi(ttlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_TOKEN_TTL_HOUR: %w", err)
	}

	return &JwtService{
		config: JwtServiceConfig{
			PrivateKey:   privateKey,
			PublicKey:    publicKey,
			TokenTTLHour: time.Duration(ttlHour) * time.Hour,
		},
	}, nil
}

// GenerateToken sekarang menggunakan RS256 dan private key.
func (s *JwtService) GenerateToken(uuid string) (string, error) {
	if uuid == "" {
		return "", errors.New("uuid cannot be empty")
	}

	expirationTime := time.Now().Add(s.config.TokenTTLHour)

	claims := &CustomClaims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "cbt-edunusa-app",
		},
	}

	// Gunakan algoritma RS256
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Tandatangani token dengan private key
	signedToken, err := token.SignedString(s.config.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

// ValidateToken sekarang menggunakan public key untuk verifikasi.
func (s *JwtService) ValidateToken(encodedToken string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma adalah RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Kembalikan public key untuk verifikasi
		return s.config.PublicKey, nil
	})

	if err != nil {
		// Penanganan error spesifik tetap relevan
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		// Error ini tidak spesifik untuk RSA, tapi baik untuk dimiliki
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
