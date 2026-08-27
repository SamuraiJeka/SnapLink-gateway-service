package dto


type Register struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Refresh struct {
	Refresh_token string `json:"refresh_token"`
}

type TokenResponse struct {
	Access_token string `json:"access_token"`
	Refresh_token string `json:"refresh_token"`
}

type ValidateRequest struct {
	Access_token string `json:"access_token"`
}

type ValidateResponse struct {
	Valid bool `json:"valid"`
	User_id uint64 `json:"user_id"`
}
