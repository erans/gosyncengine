package gosyncengine

// Participant contains information of a single participant in a message/conversation
type Participant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
