package main

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/sanyuanya/doctor/routes"

	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit:     21 * 1024 * 1024,
		AppName:       "doctor",
		CaseSensitive: true,
	})

	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${time} ${locals:requestid} ${status} - ${method} ${path} ${resBody} ${reqHeaders} ${queryParams} ${bytesSent} ${bytesReceived}\u200b\n",
		TimeZone:   "Asia/Shanghai",
		TimeFormat: time.RFC3339Nano,
	}))

	// 注册路由
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
