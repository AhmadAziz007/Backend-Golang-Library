package response

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
