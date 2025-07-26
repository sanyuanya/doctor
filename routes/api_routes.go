package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/sanyuanya/doctor/controllers"
	"github.com/sanyuanya/doctor/middlewares"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/doctor")

	api.Post("/wechat/mini/jscodeToSession", controllers.JsCodeToSession)

	api.Post("/wechat/mini/getUserPhoneNumber", controllers.GetUserPhoneNumber)

	// 需要验证 token 的接口组
	protected := api.Group("", middlewares.JWTMiddleware())

	protected.Post("/patient/index", controllers.BloodGlucoseRecordIndex)

	protected.Post("/patient/upload", controllers.BloodGlucoseRecordStore)

	protected.Post("/patient/inactive-users", controllers.BloodGlucoseRecordInactiveUsers)

	protected.Post("/user/profile-edit", controllers.UserEdit)

	protected.Post("/user/profile-info", controllers.UserInfo)

	protected.Post("/user/generateQRCode", controllers.GenerateQRCode)

	protected.Post("/file/upload", controllers.FileUpload)

	protected.Post("/feedback/create", controllers.FeedbackSave)

	api.Get("/storage/*", static.New("./public/storage"))

}
