package auth

// LoginRequest is the Login From
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
