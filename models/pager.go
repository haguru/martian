package models

type PagerRequest struct {
	From        string `json:"from"`
	Destination string `json:"destination"`
	Message     string `json:"message"`
}

type PagerResponse struct {
	Success bool `json:"sucess"`
}
