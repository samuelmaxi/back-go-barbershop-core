package config

type RestError struct {
	Message string `json:"message"`
	Err     string `json:"error"`
	Code    int    `json:"code"`
	Causes  causes `json:"causes"`
}

type causes struct {
	Field   []string `json:"field"`
	Message string   `json:"message"`
}

func NewRestError() *RestError {
	return &RestError{}
}
