package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/sanyuanya/doctor/utils"
)

// JWTMiddleware JWT 验证中间件
func JWTMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		// 从 Authorization header 获取 token
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "缺少授权令牌",
				"status":  fiber.StatusUnauthorized,
			})
		}

		// 检查 Bearer 前缀
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "授权令牌格式错误",
				"status":  fiber.StatusUnauthorized,
			})
		}

		// 提取 token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "授权令牌为空",
				"status":  fiber.StatusUnauthorized,
			})
		}

		// 验证 token
		claims, err := utils.ParseJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "授权令牌无效",
				"status":  fiber.StatusUnauthorized,
			})
		}

		// 将用户信息存储到 context 中，供后续处理器使用
		c.Locals("user_id", claims.UserID)
		c.Locals("open_id", claims.OpenID)
		c.Locals("phone_number", claims.PhoneNumber)
		c.Locals("default_role", claims.DefaultRole)
		c.Locals("active_role", claims.ActiveRole)

		// 继续执行下一个处理器
		return c.Next()
	}
}

// GetUserIDFromContext 从 context 中获取用户ID
func GetUserIDFromContext(c fiber.Ctx) uint {
	userID := c.Locals("user_id")
	if userID == nil {
		return 0
	}
	return userID.(uint)
}

// GetOpenIDFromContext 从 context 中获取 OpenID
func GetOpenIDFromContext(c fiber.Ctx) string {
	openID := c.Locals("open_id")
	if openID == nil {
		return ""
	}
	return openID.(string)
}

// GetPhoneNumberFromContext 从 context 中获取手机号
func GetPhoneNumberFromContext(c fiber.Ctx) string {
	phoneNumber := c.Locals("phone_number")
	if phoneNumber == nil {
		return ""
	}
	return phoneNumber.(string)
}

// GetDefaultRoleFromContext 从 context 中获取默认角色
func GetDefaultRoleFromContext(c fiber.Ctx) string {
	defaultRole := c.Locals("default_role")
	if defaultRole == nil {
		return ""
	}
	return defaultRole.(string)
}

// GetActiveRoleFromContext 从 context 中获取当前角色
func GetActiveRoleFromContext(c fiber.Ctx) string {
	activeRole := c.Locals("active_role")
	if activeRole == nil {
		return ""
	}
	return activeRole.(string)
}
