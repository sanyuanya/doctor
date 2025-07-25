package validators

import (
	"github.com/go-playground/validator/v10"
)

type UserEditRequest struct {
	UserID                   uint   `json:"user_id" validate:"required"`
	Avatar                   string `json:"avatar"`
	Nickname                 string `json:"nickname"`
	Gender                   string `json:"gender"`
	BirthDate                string `json:"birth_date"`
	HeightCm                 string `json:"height_cm"`
	WeightKg                 string `json:"weight_kg"`
	PhoneNumber              string `json:"phone_number"`
	OpenID                   string `json:"open_id"`
	SessionKey               string `json:"session_key"`
	UnionID                  string `json:"union_id"`
	EmergencyContactName     string `json:"emergency_contact_name"`
	EmergencyContactRelation string `json:"emergency_contact_relation"`
	EmergencyContactPhone    string `json:"emergency_contact_phone"`
	DefaultRole              string `json:"default_role"`
	ActiveRole               string `json:"active_role"`
	GroupType                string `json:"group_type"`
	RelationID               uint   `json:"relation_id"`
	PatientNotification      string `json:"patient_notification"`
	ConsultantNotification   string `json:"consultant_notification"`
}

func (m *UserEditRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(m)
}

type MiniLoginRequest struct {
	AppID       string `json:"app_id" validate:"required"`
	LoginCode   string `json:"login_code" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

func (m *MiniLoginRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(m)
}

type GetUserPhoneNumberRequest struct {
	AppID              string `json:"app_id" validate:"required"`
	GetPhoneNumberCode string `json:"get_phone_number_code" validate:"required"`
}

func (m *GetUserPhoneNumberRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(m)
}
