package errors

var (
	ReadErr          = "Read body error"
	InvalidBody      = "Invalid body"
	UserExists       = "User with this login already exists"
	NoUser           = "No user with this login"
	InvalidPassword  = "Invalid password"
	NoAuth           = "No session found"
	InvalidToken     = "Invalid token"
	InvalidTokenBody = "Invalid token body"
	InternalError    = "Internal server error"
	InvalidParams    = "Invalid url parameters"
	NoComment        = "No comment with this id"
	NoAccess         = "No access to this action"
	NoPost           = "Post not found"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func New(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}
