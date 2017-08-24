package gosyncengine

// Message contains all the details of a single message
type Message struct {
	ID        string        `json:"id"`
	AccountID string        `json:"account_id"`
	Object    string        `json:"object"`
	ThreadID  string        `json:"thread_id"`
	Date      int           `json:"date"`
	From      []Participant `json:"from"`
	To        []Participant `json:"to"`
	BCC       []Participant `json:"bcc"`
	Subject   string        `json:"subject"`
	Snippet   string        `json:"snippet"`
	Body      string        `json:"body"`
	Folder    Folder        `json:"folder"`
	ReplyTo   []Participant `json:"reply_to"`
	Starred   bool          `json:"starred"`
	Unread    bool          `json:"unread"`
}

// Messages is a list of message object
type Messages []Message
