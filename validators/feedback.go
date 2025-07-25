package validators

import "github.com/go-playground/validator/v10"

type FeedbackCreateRequest struct {
	Content string `json:"content"`
	File    string `json:"file"`
}

func (f *FeedbackCreateRequest) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(f)
}
