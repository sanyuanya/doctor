package utils

// ...existing code...

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/skip2/go-qrcode"
)

// ...existing code...

// 生成绑定码
func GenerateBindingCode() string {
	timestamp := time.Now().Format("20060102150405")
	randomNumber := fmt.Sprintf("%06d", rand.Intn(1000000)) // 生成 6 位随机数
	return timestamp + randomNumber
}

// 生成短绑定码（用于扫码）
func GenerateShortBindingCode() string {
	return fmt.Sprintf("%08d", rand.Intn(100000000)) // 8位数字码
}

// 生成二维码图片（返回 base64 编码）
func GenerateQRCodeBase64(content string, size int) (string, error) {
	if size == 0 {
		size = 256 // 默认尺寸
	}

	// 生成二维码
	qr, err := qrcode.New(content, qrcode.Medium)
	if err != nil {
		return "", err
	}

	// 转换为 PNG 字节数组
	pngBytes, err := qr.PNG(size)
	if err != nil {
		return "", err
	}

	// 转换为 base64
	base64String := base64.StdEncoding.EncodeToString(pngBytes)
	return "data:image/png;base64," + base64String, nil
}

// 生成二维码文件
func GenerateQRCodeFile(content, filename string, size int) error {
	if size == 0 {
		size = 256
	}

	return qrcode.WriteFile(content, qrcode.Medium, size, filename)
}

// 生成二维码字节数组
func GenerateQRCodeBytes(content string, size int) ([]byte, error) {
	if size == 0 {
		size = 256
	}

	return qrcode.Encode(content, qrcode.Medium, size)
}
