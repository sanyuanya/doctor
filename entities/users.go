package entities

import (
	"database/sql"
	"time"
)

type User struct {
	UserID                   uint      `json:"user_id"`
	Avatar                   string    `json:"avatar"`
	Nickname                 string    `json:"nickname"`
	Gender                   string    `json:"gender"`
	BirthDate                string    `json:"birth_date"`
	HeightCm                 string    `json:"height_cm"`
	WeightKg                 string    `json:"weight_kg"`
	PhoneNumber              string    `json:"phone_number"`
	OpenID                   string    `json:"open_id"`
	SessionKey               string    `json:"session_key"`
	UnionID                  string    `json:"union_id"`
	EmergencyContactName     string    `json:"emergency_contact_name"`
	EmergencyContactRelation string    `json:"emergency_contact_relation"`
	EmergencyContactPhone    string    `json:"emergency_contact_phone"`
	DefaultRole              string    `json:"default_role"`
	ActiveRole               string    `json:"active_role"`
	GroupType                string    `json:"group_type"`
	RelationID               uint      `json:"relation_id"`
	PatientNotification      string    `json:"patient_notification"`
	ConsultantNotification   string    `json:"consultant_notification"`
	InviteCode               string    `json:"invite_code"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

func (u *User) Register() (*User, error) {
	err := db.QueryRow(`INSERT INTO users(open_id, session_key, union_id, phone_number) VALUES($1, $2, $3, $4) RETURNING 
			user_id,
			avatar,
			nickname,
			gender,
			birth_date,
			height_cm,
			weight_kg,
			phone_number,
			open_id,
			session_key,
			union_id,
			emergency_contact_name,
			emergency_contact_relation,
			emergency_contact_phone,
			default_role,
			active_role,
			group_type,
			relation_id,
			patient_notification,
			consultant_notification,
			created_at,
			updated_at,
			invite_code`, u.OpenID, u.SessionKey, u.UnionID, u.PhoneNumber).Scan(
		u.UserID,
		u.Avatar,
		u.Nickname,
		u.Gender,
		u.BirthDate,
		u.HeightCm,
		u.WeightKg,
		u.PhoneNumber,
		u.OpenID,
		u.SessionKey,
		u.UnionID,
		u.EmergencyContactName,
		u.EmergencyContactRelation,
		u.EmergencyContactPhone,
		u.DefaultRole,
		u.ActiveRole,
		u.GroupType,
		u.RelationID,
		u.PatientNotification,
		u.ConsultantNotification,
		u.CreatedAt,
		u.UpdatedAt,
		u.InviteCode,
	)
	return u, err
}

func (u *User) Update() (sql.Result, error) {
	return db.Exec(`
		UPDATE users SET 
			avatar = $1,
			nickname = $2,
			gender = $3, 
			birth_date = $4, 
			height_cm = $5, 
			weight_kg = $6, 
			phone_number = $7, 
			open_id = $8, 
			session_key = $9, 
			union_id = $10, 
			emergency_contact_name = $11, 
			emergency_contact_relation = $12, 
			emergency_contact_phone = $13, 
			default_role = $14, 
			active_role = $15, 
			group_type = $16, 
			relation_id = $17, 
			patient_notification = $18, 
			consultant_notification = $19, 
			updated_at = NOW(),
			invite_code = $20
		WHERE user_id = $21`,
		u.Avatar,
		u.Nickname,
		u.Gender,
		u.BirthDate,
		u.HeightCm,
		u.WeightKg,
		u.PhoneNumber,
		u.OpenID,
		u.SessionKey,
		u.UnionID,
		u.EmergencyContactName,
		u.EmergencyContactRelation,
		u.EmergencyContactPhone,
		u.DefaultRole,
		u.ActiveRole,
		u.GroupType,
		u.RelationID,
		u.PatientNotification,
		u.ConsultantNotification,
		u.InviteCode,
		u.UserID,
	)
}

func (u *User) FindByID() (*User, error) {
	if err := db.QueryRow(`
		SELECT
			user_id,
			avatar,
			nickname,
			gender,
			birth_date,
			height_cm,
			weight_kg,
			phone_number,
			open_id,
			session_key,
			union_id,
			emergency_contact_name,
			emergency_contact_relation,
			emergency_contact_phone,
			default_role,
			active_role,
			group_type,
			relation_id,
			patient_notification,
			consultant_notification,
			created_at,
			updated_at,
			invite_code
		FROM
			users
		WHERE deleted_at IS NULL AND user_id = $1
	`, u.UserID).Scan(
		u.UserID,
		u.Avatar,
		u.Nickname,
		u.Gender,
		u.BirthDate,
		u.HeightCm,
		u.WeightKg,
		u.PhoneNumber,
		u.OpenID,
		u.SessionKey,
		u.UnionID,
		u.EmergencyContactName,
		u.EmergencyContactRelation,
		u.EmergencyContactPhone,
		u.DefaultRole,
		u.ActiveRole,
		u.GroupType,
		u.RelationID,
		u.PatientNotification,
		u.ConsultantNotification,
		u.CreatedAt,
		u.UpdatedAt,
		u.InviteCode,
	); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) FindByOpenID() (*User, error) {
	if err := db.QueryRow(`
		SELECT
			user_id,
			avatar,
			nickname,
			gender,
			birth_date,
			height_cm,
			weight_kg,
			phone_number,
			open_id,
			session_key,
			union_id,
			emergency_contact_name,
			emergency_contact_relation,
			emergency_contact_phone,
			default_role,
			active_role,
			group_type,
			relation_id,
			patient_notification,
			consultant_notification,
			created_at,
			updated_at,
			invite_code
		FROM
			users
		WHERE deleted_at IS NULL AND open_id = $1
	`, u.OpenID).Scan(
		u.UserID,
		u.Avatar,
		u.Nickname,
		u.Gender,
		u.BirthDate,
		u.HeightCm,
		u.WeightKg,
		u.PhoneNumber,
		u.OpenID,
		u.SessionKey,
		u.UnionID,
		u.EmergencyContactName,
		u.EmergencyContactRelation,
		u.EmergencyContactPhone,
		u.DefaultRole,
		u.ActiveRole,
		u.GroupType,
		u.RelationID,
		u.PatientNotification,
		u.ConsultantNotification,
		u.CreatedAt,
		u.UpdatedAt,
		u.InviteCode,
	); err != nil {
		return nil, err
	}

	return u, nil
}
