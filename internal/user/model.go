package user

// User stores user information
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

// Booking stores booking information
type Booking struct {
	ID       string `json:"id"`
	UserID   string `json:"userId"`
	HostelID string `json:"hostelId"`
}
