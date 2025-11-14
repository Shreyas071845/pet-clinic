package auth

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"pet-clinic/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ctxKey string

const ClaimsContextKey ctxKey = "jwt_claims"

var jwtKey []byte

func init() {
	_ = godotenv.Load()
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
}

// GenerateJWT creates a token for a username and role
func GenerateJWT(username, role string) (string, error) {
	utils.Log.WithFields(map[string]interface{}{
		"username": username,
		"role":     role,
	}).Info("Generating JWT token")

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(jwtKey)
	if err != nil {
		utils.Log.WithError(err).Error("Failed signing JWT token")
		return "", err
	}
	return signed, nil
}

// JWTMiddleware validates JWT and stores claims in request context
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.Log.Warn("Missing Authorization header")
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// allow both "Bearer <token>" and raw token
		tokenString := strings.TrimSpace(authHeader)
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = tokenString[7:]
		}

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			utils.Log.WithError(err).Error("Invalid or expired JWT token")
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// store claims in context for downstream handlers
		ctx := context.WithValue(r.Context(), ClaimsContextKey, *claims)
		utils.Log.WithField("path", r.URL.Path).Info("JWT token validated successfully")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetClaims extracts jwt.MapClaims from request context (safe wrapper)
func GetClaims(r *http.Request) (jwt.MapClaims, bool) {
	val := r.Context().Value(ClaimsContextKey)
	if val == nil {
		return nil, false
	}
	if mc, ok := val.(jwt.MapClaims); ok {
		return mc, true
	}
	// older package encoding may have map[string]interface{}; try to convert
	if raw, ok := val.(map[string]interface{}); ok {
		return jwt.MapClaims(raw), true
	}
	return nil, false
}
