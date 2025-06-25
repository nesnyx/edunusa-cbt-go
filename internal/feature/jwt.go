package feature

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UUID string `json:"uuid"`
	jwt.RegisteredClaims
}

type Service interface {
	GenerateToken(uuid string) (string, error)
	ValidateToken(tokenString string) (*CustomClaims, error) // Mengembalikan CustomClaims agar lebih mudah diakses
}

type JwtServiceConfig struct {
	SecretKey    []byte
	TokenTTLHour time.Duration // Token Time To Live dalam jam
}

type JwtService struct {
	config JwtServiceConfig
}

func NewJwtService() (Service, error) {
	secret := "60a3d29eea4"
	if secret == "" {
		return nil, errors.New("SECRET_KEY_JWT environment variable not set")
	}

	ttlStr := "48"
	if ttlStr == "" {
		ttlStr = "48" // Default TTL jika tidak diset
	}
	ttlHour, err := strconv.Atoi(ttlStr)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_TOKEN_TTL_HOUR: %w", err)
	}

	return &JwtService{
		config: JwtServiceConfig{
			SecretKey:    []byte(secret),
			TokenTTLHour: time.Duration(ttlHour) * time.Hour,
		},
	}, nil
}

// GenerateToken menghasilkan JWT baru.
func (s *JwtService) GenerateToken(uuid string) (string, error) {
	if uuid == "" {
		return "", errors.New("uuid cannot be empty") // Pesan error diperbaiki
	}

	expirationTime := time.Now().Add(s.config.TokenTTLHour)

	claims := &CustomClaims{
		UUID: uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "my-app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(s.config.SecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

// ValidateToken memvalidasi token yang diberikan.
// Mengembalikan CustomClaims untuk kemudahan akses ke data uuid.
func (s *JwtService) ValidateToken(encodedToken string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(encodedToken, claims, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma sesuai dengan yang diharapkan
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.config.SecretKey, nil
	})

	if err != nil {
		// Pustaka akan mengembalikan error spesifik jika token tidak valid karena berbagai alasan
		// (misalnya, kadaluwarsa, signature tidak cocok, format salah).
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token has expired")
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, errors.New("invalid token signature")
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Pengecekan token.Valid juga penting, meskipun ParseWithClaims sudah melakukan banyak validasi.
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// claims sudah di-populate oleh ParseWithClaims
	return claims, nil
}
