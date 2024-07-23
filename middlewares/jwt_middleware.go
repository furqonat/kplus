package middlewares

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"

	"kplus.com/utils"
)

type JwtMiddleware struct {
	env utils.Env
	jwt utils.Jwt
}

func NewJwtMiddleware(env utils.Env, jwt utils.Jwt) JwtMiddleware {
	return JwtMiddleware{
		env: env,
		jwt: jwt,
	}
}

func (j JwtMiddleware) HandleAuthWithRoles(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		idToken, err := j.getTokenFromHeaders(c)
		if err != nil {
			data := utils.ResponseError{
				Message: "no token provided",
			}
			return c.Status(fiber.StatusUnauthorized).JSON(data)
		}

		token, err := j.jwt.VerifyToken(idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError{
				Message: "invalid token",
			})
		}

		if len(roles) < 1 {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError{
				Message: "no roles provided",
			})
		}

		if len(roles) > 0 {
			if ok := j.checkRoleIsValid(roles, &token); !ok {
				return c.Status(fiber.StatusBadRequest).JSON(utils.ResponseError{
					Message: "invalid role",
				})
			}
		}
		if token.TokenType != utils.AccessToken && c.Path() != "/auth/refresh" {
			return c.Status(fiber.StatusUnauthorized).JSON(utils.ResponseError{
				Message: "invalid token type",
			})
		}
		c.Locals(utils.Token, token)
		return c.Next()
	}
}

func (j JwtMiddleware) getTokenFromHeaders(c *fiber.Ctx) (string, error) {
	bearer := c.Get("Authorization")
	if c.Path() == "/auth/refresh" {
		bearer = c.Get("Authorization-refresh")
	}
	if bearer == "" {
		return "", errors.New("no authorization header provided")
	}
	token := strings.TrimPrefix(bearer, "Bearer ")
	return token, nil
}

func (j JwtMiddleware) checkRoleIsValid(roles []string, token *utils.JwtCustomClaims) bool {
	for _, val := range roles {
		if val == token.Role {
			return true
		}
	}
	return false

}
