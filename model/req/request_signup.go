package req

type ReqSignup struct {
	Email    string `json:"email" validate:"required,email"`
	Fullname string `json:"fullname" validate:"required"`
	Password string `json:"password" validate:"required"`
}
