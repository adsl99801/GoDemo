package model

//BaseResponse o
type BaseResponse struct {
	Sus    bool
	Status int
	Msg    string
}

//LoginResponse o
type LoginResponse struct {
	BaseResponse
	Token string
}
