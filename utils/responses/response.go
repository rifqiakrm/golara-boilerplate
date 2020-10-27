package responses

type ApiResponseList interface {
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

type apiResponseList struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (a apiResponseList) GetCode() int {
	return a.Code
}

func (a apiResponseList) GetMessage() string {
	return a.Message
}

func (a apiResponseList) GetData() interface{} {
	return a.Data
}

func SuccessApiResponseList(code int, message string, data interface{}) ApiResponseList {
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ErrorApiResponse(code int, message string) ApiResponseList {
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
