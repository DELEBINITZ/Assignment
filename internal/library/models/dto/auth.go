package dto

var (
	REGULAR = "regular"
	ADMIN   = "admin"
)

type LoginReq struct {
	Email string `json:"email"`
}

type LoginRes struct {
	Email       string `json:"email"`
	Status      string `json:"status"`
	AccessToken string `json:"access_token"`
}
