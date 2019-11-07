package entities

type Error struct {
	Error     bool   `json:"error"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	error
}
