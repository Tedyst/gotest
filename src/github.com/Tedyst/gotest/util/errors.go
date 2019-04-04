package util

//ErrorString has a string
type ErrorString struct {
	Str string
}

func (e *ErrorString) Error() string {
	return e.Str
}

type ErrorJSON struct {
	Message string `json:"Message"`
	Success bool   `json:"success"`
}

type Response struct {
	Message string `json:"Message"`
	Success bool   `json:"success"`
}
