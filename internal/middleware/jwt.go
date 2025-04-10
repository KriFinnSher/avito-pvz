package middleware

import (
	"avito-pvz/internal/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// JwtMiddleware checks for user's JWT-token and sets "email" and "role" vars in context
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
		}

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
		}

		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token payload: missing email")
		}
		c.Set("email", email)

		role, ok := claims["role"].(string)
		if !ok || role == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token payload: missing role")
		}
		c.Set("role", role)

		return next(c)
	}
}
