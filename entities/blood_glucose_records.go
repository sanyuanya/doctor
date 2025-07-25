package entities

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/sanyuanya/doctor/validators"
)

type BloodGlucoseRecord struct {
	BloodGlucoseRecordID uint      `json:"blood_glucose_record_id"`
	UserID               uint      `json:"user_id"`
	UploadTime           int64     `json:"upload_time"`
	Notes                string    `json:"notes"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type BloodGlucoseRecordResponse struct {
	BloodGlucoseRecordID uint      `json:"blood_glucose_record_id"`
	UserID               uint      `json:"user_id"`
	DoctorID             uint      `json:"doctor_id"`
	UploadTime           int64     `json:"upload_time"`
	Notes                string    `json:"notes"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	Avatar               string    `json:"avatar"`
	NickName             string    `json:"nick_name"`
	GroupType            string    `json:"group_type"`
}

func (b *BloodGlucoseRecord) Insert() (*BloodGlucoseRecord, error) {
	if err := db.QueryRow(`INSERT INTO blood_glucose_record(user_id, update_time, notes) VALUES($1, $2, $3) RETURNING blood_glucose_record_id;`,
		b.UserID, b.UploadTime, b.Notes).Scan(&b.BloodGlucoseRecordID); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *BloodGlucoseRecord) Update() (sql.Result, error) {
	return db.Exec(`UPDATE blood_glucose_record SET user_id = $1, update_time = $2, notes = $3, updated_at = NOW() WHERE blood_glucose_record_id = $4`, b.UserID, b.UploadTime, b.Notes, b.BloodGlucoseRecordID)
}

func (b *BloodGlucoseRecord) FindByID() (*BloodGlucoseRecord, error) {

	if err := db.QueryRow(`
		SELECT
			blood_glucose_record_id,
		 	user_id,
			upload_time,
			notes,
			created_at,
			updated_at
		FROM
			blood_glucose_record 
		WHERE deleted_at IS NULL AND blood_glucose_record_id = $1
	`, b.BloodGlucoseRecordID).Scan(
		&b.BloodGlucoseRecordID,
		&b.UserID,
		&b.UploadTime,
		&b.Notes,
		&b.CreatedAt,
		&b.UpdatedAt); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *BloodGlucoseRecord) Index(search *validators.BloodGlucoseRecordSearchRequest) ([]*BloodGlucoseRecordResponse, error) {

	baseSQL := `
		SELECT
			b.blood_glucose_record_id,
		 	b.user_id,
			b.upload_time,
			b.notes,
			b.created_at,
			b.updated_at,
			u.avatar,
			u.nickname,
			u.group_type,
			u.relation_id
		FROM
			blood_glucose_record b
		JOIN users u ON b.user_id = u.user_id
		WHERE deleted_at IS NULL
	`

	index := 1
	execArgs := []any{}

	if search.StartTime != 0 && search.EndTime != 0 {
		baseSQL += fmt.Sprintf(" AND b.upload_time BETWEEN $%d AND $%d", index, index+1)
		index += 2
		execArgs = append(execArgs, search.StartTime, search.EndTime)
	}

	if search.UserID != 0 {
		baseSQL += fmt.Sprintf(" AND b.user_id = $%d", index)
		index++
		execArgs = append(execArgs, search.UserID)
	}

	if search.DoctorID != 0 {
		baseSQL += fmt.Sprintf(" AND u.relation_id = $%d", index)
		index++
		execArgs = append(execArgs, search.DoctorID)
	}

	baseSQL += " ORDER BY created_at DESC, blood_glucose_record_id DESC"

	// limit
	if search.Size > 0 {
		baseSQL += fmt.Sprintf(" LIMIT $%d OFFSET $%d", index, index+1)
		execArgs = append(execArgs, search.Size, (search.Page-1)*search.Size)
	}

	rows, err := db.Query(baseSQL, execArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*BloodGlucoseRecordResponse, 0)

	for rows.Next() {
		var result BloodGlucoseRecordResponse
		err := rows.Scan(
			&result.BloodGlucoseRecordID,
			&result.UserID,
			&result.UploadTime,
			&result.Notes,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.Avatar,
			&result.NickName,
			&result.GroupType,
			&result.DoctorID,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &result)
	}

	return results, nil
}

func (b *BloodGlucoseRecord) Count(search *validators.BloodGlucoseRecordSearchRequest) (uint, error) {

	baseSQL := `
		SELECT
			COUNT(*)
		FROM
			blood_glucose_record b
		JOIN users u ON b.user_id = u.user_id
		WHERE deleted_at IS NULL
	`

	index := 1
	execArgs := []any{}

	if search.StartTime != 0 && search.EndTime != 0 {
		baseSQL += fmt.Sprintf(" AND b.upload_time BETWEEN $%d AND $%d", index, index+1)
		index += 2
		execArgs = append(execArgs, search.StartTime, search.EndTime)
	}

	if search.UserID != 0 {
		baseSQL += fmt.Sprintf(" AND b.user_id = $%d", index)
		index++
		execArgs = append(execArgs, search.UserID)
	}

	if search.DoctorID != 0 {
		baseSQL += fmt.Sprintf(" AND u.relation_id = $%d", index)
		index++
		execArgs = append(execArgs, search.DoctorID)
	}

	var count uint
	err := db.QueryRow(baseSQL, execArgs...).Scan(&count)
	return count, err
}

