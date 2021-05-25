package error

type CusError struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Success bool `json:"success"`
}

func (e CusError) Error() string{
	return e.Msg
}

func NewCusError(code int, msg string) CusError{
	return CusError{
		code,
		msg,
		false,
	}
}