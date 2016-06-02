package modelUser

//User Model
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
	Token    string `json:"token, omitempty"`
}
