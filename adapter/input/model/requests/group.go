package requests

type Group struct {
	Name       string `json:"name" validate:"required,min=3,max=50"`
	TopicID    string `json:"topic_id" validate:"required,uuid4"`
	SubtopicID string `json:"subtopic_id" validate:"required,uuid4"`
}
