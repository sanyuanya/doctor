package controllers

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/sanyuanya/doctor/entities"
	"github.com/sanyuanya/doctor/middlewares"
	"github.com/sanyuanya/doctor/validators"
)

func BloodGlucoseRecordIndex(c fiber.Ctx) error {

	request := &validators.BloodGlucoseRecordSearchRequest{}

	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	bloodGlucoseRecordEntity := &entities.BloodGlucoseRecord{}

	data, err := bloodGlucoseRecordEntity.Index(request)
	if err != nil {
		log.Printf("查询血糖数据失败: %s\n", err)
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	total, err := bloodGlucoseRecordEntity.Count(request)
	if err != nil {
		log.Printf("统计血糖数据失败: %s\n", err)
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"data":  data,
			"total": total,
		},
	})
}

func BloodGlucoseRecordStore(c fiber.Ctx) error {
	request := &validators.BloodGlucoseRecordSaveRequest{}

	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	bloodGlucoseRecordEntity := &entities.BloodGlucoseRecord{
		BloodGlucoseRecordID: request.BloodGlucoseRecordID,
		UploadTime:           request.UploadTime,
		Notes:                request.Notes,
		UserID:               middlewares.GetUserIDFromContext(c),
	}

	if request.BloodGlucoseRecordID == 0 {
		data, err := bloodGlucoseRecordEntity.Insert()
		if err != nil {
			log.Printf("保存血糖记录失败: %s\n", err)
			return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusInternalServerError})
		}

		return c.JSON(fiber.Map{"message": "success", "status": fiber.StatusOK, "data": data})

	}

	if _, err := bloodGlucoseRecordEntity.FindByID(); err != nil {
		log.Printf("查询血糖记录失败: %s\n", err)
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusNotFound})
	}

	if _, err := bloodGlucoseRecordEntity.Update(); err != nil {
		log.Printf("更新血糖数据失败: %s\n", err)
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	return c.JSON(fiber.Map{"message": "success", "status": fiber.StatusOK, "data": fiber.Map{
		"upload_time":             request.UploadTime,
		"notes":                   request.Notes,
		"blood_glucose_record_id": request.BloodGlucoseRecordID,
	}})
}

func BloodGlucoseRecordInactiveUsers(c fiber.Ctx) error {
	request := &validators.InactiveUsersRequest{}
	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	bloodGlucoseRecordEntity := &entities.BloodGlucoseRecord{}

	data, err := bloodGlucoseRecordEntity.GetInactiveUsers(request)

	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	total, err := bloodGlucoseRecordEntity.GetInactiveUsersCount(request)

	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"data":  data,
			"total": total,
		},
	})
}
