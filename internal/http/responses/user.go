package responses

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func NewUserResponse(id, email, token string) UserResponse {
	return UserResponse{
		ID:    id,
		Email: email,
		Token: token,
	}
}
