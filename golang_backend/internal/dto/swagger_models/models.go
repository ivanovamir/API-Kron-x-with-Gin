package dto

type ErrorResponse struct {
	Error string `json:"error"`
}

type Token struct {
	Token string `json:"token"`
}

type DefaultData struct {
	Data interface{} `json:"data"`
}
