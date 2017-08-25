package gosyncengine

// DeltaMessages contains a delta chunk of messages
type DeltaMessages struct {
	CursorStart string   `json:"cursor_start"`
	CursorEnd   string   `json:"cursor_end"`
	Deltas      Messages `json:"deltas"`
}
