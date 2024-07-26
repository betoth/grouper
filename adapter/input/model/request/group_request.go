package request

type GroupRequest struct {
	Name   string `json:"name" validate:"required,min=3,max=50"`
	UserID string `json:"user_id" validate:"required"`
}
