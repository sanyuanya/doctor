package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/sanyuanya/doctor/entities"
	"github.com/sanyuanya/doctor/middlewares"
	"github.com/sanyuanya/doctor/utils"
	"github.com/sanyuanya/doctor/validators"
	"github.com/sanyuanya/doctor/wechat"
)

func JsCodeToSession(c fiber.Ctx) error {
	request := &validators.MiniLoginRequest{}
	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	miniEntity := &entities.Mini{
		AppID: request.AppID,
	}

	miniData, err := miniEntity.FindAccessTokenByAppID()
	if err != nil {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("获取 appId 信息失败: %s", err.Error()), "status": fiber.StatusBadRequest})
	}

	// 如果 access_token 过期，重新获取
	// 这里减去 30 秒是为了避免在后续的请求中再次过期
	if miniData.ExpiresIn-60 <= time.Now().Unix() {
		accessTokenResp, err := wechat.GetStableAccessToken(miniData.AppID, miniData.AppSecret)
		if err != nil {
			return c.JSON(fiber.Map{"message": fmt.Sprintf("获取 access_token 失败: %s", err.Error()), "status": fiber.StatusBadRequest})
		}

		miniEntity.AccessToken = accessTokenResp.AccessToken
		miniEntity.ExpiresIn = time.Now().Unix() + accessTokenResp.ExpiresIn - 60
		err = miniEntity.UpdateAccessTokenAndExpiresIn()
		if err != nil {
			return c.JSON(fiber.Map{"message": fmt.Sprintf("更新 access_token 失败: %s", err.Error()), "status": fiber.StatusBadRequest})
		}
	}

	// 获取 openid 和 session_key
	// 这里的 loginCode 是小程序登录时获取的 code
	code2SessionResp, err := wechat.Code2Session(miniData.AppID, miniData.AppSecret, request.LoginCode)
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})

	}
	if code2SessionResp.ErrCode != 0 {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("获取 openId 失败: %v", code2SessionResp.ErrMsg), "status": fiber.StatusBadGateway})
	}

	userEntity := &entities.User{
		OpenID:      code2SessionResp.OpenID,
		UnionID:     code2SessionResp.UnionID,
		SessionKey:  code2SessionResp.SessionKey,
		PhoneNumber: request.PhoneNumber,
	}

	// 查询用户是否已存在
	data, err := userEntity.FindByOpenID()

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("根据 openID 查询用户失败: %s\n", err)
		return c.JSON(fiber.Map{"message": fmt.Sprintf("查询用户失败: %v", err), "status": fiber.StatusInternalServerError})
	}

	if data == nil {
		data, err = userEntity.Register()
		if err != nil {
			log.Printf("用户注册失败: %s\n", err)
			return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
		}
	}

	// 生成 JWT token
	token, err := utils.GenerateJWT(
		data.UserID,
		data.OpenID,
		data.PhoneNumber,
		data.DefaultRole,
		data.ActiveRole,
	)
	if err != nil {
		return c.JSON(fiber.Map{"message": "生成token失败", "status": fiber.StatusInternalServerError})
	}

	// 将 token 设置到 header 中
	c.Set("Authorization", "Bearer "+token)

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"user_id":                    data.UserID,
			"avatar":                     data.Avatar,
			"nickname":                   data.Nickname,
			"gender":                     data.Gender,
			"birth_date":                 data.BirthDate,
			"height_cm":                  data.HeightCm,
			"weight_kg":                  data.WeightKg,
			"phone_number":               data.PhoneNumber,
			"open_id":                    data.OpenID,
			"session_key":                data.SessionKey,
			"union_id":                   data.UnionID,
			"emergency_contact_name":     data.EmergencyContactName,
			"emergency_contact_relation": data.EmergencyContactRelation,
			"emergency_contact_phone":    data.EmergencyContactPhone,
			"default_role":               data.DefaultRole,
			"active_role":                data.ActiveRole,
			"group_type":                 data.GroupType,
			"relation_id":                data.RelationID,
			"patient_notification":       data.PatientNotification,
			"consultant_notification":    data.ConsultantNotification,
			"created_at":                 data.CreatedAt.Format(time.DateTime),
			"updated_at":                 data.UpdatedAt.Format(time.DateTime),
			"invite_code":                data.InviteCode,
		},
	})
}

