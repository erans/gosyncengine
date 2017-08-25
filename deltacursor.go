package gosyncengine

// DeltaCursor provides a cursor to the delta API
type DeltaCursor struct {
	Cursor string `json:"cursor"`
}
