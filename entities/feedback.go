package entities

import (
	"time"
)

type Feedback struct {
	FeedbackID uint
	Content    string
	File       string
	Status     string
	UserID     uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (f *Feedback) Insert() (*Feedback, error) {

	if err := db.QueryRow(`
		insert into feedbacks (content, file, status, user_id) values($1, $2, $3, $4) RETURNING feedback_id, 
		content, 
		file, 
		status, 
		user_id, 
		created_at, 
		updated_at
	`,
		f.Content,
		f.File,
		f.Status,
		f.UserID).Scan(
		&f.FeedbackID,
		&f.Content,
		&f.File,
		&f.Status,
		&f.UserID,
		&f.CreatedAt,
		&f.UpdatedAt); err != nil {
		return nil, err
	}

	return f, nil
}