func GetUserPhoneNumber(c fiber.Ctx) error {
	request := &validators.GetUserPhoneNumberRequest{}
	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	miniEntity := &entities.Mini{
		AppID: request.AppID,
	}

	miniData, err := miniEntity.FindAccessTokenByAppID()
	if err != nil {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("获取 appId 信息失败: %v", err), "status": fiber.StatusBadRequest})
	}

	if miniData.ExpiresIn-60 <= time.Now().Unix() {
		accessTokenResp, err := wechat.GetStableAccessToken(miniData.AppID, miniData.AppSecret)
		if err != nil {
			return c.JSON(fiber.Map{"message": fmt.Sprintf("获取 access_token 失败: %s", err.Error()), "status": fiber.StatusBadRequest})
		}

		miniEntity.AccessToken = accessTokenResp.AccessToken
		// 提前 60 秒过期，防止临界点失效
		miniEntity.ExpiresIn = time.Now().Unix() + accessTokenResp.ExpiresIn - 60
		err = miniEntity.UpdateAccessTokenAndExpiresIn()
		if err != nil {
			return c.JSON(fiber.Map{"message": fmt.Sprintf("更新 access_token 失败: %s", err.Error()), "status": fiber.StatusBadRequest})
		}
	}

	phoneNumberResp, err := wechat.GetPhoneNumber(request.GetPhoneNumberCode, miniData.AccessToken)
	if err != nil {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("获取手机号失败: %s", err.Error()), "status": fiber.StatusBadRequest})
	}
	if phoneNumberResp.Errcode != 0 {
		return c.JSON(fiber.Map{"message": fmt.Sprintf("获取用户手机号失败: %s", phoneNumberResp.Errmsg), "status": fiber.StatusBadRequest})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"phone_number":      phoneNumberResp.PhoneInfo.PhoneNumber,
			"pure_phone_number": phoneNumberResp.PhoneInfo.PurePhoneNumber,
			"country_code":      phoneNumberResp.PhoneInfo.CountryCode,
			"watermark": fiber.Map{
				"timestamp": phoneNumberResp.PhoneInfo.Watermark.Timestamp,
				"app_id":    phoneNumberResp.PhoneInfo.Watermark.Appid,
			},
		},
	})
}

func UserEdit(c fiber.Ctx) error {
	request := &validators.UserEditRequest{}
	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	user := &entities.User{
		UserID: request.UserID,
	}

	data, err := user.FindByID()

	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusNotFound})
	}

	if request.Avatar != "" {
		data.Avatar = request.Avatar
	}

	if request.Nickname != "" {
		data.Nickname = request.Nickname
	}
	if request.Gender != "" {
		data.Gender = request.Gender
	}
	if request.BirthDate != "" {
		data.BirthDate = request.BirthDate
	}
	if request.HeightCm != "" {
		data.HeightCm = request.HeightCm
	}
	if request.WeightKg != "" {
		data.WeightKg = request.WeightKg
	}
	if request.PhoneNumber != "" {
		data.PhoneNumber = request.PhoneNumber
	}
	if request.OpenID != "" {
		data.OpenID = request.OpenID
	}
	if request.SessionKey != "" {
		data.SessionKey = request.SessionKey
	}
	if request.UnionID != "" {
		data.UnionID = request.UnionID
	}
	if request.EmergencyContactName != "" {
		data.EmergencyContactName = request.EmergencyContactName
	}
	if request.EmergencyContactRelation != "" {
		data.EmergencyContactRelation = request.EmergencyContactRelation
	}
	if request.EmergencyContactPhone != "" {
		data.EmergencyContactPhone = request.EmergencyContactPhone
	}
	if request.DefaultRole != "" {
		data.DefaultRole = request.DefaultRole
	}
	if request.ActiveRole != "" {
		data.ActiveRole = request.ActiveRole
	}
	if request.GroupType != "" {
		data.GroupType = request.GroupType
	}
	if request.RelationID != 0 {
		data.RelationID = request.RelationID
	}
	if request.PatientNotification != "" {
		data.PatientNotification = request.PatientNotification
	}

	if request.ConsultantNotification != "" {
		data.ConsultantNotification = request.ConsultantNotification
	}

	if request.DefaultRole == "consultant" && data.InviteCode == "" {
		data.InviteCode = utils.GenerateBindingCode()
	}

	if _, err := data.Update(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"user_id":                    data.UserID,
			"avatar":                     data.Avatar,
			"nickname":                   data.Nickname,
			"gender":                     data.Gender,
			"birth_date":                 data.BirthDate,
			"height_cm":                  data.HeightCm,
			"weight_kg":                  data.WeightKg,
			"phone_number":               data.PhoneNumber,
			"open_id":                    data.OpenID,
			"session_key":                data.SessionKey,
			"union_id":                   data.UnionID,
			"emergency_contact_name":     data.EmergencyContactName,
			"emergency_contact_relation": data.EmergencyContactRelation,
			"emergency_contact_phone":    data.EmergencyContactPhone,
			"default_role":               data.DefaultRole,
			"active_role":                data.ActiveRole,
			"group_type":                 data.GroupType,
			"relation_id":                data.RelationID,
			"patient_notification":       data.PatientNotification,
			"consultant_notification":    data.ConsultantNotification,
			"created_at":                 data.CreatedAt.Format(time.DateTime),
			"updated_at":                 data.UpdatedAt.Format(time.DateTime),
			"invite_code":                data.InviteCode,
		},
	})

}

