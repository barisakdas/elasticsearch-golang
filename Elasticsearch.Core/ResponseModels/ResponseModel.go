package responsemodels

type ResponseModel struct {
	Data       interface{}
	StatusCode int
	IsSucceded bool
	Errors     []string
}

func (r ResponseModel) Success(data interface{}) ResponseModel {
	return ResponseModel{
		Data:       data,
		StatusCode: 200,
		Errors:     nil,
		IsSucceded: true,
	}
}

func (r ResponseModel) Fail(errors []string, statusCode int) ResponseModel {
	return ResponseModel{
		Data:       nil,
		StatusCode: statusCode,
		Errors:     errors,
		IsSucceded: false,
	}
}
