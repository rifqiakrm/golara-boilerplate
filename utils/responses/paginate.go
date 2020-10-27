package responses

type PaginateResponseList interface {
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

type Links struct {
	FirstPage string `json:"first_page"`
	LastPage  string `json:"last_page"`
	NextPage  string `json:"next_page"`
	PrevPage  string `json:"prev_page"`
}

type Pagination struct {
	Total int64 `json:"total"`
}

type PaginateResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func (a PaginateResponse) GetCode() int {
	return a.Code
}

func (a PaginateResponse) GetMessage() string {
	return a.Message
}

func (a PaginateResponse) GetData() interface{} {
	return a.Data
}

func PaginateApiResponseList(code int, message string, data interface{}, pagination Pagination) ApiResponseList {
	return &PaginateResponse{
		Code:       code,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}
