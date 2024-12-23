package constant

// UserData represents the user data extracted from the JWT token
type UserData struct {
	Email string
	Name  string
	Role  string
}
