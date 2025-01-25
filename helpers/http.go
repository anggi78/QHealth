package helpers

type ErrorResponseJson struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponseJson struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponsePageJson struct {
	Status     bool               `json:"status"`
	Message    string             `json:"message"`
	Data       interface{}        `json:"data,omitempty"`
	Pagination PaginationResponse `json:"pagination"`
}

type PaginationResponse struct {
	CurrentPage int `json:"currentPage"`
	NextPage    int `json:"nextPage"`
	PrevPage    int `json:"prevPage"`
	AllPages    int `json:"allPages"`
}

func ErrorResponse(message string) ErrorResponseJson {
	return ErrorResponseJson{
		Status:  false,
		Message: message,
	}
}

func SuccessResponse(message string, data interface{}) SuccessResponseJson {
	return SuccessResponseJson{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func SuccessResponsePage(message string, data interface{}, pagination PaginationResponse) SuccessResponsePageJson {
	return SuccessResponsePageJson{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}
