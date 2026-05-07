package filters

import (
	"net/http"
	"strings"

	"smartplant-ems/common/utils"

	"github.com/beego/beego/v2/server/web"
)

func JWTAuth() web.FilterFunc {
	return func(c *web.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			token = c.GetCookie("token")
		}

		if token == "" {
			c.Abort(http.StatusUnauthorized, "Unauthorized: missing token")
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimSpace(token)

		claims, err := utils.ParseToken(token)
		if err != nil {
			c.Abort(http.StatusUnauthorized, "Unauthorized: invalid token")
			return
		}

		c.SetSession("user_id", claims.UserID)
		c.SetSession("username", claims.Username)
		c.SetSession("role", claims.Role)
	}
}

func AdminOnly() web.FilterFunc {
	return func(c *web.Context) {
		role := c.GetSession("role")
		if role != "admin" {
			c.Abort(http.StatusForbidden, "Forbidden: admin only")
			return
		}
	}
}