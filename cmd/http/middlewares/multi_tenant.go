package middlewares

import "github.com/labstack/echo/v4"

func SetTenant(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set("tenant", c.Request().Header.Get("Tenant"))
		// tenant validation Ex: get token and verify data or get tenant from session
		return next(c)
	}
}