func (b *BloodGlucoseRecord) GetInactiveUsers(search *validators.InactiveUsersRequest) ([]*User, error) {
	// 计算N天前的时间戳
	daysAgo := time.Now().AddDate(0, 0, -int(search.InactiveDays))
	cutoffTimestamp := daysAgo.Unix()

	baseSQL := `
        SELECT DISTINCT
            u.user_id,
            u.avatar,
            u.nickname,
            u.gender,
            u.birth_date,
            u.height_cm,
            u.weight_kg,
            u.phone_number,
            u.open_id,
            u.session_key,
            u.union_id,
            u.emergency_contact_name,
            u.emergency_contact_relation,
            u.emergency_contact_phone,
            u.default_role,
            u.active_role,
            u.group_type,
            u.relation_id,
            u.patient_notification,
            u.consultant_notification,
            u.created_at,
            u.updated_at,
            u.invite_code,
            COALESCE(MAX(b.upload_time), 0) as last_upload_time
        FROM
            users u
        LEFT JOIN blood_glucose_record b ON u.user_id = b.user_id AND b.deleted_at IS NULL
        WHERE 
            u.deleted_at IS NULL
    `

	index := 1
	execArgs := []any{}

	// 如果指定了医生ID，只查询该医生的患者
	if search.DoctorID != 0 {
		baseSQL += fmt.Sprintf(" AND u.relation_id = $%d", index)
		execArgs = append(execArgs, search.DoctorID)
		index++
	}

	baseSQL += " GROUP BY u.user_id"
	// 筛选出N天内没有上传数据的用户
	baseSQL += fmt.Sprintf(" HAVING COALESCE(MAX(b.upload_time), 0) < $%d", index)
	execArgs = append(execArgs, cutoffTimestamp)
	index++

	baseSQL += " ORDER BY last_upload_time ASC, u.user_id DESC"

	// 分页
	if search.Size > 0 {
		baseSQL += fmt.Sprintf(" LIMIT $%d OFFSET $%d", index, index+1)
		execArgs = append(execArgs, search.Size, (search.Page-1)*search.Size)
	}

	rows, err := db.Query(baseSQL, execArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]*User, 0)

	for rows.Next() {
		var user User
		var lastUploadTime int64

		err := rows.Scan(
			&user.UserID,
			&user.Avatar,
			&user.Nickname,
			&user.Gender,
			&user.BirthDate,
			&user.HeightCm,
			&user.WeightKg,
			&user.PhoneNumber,
			&user.OpenID,
			&user.SessionKey,
			&user.UnionID,
			&user.EmergencyContactName,
			&user.EmergencyContactRelation,
			&user.EmergencyContactPhone,
			&user.DefaultRole,
			&user.ActiveRole,
			&user.GroupType,
			&user.RelationID,
			&user.PatientNotification,
			&user.ConsultantNotification,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.InviteCode,
			&lastUploadTime,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, &user)
	}

	return results, nil
}

func (b *BloodGlucoseRecord) GetInactiveUsersCount(search *validators.InactiveUsersRequest) (uint, error) {
	// 计算N天前的时间戳
	daysAgo := time.Now().AddDate(0, 0, -int(search.InactiveDays))
	cutoffTimestamp := daysAgo.Unix()

	baseSQL := `
        SELECT
            COUNT(DISTINCT u.user_id)
        FROM
            users u
        LEFT JOIN blood_glucose_record b ON u.user_id = b.user_id AND b.deleted_at IS NULL
        WHERE 
            u.deleted_at IS NULL
    `

	index := 1
	execArgs := []any{}

	// 如果指定了医生ID，只查询该医生的患者
	if search.DoctorID != 0 {
		baseSQL += fmt.Sprintf(" AND u.relation_id = $%d", index)
		execArgs = append(execArgs, search.DoctorID)
		index++
	}

	baseSQL += " GROUP BY u.user_id"
	baseSQL += fmt.Sprintf(" HAVING COALESCE(MAX(b.upload_time), 0) < $%d", index)
	execArgs = append(execArgs, cutoffTimestamp)

	// 由于使用了 GROUP BY 和 HAVING，需要用子查询来计数
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM (%s) as inactive_users", baseSQL)

	var count uint
	err := db.QueryRow(countSQL, execArgs...).Scan(&count)
	return count, err
}
