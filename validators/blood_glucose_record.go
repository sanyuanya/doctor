package validators

import "github.com/go-playground/validator/v10"

type BloodGlucoseRecordSearchRequest struct {
	Page      uint  `json:"page"`
	Size      uint  `json:"size"`
	UserID    uint  `json:"user_id"`
	DoctorID  uint  `json:"doctor_id"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
}

func (b *BloodGlucoseRecordSearchRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(b)
}

type BloodGlucoseRecordSaveRequest struct {
	BloodGlucoseRecordID uint   `json:"blood_glucose_record_id"`
	UploadTime           int64  `json:"upload_time" validate:"required"`
	Notes                string `json:"notes" validate:"required"`
}

func (b *BloodGlucoseRecordSaveRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(b)
}

// InactiveUsersRequest 未活跃用户查询请求
type InactiveUsersRequest struct {
	InactiveDays int  `json:"inactive_days"` // 未活跃天数，默认7天
	Page         int  `json:"page"`          // 页码，从1开始
	Size         int  `json:"size"`          // 每页大小，默认20
	DoctorID     uint `json:"doctor_id"`
}

func (i *InactiveUsersRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(i)
}
