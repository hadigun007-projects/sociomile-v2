package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RBACMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {

			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: role not found"})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied: invalid role format"})
			c.Abort()
			return
		}

		for _, role := range allowedRoles {
			if role == roleStr {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied: insufficient permissions",
			"debug": map[string]interface{}{
				"user_role":     roleStr,
				"allowed_roles": allowedRoles,
			},
		})
		c.Abort()
	}
}
