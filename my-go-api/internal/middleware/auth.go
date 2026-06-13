package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"my-go-api/internal/repository" // Import your repository package

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID uint     `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

// UPDATED: Now accepts repository.UserRepository to perform DB session validation
func AuthMiddleware(jwtSecret []byte, userRepo repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &CustomClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// ─── NEW SESSION VERIFICATION LOGIC ──────────────────────────────────
		// Check PostgreSQL to see if this specific token has been logged out/deleted
		isValid, err := userRepo.IsSessionValid(tokenString)
		if err != nil || !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been invalidated by a logout action"})
			c.Abort()
			return
		}
		// ─────────────────────────────────────────────────────────────────────

		// Inject key contexts into Gin lifecycle pipeline
		c.Set("userID", claims.UserID)
		c.Set("userRoles", claims.Roles)
		c.Next()
	}
}

// RequireRoles acts as an RBAC guard filter
func RequireRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolesInterface, exists := c.Get("userRoles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied: missing privileges"})
			c.Abort()
			return
		}

		userRoles := rolesInterface.([]string)
		roleMap := make(map[string]bool)
		for _, r := range userRoles {
			roleMap[r] = true
		}

		for _, allowed := range allowedRoles {
			if roleMap[allowed] {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient role permission"})
		c.Abort()
	}
}
