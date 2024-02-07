package utils

import (
	// "context"
	"errors"
	"net/http"
	"strings"
	"time"

    "golang.org/x/crypto/bcrypt"
    "github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)


// HashPassword menghasilkan hash dari password yang diberikan
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash memeriksa apakah password yang diberikan cocok dengan hash yang diberikan
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


// JWTKey adalah kunci untuk menandatangani token JWT
var JWTKey = []byte("secret_key")

// GenerateToken menghasilkan token JWT dengan userID yang diberikan
func GenerateToken(userID uint) (string, error) {
    // Buat payload token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["authorized"] = true
    claims["userID"] = userID
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token berlaku selama 24 jam

    // Sign token dengan secret key
    tokenString, err := token.SignedString(JWTKey)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// VerifyToken memeriksa apakah token JWT valid dan mengembalikan userID jika valid
func VerifyToken(r *http.Request) (uint, error) {
    // Ambil token dari header Authorization
    tokenString := extractToken(r)
    if tokenString == "" {
        return 0, errors.New("token not found")
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validasi tipe token
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, jwt.ErrInvalidKeyType
        }
        return JWTKey, nil
    })

    if err != nil {
        return 0, err
    }

    // Periksa apakah token valid
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        userID := uint(claims["userID"].(float64))
        return userID, nil
    } else {
        return 0, jwt.ErrInvalidKey
    }
}


// extractToken mengambil token JWT dari header Authorization
func extractToken(r *http.Request) string {
    // Ambil nilai dari header Authorization
    bearerToken := r.Header.Get("Authorization")

    // Periksa apakah header Authorization sesuai dengan format "Bearer <token>"
    if len(bearerToken) > 7 && strings.ToUpper(bearerToken[0:6]) == "BEARER" {
        return bearerToken[7:]
    }
    return ""
}

func MiddlewareJWTAuth(next gin.HandlerFunc) gin.HandlerFunc {
    return func(c *gin.Context) {
        // Implementasi middleware JWT Auth di sini
        userID, err := VerifyToken(c.Request)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        // Simpan userID ke dalam context untuk digunakan di handler selanjutnya
        c.Set("userID", userID)
        next(c)
    }
}



