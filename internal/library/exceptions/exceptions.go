package exceptions

type ServerErrorCode int

type ServerError struct {
	Code      ServerErrorCode
	Msg       string
	ActualErr error
	HttpCode  int
}

func (e ServerError) Error() string {
	return e.Msg
}

func (e ServerError) GetActualErr() error {
	return e.ActualErr
}

func (e ServerError) GetCode() ServerErrorCode {
	return e.Code
}
