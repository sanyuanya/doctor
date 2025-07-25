package controllers

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/sanyuanya/doctor/entities"
	"github.com/sanyuanya/doctor/middlewares"
	"github.com/sanyuanya/doctor/validators"
)

func FeedbackSave(c fiber.Ctx) error {

	request := &validators.FeedbackCreateRequest{}

	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	feedbackEntity := &entities.Feedback{
		Content: request.Content,
		Status:  "pending",
		UserID:  middlewares.GetUserIDFromContext(c),
		File:    request.File,
	}

	data, err := feedbackEntity.Insert()
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	return c.JSON(fiber.Map{"message": "success", "status": fiber.StatusOK, "data": fiber.Map{
		"feedback_id": data.FeedbackID,
		"file":        data.File,
		"user_id":     data.UserID,
		"status":      data.Status,
		"created_at":  data.CreatedAt.Format(time.DateTime),
		"updated_at":  data.UpdatedAt.Format(time.DateTime),
	}})
}
