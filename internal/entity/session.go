package entity

// Session is used to represent a successful call to github.com/np-inprove/server/internal/auth/UseCase.Login.
// It contains the user information retrieved, as well as the token issued.
// The handlers should distribute the issued token to the client.
type Session struct {
	User  *User
	Token string
}
