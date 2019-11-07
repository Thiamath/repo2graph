package entities

type Error struct {
	ErrorFlag bool   `json:"error"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}
