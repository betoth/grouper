package responses

type Subtopic struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Topic Topic  `json:"topic"`
}
