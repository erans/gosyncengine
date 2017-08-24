package gosyncengine

// Folder contains all the details on a specific folder
type Folder struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}
