package controllers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func FileUpload(c fiber.Ctx) error {
	// 解析 multipart form
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "没有文件上传",
			"status":  fiber.StatusBadRequest,
		})
	}

	// 检查文件大小 (例如限制为 10MB)
	maxSize := int64(10 << 20) // 10MB
	if file.Size > maxSize {
		return c.JSON(fiber.Map{
			"message": "文件大小超过限制(10MB)",
			"status":  fiber.StatusBadRequest,
		})
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "打开文件失败",
			"status":  fiber.StatusBadRequest,
		})
	}
	defer src.Close()

	// 创建以日期命名的文件夹 (e.g., public/storage/2025-03-21/)
	dateFolder := time.Now().Format(time.DateOnly)
	uploadPath := filepath.Join("public/storage", dateFolder)
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return c.JSON(fiber.Map{
			"message": "创建目录失败",
			"status":  fiber.StatusInternalServerError,
		})
	}

	// 生成新的文件名，防止文件名冲突
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().UnixNano(), ext)

	// 创建目标文件
	filePath := filepath.Join(uploadPath, newFileName)
	dst, err := os.Create(filePath)
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "创建文件失败",
			"status":  fiber.StatusInternalServerError,
		})
	}
	defer dst.Close()

	// 复制文件内容
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(fiber.Map{
			"message": "文件保存失败",
			"status":  fiber.StatusInternalServerError,
		})
	}

	// 返回访问 URL
	fileURL := fmt.Sprintf("/storage/%s/%s", dateFolder, newFileName)

	return c.JSON(fiber.Map{
		"message": "文件上传成功",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"path":     fileURL,
			"filename": file.Filename,
			"size":     file.Size,
		},
	})
}
