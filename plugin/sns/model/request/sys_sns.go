package request

type SMSValidateCodeRequest struct {
	Mobile string `json:"mobile" form:"mobile" validate:"required"`
}

type SMSSendRequest struct {
	Mobile string `json:"mobile" form:"mobile" validate:"required"`
	Code   string `json:"code" form:"code" validate:"required"`
}
type LoginRequest struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Code   string `form:"code" json:"code" binding:"required"`
}
type SMSCodeRequest struct {
	Code string `json:"code" form:"code" validate:"required"`
}
