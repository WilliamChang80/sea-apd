package admin

// Admin ...
type Admin struct {
	Token    string `json:"token"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
