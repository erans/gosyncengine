package gosyncengine

// Thread contains information on a single thread
type Thread struct {
	ID                           string        `json:"id"`
	AccountID                    string        `json:"account_id"`
	Object                       string        `json:"object"`
	DraftIDs                     []string      `json:"draft_ids"`
	FirstMessageTimestamp        int           `json:"first_message_timestamp"`
	Folders                      []Folder      `json:"folders"`
	HasAttachments               bool          `json:"has_attachments"`
	LastMessageReceivedTimestamp int           `json:"last_message_received_timestamp"`
	LastMessageSentTimestamp     int           `json:"last_message_sent_timestamp"`
	LastMessageTimestamp         int           `json:"last_message_timestamp"`
	MessageIDs                   []string      `json:"message_ids"`
	Participants                 []Participant `json:"participants"`
	Snippet                      string        `json:"snippet"`
	Starred                      bool          `json:"starred"`
	Subject                      string        `json:"subject"`
	Unread                       bool          `json:"unread"`
	Version                      int           `json:"version"`
}

// Threads is a collection of Thread objects
type Threads []Thread
