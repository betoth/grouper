package request

type GroupRequest struct {
	Name       string `json:"name" validate:"required,min=3,max=50"`
	TopicID    string `json:"topic_id" validate:"required,min=3,max=50"`
	SubTopicID string `json:"subtopic_id" validate:"required,min=3,max=50"`
}