func UserInfo(c fiber.Ctx) error {

	request := &validators.UserEditRequest{}
	if err := c.Bind().Body(request); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	if err := request.Validate(); err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadRequest})
	}

	user := &entities.User{
		UserID: request.UserID,
	}

	data, err := user.FindByID()

	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusNotFound})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"user_id":                    data.UserID,
			"avatar":                     data.Avatar,
			"nickname":                   data.Nickname,
			"gender":                     data.Gender,
			"birth_date":                 data.BirthDate,
			"height_cm":                  data.HeightCm,
			"weight_kg":                  data.WeightKg,
			"phone_number":               data.PhoneNumber,
			"open_id":                    data.OpenID,
			"session_key":                data.SessionKey,
			"union_id":                   data.UnionID,
			"emergency_contact_name":     data.EmergencyContactName,
			"emergency_contact_relation": data.EmergencyContactRelation,
			"emergency_contact_phone":    data.EmergencyContactPhone,
			"default_role":               data.DefaultRole,
			"active_role":                data.ActiveRole,
			"group_type":                 data.GroupType,
			"relation_id":                data.RelationID,
			"patient_notification":       data.PatientNotification,
			"consultant_notification":    data.ConsultantNotification,
			"created_at":                 data.CreatedAt.Format(time.DateTime),
			"updated_at":                 data.UpdatedAt.Format(time.DateTime),
			"invite_code":                data.InviteCode,
		},
	})
}

func GenerateQRCode(c fiber.Ctx) error {
	user := &entities.User{
		UserID: middlewares.GetUserIDFromContext(c),
	}

	data, err := user.FindByID()

	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusNotFound})
	}

	qRCode, err := utils.GenerateQRCodeBase64(fmt.Sprintf("/page/index/index?doctor_id=%d", data.UserID), 256)
	if err != nil {
		log.Printf("二维码生成失败: %s\n", err)
		return c.JSON(fiber.Map{"message": err.Error(), "status": fiber.StatusBadGateway})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"status":  fiber.StatusOK,
		"data": fiber.Map{
			"user_id":                    data.UserID,
			"avatar":                     data.Avatar,
			"nickname":                   data.Nickname,
			"gender":                     data.Gender,
			"birth_date":                 data.BirthDate,
			"height_cm":                  data.HeightCm,
			"weight_kg":                  data.WeightKg,
			"phone_number":               data.PhoneNumber,
			"open_id":                    data.OpenID,
			"session_key":                data.SessionKey,
			"union_id":                   data.UnionID,
			"emergency_contact_name":     data.EmergencyContactName,
			"emergency_contact_relation": data.EmergencyContactRelation,
			"emergency_contact_phone":    data.EmergencyContactPhone,
			"default_role":               data.DefaultRole,
			"active_role":                data.ActiveRole,
			"group_type":                 data.GroupType,
			"relation_id":                data.RelationID,
			"patient_notification":       data.PatientNotification,
			"consultant_notification":    data.ConsultantNotification,
			"created_at":                 data.CreatedAt.Format(time.DateTime),
			"updated_at":                 data.UpdatedAt.Format(time.DateTime),
			"invite_code":                data.InviteCode,
			"qr_code":                    qRCode,
		},
	})
}
