package responses

type Topic struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Subtopic Subtopic `json:"Subtopic"`
}